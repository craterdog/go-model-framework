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

func Parameterized() ParameterizedClassLike {
	return parameterizedClass
}

// Constructor Methods

func (c *parameterizedClass_) Make(
	parameters abs.Sequential[ParameterLike],
) ParameterizedLike {
	if uti.IsUndefined(parameters) {
		panic("The parameters attribute is required by this class.")
	}
	var instance = &parameterized_{
		class_:      c,
		parameters_: parameters,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *parameterized_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

// Public Methods

func (v *parameterized_) GetClass() ParameterizedClassLike {
	return v.getClass()
}

// Private Methods

func (v *parameterized_) getClass() *parameterizedClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type parameterized_ struct {
	class_      *parameterizedClass_
	parameters_ abs.Sequential[ParameterLike]
}

// Class Structure

type parameterizedClass_ struct {
	// Define the class constants.
}

// Class Reference

var parameterizedClass = &parameterizedClass_{
	// Initialize the class constants.
}
