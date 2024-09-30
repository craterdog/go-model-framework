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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var constructorClass = &constructorClass_{
	// Initialize class constants.
}

// Function

func Constructor() ConstructorClassLike {
	return constructorClass
}

// CLASS METHODS

// Target

type constructorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constructorClass_) Make(
	name string,
	parameters abs.Sequential[ParameterLike],
	abstraction AbstractionLike,
) ConstructorLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	case col.IsUndefined(parameters):
		panic("The parameters attribute is required by this class.")
	case col.IsUndefined(abstraction):
		panic("The abstraction attribute is required by this class.")
	default:
		return &constructor_{
			// Initialize instance attributes.
			class_:       c,
			name_:        name,
			parameters_:  parameters,
			abstraction_: abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type constructor_ struct {
	// Define instance attributes.
	class_       ConstructorClassLike
	name_        string
	parameters_  abs.Sequential[ParameterLike]
	abstraction_ AbstractionLike
}

// Public

func (v *constructor_) GetClass() ConstructorClassLike {
	return v.class_
}

// Attribute

func (v *constructor_) GetName() string {
	return v.name_
}

func (v *constructor_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *constructor_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
