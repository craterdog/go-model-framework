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
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var constructorClass = &constructorClass_{
	// Initialize class constants.
}

// Function

func Constructor() ConstructorClassLike {
	return constructorClass
}

// CLASS METHODS

// Target

type constructorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constructorClass_) MakeWithAttributes(
	identifier string,
	parameters col.ListLike[ParameterLike],
	abstraction AbstractionLike,
) ConstructorLike {
	return &constructor_{
		// Initialize instance attributes.
		class_:       c,
		identifier_:  identifier,
		parameters_:  parameters,
		abstraction_: abstraction,
	}
}

// INSTANCE METHODS

// Target

type constructor_ struct {
	// Define instance attributes.
	class_       ConstructorClassLike
	identifier_  string
	parameters_  col.ListLike[ParameterLike]
	abstraction_ AbstractionLike
}

// Attributes

func (v *constructor_) GetClass() ConstructorClassLike {
	return v.class_
}

func (v *constructor_) GetIdentifier() string {
	return v.identifier_
}

func (v *constructor_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

func (v *constructor_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
