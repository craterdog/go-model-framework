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

func AdditionalArgument() AdditionalArgumentClassLike {
	return additionalArgumentReference()
}

// Constructor Methods

func (c *additionalArgumentClass_) Make(
	argument ArgumentLike,
) AdditionalArgumentLike {
	if uti.IsUndefined(argument) {
		panic("The \"argument\" attribute is required by this class.")
	}
	var instance = &additionalArgument_{
		// Initialize the instance attributes.
		argument_: argument,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *additionalArgument_) GetArgument() ArgumentLike {
	return v.argument_
}

// Public Methods

func (v *additionalArgument_) GetClass() AdditionalArgumentClassLike {
	return v.getClass()
}

// Private Methods

func (v *additionalArgument_) getClass() *additionalArgumentClass_ {
	return additionalArgumentReference()
}

// PRIVATE INTERFACE

// Instance Structure

type additionalArgument_ struct {
	// Declare the instance attributes.
	argument_ ArgumentLike
}

// Class Structure

type additionalArgumentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalArgumentReference() *additionalArgumentClass_ {
	return additionalArgumentReference_
}

var additionalArgumentReference_ = &additionalArgumentClass_{
	// Initialize the class constants.
}
