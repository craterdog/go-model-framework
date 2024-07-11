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
	mod "github.com/craterdog/go-collection-framework/v4"
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var instancesClass = &instancesClass_{
	// Initialize class constants.
}

// Function

func Instances() InstancesClassLike {
	return instancesClass
}

// CLASS METHODS

// Target

type instancesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instancesClass_) Make(
	note string,
	instances col.Sequential[InstanceLike],
) InstancesLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(note):
		panic("The note attribute is required for each Instances.")
	case mod.IsUndefined(instances):
		panic("The instances attribute is required for each Instances.")
	default:
		return &instances_{
			// Initialize instance attributes.
			class_:     c,
			note_:      note,
			instances_: instances,
		}
	}
}

// INSTANCE METHODS

// Target

type instances_ struct {
	// Define instance attributes.
	class_     InstancesClassLike
	note_      string
	instances_ col.Sequential[InstanceLike]
}

// Attributes

func (v *instances_) GetClass() InstancesClassLike {
	return v.class_
}

func (v *instances_) GetNote() string {
	return v.note_
}

func (v *instances_) GetInstances() col.Sequential[InstanceLike] {
	return v.instances_
}

// Private
