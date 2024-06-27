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

var instanceClass = &instanceClass_{
	// Initialize class constants.
}

// Function

func Instance() InstanceClassLike {
	return instanceClass
}

// CLASS METHODS

// Target

type instanceClass_ struct {
	// Define class constants.
}

// Constructors

func (c *instanceClass_) Make(
	declaration DeclarationLike,
	attributes AttributesLike,
	abstractions AbstractionsLike,
	methods MethodsLike,
) InstanceLike {
	return &instance_{
		// Initialize instance attributes.
		class_:        c,
		declaration_:  declaration,
		attributes_:   attributes,
		abstractions_: abstractions,
		methods_:      methods,
	}
}

// INSTANCE METHODS

// Target

type instance_ struct {
	// Define instance attributes.
	class_        InstanceClassLike
	declaration_  DeclarationLike
	attributes_   AttributesLike
	abstractions_ AbstractionsLike
	methods_      MethodsLike
}

// Attributes

func (v *instance_) GetClass() InstanceClassLike {
	return v.class_
}

func (v *instance_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *instance_) GetAttributes() AttributesLike {
	return v.attributes_
}

func (v *instance_) GetAbstractions() AbstractionsLike {
	return v.abstractions_
}

func (v *instance_) GetMethods() MethodsLike {
	return v.methods_
}

// Private
