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

var importsClass = &importsClass_{
	// Initialize class constants.
}

// Function

func Imports() ImportsClassLike {
	return importsClass
}

// CLASS METHODS

// Target

type importsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *importsClass_) Make(modules abs.Sequential[ModuleLike]) ImportsLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(modules):
		panic("The modules attribute is required by this class.")
	default:
		return &imports_{
			// Initialize instance attributes.
			class_:   c,
			modules_: modules,
		}
	}
}

// INSTANCE METHODS

// Target

type imports_ struct {
	// Define instance attributes.
	class_   ImportsClassLike
	modules_ abs.Sequential[ModuleLike]
}

// Public

func (v *imports_) GetClass() ImportsClassLike {
	return v.class_
}

// Attribute

func (v *imports_) GetModules() abs.Sequential[ModuleLike] {
	return v.modules_
}

// Private
