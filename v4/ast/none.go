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
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var noneClass = &noneClass_{
	// Initialize class constants.
}

// Function

func None() NoneClassLike {
	return noneClass
}

// CLASS METHODS

// Target

type noneClass_ struct {
	// Define class constants.
}

// Constructors

func (c *noneClass_) Make(newline string) NoneLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(newline):
		panic("The newline attribute is required by this class.")
	default:
		return &none_{
			// Initialize instance attributes.
			class_:   c,
			newline_: newline,
		}
	}
}

// INSTANCE METHODS

// Target

type none_ struct {
	// Define instance attributes.
	class_   NoneClassLike
	newline_ string
}

// Public

func (v *none_) GetClass() NoneClassLike {
	return v.class_
}

// Attribute

func (v *none_) GetNewline() string {
	return v.newline_
}

// Private
