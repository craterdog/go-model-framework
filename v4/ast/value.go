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

import ()

// CLASS ACCESS

// Reference

var valueClass = &valueClass_{
	// Initialize class constants.
}

// Function

func Value() ValueClassLike {
	return valueClass
}

// CLASS METHODS

// Target

type valueClass_ struct {
	// Define class constants.
}

// Constructors

func (c *valueClass_) Make(
	name string,
	abstraction AbstractionLike,
) ValueLike {
	return &value_{
		// Initialize instance attributes.
		class_:       c,
		name_:        name,
		abstraction_: abstraction,
	}
}

// INSTANCE METHODS

// Target

type value_ struct {
	// Define instance attributes.
	class_       ValueClassLike
	name_        string
	abstraction_ AbstractionLike
}

// Attributes

func (v *value_) GetClass() ValueClassLike {
	return v.class_
}

func (v *value_) GetName() string {
	return v.name_
}

func (v *value_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
