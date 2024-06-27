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
	additionalParameters col.ListLike[AdditionalParameterLike],
) AdditionalParametersLike {
	return &additionalParameters_{
		// Initialize instance attributes.
		class_:                c,
		additionalParameters_: additionalParameters,
	}
}

// INSTANCE METHODS

// Target

type additionalParameters_ struct {
	// Define instance attributes.
	class_                AdditionalParametersClassLike
	additionalParameters_ col.ListLike[AdditionalParameterLike]
}

// Attributes

func (v *additionalParameters_) GetClass() AdditionalParametersClassLike {
	return v.class_
}

func (v *additionalParameters_) GetAdditionalParameters() col.ListLike[AdditionalParameterLike] {
	return v.additionalParameters_
}

// Private
