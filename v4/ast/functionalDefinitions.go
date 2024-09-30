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

var functionalDefinitionsClass = &functionalDefinitionsClass_{
	// Initialize class constants.
}

// Function

func FunctionalDefinitions() FunctionalDefinitionsClassLike {
	return functionalDefinitionsClass
}

// CLASS METHODS

// Target

type functionalDefinitionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *functionalDefinitionsClass_) Make(functionals abs.Sequential[FunctionalLike]) FunctionalDefinitionsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(functionals):
		panic("The functionals attribute is required by this class.")
	default:
		return &functionalDefinitions_{
			// Initialize instance attributes.
			class_:       c,
			functionals_: functionals,
		}
	}
}

// INSTANCE METHODS

// Target

type functionalDefinitions_ struct {
	// Define instance attributes.
	class_       FunctionalDefinitionsClassLike
	functionals_ abs.Sequential[FunctionalLike]
}

// Public

func (v *functionalDefinitions_) GetClass() FunctionalDefinitionsClassLike {
	return v.class_
}

// Attribute

func (v *functionalDefinitions_) GetFunctionals() abs.Sequential[FunctionalLike] {
	return v.functionals_
}

// Private
