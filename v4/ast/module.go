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

import ()

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
	identifier string,
	text string,
) ModuleLike {
	return &module_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type module_ struct {
	// Define instance attributes.
	class_      ModuleClassLike
	identifier_ string
	text_       string
}

// Attributes

func (v *module_) GetClass() ModuleClassLike {
	return v.class_
}

func (v *module_) GetIdentifier() string {
	return v.identifier_
}

func (v *module_) GetText() string {
	return v.text_
}

// Private
