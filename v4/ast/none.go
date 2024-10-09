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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func None() NoneClassLike {
	return noneClass
}

// Constructor Methods

func (c *noneClass_) Make(
	newline string,
) NoneLike {
	if uti.IsUndefined(newline) {
		panic("The newline attribute is required by this class.")
	}
	var instance = &none_{
		class_:   c,
		newline_: newline,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *none_) GetNewline() string {
	return v.newline_
}

// Public Methods

func (v *none_) GetClass() NoneClassLike {
	return v.getClass()
}

// Private Methods

func (v *none_) getClass() *noneClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type none_ struct {
	class_   *noneClass_
	newline_ string
}

// Class Structure

type noneClass_ struct {
	// Define the class constants.
}

// Class Reference

var noneClass = &noneClass_{
	// Initialize the class constants.
}
