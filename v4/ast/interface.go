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

var interfaceClass = &interfaceClass_{
	// Initialize class constants.
}

// Function

func Interface() InterfaceClassLike {
	return interfaceClass
}

// CLASS METHODS

// Target

type interfaceClass_ struct {
	// Define class constants.
}

// Constructors

func (c *interfaceClass_) Make(
	abstraction AbstractionLike,
) InterfaceLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(abstraction):
		panic("The abstraction attribute is required by this class.")
	default:
		return &interface_{
			// Initialize instance attributes.
			class_:       c,
			abstraction_: abstraction,
		}
	}
}

// INSTANCE METHODS

// Target

type interface_ struct {
	// Define instance attributes.
	class_       InterfaceClassLike
	abstraction_ AbstractionLike
}

// Public

func (v *interface_) GetClass() InterfaceClassLike {
	return v.class_
}

// Attribute

func (v *interface_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// Private
