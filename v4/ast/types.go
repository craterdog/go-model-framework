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

var typesClass = &typesClass_{
	// Initialize class constants.
}

// Function

func Types() TypesClassLike {
	return typesClass
}

// CLASS METHODS

// Target

type typesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *typesClass_) Make(
	note string,
	typeIterator age.IteratorLike[TypeLike],
) TypesLike {
	return &types_{
		// Initialize instance attributes.
		class_:        c,
		note_:         note,
		typeIterator_: typeIterator,
	}
}

// INSTANCE METHODS

// Target

type types_ struct {
	// Define instance attributes.
	class_        TypesClassLike
	note_         string
	typeIterator_ age.IteratorLike[TypeLike]
}

// Attributes

func (v *types_) GetClass() TypesClassLike {
	return v.class_
}

func (v *types_) GetNote() string {
	return v.note_
}

func (v *types_) GetTypeIterator() age.IteratorLike[TypeLike] {
	return v.typeIterator_
}

// Private
