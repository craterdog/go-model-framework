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

func AspectSection() AspectSectionClassLike {
	return aspectSectionReference()
}

// Constructor Methods

func (c *aspectSectionClass_) Make(
	aspectDefinitions abs.Sequential[AspectDefinitionLike],
) AspectSectionLike {
	if uti.IsUndefined(aspectDefinitions) {
		panic("The \"aspectDefinitions\" attribute is required by this class.")
	}
	var instance = &aspectSection_{
		// Initialize the instance attributes.
		aspectDefinitions_: aspectDefinitions,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *aspectSection_) GetAspectDefinitions() abs.Sequential[AspectDefinitionLike] {
	return v.aspectDefinitions_
}

// Public Methods

func (v *aspectSection_) GetClass() AspectSectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *aspectSection_) getClass() *aspectSectionClass_ {
	return aspectSectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type aspectSection_ struct {
	// Declare the instance attributes.
	aspectDefinitions_ abs.Sequential[AspectDefinitionLike]
}

// Class Structure

type aspectSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectSectionReference() *aspectSectionClass_ {
	return aspectSectionReference_
}

var aspectSectionReference_ = &aspectSectionClass_{
	// Initialize the class constants.
}
