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

package gcmn

import (
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var abstractionClass = &abstractionClass_{
	// Any private class constants should be initialized here.
}

// Function

func Abstraction() AbstractionClassLike {
	return abstractionClass
}

// CLASS METHODS

// Target

type abstractionClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *abstractionClass_) MakeWithAttributes(
	prefix PrefixLike,
	identifier string,
	arguments col.ListLike[AbstractionLike],
) AbstractionLike {
	return &abstraction_{
		prefix_: prefix,
		identifier_: identifier,
		arguments_: arguments,
	}
}

// Functions

// INSTANCE METHODS

// Target

type abstraction_ struct {
	class_ AbstractionClassLike
	prefix_ PrefixLike
	identifier_ string
	arguments_ col.ListLike[AbstractionLike]
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

// Public

// Private
