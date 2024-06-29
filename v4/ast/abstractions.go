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
	abstractionIterator age.IteratorLike[AbstractionLike],
) AbstractionsLike {
	return &abstractions_{
		// Initialize instance attributes.
		class_:               c,
		note_:                note,
		abstractionIterator_: abstractionIterator,
	}
}

// INSTANCE METHODS

// Target

type abstractions_ struct {
	// Define instance attributes.
	class_               AbstractionsClassLike
	note_                string
	abstractionIterator_ age.IteratorLike[AbstractionLike]
}

// Attributes

func (v *abstractions_) GetClass() AbstractionsClassLike {
	return v.class_
}

func (v *abstractions_) GetNote() string {
	return v.note_
}

func (v *abstractions_) GetAbstractionIterator() age.IteratorLike[AbstractionLike] {
	return v.abstractionIterator_
}

// Private
