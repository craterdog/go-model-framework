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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PublicSubsection() PublicSubsectionClassLike {
	return publicSubsectionReference()
}

// Constructor Methods

func (c *publicSubsectionClass_) Make(
	publicMethods abs.Sequential[PublicMethodLike],
) PublicSubsectionLike {
	if uti.IsUndefined(publicMethods) {
		panic("The \"publicMethods\" attribute is required by this class.")
	}
	var instance = &publicSubsection_{
		// Initialize the instance attributes.
		publicMethods_: publicMethods,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *publicSubsection_) GetPublicMethods() abs.Sequential[PublicMethodLike] {
	return v.publicMethods_
}

// Public Methods

func (v *publicSubsection_) GetClass() PublicSubsectionClassLike {
	return v.getClass()
}

// Private Methods

func (v *publicSubsection_) getClass() *publicSubsectionClass_ {
	return publicSubsectionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type publicSubsection_ struct {
	// Declare the instance attributes.
	publicMethods_ abs.Sequential[PublicMethodLike]
}

// Class Structure

type publicSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func publicSubsectionReference() *publicSubsectionClass_ {
	return publicSubsectionReference_
}

var publicSubsectionReference_ = &publicSubsectionClass_{
	// Initialize the class constants.
}
