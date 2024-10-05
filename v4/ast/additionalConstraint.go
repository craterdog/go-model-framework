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

var additionalConstraintClass = &additionalConstraintClass_{
	// Initialize class constants.
}

// Function

func AdditionalConstraint() AdditionalConstraintClassLike {
	return additionalConstraintClass
}

// CLASS METHODS

// Target

type additionalConstraintClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalConstraintClass_) Make(constraint ConstraintLike) AdditionalConstraintLike {
	// Validate the constraints.
	switch {
	case col.IsUndefined(constraint):
		panic("The constraint attribute is required by this class.")
	default:
		return &additionalConstraint_{
			// Initialize instance attributes.
			class_:      c,
			constraint_: constraint,
		}
	}
}

// INSTANCE METHODS

// Target

type additionalConstraint_ struct {
	// Define instance attributes.
	class_      AdditionalConstraintClassLike
	constraint_ ConstraintLike
}

// Public

func (v *additionalConstraint_) GetClass() AdditionalConstraintClassLike {
	return v.class_
}

// Attribute

func (v *additionalConstraint_) GetConstraint() ConstraintLike {
	return v.constraint_
}

// Private
