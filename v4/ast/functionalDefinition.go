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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func FunctionalDefinition() FunctionalDefinitionClassLike {
	return functionalDefinitionReference()
}

// Constructor Methods

func (c *functionalDefinitionClass_) Make(
	declaration DeclarationLike,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionalDefinitionLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	if uti.IsUndefined(result) {
		panic("The \"result\" attribute is required by this class.")
	}
	var instance = &functionalDefinition_{
		// Initialize the instance attributes.
		declaration_: declaration,
		parameters_:  parameters,
		result_:      result,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *functionalDefinition_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *functionalDefinition_) GetParameters() abs.Sequential[ParameterLike] {
	return v.parameters_
}

func (v *functionalDefinition_) GetResult() ResultLike {
	return v.result_
}

// Public Methods

func (v *functionalDefinition_) GetClass() FunctionalDefinitionClassLike {
	return v.getClass()
}

// Private Methods

func (v *functionalDefinition_) getClass() *functionalDefinitionClass_ {
	return functionalDefinitionReference()
}

// PRIVATE INTERFACE

// Instance Structure

type functionalDefinition_ struct {
	// Declare the instance attributes.
	declaration_ DeclarationLike
	parameters_  abs.Sequential[ParameterLike]
	result_      ResultLike
}

// Class Structure

type functionalDefinitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionalDefinitionReference() *functionalDefinitionClass_ {
	return functionalDefinitionReference_
}

var functionalDefinitionReference_ = &functionalDefinitionClass_{
	// Initialize the class constants.
}
