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

func Arguments() ArgumentsClassLike {
	return argumentsReference()
}

// Constructor Methods

func (c *argumentsClass_) Make(
	argument ArgumentLike,
	additionalArguments abs.Sequential[AdditionalArgumentLike],
) ArgumentsLike {
	if uti.IsUndefined(argument) {
		panic("The \"argument\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalArguments) {
		panic("The \"additionalArguments\" attribute is required by this class.")
	}
	var instance = &arguments_{
		// Initialize the instance attributes.
		argument_:            argument,
		additionalArguments_: additionalArguments,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *arguments_) GetArgument() ArgumentLike {
	return v.argument_
}

func (v *arguments_) GetAdditionalArguments() abs.Sequential[AdditionalArgumentLike] {
	return v.additionalArguments_
}

// Public Methods

func (v *arguments_) GetClass() ArgumentsClassLike {
	return v.getClass()
}

// Private Methods

func (v *arguments_) getClass() *argumentsClass_ {
	return argumentsReference()
}

// PRIVATE INTERFACE

// Instance Structure

type arguments_ struct {
	// Declare the instance attributes.
	argument_            ArgumentLike
	additionalArguments_ abs.Sequential[AdditionalArgumentLike]
}

// Class Structure

type argumentsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func argumentsReference() *argumentsClass_ {
	return argumentsReference_
}

var argumentsReference_ = &argumentsClass_{
	// Initialize the class constants.
}
