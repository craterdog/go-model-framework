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

var additionalValueClass = &additionalValueClass_{
	// Initialize class constants.
}

// Function

func AdditionalValue() AdditionalValueClassLike {
	return additionalValueClass
}

// CLASS METHODS

// Target

type additionalValueClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalValueClass_) Make(name string) AdditionalValueLike {
	// Validate the arguments.
	switch {
	case isUndefined(name):
		panic("The name attribute is required for each AdditionalValue.")
	default:
		return &additionalValue_{
			// Initialize instance attributes.
			class_: c,
			name_:  name,
		}
	}
}

// INSTANCE METHODS

// Target

type additionalValue_ struct {
	// Define instance attributes.
	class_ AdditionalValueClassLike
	name_  string
}

// Attributes

func (v *additionalValue_) GetClass() AdditionalValueClassLike {
	return v.class_
}

func (v *additionalValue_) GetName() string {
	return v.name_
}

// Private
