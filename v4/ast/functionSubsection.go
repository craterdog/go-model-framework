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

func FunctionSubsection() FunctionSubsectionClassLike {
	return functionSubsectionReference()
}

// Constructor Methods

func (c *functionSubsectionClass_) Make(
	functionMethods abs.Sequential[FunctionMethodLike],
) FunctionSubsectionLike {
	if uti.IsUndefined(functionMethods) {
		panic("The \"functionMethods\" attribute is required by this class.")
	}
	var instance = &functionSubsection_{
		// Initialize the instance attributes.
		functionMethods_: functionMethods,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *functionSubsection_) GetFunctionMethods() abs.Sequential[FunctionMethodLike] {
	return v.functionMethods_
}

// Public Methods

func (v *functionSubsection_) GetClass() FunctionSubsectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *functionSubsection_) getClass() *functionSubsectionClass_ {
	return functionSubsectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type functionSubsection_ struct {
	// Declare the instance attributes.
	functionMethods_ abs.Sequential[FunctionMethodLike]
}

// Class Structure

type functionSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionSubsectionReference() *functionSubsectionClass_ {
	return functionSubsectionReference_
}

var functionSubsectionReference_ = &functionSubsectionClass_{
	// Initialize the class constants.
}
