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
	mod "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var enumerationClass = &enumerationClass_{
	// Initialize class constants.
}

// Function

func Enumeration() EnumerationClassLike {
	return enumerationClass
}

// CLASS METHODS

// Target

type enumerationClass_ struct {
	// Define class constants.
}

// Constructors

func (c *enumerationClass_) Make(values ValuesLike) EnumerationLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(values):
		panic("The values attribute is required for each Enumeration.")
	default:
		return &enumeration_{
			// Initialize instance attributes.
			class_:  c,
			values_: values,
		}
	}
}

// INSTANCE METHODS

// Target

type enumeration_ struct {
	// Define instance attributes.
	class_  EnumerationClassLike
	values_ ValuesLike
}

// Attributes

func (v *enumeration_) GetClass() EnumerationClassLike {
	return v.class_
}

func (v *enumeration_) GetValues() ValuesLike {
	return v.values_
}

// Private
