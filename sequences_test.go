package emoji_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/dmolesUC/emoji"
	"github.com/dmolesUC/emoji/data"
)

// TODO: figure out what (if anything) was *added* in each version
var samplesByVersionAndType = map[emoji.Version]map[data.SeqType]string{
	emoji.V1: {
		data.Emoji_Combining_Sequence: "9⃣", // note no variation selector
		data.Emoji_Flag_Sequence:      "🇿🇼",
	},
	emoji.V2: {
		data.Emoji_Flag_Sequence:      "🇿🇼",
		data.Emoji_Combining_Sequence: "9⃣", // note no variation selector
		data.Emoji_Modifier_Sequence:  "🤘🏿",
	},
	emoji.V3: {
		data.Emoji_Flag_Sequence:      "🇿🇼",
		data.Emoji_Combining_Sequence: "9️⃣",
		data.Emoji_Modifier_Sequence:  "🤾🏿",
		data.Emoji_ZWJ_Sequence:       "👩‍👩‍👧‍👧",
	},
	emoji.V4: {
		data.Emoji_Flag_Sequence:      "🇿🇼",
		data.Emoji_Combining_Sequence: "9️⃣",
		data.Emoji_Modifier_Sequence:  "🤾🏿",
		data.Emoji_ZWJ_Sequence:       "👁️‍🗨️",
	},
	emoji.V5: {
		data.Emoji_Flag_Sequence:     "🇿🇼",
		data.Emoji_Tag_Sequence:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		data.Emoji_Modifier_Sequence: "🧝🏿",
		data.Emoji_ZWJ_Sequence:      "👁️‍🗨️",
	},
	emoji.V11: {
		data.Emoji_Flag_Sequence:     "🇿🇼",
		data.Emoji_Tag_Sequence:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		data.Emoji_Modifier_Sequence: "🧝🏿",
		data.Emoji_ZWJ_Sequence:      "👁️‍🗨️",
	},
	emoji.V12: {
		data.Emoji_Modifier_Sequence: "🧝🏿",
		data.Emoji_Flag_Sequence:     "🇿🇼",
		data.Emoji_Tag_Sequence:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		data.Emoji_ZWJ_Sequence:      "👁️‍🗨️",
	},
}

func combinedSamples(seqType data.SeqType, v emoji.Version) []string {
	var combined []string
	for _, v2 := range emoji.AllVersions {
		if v2 >= v {
			break
		}
		if samplesByType, ok := samplesByVersionAndType[v]; ok {
			if sample, ok := samplesByType[seqType]; ok {
				combined = append(combined, sample)
			}
		}
	}
	return combined
}

func TestLegacySequences(t *testing.T) {
	for _, s := range data.AllSeqTypes {
		for _, v := range []emoji.Version{emoji.V1, emoji.V2} {
			t.Run(fmt.Sprintf("%s_%s", s, v), func(t *testing.T) {
				for _, samples := range combinedSamples(s, v) {
					ix := slices.Index(v.Sequences(s), samples)
					assert.NotEqual(t, -1, ix, "expected %v sequences for %v to include %#v (%X), but did not", t, v, s, []rune(s))
				}
			})
		}
	}
}

func TestSequences(t *testing.T) {
	for _, s := range data.AllSeqTypes {
		for _, v := range emoji.AllVersions {
			for _, samples := range combinedSamples(s, v) {
				ix := slices.Index(v.Sequences(s), samples)
				assert.NotEqual(t, -1, ix, "expected %v sequences for %v to include %#v (%X), but did not", s, v, samples, []rune(samples))
			}
		}
	}
}

func TestDisplayWidth(t *testing.T) {
	for _, v := range emoji.AllVersions {
		for _, s := range data.AllSeqTypes {
			t.Run(fmt.Sprintf("%s_%s", v, s), func(t *testing.T) {
				seqs := v.Sequences(s)
				for _, seq := range seqs {
					w := emoji.DisplayWidth(seq)
					assert.Equal(t, 1, w, "expected \"%v\" (%#v, %X) in %v (%v) to have length 1, but was %d", seq, seq, []rune(seq), v, s, w)
				}
			})
		}
	}
}
