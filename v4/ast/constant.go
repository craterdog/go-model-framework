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

var constantClass = &constantClass_{
	// Initialize class constants.
}

// Function

func Constant() ConstantClassLike {
	return constantClass
}

// CLASS METHODS

// Target

type constantClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constantClass_) Make(
	name string,
	abstraction AbstractionLike,
) ConstantLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	case col.IsUndefined(abstraction):
		panic("The abstraction attribute is required by this class.")
	default:
		return &constant_{
			// Initialize instance attributes.
			class_:       c,
			name_:        name,
			abstraction_: abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type constant_ struct {
	// Define instance attributes.
	class_       ConstantClassLike
	name_        string
	abstraction_ AbstractionLike
}

// Attributes

func (v *constant_) GetClass() ConstantClassLike {
	return v.class_
}

func (v *constant_) GetName() string {
	return v.name_
}

func (v *constant_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
