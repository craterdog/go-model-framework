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
	types col.Sequential[TypeLike],
) TypesLike {
	return &types_{
		// Initialize instance attributes.
		class_: c,
		note_:  note,
		types_: types,
	}
}

// INSTANCE METHODS

// Target

type types_ struct {
	// Define instance attributes.
	class_ TypesClassLike
	note_  string
	types_ col.Sequential[TypeLike]
}

// Attributes

func (v *types_) GetClass() TypesClassLike {
	return v.class_
}

func (v *types_) GetNote() string {
	return v.note_
}

func (v *types_) GetTypes() col.Sequential[TypeLike] {
	return v.types_
}

// Private
