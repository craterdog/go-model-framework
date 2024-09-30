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

var instanceMethodsClass = &instanceMethodsClass_{
	// Initialize class constants.
}

// Function

func InstanceMethods() InstanceMethodsClassLike {
	return instanceMethodsClass
}

// CLASS METHODS

// Target

type instanceMethodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instanceMethodsClass_) Make(
	publicMethods PublicMethodsLike,
	optionalAttributeMethods AttributeMethodsLike,
	optionalAspectInterfaces AspectInterfacesLike,
) InstanceMethodsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(publicMethods):
		panic("The publicMethods attribute is required by this class.")
	default:
		return &instanceMethods_{
			// Initialize instance attributes.
			class_:                    c,
			publicMethods_:            publicMethods,
			optionalAttributeMethods_: optionalAttributeMethods,
			optionalAspectInterfaces_: optionalAspectInterfaces,
		}
	}
}

// INSTANCE METHODS

// Target

type instanceMethods_ struct {
	// Define instance attributes.
	class_                    InstanceMethodsClassLike
	publicMethods_            PublicMethodsLike
	optionalAttributeMethods_ AttributeMethodsLike
	optionalAspectInterfaces_ AspectInterfacesLike
}

// Public

func (v *instanceMethods_) GetClass() InstanceMethodsClassLike {
	return v.class_
}

// Attribute

func (v *instanceMethods_) GetPublicMethods() PublicMethodsLike {
	return v.publicMethods_
}

func (v *instanceMethods_) GetOptionalAttributeMethods() AttributeMethodsLike {
	return v.optionalAttributeMethods_
}

func (v *instanceMethods_) GetOptionalAspectInterfaces() AspectInterfacesLike {
	return v.optionalAspectInterfaces_
}

// Private
