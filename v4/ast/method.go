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

func Method() MethodClassLike {
	return methodReference()
}

// Constructor Methods

func (c *methodClass_) Make(
	name string,
	parameters abs.Sequential[ParameterLike],
	optionalResult ResultLike,
) MethodLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	var instance = &method_{
		// Initialize the instance attributes.
		name_:           name,
		parameters_:     parameters,
		optionalResult_: optionalResult,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *method_) GetName() string {
	return v.name_
}

func (v *method_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *method_) GetOptionalResult() ResultLike {
	return v.optionalResult_
}

// Public Methods

func (v *method_) GetClass() MethodClassLike {
	return v.getClass()
}

// Private Methods

func (v *method_) getClass() *methodClass_ {
	return methodReference()
}

// PRIVATE INTERFACE

// Instance Structure

type method_ struct {
	// Declare the instance attributes.
	name_           string
	parameters_     abs.Sequential[ParameterLike]
	optionalResult_ ResultLike
}

// Class Structure

type methodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func methodReference() *methodClass_ {
	return methodReference_
}

var methodReference_ = &methodClass_{
	// Initialize the class constants.
}
