package standard

import (
	. "github.com/balzaczyy/golucene/core/analysis/tokenattributes"
	"io"
)

// standard/StandardTokenizerImpl.java

/* initial size of the lookahead buffer */
const ZZ_BUFFERSIZE = 4096

/* lexical states */
const YYINITIAL = 0

/*
This class implements Word Break rules from the Unicode Text
Segmentation algorithm, as specified in Unicode Standard Annex #29.

Tokens produced are of the following types:

	- <ALPHANUM>: A sequence of alphabetic and numeric characters
	- <NUM>: A number
	- <SOUTHEAST_ASIAN>: A sequence of characters from South and Southeast Asian languages, including Thai, Lao, Myanmar, and Khmer
	- IDEOGRAPHIC>: A single CJKV ideographic character
	- <HIRAGANA>: A single hiragana character

Technically it should auto generated by JFlex but there is no GoFlex
yet. So it's a line-by-line port.
*/
type StandardTokenizerImpl struct {
	// the input device
	zzReader io.ReadCloser

	// the current lexical state
	zzLexicalState int

	// this buffer contains the current text to be matched and is the
	// source of yytext() string
	zzBuffer []rune

	// the text position at the last accepting state
	zzMarkedPos int

	// the current text position in the buffer
	zzCurrentPos int

	// startRead marks the beginning of the yytext() string in the buffer
	zzStartRead int

	// endRead marks the last character in the buffer, that has been read from input
	zzEndRead int

	// number of newlines encountered up to the start of the matched text
	yyline int

	// the number of characters up to the start of the matched text
	_yychar int

	// the number of characters from the last newline up to the start of the matched text
	yycolumn int

	// zzAtBOL == true <=> the scanner is currently at the beginning of a line
	zzAtBOL bool

	// zzAtEOF == true <=> the scanner is at the EOF
	zzAtEOF bool

	// denotes if the user-EOF-code has already been executed
	zzEOFDone bool
}

func newStandardTokenizerImpl(in io.ReadCloser) *StandardTokenizerImpl {
	return &StandardTokenizerImpl{
		zzReader:       in,
		zzLexicalState: YYINITIAL,
		zzBuffer:       make([]rune, ZZ_BUFFERSIZE),
		zzAtBOL:        true,
	}
}

func (t *StandardTokenizerImpl) yychar() int {
	return t._yychar
}

/* Fills CharTermAttribute with the current token text. */
func (t *StandardTokenizerImpl) text(tt CharTermAttribute) {
	panic("not implemented yet")
}

/*
Resets the scanner to read from a new input stream.
Does not close the old reader.

All internal variables are reset, the old input stream
cannot be reused (internal buffer is discarded and lost).
Lexical state is set to ZZ_INITIAL.

Internal scan buffer is resized down to its initial length, if it has grown.
*/
func (t *StandardTokenizerImpl) yyreset(reader io.ReadCloser) {
	t.zzReader = reader
	t.zzAtBOL = true
	t.zzAtEOF = false
	t.zzEOFDone = false
	t.zzEndRead, t.zzStartRead = 0, 0
	t.zzCurrentPos, t.zzMarkedPos = 0, 0
	t.yyline, t._yychar, t.yycolumn = 0, 0, 0
	t.zzLexicalState = YYINITIAL
	if len(t.zzBuffer) > ZZ_BUFFERSIZE {
		t.zzBuffer = make([]rune, ZZ_BUFFERSIZE)
	}
}

/* Returns the length of the matched text region. */
func (t *StandardTokenizerImpl) yylength() int {
	return t.zzMarkedPos - t.zzStartRead
}

/*
Resumes scanning until the next regular expression is matched, the
end of input is encountered or an I/O-Error occurs.
*/
func (t *StandardTokenizerImpl) nextToken() (int, error) {
	panic("not implemented yet")
}