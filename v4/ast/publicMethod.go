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

func PublicMethod() PublicMethodClassLike {
	return publicMethodReference()
}

// Constructor Methods

func (c *publicMethodClass_) Make(
	method MethodLike,
) PublicMethodLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &publicMethod_{
		// Initialize the instance attributes.
		method_: method,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *publicMethod_) GetMethod() MethodLike {
	return v.method_
}

// Public Methods

func (v *publicMethod_) GetClass() PublicMethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *publicMethod_) getClass() *publicMethodClass_ {
	return publicMethodReference()
}

// PRIVATE INTERFACE

// Instance Structure

type publicMethod_ struct {
	// Declare the instance attributes.
	method_ MethodLike
}

// Class Structure

type publicMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func publicMethodReference() *publicMethodClass_ {
	return publicMethodReference_
}

var publicMethodReference_ = &publicMethodClass_{
	// Initialize the class constants.
}
