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

var typeDefinitionsClass = &typeDefinitionsClass_{
	// Initialize class constants.
}

// Function

func TypeDefinitions() TypeDefinitionsClassLike {
	return typeDefinitionsClass
}

// CLASS METHODS

// Target

type typeDefinitionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *typeDefinitionsClass_) Make(types abs.Sequential[TypeLike]) TypeDefinitionsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(types):
		panic("The types attribute is required by this class.")
	default:
		return &typeDefinitions_{
			// Initialize instance attributes.
			class_: c,
			types_: types,
		}
	}
}

// INSTANCE METHODS

// Target

type typeDefinitions_ struct {
	// Define instance attributes.
	class_ TypeDefinitionsClassLike
	types_ abs.Sequential[TypeLike]
}

// Public

func (v *typeDefinitions_) GetClass() TypeDefinitionsClassLike {
	return v.class_
}

// Attribute

func (v *typeDefinitions_) GetTypes() abs.Sequential[TypeLike] {
	return v.types_
}

// Private
