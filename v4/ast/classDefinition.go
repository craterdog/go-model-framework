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

func ClassDefinition() ClassDefinitionClassLike {
	return classDefinitionClass
}

// Constructor Methods

func (c *classDefinitionClass_) Make(
	declaration DeclarationLike,
	classMethods ClassMethodsLike,
) ClassDefinitionLike {
	if uti.IsUndefined(declaration) {
		panic("The declaration attribute is required by this class.")
	}
	if uti.IsUndefined(classMethods) {
		panic("The classMethods attribute is required by this class.")
	}
	var instance = &classDefinition_{
		class_:        c,
		declaration_:  declaration,
		classMethods_: classMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *classDefinition_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *classDefinition_) GetClassMethods() ClassMethodsLike {
	return v.classMethods_
}

// Public Methods

func (v *classDefinition_) GetClass() ClassDefinitionClassLike {
	return v.getClass()
}

// Private Methods

func (v *classDefinition_) getClass() *classDefinitionClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type classDefinition_ struct {
	class_        *classDefinitionClass_
	declaration_  DeclarationLike
	classMethods_ ClassMethodsLike
}

// Class Structure

type classDefinitionClass_ struct {
	// Define the class constants.
}

// Class Reference

var classDefinitionClass = &classDefinitionClass_{
	// Initialize the class constants.
}
