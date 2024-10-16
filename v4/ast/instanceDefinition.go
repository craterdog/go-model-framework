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

func InstanceDefinition() InstanceDefinitionClassLike {
	return instanceDefinitionReference()
}

// Constructor Methods

func (c *instanceDefinitionClass_) Make(
	declaration DeclarationLike,
	instanceMethods InstanceMethodsLike,
) InstanceDefinitionLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(instanceMethods) {
		panic("The \"instanceMethods\" attribute is required by this class.")
	}
	var instance = &instanceDefinition_{
		// Initialize the instance attributes.
		declaration_:     declaration,
		instanceMethods_: instanceMethods,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *instanceDefinition_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instanceDefinition_) GetInstanceMethods() InstanceMethodsLike {
	return v.instanceMethods_
}

// Public Methods

func (v *instanceDefinition_) GetClass() InstanceDefinitionClassLike {
	return v.getClass()
}

// Private Methods

func (v *instanceDefinition_) getClass() *instanceDefinitionClass_ {
	return instanceDefinitionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type instanceDefinition_ struct {
	// Declare the instance attributes.
	declaration_     DeclarationLike
	instanceMethods_ InstanceMethodsLike
}

// Class Structure

type instanceDefinitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceDefinitionReference() *instanceDefinitionClass_ {
	return instanceDefinitionReference_
}

var instanceDefinitionReference_ = &instanceDefinitionClass_{
	// Initialize the class constants.
}
