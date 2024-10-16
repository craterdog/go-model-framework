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

func InstanceSection() InstanceSectionClassLike {
	return instanceSectionReference()
}

// Constructor Methods

func (c *instanceSectionClass_) Make(
	instanceDefinitions abs.Sequential[InstanceDefinitionLike],
) InstanceSectionLike {
	if uti.IsUndefined(instanceDefinitions) {
		panic("The \"instanceDefinitions\" attribute is required by this class.")
	}
	var instance = &instanceSection_{
		// Initialize the instance attributes.
		instanceDefinitions_: instanceDefinitions,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *instanceSection_) GetInstanceDefinitions() abs.Sequential[InstanceDefinitionLike] {
	return v.instanceDefinitions_
}

// Public Methods

func (v *instanceSection_) GetClass() InstanceSectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *instanceSection_) getClass() *instanceSectionClass_ {
	return instanceSectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type instanceSection_ struct {
	// Declare the instance attributes.
	instanceDefinitions_ abs.Sequential[InstanceDefinitionLike]
}

// Class Structure

type instanceSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func instanceSectionReference() *instanceSectionClass_ {
	return instanceSectionReference_
}

var instanceSectionReference_ = &instanceSectionClass_{
	// Initialize the class constants.
}
