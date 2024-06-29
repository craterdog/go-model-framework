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

func (c *modulesClass_) Make(
	moduleIterator age.IteratorLike[ModuleLike],
) ModulesLike {
	return &modules_{
		// Initialize instance attributes.
		class_:          c,
		moduleIterator_: moduleIterator,
	}
}

// INSTANCE METHODS

// Target

type modules_ struct {
	// Define instance attributes.
	class_          ModulesClassLike
	moduleIterator_ age.IteratorLike[ModuleLike]
}

// Attributes

func (v *modules_) GetClass() ModulesClassLike {
	return v.class_
}

func (v *modules_) GetModuleIterator() age.IteratorLike[ModuleLike] {
	return v.moduleIterator_
}

// Private
