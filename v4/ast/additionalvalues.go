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
	additionalValues col.ListLike[AdditionalValueLike],
) AdditionalValuesLike {
	return &additionalValues_{
		// Initialize instance attributes.
		class_:            c,
		additionalValues_: additionalValues,
	}
}

// INSTANCE METHODS

// Target

type additionalValues_ struct {
	// Define instance attributes.
	class_            AdditionalValuesClassLike
	additionalValues_ col.ListLike[AdditionalValueLike]
}

// Attributes

func (v *additionalValues_) GetClass() AdditionalValuesClassLike {
	return v.class_
}

func (v *additionalValues_) GetAdditionalValues() col.ListLike[AdditionalValueLike] {
	return v.additionalValues_
}

// Private
