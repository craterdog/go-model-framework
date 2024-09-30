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

var classClass = &classClass_{
	// Initialize class constants.
}

// Function

func Class() ClassClassLike {
	return classClass
}

// CLASS METHODS

// Target

type classClass_ struct {
	// Define class constants.
}

// Constructors

func (c *classClass_) Make(
	declaration DeclarationLike,
	classMethods ClassMethodsLike,
) ClassLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(declaration):
		panic("The declaration attribute is required by this class.")
	case col.IsUndefined(classMethods):
		panic("The classMethods attribute is required by this class.")
	default:
		return &class_{
			// Initialize instance attributes.
			class_:        c,
			declaration_:  declaration,
			classMethods_: classMethods,
		}
	}
}

// INSTANCE METHODS

// Target

type class_ struct {
	// Define instance attributes.
	class_        ClassClassLike
	declaration_  DeclarationLike
	classMethods_ ClassMethodsLike
}

// Public

func (v *class_) GetClass() ClassClassLike {
	return v.class_
}

// Attribute

func (v *class_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *class_) GetClassMethods() ClassMethodsLike {
	return v.classMethods_
}

// Private
