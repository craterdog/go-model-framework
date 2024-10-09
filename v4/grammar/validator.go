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

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Initialize the class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	var validator = &validator_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	validator.visitor_ = Visitor().Make(validator)
	return validator
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_   *validatorClass_
	visitor_ VisitorLike

	// Define the inherited aspects.
	Methodical
}

// Public

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

func (v *validator_) ValidateToken(
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

func (v *validator_) ValidateModel(model ast.ModelLike) {
	v.visitor_.VisitModel(model)
}

// Methodical

func (v *validator_) ProcessComment(comment string) {
	v.ValidateToken(comment, CommentToken)
}

func (v *validator_) ProcessName(name string) {
	v.ValidateToken(name, NameToken)
}

func (v *validator_) ProcessNewline(newline string) {
	v.ValidateToken(newline, NewlineToken)
}

func (v *validator_) ProcessPath(path string) {
	v.ValidateToken(path, PathToken)
}

func (v *validator_) ProcessSpace(space string) {
	v.ValidateToken(space, SpaceToken)
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

func (v *validator_) PreprocessModel(model ast.ModelLike) {
}

func (v *validator_) ProcessModelSlot(slot uint) {
}

func (v *validator_) PostprocessModel(model ast.ModelLike) {
}
