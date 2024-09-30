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

var publicMethodsClass = &publicMethodsClass_{
	// Initialize class constants.
}

// Function

func PublicMethods() PublicMethodsClassLike {
	return publicMethodsClass
}

// CLASS METHODS

// Target

type publicMethodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *publicMethodsClass_) Make(methods abs.Sequential[MethodLike]) PublicMethodsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(methods):
		panic("The methods attribute is required by this class.")
	default:
		return &publicMethods_{
			// Initialize instance attributes.
			class_:   c,
			methods_: methods,
		}
	}
}

// INSTANCE METHODS

// Target

type publicMethods_ struct {
	// Define instance attributes.
	class_   PublicMethodsClassLike
	methods_ abs.Sequential[MethodLike]
}

// Public

func (v *publicMethods_) GetClass() PublicMethodsClassLike {
	return v.class_
}

// Attribute

func (v *publicMethods_) GetMethods() abs.Sequential[MethodLike] {
	return v.methods_
}

// Private
