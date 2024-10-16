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

func AspectDefinition() AspectDefinitionClassLike {
	return aspectDefinitionReference()
}

// Constructor Methods

func (c *aspectDefinitionClass_) Make(
	declaration DeclarationLike,
	aspectMethods abs.Sequential[AspectMethodLike],
) AspectDefinitionLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(aspectMethods) {
		panic("The \"aspectMethods\" attribute is required by this class.")
	}
	var instance = &aspectDefinition_{
		// Initialize the instance attributes.
		declaration_:   declaration,
		aspectMethods_: aspectMethods,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *aspectDefinition_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *aspectDefinition_) GetAspectMethods() abs.Sequential[AspectMethodLike] {
	return v.aspectMethods_
}

// Public Methods

func (v *aspectDefinition_) GetClass() AspectDefinitionClassLike {
	return v.getClass()
}

// Private Methods

func (v *aspectDefinition_) getClass() *aspectDefinitionClass_ {
	return aspectDefinitionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type aspectDefinition_ struct {
	// Declare the instance attributes.
	declaration_   DeclarationLike
	aspectMethods_ abs.Sequential[AspectMethodLike]
}

// Class Structure

type aspectDefinitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectDefinitionReference() *aspectDefinitionClass_ {
	return aspectDefinitionReference_
}

var aspectDefinitionReference_ = &aspectDefinitionClass_{
	// Initialize the class constants.
}
