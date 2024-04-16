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

package cdsn

import ()

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// This class has no private constants to initialize.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{}
}

// Functions

// INSTANCE METHODS

// Target

type validator_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

func (v *validator_) ValidateGrammar(grammar GrammarLike) {
	// TBA - Implement the method.
}

// Private
