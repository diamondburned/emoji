package data

import _ "embed"

var (
	//go:embed unicode.org/Public/emoji/1.0/emoji-data.txt
	_10EmojiDataTxtBytes []byte

	//go:embed unicode.org/Public/emoji/2.0/emoji-data.txt
	_20EmojiDataTxtBytes []byte
	//go:embed unicode.org/Public/emoji/2.0/emoji-sequences.txt
	_20EmojiSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/2.0/emoji-zwj-sequences.txt
	_20EmojiZwjSequencesTxtBytes []byte

	//go:embed unicode.org/Public/emoji/3.0/emoji-data.txt
	_30EmojiDataTxtBytes []byte
	//go:embed unicode.org/Public/emoji/3.0/emoji-sequences.txt
	_30EmojiSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/3.0/emoji-zwj-sequences.txt
	_30EmojiZwjSequencesTxtBytes []byte

	//go:embed unicode.org/Public/emoji/4.0/emoji-data.txt
	_40EmojiDataTxtBytes []byte
	//go:embed unicode.org/Public/emoji/4.0/emoji-sequences.txt
	_40EmojiSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/4.0/emoji-test.txt
	_40EmojiTestTxtBytes []byte
	//go:embed unicode.org/Public/emoji/4.0/emoji-zwj-sequences.txt
	_40EmojiZwjSequencesTxtBytes []byte

	//go:embed unicode.org/Public/emoji/5.0/emoji-data.txt
	_50EmojiDataTxtBytes []byte
	//go:embed unicode.org/Public/emoji/5.0/emoji-sequences.txt
	_50EmojiSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/5.0/emoji-test.txt
	_50EmojiTestTxtBytes []byte
	//go:embed unicode.org/Public/emoji/5.0/emoji-variation-sequences.txt
	_50EmojiVariationSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/5.0/emoji-zwj-sequences.txt
	_50EmojiZwjSequencesTxtBytes []byte

	//go:embed unicode.org/Public/emoji/11.0/emoji-data.txt
	_110EmojiDataTxtBytes []byte
	//go:embed unicode.org/Public/emoji/11.0/emoji-sequences.txt
	_110EmojiSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/11.0/emoji-test.txt
	_110EmojiTestTxtBytes []byte
	//go:embed unicode.org/Public/emoji/11.0/emoji-variation-sequences.txt
	_110EmojiVariationSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/11.0/emoji-zwj-sequences.txt
	_110EmojiZwjSequencesTxtBytes []byte

	//go:embed unicode.org/Public/emoji/12.0/emoji-data.txt
	_120EmojiDataTxtBytes []byte
	//go:embed unicode.org/Public/emoji/12.0/emoji-sequences.txt
	_120EmojiSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/12.0/emoji-test.txt
	_120EmojiTestTxtBytes []byte
	//go:embed unicode.org/Public/emoji/12.0/emoji-variation-sequences.txt
	_120EmojiVariationSequencesTxtBytes []byte
	//go:embed unicode.org/Public/emoji/12.0/emoji-zwj-sequences.txt
	_120EmojiZwjSequencesTxtBytes []byte
)

var bytesByVersionAndType = map[int]map[FileType][]byte{
	1: {
		Data: _10EmojiDataTxtBytes,
	},
	2: {
		Data:         _20EmojiDataTxtBytes,
		Sequences:    _20EmojiSequencesTxtBytes,
		ZWJSequences: _20EmojiZwjSequencesTxtBytes,
	},
	3: {
		Data:         _30EmojiDataTxtBytes,
		Sequences:    _30EmojiSequencesTxtBytes,
		ZWJSequences: _30EmojiZwjSequencesTxtBytes,
	},
	4: {
		Data:         _40EmojiDataTxtBytes,
		Sequences:    _40EmojiSequencesTxtBytes,
		Test_:        _40EmojiTestTxtBytes,
		ZWJSequences: _40EmojiZwjSequencesTxtBytes,
	},
	5: {
		Data:               _50EmojiDataTxtBytes,
		Sequences:          _50EmojiSequencesTxtBytes,
		Test_:              _50EmojiTestTxtBytes,
		VariationSequences: _50EmojiVariationSequencesTxtBytes,
		ZWJSequences:       _50EmojiZwjSequencesTxtBytes,
	},
	11: {
		Data:               _110EmojiDataTxtBytes,
		Sequences:          _110EmojiSequencesTxtBytes,
		Test_:              _110EmojiTestTxtBytes,
		VariationSequences: _110EmojiVariationSequencesTxtBytes,
		ZWJSequences:       _110EmojiZwjSequencesTxtBytes,
	},
	12: {
		Data:               _120EmojiDataTxtBytes,
		Sequences:          _120EmojiSequencesTxtBytes,
		Test_:              _120EmojiTestTxtBytes,
		VariationSequences: _120EmojiVariationSequencesTxtBytes,
		ZWJSequences:       _120EmojiZwjSequencesTxtBytes,
	},
}
