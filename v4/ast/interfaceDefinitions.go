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

var interfaceDefinitionsClass = &interfaceDefinitionsClass_{
	// Initialize class constants.
}

// Function

func InterfaceDefinitions() InterfaceDefinitionsClassLike {
	return interfaceDefinitionsClass
}

// CLASS METHODS

// Target

type interfaceDefinitionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *interfaceDefinitionsClass_) Make(
	classDefinitions ClassDefinitionsLike,
	instanceDefinitions InstanceDefinitionsLike,
	optionalAspectDefinitions AspectDefinitionsLike,
) InterfaceDefinitionsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(classDefinitions):
		panic("The classDefinitions attribute is required by this class.")
	case col.IsUndefined(instanceDefinitions):
		panic("The instanceDefinitions attribute is required by this class.")
	default:
		return &interfaceDefinitions_{
			// Initialize instance attributes.
			class_:                     c,
			classDefinitions_:          classDefinitions,
			instanceDefinitions_:       instanceDefinitions,
			optionalAspectDefinitions_: optionalAspectDefinitions,
		}
	}
}

// INSTANCE METHODS

// Target

type interfaceDefinitions_ struct {
	// Define instance attributes.
	class_                     InterfaceDefinitionsClassLike
	classDefinitions_          ClassDefinitionsLike
	instanceDefinitions_       InstanceDefinitionsLike
	optionalAspectDefinitions_ AspectDefinitionsLike
}

// Public

func (v *interfaceDefinitions_) GetClass() InterfaceDefinitionsClassLike {
	return v.class_
}

// Attribute

func (v *interfaceDefinitions_) GetClassDefinitions() ClassDefinitionsLike {
	return v.classDefinitions_
}

func (v *interfaceDefinitions_) GetInstanceDefinitions() InstanceDefinitionsLike {
	return v.instanceDefinitions_
}

func (v *interfaceDefinitions_) GetOptionalAspectDefinitions() AspectDefinitionsLike {
	return v.optionalAspectDefinitions_
}

// Private
