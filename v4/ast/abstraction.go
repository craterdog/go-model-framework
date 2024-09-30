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

var abstractionClass = &abstractionClass_{
	// Initialize class constants.
}

// Function

func Abstraction() AbstractionClassLike {
	return abstractionClass
}

// CLASS METHODS

// Target

type abstractionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *abstractionClass_) Make(
	optionalPrefix PrefixLike,
	name string,
	optionalSuffix SuffixLike,
	optionalGenericArguments GenericArgumentsLike,
) AbstractionLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(name):
		panic("The name attribute is required by this class.")
	default:
		return &abstraction_{
			// Initialize instance attributes.
			class_:                    c,
			optionalPrefix_:           optionalPrefix,
			name_:                     name,
			optionalSuffix_:           optionalSuffix,
			optionalGenericArguments_: optionalGenericArguments,
		}
	}
}

// INSTANCE METHODS

// Target

type abstraction_ struct {
	// Define instance attributes.
	class_                    AbstractionClassLike
	optionalPrefix_           PrefixLike
	name_                     string
	optionalSuffix_           SuffixLike
	optionalGenericArguments_ GenericArgumentsLike
}

// Public

func (v *abstraction_) GetClass() AbstractionClassLike {
	return v.class_
}

// Attribute

func (v *abstraction_) GetOptionalPrefix() PrefixLike {
	return v.optionalPrefix_
}

func (v *abstraction_) GetName() string {
	return v.name_
}

func (v *abstraction_) GetOptionalSuffix() SuffixLike {
	return v.optionalSuffix_
}

func (v *abstraction_) GetOptionalGenericArguments() GenericArgumentsLike {
	return v.optionalGenericArguments_
}

// Private
