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

func TypeSection() TypeSectionClassLike {
	return typeSectionReference()
}

// Constructor Methods

func (c *typeSectionClass_) Make(
	typeDefinitions abs.Sequential[TypeDefinitionLike],
) TypeSectionLike {
	if uti.IsUndefined(typeDefinitions) {
		panic("The \"typeDefinitions\" attribute is required by this class.")
	}
	var instance = &typeSection_{
		// Initialize the instance attributes.
		typeDefinitions_: typeDefinitions,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *typeSection_) GetTypeDefinitions() abs.Sequential[TypeDefinitionLike] {
	return v.typeDefinitions_
}

// Public Methods

func (v *typeSection_) GetClass() TypeSectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *typeSection_) getClass() *typeSectionClass_ {
	return typeSectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type typeSection_ struct {
	// Declare the instance attributes.
	typeDefinitions_ abs.Sequential[TypeDefinitionLike]
}

// Class Structure

type typeSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func typeSectionReference() *typeSectionClass_ {
	return typeSectionReference_
}

var typeSectionReference_ = &typeSectionClass_{
	// Initialize the class constants.
}
