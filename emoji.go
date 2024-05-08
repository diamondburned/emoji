package emoji

import (
	"fmt"
	"unicode"

	"github.com/puzpuzpuz/xsync/v3"
	"libdb.so/go-emoji/data"
)

// ZWJ is the Unicode zero-width join character
const ZWJ = '\u200d'

// IsEmoji returns true if the specified rune has the (single-character)
// Emoji property in the latest Emoji version, false otherwise
func IsEmoji(r rune) bool {
	return r > 0x7F && unicode.Is(Latest.RangeTable(data.Emoji), r)
}

// DisplayWidth attempts to guess at the display width of a string containing
// emoji, taking into account variation selectors (0xFE00-0xFE0F), zero-width
// joins (0x200D), combining diacritical marks (0x20d0-0x20ff), flags,
// and skin tone modifiers.
func DisplayWidth(str string) int {
	width := 0
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if isZeroWidth(r) {
			continue
		}
		if i > 0 && runes[i-1] == ZWJ {
			// ZWJ effectively "suppresses" the next character
			continue
		}
		if i > 0 && isRegionalIndicator(r) && isRegionalIndicator(runes[i-1]) {
			// only count first flag character in a sequence
			continue
		}
		if i > 0 && isSkinToneModifier(r) && IsEmoji(runes[i-1]) {
			// don't count skin tone modifier when it's modifying something
			continue
		}
		width += 1
	}
	return width
}

// ------------------------------------------------------------
// Version type

// Version represents an Emoji major release, e.g. V5 for Emoji version 5.0.
// Note that starting at Emoji version 11.0, the Emoji version is synchronized
// to the corresponding Unicode version, so there are no versions 6-10.
type Version int

const (
	V1 Version = 1
	V2 Version = 2
	V3 Version = 3
	V4 Version = 4
	V5 Version = 5
	// Starting at V11, Emoji version = Unicode version
	V11 Version = 11
	V12 Version = 12

	Latest = V12
)

// AllVersions lists all emoji versions in order.
var AllVersions = []Version{V1, V2, V3, V4, V5, V11, V12}

// String returns this version as a string, e.g. V4.String() -> "Emoji 4.0"
func (v Version) String() string {
	return fmt.Sprintf("Emoji %d.0", int(v))
}

// HasFile returns true if this version has a file of the specified type, false
// otherwise. E.g., ZWJ (zero width joiner) sequences were introduced only in
// Emoji version 2.0, test files in version 4.0, and variation sequences in version
// 5.0.
func (v Version) HasFile(t data.FileType) bool {
	return t.HasData(int(v))
}

// FileBytes returns the byte data of the Unicode.org source file of the specified type
// for this version, e.g. V12.FileBytes(Sequences) returns the contents of the file
// http://unicode.org/Public/emoji/12.0/emoji-sequences.txt
func (v Version) FileBytes(t data.FileType) []byte {
	bytes, err := t.GetBytes(int(v))
	if err == nil {
		return bytes
	}
	return nil
}

// RangeTable returns the Unicode range table for characters with the specified property
// in this Emoji version. Note that the range table reflects the ranges as defined in the
// source files from Unicode.org; ranges are guaranteed not to overlap, as per the RangeTable
// docs, but adjacent ranges are not coalesced.
func (v Version) RangeTable(property data.Property) *unicode.RangeTable {
	rtsByProperty, _ := rangeTables.LoadOrCompute(v, func() *xsync.MapOf[data.Property, *unicode.RangeTable] {
		return xsync.NewMapOfPresized[data.Property, *unicode.RangeTable](len(data.AllProperties))
	})
	rt, _ := rtsByProperty.LoadOrCompute(property, func() *unicode.RangeTable {
		return data.ParseRangeTable(property, v.FileBytes(data.Data))
	})
	return rt
}

// Sequences returns the Unicode emoji sequences of the specified type in this Emoji version.
func (v Version) Sequences(seqType data.SeqType) []string {
	seqsByType, _ := sequences.LoadOrCompute(v, func() *xsync.MapOf[data.SeqType, []string] {
		return xsync.NewMapOfPresized[data.SeqType, []string](len(data.AllSeqTypes))
	})
	seqs, _ := seqsByType.LoadOrCompute(seqType, func() []string {
		var parseSeq data.ParseSeq
		if v == V1 || v == V2 {
			parseSeq = data.ParseSequencesLegacy
		} else {
			parseSeq = data.ParseSequences
		}

		var fileType data.FileType
		if v == V1 {
			fileType = data.Data
		} else if seqType == data.Emoji_ZWJ_Sequence {
			fileType = data.ZWJSequences
		} else {
			fileType = data.Sequences
		}

		return parseSeq(seqType, v.FileBytes(fileType))
	})
	return seqs
}

// ------------------------------------------------------------
// Unexported symbols

var (
	rangeTables = xsync.NewMapOf[Version, *xsync.MapOf[data.Property, *unicode.RangeTable]]()
	sequences   = xsync.NewMapOf[Version, *xsync.MapOf[data.SeqType, []string]]()
)

func isRegionalIndicator(r rune) bool {
	return unicode.Is(data.RegionalIndicator, r)
}

func isSkinToneModifier(r rune) bool {
	return unicode.Is(data.EmojiSkinToneModifier, r)
}

func isZeroWidth(r rune) bool {
	return r == ZWJ ||
		unicode.Is(unicode.Variation_Selector, r) ||
		unicode.Is(data.CombiningDiacritical, r) ||
		unicode.Is(data.Tag, r)
}
