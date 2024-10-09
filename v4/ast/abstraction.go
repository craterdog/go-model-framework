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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Abstraction() AbstractionClassLike {
	return abstractionClass
}

// Constructor Methods

func (c *abstractionClass_) Make(
	optionalPrefix PrefixLike,
	name string,
	optionalSuffix SuffixLike,
	optionalArguments ArgumentsLike,
) AbstractionLike {
	if uti.IsUndefined(name) {
		panic("The name attribute is required by this class.")
	}
	var instance = &abstraction_{
		class_:             c,
		optionalPrefix_:    optionalPrefix,
		name_:              name,
		optionalSuffix_:    optionalSuffix,
		optionalArguments_: optionalArguments,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *abstraction_) GetOptionalPrefix() PrefixLike {
	return v.optionalPrefix_
}

func (v *abstraction_) GetName() string {
	return v.name_
}

func (v *abstraction_) GetOptionalSuffix() SuffixLike {
	return v.optionalSuffix_
}

func (v *abstraction_) GetOptionalArguments() ArgumentsLike {
	return v.optionalArguments_
}

// Public Methods

func (v *abstraction_) GetClass() AbstractionClassLike {
	return v.getClass()
}

// Private Methods

func (v *abstraction_) getClass() *abstractionClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type abstraction_ struct {
	class_             *abstractionClass_
	optionalPrefix_    PrefixLike
	name_              string
	optionalSuffix_    SuffixLike
	optionalArguments_ ArgumentsLike
}

// Class Structure

type abstractionClass_ struct {
	// Define the class constants.
}

// Class Reference

var abstractionClass = &abstractionClass_{
	// Initialize the class constants.
}
