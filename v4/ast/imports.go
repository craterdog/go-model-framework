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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Imports() ImportsClassLike {
	return importsReference()
}

// Constructor Methods

func (c *importsClass_) Make(
	modules abs.Sequential[ModuleLike],
) ImportsLike {
	if uti.IsUndefined(modules) {
		panic("The \"modules\" attribute is required by this class.")
	}
	var instance = &imports_{
		// Initialize the instance attributes.
		modules_: modules,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *imports_) GetModules() abs.Sequential[ModuleLike] {
	return v.modules_
}

// Public Methods

func (v *imports_) GetClass() ImportsClassLike {
	return v.getClass()
}

// Private Methods

func (v *imports_) getClass() *importsClass_ {
	return importsReference()
}

// PRIVATE INTERFACE

// Instance Structure

type imports_ struct {
	// Declare the instance attributes.
	modules_ abs.Sequential[ModuleLike]
}

// Class Structure

type importsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func importsReference() *importsClass_ {
	return importsReference_
}

var importsReference_ = &importsClass_{
	// Initialize the class constants.
}
