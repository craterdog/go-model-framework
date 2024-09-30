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

var constantMethodsClass = &constantMethodsClass_{
	// Initialize class constants.
}

// Function

func ConstantMethods() ConstantMethodsClassLike {
	return constantMethodsClass
}

// CLASS METHODS

// Target

type constantMethodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constantMethodsClass_) Make(constants abs.Sequential[ConstantLike]) ConstantMethodsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(constants):
		panic("The constants attribute is required by this class.")
	default:
		return &constantMethods_{
			// Initialize instance attributes.
			class_:     c,
			constants_: constants,
		}
	}
}

// INSTANCE METHODS

// Target

type constantMethods_ struct {
	// Define instance attributes.
	class_     ConstantMethodsClassLike
	constants_ abs.Sequential[ConstantLike]
}

// Public

func (v *constantMethods_) GetClass() ConstantMethodsClassLike {
	return v.class_
}

// Attribute

func (v *constantMethods_) GetConstants() abs.Sequential[ConstantLike] {
	return v.constants_
}

// Private
