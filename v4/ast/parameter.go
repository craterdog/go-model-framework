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
)

// CLASS ACCESS

// Reference

var parameterClass = &parameterClass_{
	// Initialize class constants.
}

// Function

func Parameter() ParameterClassLike {
	return parameterClass
}

// CLASS METHODS

// Target

type parameterClass_ struct {
	// Define class constants.
}

// Constructors

func (c *parameterClass_) Make(
	name string,
	abstraction AbstractionLike,
) ParameterLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required for each Parameter.")
	case col.IsUndefined(abstraction):
		panic("The abstraction attribute is required for each Parameter.")
	default:
		return &parameter_{
			// Initialize instance attributes.
			class_:       c,
			name_:        name,
			abstraction_: abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type parameter_ struct {
	// Define instance attributes.
	class_       ParameterClassLike
	name_        string
	abstraction_ AbstractionLike
}

// Attributes

func (v *parameter_) GetClass() ParameterClassLike {
	return v.class_
}

func (v *parameter_) GetName() string {
	return v.name_
}

func (v *parameter_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
