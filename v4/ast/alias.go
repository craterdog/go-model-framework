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

var aliasClass = &aliasClass_{
	// Initialize class constants.
}

// Function

func Alias() AliasClassLike {
	return aliasClass
}

// CLASS METHODS

// Target

type aliasClass_ struct {
	// Define class constants.
}

// Constructors

func (c *aliasClass_) Make(identifier string) AliasLike {
	return &alias_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type alias_ struct {
	// Define instance attributes.
	class_      AliasClassLike
	identifier_ string
}

// Attributes

func (v *alias_) GetClass() AliasClassLike {
	return v.class_
}

func (v *alias_) GetIdentifier() string {
	return v.identifier_
}

// Private
