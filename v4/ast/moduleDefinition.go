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
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var moduleDefinitionClass = &moduleDefinitionClass_{
	// Initialize class constants.
}

// Function

func ModuleDefinition() ModuleDefinitionClassLike {
	return moduleDefinitionClass
}

// CLASS METHODS

// Target

type moduleDefinitionClass_ struct {
	// Define class constants.
}

// Constructors

func (c *moduleDefinitionClass_) Make(
	notice NoticeLike,
	header HeaderLike,
	optionalImports ImportsLike,
) ModuleDefinitionLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(notice):
		panic("The notice attribute is required by this class.")
	case col.IsUndefined(header):
		panic("The header attribute is required by this class.")
	default:
		return &moduleDefinition_{
			// Initialize instance attributes.
			class_:           c,
			notice_:          notice,
			header_:          header,
			optionalImports_: optionalImports,
		}
	}
}

// INSTANCE METHODS

// Target

type moduleDefinition_ struct {
	// Define instance attributes.
	class_           ModuleDefinitionClassLike
	notice_          NoticeLike
	header_          HeaderLike
	optionalImports_ ImportsLike
}

// Public

func (v *moduleDefinition_) GetClass() ModuleDefinitionClassLike {
	return v.class_
}

// Attribute

func (v *moduleDefinition_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *moduleDefinition_) GetHeader() HeaderLike {
	return v.header_
}

func (v *moduleDefinition_) GetOptionalImports() ImportsLike {
	return v.optionalImports_
}

// Private
