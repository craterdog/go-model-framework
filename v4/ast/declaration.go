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

func Declaration() DeclarationClassLike {
	return declarationReference()
}

// Constructor Methods

func (c *declarationClass_) Make(
	comment string,
	name string,
	optionalConstraints ConstraintsLike,
) DeclarationLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &declaration_{
		// Initialize the instance attributes.
		comment_:             comment,
		name_:                name,
		optionalConstraints_: optionalConstraints,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *declaration_) GetComment() string {
	return v.comment_
}

func (v *declaration_) GetName() string {
	return v.name_
}

func (v *declaration_) GetOptionalConstraints() ConstraintsLike {
	return v.optionalConstraints_
}

// Public Methods

func (v *declaration_) GetClass() DeclarationClassLike {
	return v.getClass()
}

// Private Methods

func (v *declaration_) getClass() *declarationClass_ {
	return declarationReference()
}

// PRIVATE INTERFACE

// Instance Structure

type declaration_ struct {
	// Declare the instance attributes.
	comment_             string
	name_                string
	optionalConstraints_ ConstraintsLike
}

// Class Structure

type declarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func declarationReference() *declarationClass_ {
	return declarationReference_
}

var declarationReference_ = &declarationClass_{
	// Initialize the class constants.
}
