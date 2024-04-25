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

package agent

import (
	gcm "github.com/craterdog/go-model-framework/v4/gcmn"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
	// Any private class constants should be initialized here.
}

// Function

func Parser() ParserClassLike {
	return parserClass
}

// CLASS METHODS

// Target

type parserClass_ struct {
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
	class_ ParserClassLike
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) ParseSource(source string) gcm.ModelLike {
	var result_ gcm.ModelLike
	// TBA - Implement the method.
	return result_
}

// Private
