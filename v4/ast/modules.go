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

var modulesClass = &modulesClass_{
	// Initialize class constants.
}

// Function

func Modules() ModulesClassLike {
	return modulesClass
}

// CLASS METHODS

// Target

type modulesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *modulesClass_) Make(modules col.Sequential[ModuleLike]) ModulesLike {
	return &modules_{
		// Initialize instance attributes.
		class_:   c,
		modules_: modules,
	}
}

// INSTANCE METHODS

// Target

type modules_ struct {
	// Define instance attributes.
	class_   ModulesClassLike
	modules_ col.Sequential[ModuleLike]
}

// Attributes

func (v *modules_) GetClass() ModulesClassLike {
	return v.class_
}

func (v *modules_) GetModules() col.Sequential[ModuleLike] {
	return v.modules_
}

// Private
