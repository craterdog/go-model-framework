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

var suffixClass = &suffixClass_{
	// Initialize class constants.
}

// Function

func Suffix() SuffixClassLike {
	return suffixClass
}

// CLASS METHODS

// Target

type suffixClass_ struct {
	// Define class constants.
}

// Constructors

func (c *suffixClass_) Make(name string) SuffixLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	default:
		return &suffix_{
			// Initialize instance attributes.
			class_: c,
			name_:  name,
		}
	}
}

// INSTANCE METHODS

// Target

type suffix_ struct {
	// Define instance attributes.
	class_ SuffixClassLike
	name_  string
}

// Public

func (v *suffix_) GetClass() SuffixClassLike {
	return v.class_
}

// Attribute

func (v *suffix_) GetName() string {
	return v.name_
}

// Private
