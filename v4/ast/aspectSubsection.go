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

func AspectSubsection() AspectSubsectionClassLike {
	return aspectSubsectionReference()
}

// Constructor Methods

func (c *aspectSubsectionClass_) Make(
	aspectInterfaces abs.Sequential[AspectInterfaceLike],
) AspectSubsectionLike {
	if uti.IsUndefined(aspectInterfaces) {
		panic("The \"aspectInterfaces\" attribute is required by this class.")
	}
	var instance = &aspectSubsection_{
		// Initialize the instance attributes.
		aspectInterfaces_: aspectInterfaces,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *aspectSubsection_) GetAspectInterfaces() abs.Sequential[AspectInterfaceLike] {
	return v.aspectInterfaces_
}

// Public Methods

func (v *aspectSubsection_) GetClass() AspectSubsectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *aspectSubsection_) getClass() *aspectSubsectionClass_ {
	return aspectSubsectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type aspectSubsection_ struct {
	// Declare the instance attributes.
	aspectInterfaces_ abs.Sequential[AspectInterfaceLike]
}

// Class Structure

type aspectSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectSubsectionReference() *aspectSubsectionClass_ {
	return aspectSubsectionReference_
}

var aspectSubsectionReference_ = &aspectSubsectionClass_{
	// Initialize the class constants.
}
