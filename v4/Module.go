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
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
the commonly used classes that are exported by this module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in a package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-model-framework/wiki
*/
package module

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	gen "github.com/craterdog/go-model-framework/v4/generator"
	gra "github.com/craterdog/go-model-framework/v4/grammar"
)

// TYPE ALIASES

// AST

type (
	AbstractionLike           = ast.AbstractionLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalConstraintLike  = ast.AdditionalConstraintLike
	AdditionalValueLike       = ast.AdditionalValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	ArrayLike                 = ast.ArrayLike
	AspectDefinitionLike      = ast.AspectDefinitionLike
	AspectInterfaceLike       = ast.AspectInterfaceLike
	AspectMethodLike          = ast.AspectMethodLike
	AspectSectionLike         = ast.AspectSectionLike
	AspectSubsectionLike      = ast.AspectSubsectionLike
	AttributeMethodLike       = ast.AttributeMethodLike
	AttributeSubsectionLike   = ast.AttributeSubsectionLike
	ChannelLike               = ast.ChannelLike
	ClassDefinitionLike       = ast.ClassDefinitionLike
	ClassMethodsLike          = ast.ClassMethodsLike
	ClassSectionLike          = ast.ClassSectionLike
	ConstantMethodLike        = ast.ConstantMethodLike
	ConstantSubsectionLike    = ast.ConstantSubsectionLike
	ConstraintLike            = ast.ConstraintLike
	ConstraintsLike           = ast.ConstraintsLike
	ConstructorMethodLike     = ast.ConstructorMethodLike
	ConstructorSubsectionLike = ast.ConstructorSubsectionLike
	DeclarationLike           = ast.DeclarationLike
	EnumerationLike           = ast.EnumerationLike
	FunctionMethodLike        = ast.FunctionMethodLike
	FunctionSubsectionLike    = ast.FunctionSubsectionLike
	FunctionalDefinitionLike  = ast.FunctionalDefinitionLike
	FunctionalSectionLike     = ast.FunctionalSectionLike
	GetterMethodLike          = ast.GetterMethodLike
	HeaderLike                = ast.HeaderLike
	ImportsLike               = ast.ImportsLike
	InstanceDefinitionLike    = ast.InstanceDefinitionLike
	InstanceMethodsLike       = ast.InstanceMethodsLike
	InstanceSectionLike       = ast.InstanceSectionLike
	InterfaceDefinitionsLike  = ast.InterfaceDefinitionsLike
	MapLike                   = ast.MapLike
	MethodLike                = ast.MethodLike
	ModelLike                 = ast.ModelLike
	ModuleLike                = ast.ModuleLike
	ModuleDefinitionLike      = ast.ModuleDefinitionLike
	NoneLike                  = ast.NoneLike
	NoticeLike                = ast.NoticeLike
	ParameterLike             = ast.ParameterLike
	ParameterizedLike         = ast.ParameterizedLike
	PrefixLike                = ast.PrefixLike
	PrimitiveDefinitionsLike  = ast.PrimitiveDefinitionsLike
	PublicMethodLike          = ast.PublicMethodLike
	PublicSubsectionLike      = ast.PublicSubsectionLike
	ResultLike                = ast.ResultLike
	SetterMethodLike          = ast.SetterMethodLike
	SuffixLike                = ast.SuffixLike
	TypeDefinitionLike        = ast.TypeDefinitionLike
	TypeSectionLike           = ast.TypeSectionLike
	ValueLike                 = ast.ValueLike
)

// Grammar

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike

	Methodical = gra.Methodical
)

// Generator

type (
	ClassesLike = gen.ClassesLike
)

// UNIVERSAL CONSTRUCTORS

// Grammar

func Formatter(args ...any) FormatterLike {
	if len(args) > 0 {
		panic("The \"formatter\" constructor does not take any arguments.")
	}
	var formatter = gra.Formatter().Make()
	return formatter
}

func Parser(args ...any) ParserLike {
	if len(args) > 0 {
		panic("The \"parser\" constructor does not take any arguments.")
	}
	var parser = gra.Parser().Make()
	return parser
}

func Validator(args ...any) ValidatorLike {
	if len(args) > 0 {
		panic("The \"validator\" constructor does not take any arguments.")
	}
	var validator = gra.Validator().Make()
	return validator
}

func Visitor(args ...any) VisitorLike {
	// Initialize the possible arguments.
	var processor Methodical

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case Methodical:
			processor = actual
		default:
			if uti.IsDefined(arg) {
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the \"visitor\" constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var visitor = gra.Visitor().Make(
		processor,
	)
	return visitor
}

// Generator

func Classes(args ...any) ClassesLike {
	if len(args) > 0 {
		panic("The \"classes\" constructor does not take any arguments.")
	}
	var classes = gen.Classes().Make()
	return classes
}
