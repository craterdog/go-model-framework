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

func FunctionMethod() FunctionMethodClassLike {
	return functionMethodReference()
}

// Constructor Methods

func (c *functionMethodClass_) Make(
	name string,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionMethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	if uti.IsUndefined(result) {
		panic("The \"result\" attribute is required by this class.")
	}
	var instance = &functionMethod_{
		// Initialize the instance attributes.
		name_:       name,
		parameters_: parameters,
		result_:     result,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *functionMethod_) GetName() string {
	return v.name_
}

func (v *functionMethod_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *functionMethod_) GetResult() ResultLike {
	return v.result_
}

// Public Methods

func (v *functionMethod_) GetClass() FunctionMethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *functionMethod_) getClass() *functionMethodClass_ {
	return functionMethodReference()
}

// PRIVATE INTERFACE

// Instance Structure

type functionMethod_ struct {
	// Declare the instance attributes.
	name_       string
	parameters_ abs.Sequential[ParameterLike]
	result_     ResultLike
}

// Class Structure

type functionMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionMethodReference() *functionMethodClass_ {
	return functionMethodReference_
}

var functionMethodReference_ = &functionMethodClass_{
	// Initialize the class constants.
}
