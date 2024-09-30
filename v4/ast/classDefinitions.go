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

var classDefinitionsClass = &classDefinitionsClass_{
	// Initialize class constants.
}

// Function

func ClassDefinitions() ClassDefinitionsClassLike {
	return classDefinitionsClass
}

// CLASS METHODS

// Target

type classDefinitionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *classDefinitionsClass_) Make(classes abs.Sequential[ClassLike]) ClassDefinitionsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(classes):
		panic("The classes attribute is required by this class.")
	default:
		return &classDefinitions_{
			// Initialize instance attributes.
			class_:   c,
			classes_: classes,
		}
	}
}

// INSTANCE METHODS

// Target

type classDefinitions_ struct {
	// Define instance attributes.
	class_   ClassDefinitionsClassLike
	classes_ abs.Sequential[ClassLike]
}

// Public

func (v *classDefinitions_) GetClass() ClassDefinitionsClassLike {
	return v.class_
}

// Attribute

func (v *classDefinitions_) GetClasses() abs.Sequential[ClassLike] {
	return v.classes_
}

// Private
