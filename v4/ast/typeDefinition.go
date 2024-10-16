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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func TypeDefinition() TypeDefinitionClassLike {
	return typeDefinitionReference()
}

// Constructor Methods

func (c *typeDefinitionClass_) Make(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	optionalEnumeration EnumerationLike,
) TypeDefinitionLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &typeDefinition_{
		// Initialize the instance attributes.
		declaration_:         declaration,
		abstraction_:         abstraction,
		optionalEnumeration_: optionalEnumeration,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *typeDefinition_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *typeDefinition_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *typeDefinition_) GetOptionalEnumeration() EnumerationLike {
	return v.optionalEnumeration_
}

// Public Methods

func (v *typeDefinition_) GetClass() TypeDefinitionClassLike {
	return v.getClass()
}

// Private Methods

func (v *typeDefinition_) getClass() *typeDefinitionClass_ {
	return typeDefinitionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type typeDefinition_ struct {
	// Declare the instance attributes.
	declaration_         DeclarationLike
	abstraction_         AbstractionLike
	optionalEnumeration_ EnumerationLike
}

// Class Structure

type typeDefinitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func typeDefinitionReference() *typeDefinitionClass_ {
	return typeDefinitionReference_
}

var typeDefinitionReference_ = &typeDefinitionClass_{
	// Initialize the class constants.
}
