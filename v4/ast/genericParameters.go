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

var genericParametersClass = &genericParametersClass_{
	// Initialize class constants.
}

// Function

func GenericParameters() GenericParametersClassLike {
	return genericParametersClass
}

// CLASS METHODS

// Target

type genericParametersClass_ struct {
	// Define class constants.
}

// Constructors

func (c *genericParametersClass_) Make(parameters abs.Sequential[ParameterLike]) GenericParametersLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(parameters):
		panic("The parameters attribute is required by this class.")
	default:
		return &genericParameters_{
			// Initialize instance attributes.
			class_:      c,
			parameters_: parameters,
		}
	}
}

// INSTANCE METHODS

// Target

type genericParameters_ struct {
	// Define instance attributes.
	class_      GenericParametersClassLike
	parameters_ abs.Sequential[ParameterLike]
}

// Public

func (v *genericParameters_) GetClass() GenericParametersClassLike {
	return v.class_
}

// Attribute

func (v *genericParameters_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

// Private
