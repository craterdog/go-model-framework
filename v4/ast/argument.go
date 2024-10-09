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

func Argument() ArgumentClassLike {
	return argumentClass
}

// Constructor Methods

func (c *argumentClass_) Make(
	abstraction AbstractionLike,
) ArgumentLike {
	if uti.IsUndefined(abstraction) {
		panic("The abstraction attribute is required by this class.")
	}
	var instance = &argument_{
		class_:       c,
		abstraction_: abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *argument_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Public Methods

func (v *argument_) GetClass() ArgumentClassLike {
	return v.getClass()
}

// Private Methods

func (v *argument_) getClass() *argumentClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type argument_ struct {
	class_       *argumentClass_
	abstraction_ AbstractionLike
}

// Class Structure

type argumentClass_ struct {
	// Define the class constants.
}

// Class Reference

var argumentClass = &argumentClass_{
	// Initialize the class constants.
}
