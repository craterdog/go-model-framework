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
	prefix PrefixLike,
	alias AliasLike,
	name string,
	genericArguments GenericArgumentsLike,
) AbstractionLike {
	switch {
	case len(name) == 0:
		panic("The name attribute is required for each abstraction.")
	default:
		return &abstraction_{
			// Initialize instance attributes.
			class_:            c,
			prefix_:           prefix,
			alias_:            alias,
			name_:             name,
			genericArguments_: genericArguments,
		}
	}
}

// INSTANCE METHODS

// Target

type abstraction_ struct {
	// Define instance attributes.
	class_            AbstractionClassLike
	prefix_           PrefixLike
	alias_            AliasLike
	name_             string
	genericArguments_ GenericArgumentsLike
}

// Attributes

func (v *abstraction_) GetClass() AbstractionClassLike {
	return v.class_
}

func (v *abstraction_) GetOptionalPrefix() PrefixLike {
	return v.prefix_
}

func (v *abstraction_) GetOptionalAlias() AliasLike {
	return v.alias_
}

func (v *abstraction_) GetName() string {
	return v.name_
}

func (v *abstraction_) GetOptionalGenericArguments() GenericArgumentsLike {
	return v.genericArguments_
}

// Private
