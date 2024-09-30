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

var aspectDefinitionsClass = &aspectDefinitionsClass_{
	// Initialize class constants.
}

// Function

func AspectDefinitions() AspectDefinitionsClassLike {
	return aspectDefinitionsClass
}

// CLASS METHODS

// Target

type aspectDefinitionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *aspectDefinitionsClass_) Make(aspects abs.Sequential[AspectLike]) AspectDefinitionsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(aspects):
		panic("The aspects attribute is required by this class.")
	default:
		return &aspectDefinitions_{
			// Initialize instance attributes.
			class_:   c,
			aspects_: aspects,
		}
	}
}

// INSTANCE METHODS

// Target

type aspectDefinitions_ struct {
	// Define instance attributes.
	class_   AspectDefinitionsClassLike
	aspects_ abs.Sequential[AspectLike]
}

// Public

func (v *aspectDefinitions_) GetClass() AspectDefinitionsClassLike {
	return v.class_
}

// Attribute

func (v *aspectDefinitions_) GetAspects() abs.Sequential[AspectLike] {
	return v.aspects_
}

// Private
