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

import ()

// CLASS ACCESS

// Reference

var prefixClass = &prefixClass_{
	// Any private class constants should be initialized here.
}

// Function

func Prefix() PrefixClassLike {
	return prefixClass
}

// CLASS METHODS

// Target

type prefixClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *prefixClass_) MakeWithAttributes(
	identifier string,
	type_ PrefixType,
) PrefixLike {
	return &prefix_{
		identifier_: identifier,
		type_: type_,
	}
}

// Functions

// INSTANCE METHODS

// Target

type prefix_ struct {
	type_ PrefixType
	identifier_ string
}

// Attributes

func (v *prefix_) GetType() PrefixType {
	return v.type_
}

func (v *prefix_) GetIdentifier() string {
	return v.identifier_
}

// Public

// Private
