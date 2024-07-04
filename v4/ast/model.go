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
	imports ImportsLike,
	types TypesLike,
	functionals FunctionalsLike,
	classes ClassesLike,
	instances InstancesLike,
	aspects AspectsLike,
) ModelLike {
	return &model_{
		// Initialize instance attributes.
		class_:       c,
		notice_:      notice,
		header_:      header,
		imports_:     imports,
		types_:       types,
		functionals_: functionals,
		classes_:     classes,
		instances_:   instances,
		aspects_:     aspects,
	}
}

// INSTANCE METHODS

// Target

type model_ struct {
	// Define instance attributes.
	class_       ModelClassLike
	notice_      NoticeLike
	header_      HeaderLike
	imports_     ImportsLike
	types_       TypesLike
	functionals_ FunctionalsLike
	classes_     ClassesLike
	instances_   InstancesLike
	aspects_     AspectsLike
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

func (v *model_) GetImports() ImportsLike {
	return v.imports_
}

func (v *model_) GetTypes() TypesLike {
	return v.types_
}

func (v *model_) GetFunctionals() FunctionalsLike {
	return v.functionals_
}

func (v *model_) GetClasses() ClassesLike {
	return v.classes_
}

func (v *model_) GetInstances() InstancesLike {
	return v.instances_
}

func (v *model_) GetAspects() AspectsLike {
	return v.aspects_
}

// Private
