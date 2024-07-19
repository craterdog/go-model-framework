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

var attributeClass = &attributeClass_{
	// Initialize class constants.
}

// Function

func Attribute() AttributeClassLike {
	return attributeClass
}

// CLASS METHODS

// Target

type attributeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *attributeClass_) Make(
	name string,
	optionalParameter ParameterLike,
	optionalAbstraction AbstractionLike,
) AttributeLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	default:
		return &attribute_{
			// Initialize instance attributes.
			class_:               c,
			name_:                name,
			optionalParameter_:   optionalParameter,
			optionalAbstraction_: optionalAbstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type attribute_ struct {
	// Define instance attributes.
	class_               AttributeClassLike
	name_                string
	optionalParameter_   ParameterLike
	optionalAbstraction_ AbstractionLike
}

// Attributes

func (v *attribute_) GetClass() AttributeClassLike {
	return v.class_
}

func (v *attribute_) GetName() string {
	return v.name_
}

func (v *attribute_) GetOptionalParameter() ParameterLike {
	return v.optionalParameter_
}

func (v *attribute_) GetOptionalAbstraction() AbstractionLike {
	return v.optionalAbstraction_
}

// Private
