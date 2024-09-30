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

var primitiveDefinitionsClass = &primitiveDefinitionsClass_{
	// Initialize class constants.
}

// Function

func PrimitiveDefinitions() PrimitiveDefinitionsClassLike {
	return primitiveDefinitionsClass
}

// CLASS METHODS

// Target

type primitiveDefinitionsClass_ struct {
	// Define class constants.
}

// Constructors

func (c *primitiveDefinitionsClass_) Make(
	optionalTypeDefinitions TypeDefinitionsLike,
	optionalFunctionalDefinitions FunctionalDefinitionsLike,
) PrimitiveDefinitionsLike {
	// Validate the arguments.
	switch {
	default:
		return &primitiveDefinitions_{
			// Initialize instance attributes.
			class_:                         c,
			optionalTypeDefinitions_:       optionalTypeDefinitions,
			optionalFunctionalDefinitions_: optionalFunctionalDefinitions,
		}
	}
}

// INSTANCE METHODS

// Target

type primitiveDefinitions_ struct {
	// Define instance attributes.
	class_                         PrimitiveDefinitionsClassLike
	optionalTypeDefinitions_       TypeDefinitionsLike
	optionalFunctionalDefinitions_ FunctionalDefinitionsLike
}

// Public

func (v *primitiveDefinitions_) GetClass() PrimitiveDefinitionsClassLike {
	return v.class_
}

// Attribute

func (v *primitiveDefinitions_) GetOptionalTypeDefinitions() TypeDefinitionsLike {
	return v.optionalTypeDefinitions_
}

func (v *primitiveDefinitions_) GetOptionalFunctionalDefinitions() FunctionalDefinitionsLike {
	return v.optionalFunctionalDefinitions_
}

// Private
