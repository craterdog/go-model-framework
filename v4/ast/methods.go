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

var methodsClass = &methodsClass_{
	// Initialize class constants.
}

// Function

func Methods() MethodsClassLike {
	return methodsClass
}

// CLASS METHODS

// Target

type methodsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *methodsClass_) Make(
	note string,
	methods col.ListLike[MethodLike],
) MethodsLike {
	return &methods_{
		// Initialize instance attributes.
		class_:   c,
		note_:    note,
		methods_: methods,
	}
}

// INSTANCE METHODS

// Target

type methods_ struct {
	// Define instance attributes.
	class_   MethodsClassLike
	note_    string
	methods_ col.ListLike[MethodLike]
}

// Attributes

func (v *methods_) GetClass() MethodsClassLike {
	return v.class_
}

func (v *methods_) GetNote() string {
	return v.note_
}

func (v *methods_) GetMethods() col.ListLike[MethodLike] {
	return v.methods_
}

// Private
