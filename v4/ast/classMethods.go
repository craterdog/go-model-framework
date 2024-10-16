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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ClassMethods() ClassMethodsClassLike {
	return classMethodsReference()
}

// Constructor Methods

func (c *classMethodsClass_) Make(
	constructorSubsection ConstructorSubsectionLike,
	optionalConstantSubsection ConstantSubsectionLike,
	optionalFunctionSubsection FunctionSubsectionLike,
) ClassMethodsLike {
	if uti.IsUndefined(constructorSubsection) {
		panic("The \"constructorSubsection\" attribute is required by this class.")
	}
	var instance = &classMethods_{
		// Initialize the instance attributes.
		constructorSubsection_:      constructorSubsection,
		optionalConstantSubsection_: optionalConstantSubsection,
		optionalFunctionSubsection_: optionalFunctionSubsection,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *classMethods_) GetConstructorSubsection() ConstructorSubsectionLike {
	return v.constructorSubsection_
}

func (v *classMethods_) GetOptionalConstantSubsection() ConstantSubsectionLike {
	return v.optionalConstantSubsection_
}

func (v *classMethods_) GetOptionalFunctionSubsection() FunctionSubsectionLike {
	return v.optionalFunctionSubsection_
}

// Public Methods

func (v *classMethods_) GetClass() ClassMethodsClassLike {
	return v.getClass()
}

// Private Methods

func (v *classMethods_) getClass() *classMethodsClass_ {
	return classMethodsReference()
}

// PRIVATE INTERFACE

// Instance Structure

type classMethods_ struct {
	// Declare the instance attributes.
	constructorSubsection_      ConstructorSubsectionLike
	optionalConstantSubsection_ ConstantSubsectionLike
	optionalFunctionSubsection_ FunctionSubsectionLike
}

// Class Structure

type classMethodsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classMethodsReference() *classMethodsClass_ {
	return classMethodsReference_
}

var classMethodsReference_ = &classMethodsClass_{
	// Initialize the class constants.
}
