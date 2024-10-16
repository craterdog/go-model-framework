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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ConstructorMethod() ConstructorMethodClassLike {
	return constructorMethodReference()
}

// Constructor Methods

func (c *constructorMethodClass_) Make(
	name string,
	parameters abs.Sequential[ParameterLike],
	abstraction AbstractionLike,
) ConstructorMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &constructorMethod_{
		// Initialize the instance attributes.
		name_:        name,
		parameters_:  parameters,
		abstraction_: abstraction,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *constructorMethod_) GetName() string {
	return v.name_
}

func (v *constructorMethod_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *constructorMethod_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Public Methods

func (v *constructorMethod_) GetClass() ConstructorMethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *constructorMethod_) getClass() *constructorMethodClass_ {
	return constructorMethodReference()
}

// PRIVATE INTERFACE

// Instance Structure

type constructorMethod_ struct {
	// Declare the instance attributes.
	name_        string
	parameters_  abs.Sequential[ParameterLike]
	abstraction_ AbstractionLike
}

// Class Structure

type constructorMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constructorMethodReference() *constructorMethodClass_ {
	return constructorMethodReference_
}

var constructorMethodReference_ = &constructorMethodClass_{
	// Initialize the class constants.
}
