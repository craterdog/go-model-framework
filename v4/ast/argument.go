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

var argumentClass = &argumentClass_{
	// Initialize class constants.
}

// Function

func Argument() ArgumentClassLike {
	return argumentClass
}

// CLASS METHODS

// Target

type argumentClass_ struct {
	// Define class constants.
}

// Constructors

func (c *argumentClass_) Make(abstraction AbstractionLike) ArgumentLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(abstraction):
		panic("The abstraction attribute is required by this class.")
	default:
		return &argument_{
			// Initialize instance attributes.
			class_:       c,
			abstraction_: abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type argument_ struct {
	// Define instance attributes.
	class_       ArgumentClassLike
	abstraction_ AbstractionLike
}

// Attributes

func (v *argument_) GetClass() ArgumentClassLike {
	return v.class_
}

func (v *argument_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
