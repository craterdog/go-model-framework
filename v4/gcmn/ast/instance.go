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
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// CLASS ACCESS

// Reference

var instanceClass = &instanceClass_{
	// Any private class constants should be initialized here.
}

// Function

func Instance() InstanceClassLike {
	return instanceClass
}

// CLASS METHODS

// Target

type instanceClass_ struct {
}

// Constants

// Constructors

func (c *instanceClass_) MakeWithAttributes(
	declaration DeclarationLike,
	attributes col.ListLike[AttributeLike],
	abstractions col.ListLike[AbstractionLike],
	methods col.ListLike[MethodLike],
) InstanceLike {
	return &instance_{
		declaration_:  declaration,
		attributes_:   attributes,
		abstractions_: abstractions,
		methods_:      methods,
	}
}

// Functions

// INSTANCE METHODS

// Target

type instance_ struct {
	class_        InstanceClassLike
	declaration_  DeclarationLike
	attributes_   col.ListLike[AttributeLike]
	abstractions_ col.ListLike[AbstractionLike]
	methods_      col.ListLike[MethodLike]
}

// Attributes

func (v *instance_) GetClass() InstanceClassLike {
	return v.class_
}

func (v *instance_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instance_) GetAttributes() col.ListLike[AttributeLike] {
	return v.attributes_
}

func (v *instance_) GetAbstractions() col.ListLike[AbstractionLike] {
	return v.abstractions_
}

func (v *instance_) GetMethods() col.ListLike[MethodLike] {
	return v.methods_
}

// Public

// Private
