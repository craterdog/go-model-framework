/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package grammar

import (
	fmt "fmt"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	reg "regexp"
	sts "strings"
	uni "unicode"
)

// CLASS INTERFACE

// Access Function

func Scanner() ScannerClassLike {
	return scannerReference()
}

// Constructor Methods

func (c *scannerClass_) Make(
	source string,
	tokens abs.QueueLike[TokenLike],
) ScannerLike {
	if uti.IsUndefined(source) {
		panic("The \"source\" attribute is required by this class.")
	}
	if uti.IsUndefined(tokens) {
		panic("The \"tokens\" attribute is required by this class.")
	}
	var instance = &scanner_{
		// Initialize the instance attributes.
		line_:     1,
		position_: 1,
		runes_:    []rune(source),
		tokens_:   tokens,
	}
	go instance.scanTokens() // Do scanning in the background...
	return instance
}

// Function Methods

func (c *scannerClass_) FormatToken(token TokenLike) string {
	var result_ string
	var value = token.GetValue()
	value = fmt.Sprintf("%q", value)
	if len(value) > 40 {
		value = fmt.Sprintf("%.40q...", value)
	}
	result_ = fmt.Sprintf(
		"Token [type: %s, line: %d, position: %d]: %s",
		c.tokens_[token.GetType()],
		token.GetLine(),
		token.GetPosition(),
		value,
	)
	return result_
}

func (c *scannerClass_) FormatType(tokenType TokenType) string {
	var result_ = c.tokens_[tokenType]
	return result_
}

func (c *scannerClass_) MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var result_ bool
	var matcher = c.matchers_[tokenType]
	var match = matcher.FindString(tokenValue)
	result_ = uti.IsDefined(match)
	return result_
}

// INSTANCE INTERFACE

// Public Methods

func (v *scanner_) GetClass() ScannerClassLike {
	return v.getClass()
}

// Private Methods

func (v *scanner_) getClass() *scannerClass_ {
	return scannerReference()
}

func (v *scanner_) emitToken(tokenType TokenType) {
	var value = string(v.runes_[v.first_:v.next_])
	switch value {
	case "\x00":
		value = "<NULL>"
	case "\a":
		value = "<BELL>"
	case "\b":
		value = "<BKSP>"
	case "\t":
		value = "<HTAB>"
	case "\f":
		value = "<FMFD>"
	case "\r":
		value = "<CRTN>"
	case "\v":
		value = "<VTAB>"
	}
	var token = Token().Make(v.line_, v.position_, tokenType, value)
	//fmt.Println(Scanner().FormatToken(token)) // Uncomment when debugging.
	v.tokens_.AddValue(token) // This will block if the queue is full.
}

func (v *scanner_) foundError() {
	v.next_++
	v.emitToken(ErrorToken)
}

func (v *scanner_) foundToken(tokenType TokenType) bool {
	// Attempt to match the specified token type.
	var text = string(v.runes_[v.next_:])
	var matcher = v.getClass().matchers_[tokenType]
	var match = matcher.FindString(text)
	if len(match) == 0 {
		return false
	}

	// Check for false delimiter matches.
	var token = []rune(match)
	var length = uint(len(token))
	var previous = token[length-1]
	if tokenType == DelimiterToken && uint(len(v.runes_)) > v.next_+length {
		var next = v.runes_[v.next_+length]
		if (uni.IsLetter(previous) || uni.IsNumber(previous)) &&
			(uni.IsLetter(next) || uni.IsNumber(next) || next == '_') {
			return false
		}
	}

	// Found the requested token type.
	v.next_ += length
	v.emitToken(tokenType)
	var count = uint(sts.Count(match, "\n"))
	if count > 0 {
		v.line_ += count
		v.position_ = v.indexOfLastEol(token)
	} else {
		v.position_ += v.next_ - v.first_
	}
	v.first_ = v.next_
	return true
}

func (v *scanner_) indexOfLastEol(runes []rune) uint {
	var length = uint(len(runes))
	for index := length; index > 0; index-- {
		if runes[index-1] == '\n' {
			return length - index + 1
		}
	}
	return 0
}

func (v *scanner_) scanTokens() {
loop:
	for v.next_ < uint(len(v.runes_)) {
		switch {
		// Find the next token type.
		case v.foundToken(CommentToken):
		case v.foundToken(DelimiterToken):
		case v.foundToken(NameToken):
		case v.foundToken(NewlineToken):
		case v.foundToken(PathToken):
		case v.foundToken(SpaceToken):
		default:
			v.foundError()
			break loop
		}
	}
	v.tokens_.CloseQueue()
}

// PRIVATE INTERFACE

// Instance Structure

type scanner_ struct {
	// Declare the instance attributes.
	first_    uint // A zero based index of the first possible rune in the next token.
	next_     uint // A zero based index of the next possible rune in the next token.
	line_     uint // The line number in the source string of the next rune.
	position_ uint // The position in the current line of the next rune.
	runes_    []rune
	tokens_   abs.QueueLike[TokenLike]
}

// Class Structure

type scannerClass_ struct {
	// Declare the class constants.
	tokens_   map[TokenType]string
	matchers_ map[TokenType]*reg.Regexp
}

// Class Reference

func scannerReference() *scannerClass_ {
	return scannerReference_
}

var scannerReference_ = &scannerClass_{
	// Initialize the class constants.
	tokens_: map[TokenType]string{
		ErrorToken:     "error",
		CommentToken:   "comment",
		DelimiterToken: "delimiter",
		NameToken:      "name",
		NewlineToken:   "newline",
		PathToken:      "path",
		SpaceToken:     "space",
	},
	matchers_: map[TokenType]*reg.Regexp{
		CommentToken:   reg.MustCompile("^" + comment_),
		DelimiterToken: reg.MustCompile("^" + delimiter_),
		NameToken:      reg.MustCompile("^" + name_),
		NewlineToken:   reg.MustCompile("^" + newline_),
		PathToken:      reg.MustCompile("^" + path_),
		SpaceToken:     reg.MustCompile("^" + space_),
	},
}

// Private Constants

/*
NOTE:
These private constants define the regular expression sub-patterns that make up
the intrinsic types and token types.  Unfortunately there is no way to make them
private to the scanner class since they must be TRUE Go constants to be combined
in this way.  We append an underscore to each name to lessen the chance of a
name collision with other private Go class constants in this package.
*/
const (
	// Define the regular expression patterns for each intrinsic type.
	any_     = "." // This does NOT include newline characters.
	control_ = "\\p{Cc}"
	digit_   = "\\p{Nd}"
	eol_     = "\\r?\\n"
	lower_   = "\\p{Ll}"
	upper_   = "\\p{Lu}"

	// Define the regular expression patterns for each token type.
	comment_   = "(?:/\\*" + eol_ + "(" + any_ + "|" + eol_ + ")*?" + eol_ + "\\*/" + eol_ + ")"
	delimiter_ = "(?:type|package|map|iota|interface|import|func|const|chan|\\}|\\{|\\]|\\[|\\.|\\)|\\(|=|// Type Definitions|// Public Methods|// Instance Definitions|// Functional Definitions|// Function Methods|// Constructor Methods|// Constant Methods|// Class Definitions|// Attribute Methods|// Aspect Definitions|// Aspect Methods|,)"
	name_      = "(?:(" + lower_ + "|" + upper_ + ")(" + lower_ + "|" + upper_ + "|" + digit_ + ")*_?)"
	newline_   = "(?:\\r?\\n)"
	path_      = "(?:\"" + any_ + "*?\")"
	space_     = "(?:[ \\t]+)"
)
