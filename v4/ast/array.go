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

// CLASS ACCESS

// Reference

var arrayClass = &arrayClass_{
	// Initialize class constants.
}

// Function

func Array() ArrayClassLike {
	return arrayClass
}

// CLASS METHODS

// Target

type arrayClass_ struct {
	// Define class constants.
}

// Constructors

func (c *arrayClass_) Make() ArrayLike {
	// Validate the arguments.
	switch {
	default:
		return &array_{
			// Initialize instance attributes.
			class_: c,
		}
	}
}

// INSTANCE METHODS

// Target

type array_ struct {
	// Define instance attributes.
	class_ ArrayClassLike
}

// Public

func (v *array_) GetClass() ArrayClassLike {
	return v.class_
}

// Attribute

// Private
