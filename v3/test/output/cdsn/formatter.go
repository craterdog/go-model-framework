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

package cdsn

import ()

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Any private class constants should be initialized here.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{}
}

// Functions

// INSTANCE METHODS

// Target

type formatter_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

func (v *formatter_) FormatDefinition(definition DefinitionLike) string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

func (v *formatter_) FormatSyntax(syntax SyntaxLike) string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Private
