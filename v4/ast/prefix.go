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

import ()

// CLASS ACCESS

// Reference

var prefixClass = &prefixClass_{
	// Initialize class constants.
}

// Function

func Prefix() PrefixClassLike {
	return prefixClass
}

// CLASS METHODS

// Target

type prefixClass_ struct {
	// Define class constants.
}

// Constructors

func (c *prefixClass_) MakeWithAttributes(
	identifier string,
	type_ PrefixType,
) PrefixLike {
	return &prefix_{
		// Initialize instance attributes.
		class_:      c,
		identifier_: identifier,
		type_:       type_,
	}
}

// INSTANCE METHODS

// Target

type prefix_ struct {
	// Define instance attributes.
	class_      PrefixClassLike
	type_       PrefixType
	identifier_ string
}

// Attributes

func (v *prefix_) GetClass() PrefixClassLike {
	return v.class_
}

func (v *prefix_) GetType() PrefixType {
	return v.type_
}

func (v *prefix_) GetIdentifier() string {
	return v.identifier_
}

// Private
