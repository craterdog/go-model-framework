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

var additionalArgumentClass = &additionalArgumentClass_{
	// Initialize class constants.
}

// Function

func AdditionalArgument() AdditionalArgumentClassLike {
	return additionalArgumentClass
}

// CLASS METHODS

// Target

type additionalArgumentClass_ struct {
	// Define class constants.
}

// Constructors

func (c *additionalArgumentClass_) Make(argument ArgumentLike) AdditionalArgumentLike {
	return &additionalArgument_{
		// Initialize instance attributes.
		class_:    c,
		argument_: argument,
	}
}

// INSTANCE METHODS

// Target

type additionalArgument_ struct {
	// Define instance attributes.
	class_    AdditionalArgumentClassLike
	argument_ ArgumentLike
}

// Attributes

func (v *additionalArgument_) GetClass() AdditionalArgumentClassLike {
	return v.class_
}

func (v *additionalArgument_) GetArgument() ArgumentLike {
	return v.argument_
}

// Private
