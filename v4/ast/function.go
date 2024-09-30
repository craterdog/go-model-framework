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
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
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

func (c *functionClass_) Make(
	name string,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	case col.IsUndefined(parameters):
		panic("The parameters attribute is required by this class.")
	case col.IsUndefined(result):
		panic("The result attribute is required by this class.")
	default:
		return &function_{
			// Initialize instance attributes.
			class_:      c,
			name_:       name,
			parameters_: parameters,
			result_:     result,
		}
	}
}

// INSTANCE METHODS

// Target

type function_ struct {
	// Define instance attributes.
	class_      FunctionClassLike
	name_       string
	parameters_ abs.Sequential[ParameterLike]
	result_     ResultLike
}

// Public

func (v *function_) GetClass() FunctionClassLike {
	return v.class_
}

// Attribute

func (v *function_) GetName() string {
	return v.name_
}

func (v *function_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *function_) GetResult() ResultLike {
	return v.result_
}

// Private
