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

var constructorMethodsClass = &constructorMethodsClass_{
	// Initialize class constants.
}

// Function

func ConstructorMethods() ConstructorMethodsClassLike {
	return constructorMethodsClass
}

// CLASS METHODS

// Target

type constructorMethodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constructorMethodsClass_) Make(constructors abs.Sequential[ConstructorLike]) ConstructorMethodsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(constructors):
		panic("The constructors attribute is required by this class.")
	default:
		return &constructorMethods_{
			// Initialize instance attributes.
			class_:        c,
			constructors_: constructors,
		}
	}
}

// INSTANCE METHODS

// Target

type constructorMethods_ struct {
	// Define instance attributes.
	class_        ConstructorMethodsClassLike
	constructors_ abs.Sequential[ConstructorLike]
}

// Public

func (v *constructorMethods_) GetClass() ConstructorMethodsClassLike {
	return v.class_
}

// Attribute

func (v *constructorMethods_) GetConstructors() abs.Sequential[ConstructorLike] {
	return v.constructors_
}

// Private
