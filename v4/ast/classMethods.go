// Public

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
)

// CLASS ACCESS

// Reference

var classMethodsClass = &classMethodsClass_{
	// Initialize class constants.
}

// Function

func ClassMethods() ClassMethodsClassLike {
	return classMethodsClass
}

// CLASS METHODS

// Target

type classMethodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *classMethodsClass_) Make(
	constructorMethods ConstructorMethodsLike,
	optionalConstantMethods ConstantMethodsLike,
	optionalFunctionMethods FunctionMethodsLike,
) ClassMethodsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(constructorMethods):
		panic("The constructorMethods attribute is required by this class.")
	default:
		return &classMethods_{
			// Initialize instance attributes.
			class_:                   c,
			constructorMethods_:      constructorMethods,
			optionalConstantMethods_: optionalConstantMethods,
			optionalFunctionMethods_: optionalFunctionMethods,
		}
	}
}

// INSTANCE METHODS

// Target

type classMethods_ struct {
	// Define instance attributes.
	class_                   ClassMethodsClassLike
	constructorMethods_      ConstructorMethodsLike
	optionalConstantMethods_ ConstantMethodsLike
	optionalFunctionMethods_ FunctionMethodsLike
}

func (v *classMethods_) GetClass() ClassMethodsClassLike {
	return v.class_
}

// Attribute

func (v *classMethods_) GetConstructorMethods() ConstructorMethodsLike {
	return v.constructorMethods_
}

func (v *classMethods_) GetOptionalConstantMethods() ConstantMethodsLike {
	return v.optionalConstantMethods_
}

func (v *classMethods_) GetOptionalFunctionMethods() FunctionMethodsLike {
	return v.optionalFunctionMethods_
}

// Private
