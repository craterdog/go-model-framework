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

var modelClass = &modelClass_{
	// Initialize class constants.
}

// Function

func Model() ModelClassLike {
	return modelClass
}

// CLASS METHODS

// Target

type modelClass_ struct {
	// Define class constants.
}

// Constructors

func (c *modelClass_) Make(
	notice NoticeLike,
	header HeaderLike,
	optionalImports ImportsLike,
	optionalTypes TypesLike,
	optionalFunctionals FunctionalsLike,
	classes ClassesLike,
	instances InstancesLike,
	optionalAspects AspectsLike,
) ModelLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(notice):
		panic("The notice attribute is required by this class.")
	case col.IsUndefined(header):
		panic("The header attribute is required by this class.")
	case col.IsUndefined(classes):
		panic("The classes attribute is required by this class.")
	case col.IsUndefined(instances):
		panic("The instances attribute is required by this class.")
	default:
		return &model_{
			// Initialize instance attributes.
			class_:               c,
			notice_:              notice,
			header_:              header,
			optionalImports_:     optionalImports,
			optionalTypes_:       optionalTypes,
			optionalFunctionals_: optionalFunctionals,
			classes_:             classes,
			instances_:           instances,
			optionalAspects_:     optionalAspects,
		}
	}
}

// INSTANCE METHODS

// Target

type model_ struct {
	// Define instance attributes.
	class_               ModelClassLike
	notice_              NoticeLike
	header_              HeaderLike
	optionalImports_     ImportsLike
	optionalTypes_       TypesLike
	optionalFunctionals_ FunctionalsLike
	classes_             ClassesLike
	instances_           InstancesLike
	optionalAspects_     AspectsLike
}

// Attributes

func (v *model_) GetClass() ModelClassLike {
	return v.class_
}

func (v *model_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *model_) GetHeader() HeaderLike {
	return v.header_
}

func (v *model_) GetOptionalImports() ImportsLike {
	return v.optionalImports_
}

func (v *model_) GetOptionalTypes() TypesLike {
	return v.optionalTypes_
}

func (v *model_) GetOptionalFunctionals() FunctionalsLike {
	return v.optionalFunctionals_
}

func (v *model_) GetClasses() ClassesLike {
	return v.classes_
}

func (v *model_) GetInstances() InstancesLike {
	return v.instances_
}

func (v *model_) GetOptionalAspects() AspectsLike {
	return v.optionalAspects_
}

// Private
