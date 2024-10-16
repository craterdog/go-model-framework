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

func Value() ValueClassLike {
	return valueReference()
}

// Constructor Methods

func (c *valueClass_) Make(
	name string,
	abstraction AbstractionLike,
) ValueLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &value_{
		// Initialize the instance attributes.
		name_:        name,
		abstraction_: abstraction,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *value_) GetName() string {
	return v.name_
}

func (v *value_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Public Methods

func (v *value_) GetClass() ValueClassLike {
	return v.getClass()
}

// Private Methods

func (v *value_) getClass() *valueClass_ {
	return valueReference()
}

// PRIVATE INTERFACE

// Instance Structure

type value_ struct {
	// Declare the instance attributes.
	name_        string
	abstraction_ AbstractionLike
}

// Class Structure

type valueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func valueReference() *valueClass_ {
	return valueReference_
}

var valueReference_ = &valueClass_{
	// Initialize the class constants.
}
