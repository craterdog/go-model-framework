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

package gcmn

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var modelClass = &modelClass_{
	// Any private class constants should be initialized here.
}

// Function

func Model() ModelClassLike {
	return modelClass
}

// CLASS METHODS

// Target

type modelClass_ struct {
}

// Constants

// Constructors

func (c *modelClass_) MakeWithAttributes(
	notice NoticeLike,
	header HeaderLike,
	modules col.ListLike[ModuleLike],
	types col.ListLike[TypeLike],
	functionals col.ListLike[FunctionalLike],
	aspects col.ListLike[AspectLike],
	classes col.ListLike[ClassLike],
	instances col.ListLike[InstanceLike],
) ModelLike {
	return &model_{
		notice_:      notice,
		header_:      header,
		modules_:     modules,
		types_:       types,
		functionals_: functionals,
		aspects_:     aspects,
		classes_:     classes,
		instances_:   instances,
	}
}

// Functions

// INSTANCE METHODS

// Target

type model_ struct {
	class_       ModelClassLike
	notice_      NoticeLike
	header_      HeaderLike
	modules_     col.ListLike[ModuleLike]
	types_       col.ListLike[TypeLike]
	functionals_ col.ListLike[FunctionalLike]
	aspects_     col.ListLike[AspectLike]
	classes_     col.ListLike[ClassLike]
	instances_   col.ListLike[InstanceLike]
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

func (v *model_) GetModules() col.ListLike[ModuleLike] {
	return v.modules_
}

func (v *model_) GetTypes() col.ListLike[TypeLike] {
	return v.types_
}

func (v *model_) GetFunctionals() col.ListLike[FunctionalLike] {
	return v.functionals_
}

func (v *model_) GetAspects() col.ListLike[AspectLike] {
	return v.aspects_
}

func (v *model_) GetClasses() col.ListLike[ClassLike] {
	return v.classes_
}

func (v *model_) GetInstances() col.ListLike[InstanceLike] {
	return v.instances_
}

// Public

// Private
