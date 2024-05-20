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
	col "github.com/craterdog/go-collection-framework/v4"
)

// CLASS ACCESS

// Reference

var classClass = &classClass_{
	// Any private class constants should be initialized here.
}

// Function

func Class() ClassClassLike {
	return classClass
}

// CLASS METHODS

// Target

type classClass_ struct {
}

// Constants

// Constructors

func (c *classClass_) MakeWithAttributes(
	declaration DeclarationLike,
	constants col.ListLike[ConstantLike],
	constructors col.ListLike[ConstructorLike],
	functions col.ListLike[FunctionLike],
) ClassLike {
	return &class_{
		declaration_:  declaration,
		constants_:    constants,
		constructors_: constructors,
		functions_:    functions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type class_ struct {
	class_        ClassClassLike
	declaration_  DeclarationLike
	constants_    col.ListLike[ConstantLike]
	constructors_ col.ListLike[ConstructorLike]
	functions_    col.ListLike[FunctionLike]
}

// Attributes

func (v *class_) GetClass() ClassClassLike {
	return v.class_
}

func (v *class_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *class_) GetConstants() col.ListLike[ConstantLike] {
	return v.constants_
}

func (v *class_) GetConstructors() col.ListLike[ConstructorLike] {
	return v.constructors_
}

func (v *class_) GetFunctions() col.ListLike[FunctionLike] {
	return v.functions_
}

// Public

// Private
