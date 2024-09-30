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

var functionMethodsClass = &functionMethodsClass_{
	// Initialize class constants.
}

// Function

func FunctionMethods() FunctionMethodsClassLike {
	return functionMethodsClass
}

// CLASS METHODS

// Target

type functionMethodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *functionMethodsClass_) Make(functions abs.Sequential[FunctionLike]) FunctionMethodsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(functions):
		panic("The functions attribute is required by this class.")
	default:
		return &functionMethods_{
			// Initialize instance attributes.
			class_:     c,
			functions_: functions,
		}
	}
}

// INSTANCE METHODS

// Target

type functionMethods_ struct {
	// Define instance attributes.
	class_     FunctionMethodsClassLike
	functions_ abs.Sequential[FunctionLike]
}

// Public

func (v *functionMethods_) GetClass() FunctionMethodsClassLike {
	return v.class_
}

// Attribute

func (v *functionMethods_) GetFunctions() abs.Sequential[FunctionLike] {
	return v.functions_
}

// Private
