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
)

// CLASS ACCESS

// Reference

var modelClass = &modelClass_{
	// Initialize class constants.
}

// Function

func Model() ModelClassLike {
	return modelClass
}

// CLASS METHODS

// Target

type modelClass_ struct {
	// Define class constants.
}

// Constructors

func (c *modelClass_) Make(
	moduleDefinition ModuleDefinitionLike,
	primitiveDefinitions PrimitiveDefinitionsLike,
	interfaceDefinitions InterfaceDefinitionsLike,
) ModelLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(moduleDefinition):
		panic("The moduleDefinition attribute is required by this class.")
	case col.IsUndefined(primitiveDefinitions):
		panic("The primitiveDefinitions attribute is required by this class.")
	case col.IsUndefined(interfaceDefinitions):
		panic("The interfaceDefinitions attribute is required by this class.")
	default:
		return &model_{
			// Initialize instance attributes.
			class_:                c,
			moduleDefinition_:     moduleDefinition,
			primitiveDefinitions_: primitiveDefinitions,
			interfaceDefinitions_: interfaceDefinitions,
		}
	}
}

// INSTANCE METHODS

// Target

type model_ struct {
	// Define instance attributes.
	class_                ModelClassLike
	moduleDefinition_     ModuleDefinitionLike
	primitiveDefinitions_ PrimitiveDefinitionsLike
	interfaceDefinitions_ InterfaceDefinitionsLike
}

// Public

func (v *model_) GetClass() ModelClassLike {
	return v.class_
}

// Attribute

func (v *model_) GetModuleDefinition() ModuleDefinitionLike {
	return v.moduleDefinition_
}

func (v *model_) GetPrimitiveDefinitions() PrimitiveDefinitionsLike {
	return v.primitiveDefinitions_
}

func (v *model_) GetInterfaceDefinitions() InterfaceDefinitionsLike {
	return v.interfaceDefinitions_
}

// Private
