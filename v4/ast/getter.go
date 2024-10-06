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

var getterClass = &getterClass_{
	// Initialize class constants.
}

// Function

func Getter() GetterClassLike {
	return getterClass
}

// CLASS METHODS

// Target

type getterClass_ struct {
	// Define class constants.
}

// Constructors

func (c *getterClass_) Make(
	name string,
	abstraction AbstractionLike,
) GetterLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	default:
		return &getter_{
			// Initialize instance attributes.
			class_:       c,
			name_:        name,
			abstraction_: abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type getter_ struct {
	// Define instance attributes.
	class_       GetterClassLike
	name_        string
	abstraction_ AbstractionLike
}

// Public

func (v *getter_) GetClass() GetterClassLike {
	return v.class_
}

// Attribute

func (v *getter_) GetName() string {
	return v.name_
}

func (v *getter_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
