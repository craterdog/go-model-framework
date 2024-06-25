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

import ()

// CLASS ACCESS

// Reference

var aspectClass = &aspectClass_{
	// Initialize class constants.
}

// Function

func Aspect() AspectClassLike {
	return aspectClass
}

// CLASS METHODS

// Target

type aspectClass_ struct {
	// Define class constants.
}

// Constructors

func (c *aspectClass_) Make(
	declaration DeclarationLike,
	methods MethodsLike,
) AspectLike {
	return &aspect_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type aspect_ struct {
	// Define instance attributes.
	class_       AspectClassLike
	declaration_ DeclarationLike
	methods_     MethodsLike
}

// Attributes

func (v *aspect_) GetClass() AspectClassLike {
	return v.class_
}

func (v *aspect_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *aspect_) GetMethods() MethodsLike {
	return v.methods_
}

// Private
