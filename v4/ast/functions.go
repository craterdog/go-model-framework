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
	functions col.ListLike[FunctionLike],
) FunctionsLike {
	return &functions_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type functions_ struct {
	// Define instance attributes.
	class_     FunctionsClassLike
	note_      string
	functions_ col.ListLike[FunctionLike]
}

// Attributes

func (v *functions_) GetClass() FunctionsClassLike {
	return v.class_
}

func (v *functions_) GetNote() string {
	return v.note_
}

func (v *functions_) GetFunctions() col.ListLike[FunctionLike] {
	return v.functions_
}

// Private
