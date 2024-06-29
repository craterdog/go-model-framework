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

var additionalValuesClass = &additionalValuesClass_{
	// Initialize class constants.
}

// Function

func AdditionalValues() AdditionalValuesClassLike {
	return additionalValuesClass
}

// CLASS METHODS

// Target

type additionalValuesClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalValuesClass_) Make(
	additionalValueIterator age.IteratorLike[AdditionalValueLike],
) AdditionalValuesLike {
	return &additionalValues_{
		// Initialize instance attributes.
		class_:                   c,
		additionalValueIterator_: additionalValueIterator,
	}
}

// INSTANCE METHODS

// Target

type additionalValues_ struct {
	// Define instance attributes.
	class_                   AdditionalValuesClassLike
	additionalValueIterator_ age.IteratorLike[AdditionalValueLike]
}

// Attributes

func (v *additionalValues_) GetClass() AdditionalValuesClassLike {
	return v.class_
}

func (v *additionalValues_) GetAdditionalValueIterator() age.IteratorLike[AdditionalValueLike] {
	return v.additionalValueIterator_
}

// Private
