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
	constants ConstantsLike,
	functions FunctionsLike,
) ClassLike {
	switch {
	case declaration == nil:
		panic("The declaration attribute is required for each class.")
	case constructors == nil:
		panic("At least one constructor is required for each class.")
	default:
		return &class_{
			// Initialize instance attributes.
			class_:        c,
			declaration_:  declaration,
			constructors_: constructors,
			constants_:    constants,
			functions_:    functions,
		}
	}
}

// INSTANCE METHODS

// Target

type class_ struct {
	// Define instance attributes.
	class_        ClassClassLike
	declaration_  DeclarationLike
	constructors_ ConstructorsLike
	constants_    ConstantsLike
	functions_    FunctionsLike
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
	return v.constants_
}

func (v *class_) GetOptionalFunctions() FunctionsLike {
	return v.functions_
}

// Private
