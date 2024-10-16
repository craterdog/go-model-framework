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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Constraint() ConstraintClassLike {
	return constraintReference()
}

// Constructor Methods

func (c *constraintClass_) Make(
	name string,
	abstraction AbstractionLike,
) ConstraintLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &constraint_{
		// Initialize the instance attributes.
		name_:        name,
		abstraction_: abstraction,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *constraint_) GetName() string {
	return v.name_
}

func (v *constraint_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Public Methods

func (v *constraint_) GetClass() ConstraintClassLike {
	return v.getClass()
}

// Private Methods

func (v *constraint_) getClass() *constraintClass_ {
	return constraintReference()
}

// PRIVATE INTERFACE

// Instance Structure

type constraint_ struct {
	// Declare the instance attributes.
	name_        string
	abstraction_ AbstractionLike
}

// Class Structure

type constraintClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constraintReference() *constraintClass_ {
	return constraintReference_
}

var constraintReference_ = &constraintClass_{
	// Initialize the class constants.
}
