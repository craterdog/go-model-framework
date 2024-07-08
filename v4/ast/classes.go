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

var classesClass = &classesClass_{
	// Initialize class constants.
}

// Function

func Classes() ClassesClassLike {
	return classesClass
}

// CLASS METHODS

// Target

type classesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *classesClass_) Make(
	note string,
	classes col.Sequential[ClassLike],
) ClassesLike {
	// Validate the arguments.
	switch {
	case isUndefined(note):
		panic("The note attribute is required for each Classes.")
	case isUndefined(classes):
		panic("The classes attribute is required for each Classes.")
	default:
		return &classes_{
			// Initialize instance attributes.
			class_:   c,
			note_:    note,
			classes_: classes,
		}
	}
}

// INSTANCE METHODS

// Target

type classes_ struct {
	// Define instance attributes.
	class_   ClassesClassLike
	note_    string
	classes_ col.Sequential[ClassLike]
}

// Attributes

func (v *classes_) GetClass() ClassesClassLike {
	return v.class_
}

func (v *classes_) GetNote() string {
	return v.note_
}

func (v *classes_) GetClasses() col.Sequential[ClassLike] {
	return v.classes_
}

// Private
