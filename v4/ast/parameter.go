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

import ()

// CLASS ACCESS

// Reference

var parameterClass = &parameterClass_{
	// Initialize class constants.
}

// Function

func Parameter() ParameterClassLike {
	return parameterClass
}

// CLASS METHODS

// Target

type parameterClass_ struct {
	// Define class constants.
}

// Constructors

func (c *parameterClass_) MakeWithAttributes(
	identifier string,
	abstraction AbstractionLike,
) ParameterLike {
	return &parameter_{
		// Initialize instance attributes.
		class_:       c,
		identifier_:  identifier,
		abstraction_: abstraction,
	}
}

// INSTANCE METHODS

// Target

type parameter_ struct {
	// Define instance attributes.
	class_       ParameterClassLike
	identifier_  string
	abstraction_ AbstractionLike
}

// Attributes

func (v *parameter_) GetClass() ParameterClassLike {
	return v.class_
}

func (v *parameter_) GetIdentifier() string {
	return v.identifier_
}

func (v *parameter_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
