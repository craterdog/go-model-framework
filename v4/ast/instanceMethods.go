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

func InstanceMethods() InstanceMethodsClassLike {
	return instanceMethodsClass
}

// Constructor Methods

func (c *instanceMethodsClass_) Make(
	publicSubsection PublicSubsectionLike,
	optionalAttributeSubsection AttributeSubsectionLike,
	optionalAspectSubsection AspectSubsectionLike,
) InstanceMethodsLike {
	if uti.IsUndefined(publicSubsection) {
		panic("The \"publicSubsection\" attribute is required by this class.")
	}
	var instance = &instanceMethods_{
		// Initialize the instance attributes.
		class_:                       c,
		publicSubsection_:            publicSubsection,
		optionalAttributeSubsection_: optionalAttributeSubsection,
		optionalAspectSubsection_:    optionalAspectSubsection,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *instanceMethods_) GetPublicSubsection() PublicSubsectionLike {
	return v.publicSubsection_
}

func (v *instanceMethods_) GetOptionalAttributeSubsection() AttributeSubsectionLike {
	return v.optionalAttributeSubsection_
}

func (v *instanceMethods_) GetOptionalAspectSubsection() AspectSubsectionLike {
	return v.optionalAspectSubsection_
}

// Public Methods

func (v *instanceMethods_) GetClass() InstanceMethodsClassLike {
	return v.getClass()
}

// Private Methods

func (v *instanceMethods_) getClass() *instanceMethodsClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type instanceMethods_ struct {
	// Declare the instance attributes.
	class_                       *instanceMethodsClass_
	publicSubsection_            PublicSubsectionLike
	optionalAttributeSubsection_ AttributeSubsectionLike
	optionalAspectSubsection_    AspectSubsectionLike
}

// Class Structure

type instanceMethodsClass_ struct {
	// Declare the class constants.
}

// Class Reference

var instanceMethodsClass = &instanceMethodsClass_{
	// Initialize the class constants.
}
