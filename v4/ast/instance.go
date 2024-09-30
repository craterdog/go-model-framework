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

var instanceClass = &instanceClass_{
	// Initialize class constants.
}

// Function

func Instance() InstanceClassLike {
	return instanceClass
}

// CLASS METHODS

// Target

type instanceClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instanceClass_) Make(
	declaration DeclarationLike,
	instanceMethods InstanceMethodsLike,
) InstanceLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(declaration):
		panic("The declaration attribute is required by this class.")
	case col.IsUndefined(instanceMethods):
		panic("The instanceMethods attribute is required by this class.")
	default:
		return &instance_{
			// Initialize instance attributes.
			class_:           c,
			declaration_:     declaration,
			instanceMethods_: instanceMethods,
		}
	}
}

// INSTANCE METHODS

// Target

type instance_ struct {
	// Define instance attributes.
	class_           InstanceClassLike
	declaration_     DeclarationLike
	instanceMethods_ InstanceMethodsLike
}

// Public

func (v *instance_) GetClass() InstanceClassLike {
	return v.class_
}

// Attribute

func (v *instance_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instance_) GetInstanceMethods() InstanceMethodsLike {
	return v.instanceMethods_
}

// Private
