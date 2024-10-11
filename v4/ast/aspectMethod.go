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

func AspectMethod() AspectMethodClassLike {
	return aspectMethodClass
}

// Constructor Methods

func (c *aspectMethodClass_) Make(
	method MethodLike,
) AspectMethodLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &aspectMethod_{
		// Initialize the instance attributes.
		class_:  c,
		method_: method,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *aspectMethod_) GetMethod() MethodLike {
	return v.method_
}

// Public Methods

func (v *aspectMethod_) GetClass() AspectMethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *aspectMethod_) getClass() *aspectMethodClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type aspectMethod_ struct {
	// Declare the instance attributes.
	class_  *aspectMethodClass_
	method_ MethodLike
}

// Class Structure

type aspectMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

var aspectMethodClass = &aspectMethodClass_{
	// Initialize the class constants.
}
