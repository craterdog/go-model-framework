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
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
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
	constructors abs.Sequential[ConstructorLike],
) ConstructorsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(note):
		panic("The note attribute is required by this class.")
	case col.IsUndefined(constructors):
		panic("The constructors attribute is required by this class.")
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
	constructors_ abs.Sequential[ConstructorLike]
}

// Attributes

func (v *constructors_) GetClass() ConstructorsClassLike {
	return v.class_
}

func (v *constructors_) GetNote() string {
	return v.note_
}

func (v *constructors_) GetConstructors() abs.Sequential[ConstructorLike] {
	return v.constructors_
}

// Private
