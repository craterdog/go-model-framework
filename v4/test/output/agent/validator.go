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

package agent

import (
	ast "github.com/craterdog/go-model-framework/v4/gcmn/ast"
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Any private class constants should be initialized here.
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
	class_ ValidatorClassLike
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateModel(model ast.ModelLike) {
	// TBA - Implement the method.
}

// Private
