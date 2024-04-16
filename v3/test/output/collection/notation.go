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

package collection

import ()

// CLASS ACCESS

// Reference

var notationClass = &notationClass_{
	// This class has no private constants to initialize.
}

// Function

func Notation() NotationClassLike {
	return notationClass
}

// CLASS METHODS

// Target

type notationClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *notationClass_) Make() NotationLike {
	return &notation_{}
}

// Functions

// INSTANCE METHODS

// Target

type notation_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Canonical

func (v *notation_) FormatCollection(collection Collection) string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

func (v *notation_) ParseSource(source string) Collection {
	var result_ Collection
	// TBA - Implement the method.
	return result_
}

// Public

// Private
