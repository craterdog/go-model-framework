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

import ()

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
	modules ModulesLike,
	types TypesLike,
	functionals FunctionalsLike,
	aspects AspectsLike,
	classes ClassesLike,
	instances InstancesLike,
) ModelLike {
	return &model_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type model_ struct {
	// Define instance attributes.
	class_       ModelClassLike
	notice_      NoticeLike
	header_      HeaderLike
	modules_     ModulesLike
	types_       TypesLike
	functionals_ FunctionalsLike
	aspects_     AspectsLike
	classes_     ClassesLike
	instances_   InstancesLike
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

func (v *model_) GetModules() ModulesLike {
	return v.modules_
}

func (v *model_) GetTypes() TypesLike {
	return v.types_
}

func (v *model_) GetFunctionals() FunctionalsLike {
	return v.functionals_
}

func (v *model_) GetAspects() AspectsLike {
	return v.aspects_
}

func (v *model_) GetClasses() ClassesLike {
	return v.classes_
}

func (v *model_) GetInstances() InstancesLike {
	return v.instances_
}

// Private
