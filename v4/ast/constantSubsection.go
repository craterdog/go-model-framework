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

func ConstantSubsection() ConstantSubsectionClassLike {
	return constantSubsectionReference()
}

// Constructor Methods

func (c *constantSubsectionClass_) Make(
	constantMethods abs.Sequential[ConstantMethodLike],
) ConstantSubsectionLike {
	if uti.IsUndefined(constantMethods) {
		panic("The \"constantMethods\" attribute is required by this class.")
	}
	var instance = &constantSubsection_{
		// Initialize the instance attributes.
		constantMethods_: constantMethods,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *constantSubsection_) GetConstantMethods() abs.Sequential[ConstantMethodLike] {
	return v.constantMethods_
}

// Public Methods

func (v *constantSubsection_) GetClass() ConstantSubsectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *constantSubsection_) getClass() *constantSubsectionClass_ {
	return constantSubsectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type constantSubsection_ struct {
	// Declare the instance attributes.
	constantMethods_ abs.Sequential[ConstantMethodLike]
}

// Class Structure

type constantSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constantSubsectionReference() *constantSubsectionClass_ {
	return constantSubsectionReference_
}

var constantSubsectionReference_ = &constantSubsectionClass_{
	// Initialize the class constants.
}
