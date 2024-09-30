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
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
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

// Public

func (v *additionalValue_) GetClass() AdditionalValueClassLike {
	return v.class_
}

// Attribute

func (v *additionalValue_) GetName() string {
	return v.name_
}

// Private
