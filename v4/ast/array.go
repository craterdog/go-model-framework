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

// CLASS INTERFACE

// Access Function

func Array() ArrayClassLike {
	return arrayClass
}

// Constructor Methods

func (c *arrayClass_) Make() ArrayLike {
	var instance = &array_{
		// Initialize the instance attributes.
		class_: c,
	}
	return instance
}

// INSTANCE INTERFACE

// Public Methods

func (v *array_) GetClass() ArrayClassLike {
	return v.getClass()
}

// Private Methods

func (v *array_) getClass() *arrayClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type array_ struct {
	// Declare the instance attributes.
	class_ *arrayClass_
}

// Class Structure

type arrayClass_ struct {
	// Declare the class constants.
}

// Class Reference

var arrayClass = &arrayClass_{
	// Initialize the class constants.
}
