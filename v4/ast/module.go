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

func Module() ModuleClassLike {
	return moduleReference()
}

// Constructor Methods

func (c *moduleClass_) Make(
	name string,
	path string,
) ModuleLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(path) {
		panic("The \"path\" attribute is required by this class.")
	}
	var instance = &module_{
		// Initialize the instance attributes.
		name_: name,
		path_: path,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *module_) GetName() string {
	return v.name_
}

func (v *module_) GetPath() string {
	return v.path_
}

// Public Methods

func (v *module_) GetClass() ModuleClassLike {
	return v.getClass()
}

// Private Methods

func (v *module_) getClass() *moduleClass_ {
	return moduleReference()
}

// PRIVATE INTERFACE

// Instance Structure

type module_ struct {
	// Declare the instance attributes.
	name_ string
	path_ string
}

// Class Structure

type moduleClass_ struct {
	// Declare the class constants.
}

// Class Reference

func moduleReference() *moduleClass_ {
	return moduleReference_
}

var moduleReference_ = &moduleClass_{
	// Initialize the class constants.
}
