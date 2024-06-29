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

var additionalArgumentsClass = &additionalArgumentsClass_{
	// Initialize class constants.
}

// Function

func AdditionalArguments() AdditionalArgumentsClassLike {
	return additionalArgumentsClass
}

// CLASS METHODS

// Target

type additionalArgumentsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalArgumentsClass_) Make(additionalArgumentIterator age.IteratorLike[AdditionalArgumentLike]) AdditionalArgumentsLike {
	return &additionalArguments_{
		// Initialize instance attributes.
		class_:                      c,
		additionalArgumentIterator_: additionalArgumentIterator,
	}
}

// INSTANCE METHODS

// Target

type additionalArguments_ struct {
	// Define instance attributes.
	class_                      AdditionalArgumentsClassLike
	additionalArgumentIterator_ age.IteratorLike[AdditionalArgumentLike]
}

// Attributes

func (v *additionalArguments_) GetClass() AdditionalArgumentsClassLike {
	return v.class_
}

func (v *additionalArguments_) GetAdditionalArgumentIterator() age.IteratorLike[AdditionalArgumentLike] {
	return v.additionalArgumentIterator_
}

// Private
