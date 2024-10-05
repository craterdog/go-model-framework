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

var constraintClass = &constraintClass_{
	// Initialize class constants.
}

// Function

func Constraint() ConstraintClassLike {
	return constraintClass
}

// CLASS METHODS

// Target

type constraintClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constraintClass_) Make(
	name string,
	abstraction AbstractionLike,
) ConstraintLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	case col.IsUndefined(abstraction):
		panic("The abstraction attribute is required by this class.")
	default:
		return &constraint_{
			// Initialize instance attributes.
			class_:       c,
			name_:        name,
			abstraction_: abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type constraint_ struct {
	// Define instance attributes.
	class_       ConstraintClassLike
	name_        string
	abstraction_ AbstractionLike
}

// Public

func (v *constraint_) GetClass() ConstraintClassLike {
	return v.class_
}

// Attribute

func (v *constraint_) GetName() string {
	return v.name_
}

func (v *constraint_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
