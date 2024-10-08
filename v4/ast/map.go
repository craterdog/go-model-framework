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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Map() MapClassLike {
	return mapClass
}

// Constructor Methods

func (c *mapClass_) Make(
	name string,
) MapLike {
	if uti.IsUndefined(name) {
		panic("The name attribute is required by this class.")
	}
	var instance = &map_{
		class_: c,
		name_:  name,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *map_) GetName() string {
	return v.name_
}

// Public Methods

func (v *map_) GetClass() MapClassLike {
	return v.getClass()
}

// Private Methods

func (v *map_) getClass() *mapClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type map_ struct {
	class_ *mapClass_
	name_  string
}

// Class Structure

type mapClass_ struct {
	// Define the class constants.
}

// Class Reference

var mapClass = &mapClass_{
	// Initialize the class constants.
}
