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

var parameterizedClass = &parameterizedClass_{
	// Initialize class constants.
}

// Function

func Parameterized() ParameterizedClassLike {
	return parameterizedClass
}

// CLASS METHODS

// Target

type parameterizedClass_ struct {
	// Define class constants.
}

// Constructors

func (c *parameterizedClass_) Make(parameters ParametersLike) ParameterizedLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(parameters):
		panic("The parameters attribute is required for each Parameterized.")
	default:
		return &parameterized_{
			// Initialize instance attributes.
			class_:      c,
			parameters_: parameters,
		}
	}
}

// INSTANCE METHODS

// Target

type parameterized_ struct {
	// Define instance attributes.
	class_      ParameterizedClassLike
	parameters_ ParametersLike
}

// Attributes

func (v *parameterized_) GetClass() ParameterizedClassLike {
	return v.class_
}

func (v *parameterized_) GetParameters() ParametersLike {
	return v.parameters_
}

// Private
