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

func GetterMethod() GetterMethodClassLike {
	return getterMethodClass
}

// Constructor Methods

func (c *getterMethodClass_) Make(
	name string,
	abstraction AbstractionLike,
) GetterMethodLike {
	if uti.IsUndefined(name) {
		panic("The name attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The abstraction attribute is required by this class.")
	}
	var instance = &getterMethod_{
		class_:       c,
		name_:        name,
		abstraction_: abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *getterMethod_) GetName() string {
	return v.name_
}

func (v *getterMethod_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Public Methods

func (v *getterMethod_) GetClass() GetterMethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *getterMethod_) getClass() *getterMethodClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type getterMethod_ struct {
	class_       *getterMethodClass_
	name_        string
	abstraction_ AbstractionLike
}

// Class Structure

type getterMethodClass_ struct {
	// Define the class constants.
}

// Class Reference

var getterMethodClass = &getterMethodClass_{
	// Initialize the class constants.
}