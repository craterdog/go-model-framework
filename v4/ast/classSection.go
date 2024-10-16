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

func ClassSection() ClassSectionClassLike {
	return classSectionReference()
}

// Constructor Methods

func (c *classSectionClass_) Make(
	classDefinitions abs.Sequential[ClassDefinitionLike],
) ClassSectionLike {
	if uti.IsUndefined(classDefinitions) {
		panic("The \"classDefinitions\" attribute is required by this class.")
	}
	var instance = &classSection_{
		// Initialize the instance attributes.
		classDefinitions_: classDefinitions,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *classSection_) GetClassDefinitions() abs.Sequential[ClassDefinitionLike] {
	return v.classDefinitions_
}

// Public Methods

func (v *classSection_) GetClass() ClassSectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *classSection_) getClass() *classSectionClass_ {
	return classSectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type classSection_ struct {
	// Declare the instance attributes.
	classDefinitions_ abs.Sequential[ClassDefinitionLike]
}

// Class Structure

type classSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classSectionReference() *classSectionClass_ {
	return classSectionReference_
}

var classSectionReference_ = &classSectionClass_{
	// Initialize the class constants.
}
