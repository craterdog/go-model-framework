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

var abstractionsClass = &abstractionsClass_{
	// Initialize class constants.
}

// Function

func Abstractions() AbstractionsClassLike {
	return abstractionsClass
}

// CLASS METHODS

// Target

type abstractionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *abstractionsClass_) Make(
	note string,
	abstractions col.Sequential[AbstractionLike],
) AbstractionsLike {
	switch {
	case len(note) == 0:
		panic("The note attribute is required for each sequence of abstractions.")
	case abstractions == nil || abstractions.IsEmpty():
		panic("At least one abstraction is required for each sequence of abstractions.")
	default:
		return &abstractions_{
			// Initialize instance attributes.
			class_:        c,
			note_:         note,
			abstractions_: abstractions,
		}
	}
}

// INSTANCE METHODS

// Target

type abstractions_ struct {
	// Define instance attributes.
	class_        AbstractionsClassLike
	note_         string
	abstractions_ col.Sequential[AbstractionLike]
}

// Attributes

func (v *abstractions_) GetClass() AbstractionsClassLike {
	return v.class_
}

func (v *abstractions_) GetNote() string {
	return v.note_
}

func (v *abstractions_) GetAbstractions() col.Sequential[AbstractionLike] {
	return v.abstractions_
}

// Private
