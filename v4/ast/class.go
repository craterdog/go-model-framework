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

var classClass = &classClass_{
	// Initialize class constants.
}

// Function

func Class() ClassClassLike {
	return classClass
}

// CLASS METHODS

// Target

type classClass_ struct {
	// Define class constants.
}

// Constructors

func (c *classClass_) Make(
	declaration DeclarationLike,
	constructors ConstructorsLike,
	optionalConstants ConstantsLike,
	optionalFunctions FunctionsLike,
) ClassLike {
	// Validate the arguments.
	switch {
	case col.IsUndefined(declaration):
		panic("The declaration attribute is required for each Class.")
	case col.IsUndefined(constructors):
		panic("The constructors attribute is required for each Class.")
	default:
		return &class_{
			// Initialize instance attributes.
			class_:             c,
			declaration_:       declaration,
			constructors_:      constructors,
			optionalConstants_: optionalConstants,
			optionalFunctions_: optionalFunctions,
		}
	}
}

// INSTANCE METHODS

// Target

type class_ struct {
	// Define instance attributes.
	class_             ClassClassLike
	declaration_       DeclarationLike
	constructors_      ConstructorsLike
	optionalConstants_ ConstantsLike
	optionalFunctions_ FunctionsLike
}

// Attributes

func (v *class_) GetClass() ClassClassLike {
	return v.class_
}

func (v *class_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *class_) GetConstructors() ConstructorsLike {
	return v.constructors_
}

func (v *class_) GetOptionalConstants() ConstantsLike {
	return v.optionalConstants_
}

func (v *class_) GetOptionalFunctions() FunctionsLike {
	return v.optionalFunctions_
}

// Private
