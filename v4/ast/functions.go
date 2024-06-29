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

var functionsClass = &functionsClass_{
	// Initialize class constants.
}

// Function

func Functions() FunctionsClassLike {
	return functionsClass
}

// CLASS METHODS

// Target

type functionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *functionsClass_) Make(
	note string,
	functionIterator age.IteratorLike[FunctionLike],
) FunctionsLike {
	return &functions_{
		// Initialize instance attributes.
		class_:            c,
		note_:             note,
		functionIterator_: functionIterator,
	}
}

// INSTANCE METHODS

// Target

type functions_ struct {
	// Define instance attributes.
	class_            FunctionsClassLike
	note_             string
	functionIterator_ age.IteratorLike[FunctionLike]
}

// Attributes

func (v *functions_) GetClass() FunctionsClassLike {
	return v.class_
}

func (v *functions_) GetNote() string {
	return v.note_
}

func (v *functions_) GetFunctionIterator() age.IteratorLike[FunctionLike] {
	return v.functionIterator_
}

// Private
