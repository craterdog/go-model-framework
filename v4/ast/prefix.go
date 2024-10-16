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

func Prefix() PrefixClassLike {
	return prefixReference()
}

// Constructor Methods

func (c *prefixClass_) Make(
	any_ any,
) PrefixLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &prefix_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *prefix_) GetAny() any {
	return v.any_
}

// Public Methods

func (v *prefix_) GetClass() PrefixClassLike {
	return v.getClass()
}

// Private Methods

func (v *prefix_) getClass() *prefixClass_ {
	return prefixReference()
}

// PRIVATE INTERFACE

// Instance Structure

type prefix_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type prefixClass_ struct {
	// Declare the class constants.
}

// Class Reference

func prefixReference() *prefixClass_ {
	return prefixReference_
}

var prefixReference_ = &prefixClass_{
	// Initialize the class constants.
}
