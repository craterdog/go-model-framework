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

var aspectClass = &aspectClass_{
	// Any private class constants should be initialized here.
}

// Function

func Aspect() AspectClassLike {
	return aspectClass
}

// CLASS METHODS

// Target

type aspectClass_ struct {
}

// Constants

// Constructors

func (c *aspectClass_) MakeWithAttributes(
	declaration DeclarationLike,
	methods col.ListLike[MethodLike],
) AspectLike {
	return &aspect_{
		declaration_: declaration,
		methods_:     methods,
	}
}

// Functions

// INSTANCE METHODS

// Target

type aspect_ struct {
	class_       AspectClassLike
	declaration_ DeclarationLike
	methods_     col.ListLike[MethodLike]
}

// Attributes

func (v *aspect_) GetClass() AspectClassLike {
	return v.class_
}

func (v *aspect_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *aspect_) GetMethods() col.ListLike[MethodLike] {
	return v.methods_
}

// Public

// Private
