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

var scannerClass = &scannerClass_{
	// Any private class constants should be initialized here.
}

// Function

func Scanner() ScannerClassLike {
	return scannerClass
}

// CLASS METHODS

// Target

type scannerClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *scannerClass_) Make(
	source string,
	tokens col.QueueLike[TokenLike],
) ScannerLike {
	return &scanner_{}
}

// Functions

func (c *scannerClass_) FormatToken(token TokenLike) string {
	var result_ string
	// TBA - Implement the function.
	return result_
}

func (c *scannerClass_) MatchToken(
	type_ TokenType,
	text string,
) col.ListLike[string] {
	var result_ col.ListLike[string]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type scanner_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

// Private
