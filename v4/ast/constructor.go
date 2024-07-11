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
	mod "github.com/craterdog/go-collection-framework/v4"
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
	optionalParameters ParametersLike,
	abstraction AbstractionLike,
) ConstructorLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(name):
		panic("The name attribute is required for each Constructor.")
	case mod.IsUndefined(abstraction):
		panic("The abstraction attribute is required for each Constructor.")
	default:
		return &constructor_{
			// Initialize instance attributes.
			class_:              c,
			name_:               name,
			optionalParameters_: optionalParameters,
			abstraction_:        abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type constructor_ struct {
	// Define instance attributes.
	class_              ConstructorClassLike
	name_               string
	optionalParameters_ ParametersLike
	abstraction_        AbstractionLike
}

// Attributes

func (v *constructor_) GetClass() ConstructorClassLike {
	return v.class_
}

func (v *constructor_) GetName() string {
	return v.name_
}

func (v *constructor_) GetOptionalParameters() ParametersLike {
	return v.optionalParameters_
}

func (v *constructor_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
