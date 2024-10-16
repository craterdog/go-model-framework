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

func SetterMethod() SetterMethodClassLike {
	return setterMethodReference()
}

// Constructor Methods

func (c *setterMethodClass_) Make(
	name string,
	parameter ParameterLike,
) SetterMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameter) {
		panic("The \"parameter\" attribute is required by this class.")
	}
	var instance = &setterMethod_{
		// Initialize the instance attributes.
		name_:      name,
		parameter_: parameter,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *setterMethod_) GetName() string {
	return v.name_
}

func (v *setterMethod_) GetParameter() ParameterLike {
	return v.parameter_
}

// Public Methods

func (v *setterMethod_) GetClass() SetterMethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *setterMethod_) getClass() *setterMethodClass_ {
	return setterMethodReference()
}

// PRIVATE INTERFACE

// Instance Structure

type setterMethod_ struct {
	// Declare the instance attributes.
	name_      string
	parameter_ ParameterLike
}

// Class Structure

type setterMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func setterMethodReference() *setterMethodClass_ {
	return setterMethodReference_
}

var setterMethodReference_ = &setterMethodClass_{
	// Initialize the class constants.
}
