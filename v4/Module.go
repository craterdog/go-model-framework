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

/*
Package "module" defines a universal constructor for each class that is exported
by this module.  Each constructor delegates the actual construction process to
one of the classes defined in a subpackage for this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-model-framework/wiki

This package follows the Crater Dog Technologies™ (craterdog) Go Coding
Conventions located here:
  - https://github.com/craterdog/go-model-framework/wiki

The classes defined in this module provide the ability to parse, validate and
format Go Class Model Notation (GCMN).  They can also generate concrete class
implementation files for each abstract class defined in the Package.go file.
*/
package module

import (
	col "github.com/craterdog/go-collection-framework/v4"
	age "github.com/craterdog/go-model-framework/v4/gcmn/agent"
	ast "github.com/craterdog/go-model-framework/v4/gcmn/ast"
	sts "strings"
)

// TYPE PROMOTIONS

// AST

type (
	PrefixType ast.PrefixType
)

const (
	ErrorPrefix   = ast.ErrorPrefix
	AliasPrefix   = ast.AliasPrefix
	ArrayPrefix   = ast.ArrayPrefix
	ChannelPrefix = ast.ChannelPrefix
	MapPrefix     = ast.MapPrefix
)

type (
	AbstractionLike ast.AbstractionLike
	AspectLike      ast.AspectLike
	AttributeLike   ast.AttributeLike
	ClassLike       ast.ClassLike
	ConstantLike    ast.ConstantLike
	ConstructorLike ast.ConstructorLike
	DeclarationLike ast.DeclarationLike
	EnumerationLike ast.EnumerationLike
	FunctionLike    ast.FunctionLike
	FunctionalLike  ast.FunctionalLike
	HeaderLike      ast.HeaderLike
	InstanceLike    ast.InstanceLike
	MethodLike      ast.MethodLike
	ModelLike       ast.ModelLike
	ModuleLike      ast.ModuleLike
	NoticeLike      ast.NoticeLike
	ParameterLike   ast.ParameterLike
	PrefixLike      ast.PrefixLike
	ResultLike      ast.ResultLike
	TypeLike        ast.TypeLike
)

// Agents

type (
	FormatterLike age.FormatterLike
	GeneratorLike age.GeneratorLike
	ParserLike    age.ParserLike
	ValidatorLike age.FormatterLike
)

// UNIVERSAL FUNCTIONS

// Agents

func CreateModel(
	name string,
	copyright string,
) string {
	var generator = age.Generator().Make()
	return generator.CreateModel(name, copyright)
}

func FormatModel(model ModelLike) string {
	var formatter = age.Formatter().Make()
	return formatter.FormatModel(model)
}

func GenerateClasses(model ModelLike) col.CatalogLike[string, string] {
	var classes = col.Catalog[string, string]()
	var generator = age.Generator().Make()
	var iterator = model.GetClasses().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var className = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetIdentifier(),
			"ClassLike",
		))
		var classSource = generator.GenerateClass(model, className)
		classes.SetValue(className, classSource)
	}
	return classes
}

func ParseSource(source string) ModelLike {
	var parser = age.Parser().Make()
	return parser.ParseSource(source)
}

func ValidateModel(model ModelLike) {
	var validator = age.Validator().Make()
	validator.ValidateModel(model)
}
