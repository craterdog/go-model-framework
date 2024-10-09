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

func AttributeSubsection() AttributeSubsectionClassLike {
	return attributeSubsectionClass
}

// Constructor Methods

func (c *attributeSubsectionClass_) Make(
	attributeMethods abs.Sequential[AttributeMethodLike],
) AttributeSubsectionLike {
	if uti.IsUndefined(attributeMethods) {
		panic("The attributeMethods attribute is required by this class.")
	}
	var instance = &attributeSubsection_{
		class_:            c,
		attributeMethods_: attributeMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *attributeSubsection_) GetAttributeMethods() abs.Sequential[AttributeMethodLike] {
	return v.attributeMethods_
}

// Public Methods

func (v *attributeSubsection_) GetClass() AttributeSubsectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *attributeSubsection_) getClass() *attributeSubsectionClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type attributeSubsection_ struct {
	class_            *attributeSubsectionClass_
	attributeMethods_ abs.Sequential[AttributeMethodLike]
}

// Class Structure

type attributeSubsectionClass_ struct {
	// Define the class constants.
}

// Class Reference

var attributeSubsectionClass = &attributeSubsectionClass_{
	// Initialize the class constants.
}
