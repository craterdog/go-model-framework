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

var enumerationClass = &enumerationClass_{
	// Initialize class constants.
}

// Function

func Enumeration() EnumerationClassLike {
	return enumerationClass
}

// CLASS METHODS

// Target

type enumerationClass_ struct {
	// Define class constants.
}

// Constructors

func (c *enumerationClass_) Make(
	value ValueLike,
	additionalValues col.ListLike[AdditionalValueLike],
) EnumerationLike {
	return &enumeration_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type enumeration_ struct {
	// Define instance attributes.
	class_            EnumerationClassLike
	value_            ValueLike
	additionalValues_ col.ListLike[AdditionalValueLike]
}

// Attributes

func (v *enumeration_) GetClass() EnumerationClassLike {
	return v.class_
}

func (v *enumeration_) GetValue() ValueLike {
	return v.value_
}

func (v *enumeration_) GetAdditionalValues() col.ListLike[AdditionalValueLike] {
	return v.additionalValues_
}

// Private
