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
	identifier string,
	abstraction AbstractionLike,
) ConstantLike {
	return &constant_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type constant_ struct {
	// Define instance attributes.
	class_       ConstantClassLike
	identifier_  string
	abstraction_ AbstractionLike
}

// Attributes

func (v *constant_) GetClass() ConstantClassLike {
	return v.class_
}

func (v *constant_) GetIdentifier() string {
	return v.identifier_
}

func (v *constant_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
