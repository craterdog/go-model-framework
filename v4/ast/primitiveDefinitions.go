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

// CLASS INTERFACE

// Access Function

func PrimitiveDefinitions() PrimitiveDefinitionsClassLike {
	return primitiveDefinitionsReference()
}

// Constructor Methods

func (c *primitiveDefinitionsClass_) Make(
	optionalTypeSection TypeSectionLike,
	optionalFunctionalSection FunctionalSectionLike,
) PrimitiveDefinitionsLike {
	var instance = &primitiveDefinitions_{
		// Initialize the instance attributes.
		optionalTypeSection_:       optionalTypeSection,
		optionalFunctionalSection_: optionalFunctionalSection,
	}
	return instance

}

// INSTANCE INTERFACE

// Attribute Methods

func (v *primitiveDefinitions_) GetOptionalTypeSection() TypeSectionLike {
	return v.optionalTypeSection_
}

func (v *primitiveDefinitions_) GetOptionalFunctionalSection() FunctionalSectionLike {
	return v.optionalFunctionalSection_
}

// Public Methods

func (v *primitiveDefinitions_) GetClass() PrimitiveDefinitionsClassLike {
	return v.getClass()
}

// Private Methods

func (v *primitiveDefinitions_) getClass() *primitiveDefinitionsClass_ {
	return primitiveDefinitionsReference()
}

// PRIVATE INTERFACE

// Instance Structure

type primitiveDefinitions_ struct {
	// Declare the instance attributes.
	optionalTypeSection_       TypeSectionLike
	optionalFunctionalSection_ FunctionalSectionLike
}

// Class Structure

type primitiveDefinitionsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func primitiveDefinitionsReference() *primitiveDefinitionsClass_ {
	return primitiveDefinitionsReference_
}

var primitiveDefinitionsReference_ = &primitiveDefinitionsClass_{
	// Initialize the class constants.
}
