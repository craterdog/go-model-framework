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

func InterfaceDefinitions() InterfaceDefinitionsClassLike {
	return interfaceDefinitionsReference()
}

// Constructor Methods

func (c *interfaceDefinitionsClass_) Make(
	classSection ClassSectionLike,
	instanceSection InstanceSectionLike,
	optionalAspectSection AspectSectionLike,
) InterfaceDefinitionsLike {
	if uti.IsUndefined(classSection) {
		panic("The \"classSection\" attribute is required by this class.")
	}
	if uti.IsUndefined(instanceSection) {
		panic("The \"instanceSection\" attribute is required by this class.")
	}
	var instance = &interfaceDefinitions_{
		// Initialize the instance attributes.
		classSection_:          classSection,
		instanceSection_:       instanceSection,
		optionalAspectSection_: optionalAspectSection,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *interfaceDefinitions_) GetClassSection() ClassSectionLike {
	return v.classSection_
}

func (v *interfaceDefinitions_) GetInstanceSection() InstanceSectionLike {
	return v.instanceSection_
}

func (v *interfaceDefinitions_) GetOptionalAspectSection() AspectSectionLike {
	return v.optionalAspectSection_
}

// Public Methods

func (v *interfaceDefinitions_) GetClass() InterfaceDefinitionsClassLike {
	return v.getClass()
}

// Private Methods

func (v *interfaceDefinitions_) getClass() *interfaceDefinitionsClass_ {
	return interfaceDefinitionsReference()
}

// PRIVATE INTERFACE

// Instance Structure

type interfaceDefinitions_ struct {
	// Declare the instance attributes.
	classSection_          ClassSectionLike
	instanceSection_       InstanceSectionLike
	optionalAspectSection_ AspectSectionLike
}

// Class Structure

type interfaceDefinitionsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func interfaceDefinitionsReference() *interfaceDefinitionsClass_ {
	return interfaceDefinitionsReference_
}

var interfaceDefinitionsReference_ = &interfaceDefinitionsClass_{
	// Initialize the class constants.
}
