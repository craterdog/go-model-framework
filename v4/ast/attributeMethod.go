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

func AttributeMethod() AttributeMethodClassLike {
	return attributeMethodClass
}

// Constructor Methods

func (c *attributeMethodClass_) Make(
	any_ any,
) AttributeMethodLike {
	if uti.IsUndefined(any_) {
		panic("The any attribute is required by this class.")
	}
	var instance = &attributeMethod_{
		class_: c,
		any_:   any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *attributeMethod_) GetAny() any {
	return v.any_
}

// Public Methods

func (v *attributeMethod_) GetClass() AttributeMethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *attributeMethod_) getClass() *attributeMethodClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type attributeMethod_ struct {
	class_ *attributeMethodClass_
	any_   any
}

// Class Structure

type attributeMethodClass_ struct {
	// Define the class constants.
}

// Class Reference

var attributeMethodClass = &attributeMethodClass_{
	// Initialize the class constants.
}
