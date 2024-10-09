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

func Header() HeaderClassLike {
	return headerClass
}

// Constructor Methods

func (c *headerClass_) Make(
	comment string,
	name string,
) HeaderLike {
	if uti.IsUndefined(comment) {
		panic("The comment attribute is required by this class.")
	}
	if uti.IsUndefined(name) {
		panic("The name attribute is required by this class.")
	}
	var instance = &header_{
		class_:   c,
		comment_: comment,
		name_:    name,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *header_) GetComment() string {
	return v.comment_
}

func (v *header_) GetName() string {
	return v.name_
}

// Public Methods

func (v *header_) GetClass() HeaderClassLike {
	return v.getClass()
}

// Private Methods

func (v *header_) getClass() *headerClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type header_ struct {
	class_   *headerClass_
	comment_ string
	name_    string
}

// Class Structure

type headerClass_ struct {
	// Define the class constants.
}

// Class Reference

var headerClass = &headerClass_{
	// Initialize the class constants.
}
