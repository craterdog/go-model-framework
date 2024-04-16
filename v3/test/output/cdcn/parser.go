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

package cdcn

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
	// This class has no private constants to initialize.
}

// Function

func Parser() ParserClassLike {
	return parserClass
}

// CLASS METHODS

// Target

type parserClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{}
}

// Functions

// INSTANCE METHODS

// Target

type parser_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

func (v *parser_) ParseSource(source string) col.Collection {
	var result_ col.Collection
	// TBA - Implement the method.
	return result_
}

// Private
