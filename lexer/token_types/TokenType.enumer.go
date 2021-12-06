//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package token_types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type TokenType string
type TokenTypes []TokenType

func Slice(es ...TokenType) TokenTypes {
	var s TokenTypes

	for _, e := range es {
		s = append(s, e)
	}

	return s
}

const (
	AlternationEnd   TokenType = "alternation-end"
	AlternationStart TokenType = "alternation-start"
	BinVal           TokenType = "bin-val"
	CNl              TokenType = "c-nl"
	CWsp             TokenType = "c-wsp"
	CharVal          TokenType = "char-val"
	Comment          TokenType = "comment"
	Concatentation   TokenType = "concatentation"
	DecVal           TokenType = "dec-val"
	DefinedAs        TokenType = "defined-as"
	Eof              TokenType = "eof"
	GroupEnd         TokenType = "group-end"
	GroupStart       TokenType = "group-start"
	HexVal           TokenType = "hex-val"
	LineEnd          TokenType = "line-end"
	LineStart        TokenType = "line-start"
	NumVal           TokenType = "num-val"
	OptionEnd        TokenType = "option-end"
	OptionStart      TokenType = "option-start"
	ProseVal         TokenType = "prose-val"
	Repeat           TokenType = "repeat"
	Repetition       TokenType = "repetition"
	Rulename         TokenType = "rulename"
	SyntaxError      TokenType = "syntax-error"
)

var (
	All = TokenTypes{
		AlternationEnd,
		AlternationStart,
		BinVal,
		CNl,
		CWsp,
		CharVal,
		Comment,
		Concatentation,
		DecVal,
		DefinedAs,
		Eof,
		GroupEnd,
		GroupStart,
		HexVal,
		LineEnd,
		LineStart,
		NumVal,
		OptionEnd,
		OptionStart,
		ProseVal,
		Repeat,
		Repetition,
		Rulename,
		SyntaxError,
	}
)

func Is(v string) bool {
	return All.Is(v)
}

func Parse(v string) (TokenType, error) {
	return All.Parse(v)
}

func Strings() []string {
	return All.Strings()
}

func (t TokenType) String() string {
	return string(t)
}

var ErrTokenTypeInvalid = errors.New("invalid enumeration type")

func Error(vs TokenTypes, v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrTokenTypeInvalid, v, strings.Join(vs.Strings(), ","),
	)
}

func (t TokenTypes) Strings() []string {
	var ss []string

	for _, v := range t {
		ss = append(ss, v.String())
	}

	return ss
}

func (t TokenTypes) Parse(v string) (TokenType, error) {
	var o TokenType
	var f bool
	n := strings.ToLower(v)

	for _, e := range t {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, Error(t, v)
	}

	return o, nil
}

func (t TokenTypes) Is(v string) bool {
	var f bool

	for _, e := range t {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func (t TokenTypes) Contains(v TokenType) bool {
	for _, e := range t {
		if e == v {
			return true
		}
	}

	return false
}

func (t TokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *TokenType) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}