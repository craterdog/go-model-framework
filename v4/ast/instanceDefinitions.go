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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var instanceDefinitionsClass = &instanceDefinitionsClass_{
	// Initialize class constants.
}

// Function

func InstanceDefinitions() InstanceDefinitionsClassLike {
	return instanceDefinitionsClass
}

// CLASS METHODS

// Target

type instanceDefinitionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instanceDefinitionsClass_) Make(instances abs.Sequential[InstanceLike]) InstanceDefinitionsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(instances):
		panic("The instances attribute is required by this class.")
	default:
		return &instanceDefinitions_{
			// Initialize instance attributes.
			class_:     c,
			instances_: instances,
		}
	}
}

// INSTANCE METHODS

// Target

type instanceDefinitions_ struct {
	// Define instance attributes.
	class_     InstanceDefinitionsClassLike
	instances_ abs.Sequential[InstanceLike]
}

// Public

func (v *instanceDefinitions_) GetClass() InstanceDefinitionsClassLike {
	return v.class_
}

// Attribute

func (v *instanceDefinitions_) GetInstances() abs.Sequential[InstanceLike] {
	return v.instances_
}

// Private
