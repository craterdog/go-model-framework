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
	constructorIterator age.IteratorLike[ConstructorLike],
) ConstructorsLike {
	return &constructors_{
		// Initialize instance attributes.
		class_:               c,
		note_:                note,
		constructorIterator_: constructorIterator,
	}
}

// INSTANCE METHODS

// Target

type constructors_ struct {
	// Define instance attributes.
	class_               ConstructorsClassLike
	note_                string
	constructorIterator_ age.IteratorLike[ConstructorLike]
}

// Attributes

func (v *constructors_) GetClass() ConstructorsClassLike {
	return v.class_
}

func (v *constructors_) GetNote() string {
	return v.note_
}

func (v *constructors_) GetConstructorIterator() age.IteratorLike[ConstructorLike] {
	return v.constructorIterator_
}

// Private
