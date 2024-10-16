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

func Model() ModelClassLike {
	return modelReference()
}

// Constructor Methods

func (c *modelClass_) Make(
	moduleDefinition ModuleDefinitionLike,
	primitiveDefinitions PrimitiveDefinitionsLike,
	interfaceDefinitions InterfaceDefinitionsLike,
) ModelLike {
	if uti.IsUndefined(moduleDefinition) {
		panic("The \"moduleDefinition\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitiveDefinitions) {
		panic("The \"primitiveDefinitions\" attribute is required by this class.")
	}
	if uti.IsUndefined(interfaceDefinitions) {
		panic("The \"interfaceDefinitions\" attribute is required by this class.")
	}
	var instance = &model_{
		// Initialize the instance attributes.
		moduleDefinition_:     moduleDefinition,
		primitiveDefinitions_: primitiveDefinitions,
		interfaceDefinitions_: interfaceDefinitions,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *model_) GetModuleDefinition() ModuleDefinitionLike {
	return v.moduleDefinition_
}

func (v *model_) GetPrimitiveDefinitions() PrimitiveDefinitionsLike {
	return v.primitiveDefinitions_
}

func (v *model_) GetInterfaceDefinitions() InterfaceDefinitionsLike {
	return v.interfaceDefinitions_
}

// Public Methods

func (v *model_) GetClass() ModelClassLike {
	return v.getClass()
}

// Private Methods

func (v *model_) getClass() *modelClass_ {
	return modelReference()
}

// PRIVATE INTERFACE

// Instance Structure

type model_ struct {
	// Declare the instance attributes.
	moduleDefinition_     ModuleDefinitionLike
	primitiveDefinitions_ PrimitiveDefinitionsLike
	interfaceDefinitions_ InterfaceDefinitionsLike
}

// Class Structure

type modelClass_ struct {
	// Declare the class constants.
}

// Class Reference

func modelReference() *modelClass_ {
	return modelReference_
}

var modelReference_ = &modelClass_{
	// Initialize the class constants.
}
