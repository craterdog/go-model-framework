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
	instances col.ListLike[InstanceLike],
) InstancesLike {
	return &instances_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type instances_ struct {
	// Define instance attributes.
	class_     InstancesClassLike
	note_      string
	instances_ col.ListLike[InstanceLike]
}

// Attributes

func (v *instances_) GetClass() InstancesClassLike {
	return v.class_
}

func (v *instances_) GetNote() string {
	return v.note_
}

func (v *instances_) GetInstances() col.ListLike[InstanceLike] {
	return v.instances_
}

// Private
