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

func Result() ResultClassLike {
	return resultReference()
}

// Constructor Methods

func (c *resultClass_) Make(
	any_ any,
) ResultLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &result_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *result_) GetAny() any {
	return v.any_
}

// Public Methods

func (v *result_) GetClass() ResultClassLike {
	return v.getClass()
}

// Private Methods

func (v *result_) getClass() *resultClass_ {
	return resultReference()
}

// PRIVATE INTERFACE

// Instance Structure

type result_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type resultClass_ struct {
	// Declare the class constants.
}

// Class Reference

func resultReference() *resultClass_ {
	return resultReference_
}

var resultReference_ = &resultClass_{
	// Initialize the class constants.
}
