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

var attributeMethodsClass = &attributeMethodsClass_{
	// Initialize class constants.
}

// Function

func AttributeMethods() AttributeMethodsClassLike {
	return attributeMethodsClass
}

// CLASS METHODS

// Target

type attributeMethodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *attributeMethodsClass_) Make(attributes abs.Sequential[AttributeLike]) AttributeMethodsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(attributes):
		panic("The attributes attribute is required by this class.")
	default:
		return &attributeMethods_{
			// Initialize instance attributes.
			class_:      c,
			attributes_: attributes,
		}
	}
}

// INSTANCE METHODS

// Target

type attributeMethods_ struct {
	// Define instance attributes.
	class_      AttributeMethodsClassLike
	attributes_ abs.Sequential[AttributeLike]
}

// Public

func (v *attributeMethods_) GetClass() AttributeMethodsClassLike {
	return v.class_
}

// Attribute

func (v *attributeMethods_) GetAttributes() abs.Sequential[AttributeLike] {
	return v.attributes_
}

// Private
