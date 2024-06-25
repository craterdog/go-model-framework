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

var resultClass = &resultClass_{
	// Initialize class constants.
}

// Function

func Result() ResultClassLike {
	return resultClass
}

// CLASS METHODS

// Target

type resultClass_ struct {
	// Define class constants.
}

// Constructors

func (c *resultClass_) Make(any any) ResultLike {
	return &result_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type result_ struct {
	// Define instance attributes.
	class_ ResultClassLike
	any_   any
}

// Attributes

func (v *result_) GetClass() ResultClassLike {
	return v.class_
}

func (v *result_) GetAny() any {
	return v.any_
}

// Private
