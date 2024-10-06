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

var setterClass = &setterClass_{
	// Initialize class constants.
}

// Function

func Setter() SetterClassLike {
	return setterClass
}

// CLASS METHODS

// Target

type setterClass_ struct {
	// Define class constants.
}

// Constructors

func (c *setterClass_) Make(
	name string,
	parameter ParameterLike,
) SetterLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	default:
		return &setter_{
			// Initialize instance attributes.
			class_:     c,
			name_:      name,
			parameter_: parameter,
		}
	}
}

// INSTANCE METHODS

// Target

type setter_ struct {
	// Define instance attributes.
	class_     SetterClassLike
	name_      string
	parameter_ ParameterLike
}

// Public

func (v *setter_) GetClass() SetterClassLike {
	return v.class_
}

// Attribute

func (v *setter_) GetName() string {
	return v.name_
}

func (v *setter_) GetParameter() ParameterLike {
	return v.parameter_
}

// Private
