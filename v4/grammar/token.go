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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Token() TokenClassLike {
	return tokenReference()
}

// Constructor Methods

func (c *tokenClass_) Make(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	if uti.IsUndefined(line) {
		panic("The \"line\" attribute is required by this class.")
	}
	if uti.IsUndefined(position) {
		panic("The \"position\" attribute is required by this class.")
	}
	if uti.IsUndefined(type_) {
		panic("The \"type\" attribute is required by this class.")
	}
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	var instance = &token_{
		// Initialize the instance attributes.
		line_:     line,
		position_: position,
		type_:     type_,
		value_:    value,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *token_) GetLine() uint {
	return v.line_
}

func (v *token_) GetPosition() uint {
	return v.position_
}

func (v *token_) GetType() TokenType {
	return v.type_
}

func (v *token_) GetValue() string {
	return v.value_
}

// Public Methods

func (v *token_) GetClass() TokenClassLike {
	return v.getClass()
}

// Private Methods

func (v *token_) getClass() *tokenClass_ {
	return tokenReference()
}

// PRIVATE INTERFACE

// Instance Structure

type token_ struct {
	// Declare the instance attributes.
	line_     uint
	position_ uint
	type_     TokenType
	value_    string
}

// Class Structure

type tokenClass_ struct {
	// Declare the class constants.
}

// Class Reference

func tokenReference() *tokenClass_ {
	return tokenReference_
}

var tokenReference_ = &tokenClass_{
	// Initialize the class constants.
}
