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

func (c *abstractionClass_) MakeWithAttributes(
	prefix PrefixLike,
	identifier string,
	arguments col.ListLike[AbstractionLike],
) AbstractionLike {
	return &abstraction_{
		// Initialize instance attributes.
		class_:      c,
		prefix_:     prefix,
		identifier_: identifier,
		arguments_:  arguments,
	}
}

// INSTANCE METHODS

// Target

type abstraction_ struct {
	// Define instance attributes.
	class_      AbstractionClassLike
	prefix_     PrefixLike
	identifier_ string
	arguments_  col.ListLike[AbstractionLike]
}

// Attributes

func (v *abstraction_) GetClass() AbstractionClassLike {
	return v.class_
}

func (v *abstraction_) GetPrefix() PrefixLike {
	return v.prefix_
}

func (v *abstraction_) GetIdentifier() string {
	return v.identifier_
}

func (v *abstraction_) GetArguments() col.ListLike[AbstractionLike] {
	return v.arguments_
}

// Private
