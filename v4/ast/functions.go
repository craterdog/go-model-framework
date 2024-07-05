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
	functions col.Sequential[FunctionLike],
) FunctionsLike {
	switch {
	case len(note) == 0:
		panic("The note attribute is required for each sequence of functions.")
	case functions == nil || functions.IsEmpty():
		panic("At least one function is required for each sequence of functions.")
	default:
		return &functions_{
			// Initialize instance attributes.
			class_:     c,
			note_:      note,
			functions_: functions,
		}
	}
}

// INSTANCE METHODS

// Target

type functions_ struct {
	// Define instance attributes.
	class_     FunctionsClassLike
	note_      string
	functions_ col.Sequential[FunctionLike]
}

// Attributes

func (v *functions_) GetClass() FunctionsClassLike {
	return v.class_
}

func (v *functions_) GetNote() string {
	return v.note_
}

func (v *functions_) GetFunctions() col.Sequential[FunctionLike] {
	return v.functions_
}

// Private
