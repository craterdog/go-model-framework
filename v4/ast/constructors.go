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

var constructorsClass = &constructorsClass_{
	// Initialize class constants.
}

// Function

func Constructors() ConstructorsClassLike {
	return constructorsClass
}

// CLASS METHODS

// Target

type constructorsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *constructorsClass_) Make(
	note string,
	constructors col.Sequential[ConstructorLike],
) ConstructorsLike {
	switch {
	case len(note) == 0:
		panic("The note attribute is required for each sequence of constructors.")
	case constructors == nil || constructors.IsEmpty():
		panic("At least one constructor is required for each sequence of constructors.")
	default:
		return &constructors_{
			// Initialize instance attributes.
			class_:        c,
			note_:         note,
			constructors_: constructors,
		}
	}
}

// INSTANCE METHODS

// Target

type constructors_ struct {
	// Define instance attributes.
	class_        ConstructorsClassLike
	note_         string
	constructors_ col.Sequential[ConstructorLike]
}

// Attributes

func (v *constructors_) GetClass() ConstructorsClassLike {
	return v.class_
}

func (v *constructors_) GetNote() string {
	return v.note_
}

func (v *constructors_) GetConstructors() col.Sequential[ConstructorLike] {
	return v.constructors_
}

// Private
