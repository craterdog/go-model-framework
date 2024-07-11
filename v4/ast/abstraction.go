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
	mod "github.com/craterdog/go-collection-framework/v4"
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
	optionalAlias AliasLike,
	name string,
	optionalGenericArguments GenericArgumentsLike,
) AbstractionLike {
	// Validate the arguments.
	switch {
	case mod.IsUndefined(name):
		panic("The name attribute is required for each Abstraction.")
	default:
		return &abstraction_{
			// Initialize instance attributes.
			class_:                    c,
			optionalPrefix_:           optionalPrefix,
			optionalAlias_:            optionalAlias,
			name_:                     name,
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
	optionalAlias_            AliasLike
	name_                     string
	optionalGenericArguments_ GenericArgumentsLike
}

// Attributes

func (v *abstraction_) GetClass() AbstractionClassLike {
	return v.class_
}

func (v *abstraction_) GetOptionalPrefix() PrefixLike {
	return v.optionalPrefix_
}

func (v *abstraction_) GetOptionalAlias() AliasLike {
	return v.optionalAlias_
}

func (v *abstraction_) GetName() string {
	return v.name_
}

func (v *abstraction_) GetOptionalGenericArguments() GenericArgumentsLike {
	return v.optionalGenericArguments_
}

// Private
