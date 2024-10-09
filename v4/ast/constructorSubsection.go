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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ConstructorSubsection() ConstructorSubsectionClassLike {
	return constructorSubsectionClass
}

// Constructor Methods

func (c *constructorSubsectionClass_) Make(
	constructorMethods abs.Sequential[ConstructorMethodLike],
) ConstructorSubsectionLike {
	if uti.IsUndefined(constructorMethods) {
		panic("The constructorMethods attribute is required by this class.")
	}
	var instance = &constructorSubsection_{
		class_:              c,
		constructorMethods_: constructorMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *constructorSubsection_) GetConstructorMethods() abs.Sequential[ConstructorMethodLike] {
	return v.constructorMethods_
}

// Public Methods

func (v *constructorSubsection_) GetClass() ConstructorSubsectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *constructorSubsection_) getClass() *constructorSubsectionClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type constructorSubsection_ struct {
	class_              *constructorSubsectionClass_
	constructorMethods_ abs.Sequential[ConstructorMethodLike]
}

// Class Structure

type constructorSubsectionClass_ struct {
	// Define the class constants.
}

// Class Reference

var constructorSubsectionClass = &constructorSubsectionClass_{
	// Initialize the class constants.
}
