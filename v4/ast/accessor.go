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

var accessorClass = &accessorClass_{
	// Initialize class constants.
}

// Function

func Accessor() AccessorClassLike {
	return accessorClass
}

// CLASS METHODS

// Target

type accessorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *accessorClass_) Make(
	any_ any,
) AccessorLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(any_):
		panic("The any attribute is required by this class.")
	default:
		return &accessor_{
			// Initialize instance attributes.
			class_: c,
			any_:   any_,
		}
	}
}

// INSTANCE METHODS

// Target

type accessor_ struct {
	// Define instance attributes.
	class_ AccessorClassLike
	any_   any
}

// Public

func (v *accessor_) GetClass() AccessorClassLike {
	return v.class_
}

// Attribute

func (v *accessor_) GetAny() any {
	return v.any_
}

// Private
