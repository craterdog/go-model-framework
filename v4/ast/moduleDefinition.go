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

func ModuleDefinition() ModuleDefinitionClassLike {
	return moduleDefinitionClass
}

// Constructor Methods

func (c *moduleDefinitionClass_) Make(
	notice NoticeLike,
	header HeaderLike,
	optionalImports ImportsLike,
) ModuleDefinitionLike {
	if uti.IsUndefined(notice) {
		panic("The \"notice\" attribute is required by this class.")
	}
	if uti.IsUndefined(header) {
		panic("The \"header\" attribute is required by this class.")
	}
	var instance = &moduleDefinition_{
		// Initialize the instance attributes.
		class_:           c,
		notice_:          notice,
		header_:          header,
		optionalImports_: optionalImports,
	}
	return instance
}

// INSTANCE INTERFACE

// Attribute Methods

func (v *moduleDefinition_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *moduleDefinition_) GetHeader() HeaderLike {
	return v.header_
}

func (v *moduleDefinition_) GetOptionalImports() ImportsLike {
	return v.optionalImports_
}

// Public Methods

func (v *moduleDefinition_) GetClass() ModuleDefinitionClassLike {
	return v.getClass()
}

// Private Methods

func (v *moduleDefinition_) getClass() *moduleDefinitionClass_ {
	return v.class_
}

// PRIVATE INTERFACE

// Instance Structure

type moduleDefinition_ struct {
	// Declare the instance attributes.
	class_           *moduleDefinitionClass_
	notice_          NoticeLike
	header_          HeaderLike
	optionalImports_ ImportsLike
}

// Class Structure

type moduleDefinitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

var moduleDefinitionClass = &moduleDefinitionClass_{
	// Initialize the class constants.
}
