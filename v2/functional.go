/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package models

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// CLASS ACCESS

// Reference

var functionalClass = &functionalClass_{
	// TBA - Assign class constant values.
}

// Function

func Functional() FunctionalClassLike {
	return functionalClass
}

// CLASS METHODS

// Target

type functionalClass_ struct {
	// TBA - Add private class constants.
}

// Constants

// Constructors

func (c *functionalClass_) MakeWithAttributes(
	declaration DeclarationLike,
	parameters col.ListLike[ParameterLike],
	result ResultLike,
) FunctionalLike {
	return &functional_{
		declaration_: declaration,
		parameters_:  parameters,
		result_:      result,
	}
}

// Functions

// INSTANCE METHODS

// Target

type functional_ struct {
	declaration_ DeclarationLike
	parameters_  col.ListLike[ParameterLike]
	result_      ResultLike
}

// Attributes

func (v *functional_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *functional_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

func (v *functional_) GetResult() ResultLike {
	return v.result_
}

// Public

// Private
