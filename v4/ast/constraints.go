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

var constraintsClass = &constraintsClass_{
	// Initialize class constants.
}

// Function

func Constraints() ConstraintsClassLike {
	return constraintsClass
}

// CLASS METHODS

// Target

type constraintsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constraintsClass_) Make(
	constraint ConstraintLike,
	additionalConstraints abs.Sequential[AdditionalConstraintLike],
) ConstraintsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(constraint):
		panic("The constraint attribute is required by this class.")
	case col.IsUndefined(additionalConstraints):
		panic("The additional constraints attribute is required by this class.")
	default:
		return &constraints_{
			// Initialize instance attributes.
			class_:                 c,
			constraint_:            constraint,
			additionalConstraints_: additionalConstraints,
		}
	}
}

// INSTANCE METHODS

// Target

type constraints_ struct {
	// Define instance attributes.
	class_                 ConstraintsClassLike
	constraint_            ConstraintLike
	additionalConstraints_ abs.Sequential[AdditionalConstraintLike]
}

// Public

func (v *constraints_) GetClass() ConstraintsClassLike {
	return v.class_
}

// Attribute

func (v *constraints_) GetConstraint() ConstraintLike {
	return v.constraint_
}

func (v *constraints_) GetAdditionalConstraints() abs.Sequential[AdditionalConstraintLike] {
	return v.additionalConstraints_
}

// Private
