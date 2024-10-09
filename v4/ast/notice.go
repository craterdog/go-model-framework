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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Notice() NoticeClassLike {
	return noticeClass
}

// Constructor Methods

func (c *noticeClass_) Make(
	comment string,
) NoticeLike {
	if uti.IsUndefined(comment) {
		panic("The comment attribute is required by this class.")
	}
	var instance = &notice_{
		class_:   c,
		comment_: comment,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *notice_) GetComment() string {
	return v.comment_
}

// Public Methods

func (v *notice_) GetClass() NoticeClassLike {
	return v.getClass()
}

// Private Methods

func (v *notice_) getClass() *noticeClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type notice_ struct {
	class_   *noticeClass_
	comment_ string
}

// Class Structure

type noticeClass_ struct {
	// Define the class constants.
}

// Class Reference

var noticeClass = &noticeClass_{
	// Initialize the class constants.
}
