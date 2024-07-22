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

var prefixClass = &prefixClass_{
	// Initialize class constants.
}

// Function

func Prefix() PrefixClassLike {
	return prefixClass
}

// CLASS METHODS

// Target

type prefixClass_ struct {
	// Define class constants.
}

// Constructors

func (c *prefixClass_) Make(any_ any) PrefixLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &prefix_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type prefix_ struct {
	// Define instance attributes.
	class_ PrefixClassLike
	any_   any
}

// Attributes

func (v *prefix_) GetClass() PrefixClassLike {
	return v.class_
}

func (v *prefix_) GetAny() any {
	return v.any_
}

// Private
