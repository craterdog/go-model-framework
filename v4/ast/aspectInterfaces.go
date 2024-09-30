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

var aspectInterfacesClass = &aspectInterfacesClass_{
	// Initialize class constants.
}

// Function

func AspectInterfaces() AspectInterfacesClassLike {
	return aspectInterfacesClass
}

// CLASS METHODS

// Target

type aspectInterfacesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *aspectInterfacesClass_) Make(interfaces abs.Sequential[InterfaceLike]) AspectInterfacesLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(interfaces):
		panic("The interfaces attribute is required by this class.")
	default:
		return &aspectInterfaces_{
			// Initialize instance attributes.
			class_:      c,
			interfaces_: interfaces,
		}
	}
}

// INSTANCE METHODS

// Target

type aspectInterfaces_ struct {
	// Define instance attributes.
	class_      AspectInterfacesClassLike
	interfaces_ abs.Sequential[InterfaceLike]
}

// Public

func (v *aspectInterfaces_) GetClass() AspectInterfacesClassLike {
	return v.class_
}

// Attribute

func (v *aspectInterfaces_) GetInterfaces() abs.Sequential[InterfaceLike] {
	return v.interfaces_
}

// Private
