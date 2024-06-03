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
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var functionClass = &functionClass_{
	// Initialize class constants.
}

// Function

func Function() FunctionClassLike {
	return functionClass
}

// CLASS METHODS

// Target

type functionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *functionClass_) MakeWithAttributes(
	identifier string,
	parameters col.ListLike[ParameterLike],
	result ResultLike,
) FunctionLike {
	return &function_{
		// Initialize instance attributes.
		class_:      c,
		identifier_: identifier,
		parameters_: parameters,
		result_:     result,
	}
}

// INSTANCE METHODS

// Target

type function_ struct {
	// Define instance attributes.
	class_      FunctionClassLike
	identifier_ string
	parameters_ col.ListLike[ParameterLike]
	result_     ResultLike
}

// Attributes

func (v *function_) GetClass() FunctionClassLike {
	return v.class_
}

func (v *function_) GetIdentifier() string {
	return v.identifier_
}

func (v *function_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

func (v *function_) GetResult() ResultLike {
	return v.result_
}

// Private
