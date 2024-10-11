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

package grammar

import (
	fmt "fmt"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// Constructor Methods

func (c *validatorClass_) Make() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	instance.visitor_ = Visitor().Make(instance)
	return instance
}

// INSTANCE INTERFACE

// Methodical Methods

func (v *validator_) ProcessComment(comment string) {
	v.validateToken(comment, CommentToken)
}

func (v *validator_) ProcessName(name string) {
	v.validateToken(name, NameToken)
}

func (v *validator_) ProcessNewline(newline string) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessPath(path string) {
	v.validateToken(path, PathToken)
}

func (v *validator_) ProcessSpace(space string) {
	v.validateToken(space, SpaceToken)
}

func (v *validator_) PreprocessInterfaceDefinitions(
	interfaceDefinition ast.InterfaceDefinitionsLike,
) {
	var classSection = interfaceDefinition.GetClassSection()
	var instanceSection = interfaceDefinition.GetInstanceSection()
	var classes = classSection.GetClassDefinitions().GetIterator()
	var instances = instanceSection.GetInstanceDefinitions().GetIterator()
	if classes.GetSize() != instances.GetSize() {
		panic("The class list and instance list are different sizes.")
	}
	for classes.HasNext() && instances.HasNext() {
		var class = classes.GetNext()
		var className = sts.TrimSuffix(class.GetDeclaration().GetName(), "ClassLike")
		var instance = instances.GetNext()
		var instanceName = sts.TrimSuffix(instance.GetDeclaration().GetName(), "Like")
		if className != instanceName {
			var message = fmt.Sprintf(
				"The following class name and instance name don't match: %q, %q",
				className,
				instanceName,
			)
			panic(message)
		}
	}
}

// Public Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return v.getClass()
}

func (v *validator_) ValidateModel(
	model ast.ModelLike,
) {
	v.visitor_.VisitModel(model)
}

// Private Methods

func (v *validator_) getClass() *validatorClass_ {
	return v.class_
}

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	if !Scanner().MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			Scanner().FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// PRIVATE INTERFACE

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	class_   *validatorClass_
	visitor_ VisitorLike

	// Define the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

var validatorClass = &validatorClass_{
	// Initialize the class constants.
}
