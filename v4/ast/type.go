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

var typeClass = &typeClass_{
	// Initialize class constants.
}

// Function

func Type() TypeClassLike {
	return typeClass
}

// CLASS METHODS

// Target

type typeClass_ struct {
	// Define class constants.
}

// Constructors

func (c *typeClass_) MakeWithAttributes(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	enumeration EnumerationLike,
) TypeLike {
	return &type_{
		// Initialize instance attributes.
		class_:       c,
		declaration_: declaration,
		abstraction_: abstraction,
		enumeration_: enumeration,
	}
}

// INSTANCE METHODS

// Target

type type_ struct {
	// Define instance attributes.
	class_       TypeClassLike
	declaration_ DeclarationLike
	abstraction_ AbstractionLike
	enumeration_ EnumerationLike
}

// Attributes

func (v *type_) GetClass() TypeClassLike {
	return v.class_
}

func (v *type_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *type_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

func (v *type_) GetEnumeration() EnumerationLike {
	return v.enumeration_
}

// Private
