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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Constraints() ConstraintsClassLike {
	return constraintsReference()
}

// Constructor Methods

func (c *constraintsClass_) Make(
	constraint ConstraintLike,
	additionalConstraints abs.Sequential[AdditionalConstraintLike],
) ConstraintsLike {
	if uti.IsUndefined(constraint) {
		panic("The \"constraint\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalConstraints) {
		panic("The \"additionalConstraints\" attribute is required by this class.")
	}
	var instance = &constraints_{
		// Initialize the instance attributes.
		constraint_:            constraint,
		additionalConstraints_: additionalConstraints,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *constraints_) GetConstraint() ConstraintLike {
	return v.constraint_
}

func (v *constraints_) GetAdditionalConstraints() abs.Sequential[AdditionalConstraintLike] {
	return v.additionalConstraints_
}

// Public Methods

func (v *constraints_) GetClass() ConstraintsClassLike {
	return v.getClass()
}

// Private Methods

func (v *constraints_) getClass() *constraintsClass_ {
	return constraintsReference()
}

// PRIVATE INTERFACE

// Instance Structure

type constraints_ struct {
	// Declare the instance attributes.
	constraint_            ConstraintLike
	additionalConstraints_ abs.Sequential[AdditionalConstraintLike]
}

// Class Structure

type constraintsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constraintsReference() *constraintsClass_ {
	return constraintsReference_
}

var constraintsReference_ = &constraintsClass_{
	// Initialize the class constants.
}
