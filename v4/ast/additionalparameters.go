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

var additionalParametersClass = &additionalParametersClass_{
	// Initialize class constants.
}

// Function

func AdditionalParameters() AdditionalParametersClassLike {
	return additionalParametersClass
}

// CLASS METHODS

// Target

type additionalParametersClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalParametersClass_) Make(
	additionalParameterIterator age.IteratorLike[AdditionalParameterLike],
) AdditionalParametersLike {
	return &additionalParameters_{
		// Initialize instance attributes.
		class_:                       c,
		additionalParameterIterator_: additionalParameterIterator,
	}
}

// INSTANCE METHODS

// Target

type additionalParameters_ struct {
	// Define instance attributes.
	class_                       AdditionalParametersClassLike
	additionalParameterIterator_ age.IteratorLike[AdditionalParameterLike]
}

// Attributes

func (v *additionalParameters_) GetClass() AdditionalParametersClassLike {
	return v.class_
}

func (v *additionalParameters_) GetAdditionalParameterIterator() age.IteratorLike[AdditionalParameterLike] {
	return v.additionalParameterIterator_
}

// Private
