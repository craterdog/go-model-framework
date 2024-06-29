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
	age "github.com/craterdog/go-collection-framework/v4/agent"
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
	methodIterator age.IteratorLike[MethodLike],
) MethodsLike {
	return &methods_{
		// Initialize instance attributes.
		class_:          c,
		note_:           note,
		methodIterator_: methodIterator,
	}
}

// INSTANCE METHODS

// Target

type methods_ struct {
	// Define instance attributes.
	class_          MethodsClassLike
	note_           string
	methodIterator_ age.IteratorLike[MethodLike]
}

// Attributes

func (v *methods_) GetClass() MethodsClassLike {
	return v.class_
}

func (v *methods_) GetNote() string {
	return v.note_
}

func (v *methods_) GetMethodIterator() age.IteratorLike[MethodLike] {
	return v.methodIterator_
}

// Private
