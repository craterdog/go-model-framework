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

func AdditionalValue() AdditionalValueClassLike {
	return additionalValueClass
}

// Constructor Methods

func (c *additionalValueClass_) Make(
	name string,
) AdditionalValueLike {
	if uti.IsUndefined(name) {
		panic("The name attribute is required by this class.")
	}
	var instance = &additionalValue_{
		class_: c,
		name_:  name,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *additionalValue_) GetName() string {
	return v.name_
}

// Public Methods

func (v *additionalValue_) GetClass() AdditionalValueClassLike {
	return v.getClass()
}

// Private Methods

func (v *additionalValue_) getClass() *additionalValueClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type additionalValue_ struct {
	class_ *additionalValueClass_
	name_  string
}

// Class Structure

type additionalValueClass_ struct {
	// Define the class constants.
}

// Class Reference

var additionalValueClass = &additionalValueClass_{
	// Initialize the class constants.
}
