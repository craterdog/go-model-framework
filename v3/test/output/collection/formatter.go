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

package collection

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
	defaultMaximum_ int
}

// Constants

func (c *formatterClass_) DefaultMaximum() int {
	return c.defaultMaximum_
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{}
}

func (c *formatterClass_) MakeWithMaximum(maximum int) FormatterLike {
	return &formatter_{
		maximum_: maximum,
	}
}

// Functions

// INSTANCE METHODS

// Target

type formatter_ struct {
	depth_ int
	maximum_ int
}

// Attributes

func (v *formatter_) GetDepth() int {
	return v.depth_
}

func (v *formatter_) GetMaximum() int {
	return v.maximum_
}

// Public

func (v *formatter_) FormatCollection(collection Collection) string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Private
