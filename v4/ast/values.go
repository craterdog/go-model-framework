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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var valuesClass = &valuesClass_{
	// Initialize class constants.
}

// Function

func Values() ValuesClassLike {
	return valuesClass
}

// CLASS METHODS

// Target

type valuesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *valuesClass_) Make(
	value ValueLike,
	additionalValues abs.Sequential[AdditionalValueLike],
) ValuesLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(value):
		panic("The value attribute is required by this class.")
	case col.IsUndefined(additionalValues):
		panic("The additionalValues attribute is required by this class.")
	default:
		return &values_{
			// Initialize instance attributes.
			class_:            c,
			value_:            value,
			additionalValues_: additionalValues,
		}
	}
}

// INSTANCE METHODS

// Target

type values_ struct {
	// Define instance attributes.
	class_            ValuesClassLike
	value_            ValueLike
	additionalValues_ abs.Sequential[AdditionalValueLike]
}

// Attributes

func (v *values_) GetClass() ValuesClassLike {
	return v.class_
}

func (v *values_) GetValue() ValueLike {
	return v.value_
}

func (v *values_) GetAdditionalValues() abs.Sequential[AdditionalValueLike] {
	return v.additionalValues_
}

// Private
