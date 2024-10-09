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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func AspectInterface() AspectInterfaceClassLike {
	return aspectInterfaceClass
}

// Constructor Methods

func (c *aspectInterfaceClass_) Make(
	abstraction AbstractionLike,
) AspectInterfaceLike {
	if uti.IsUndefined(abstraction) {
		panic("The abstraction attribute is required by this class.")
	}
	var instance = &aspectInterface_{
		class_:       c,
		abstraction_: abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *aspectInterface_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Public Methods

func (v *aspectInterface_) GetClass() AspectInterfaceClassLike {
	return v.getClass()
}

// Private Methods

func (v *aspectInterface_) getClass() *aspectInterfaceClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type aspectInterface_ struct {
	class_       *aspectInterfaceClass_
	abstraction_ AbstractionLike
}

// Class Structure

type aspectInterfaceClass_ struct {
	// Define the class constants.
}

// Class Reference

var aspectInterfaceClass = &aspectInterfaceClass_{
	// Initialize the class constants.
}
