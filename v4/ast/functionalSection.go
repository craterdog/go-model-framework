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

func FunctionalSection() FunctionalSectionClassLike {
	return functionalSectionClass
}

// Constructor Methods

func (c *functionalSectionClass_) Make(
	functionalDefinitions abs.Sequential[FunctionalDefinitionLike],
) FunctionalSectionLike {
	if uti.IsUndefined(functionalDefinitions) {
		panic("The functionalDefinitions attribute is required by this class.")
	}
	var instance = &functionalSection_{
		class_:                 c,
		functionalDefinitions_: functionalDefinitions,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *functionalSection_) GetFunctionalDefinitions() abs.Sequential[FunctionalDefinitionLike] {
	return v.functionalDefinitions_
}

// Public Methods

func (v *functionalSection_) GetClass() FunctionalSectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *functionalSection_) getClass() *functionalSectionClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type functionalSection_ struct {
	class_                 *functionalSectionClass_
	functionalDefinitions_ abs.Sequential[FunctionalDefinitionLike]
}

// Class Structure

type functionalSectionClass_ struct {
	// Define the class constants.
}

// Class Reference

var functionalSectionClass = &functionalSectionClass_{
	// Initialize the class constants.
}
