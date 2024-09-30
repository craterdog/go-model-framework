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
)

// CLASS ACCESS

// Reference

var moduleClass = &moduleClass_{
	// Initialize class constants.
}

// Function

func Module() ModuleClassLike {
	return moduleClass
}

// CLASS METHODS

// Target

type moduleClass_ struct {
	// Define class constants.
}

// Constructors

func (c *moduleClass_) Make(
	name string,
	path string,
) ModuleLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	case col.IsUndefined(path):
		panic("The path attribute is required by this class.")
	default:
		return &module_{
			// Initialize instance attributes.
			class_: c,
			name_:  name,
			path_:  path,
		}
	}
}

// INSTANCE METHODS

// Target

type module_ struct {
	// Define instance attributes.
	class_ ModuleClassLike
	name_  string
	path_  string
}

// Public

func (v *module_) GetClass() ModuleClassLike {
	return v.class_
}

// Attribute

func (v *module_) GetName() string {
	return v.name_
}

func (v *module_) GetPath() string {
	return v.path_
}

// Private
