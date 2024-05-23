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
	fmt "fmt"
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

// FUNCTION EXPORTS

// AST

func Abstraction(arguments ...any) AbstractionLike {
	// Initialize the possible arguments.
	var prefix PrefixLike
	var identifier string
	var arguments_ col.ListLike[ast.AbstractionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PrefixLike:
			prefix = actual
		case string:
			identifier = actual
		case col.ListLike[ast.AbstractionLike]:
			arguments_ = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the abstraction constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var abstraction = ast.Abstraction().MakeWithAttributes(
		prefix,
		identifier,
		arguments_,
	)
	return abstraction
}

func Aspect(arguments ...any) AspectLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var methods col.ListLike[ast.MethodLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			declaration = actual
		case col.ListLike[ast.MethodLike]:
			methods = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the aspect constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var aspect = ast.Aspect().MakeWithAttributes(
		declaration,
		methods,
	)
	return aspect
}

func Attribute(arguments ...any) AttributeLike {
	// Initialize the possible arguments.
	var identifier string
	var parameter ParameterLike
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case ParameterLike:
			parameter = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the attribute constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var attribute = ast.Attribute().MakeWithAttributes(
		identifier,
		parameter,
		abstraction,
	)
	return attribute
}

func Class(arguments ...any) ClassLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var constants col.ListLike[ast.ConstantLike]
	var constructors col.ListLike[ast.ConstructorLike]
	var functions col.ListLike[ast.FunctionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			declaration = actual
		case col.ListLike[ast.ConstantLike]:
			constants = actual
		case col.ListLike[ast.ConstructorLike]:
			constructors = actual
		case col.ListLike[ast.FunctionLike]:
			functions = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the class constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var class = ast.Class().MakeWithAttributes(
		declaration,
		constants,
		constructors,
		functions,
	)
	return class
}

func Constant(arguments ...any) ConstantLike {
	// Initialize the possible arguments.
	var identifier string
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the constant constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constant = ast.Constant().MakeWithAttributes(
		identifier,
		abstraction,
	)
	return constant
}

func Constructor(arguments ...any) ConstructorLike {
	// Initialize the possible arguments.
	var identifier string
	var parameters col.ListLike[ast.ParameterLike]
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the constructor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constructor = ast.Constructor().MakeWithAttributes(
		identifier,
		parameters,
		abstraction,
	)
	return constructor
}

func Declaration(arguments ...any) DeclarationLike {
	// Initialize the possible arguments.
	var comment string
	var identifier string
	var parameters col.ListLike[ast.ParameterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			if len(comment) == 0 {
				comment = actual
			} else {
				identifier = actual
			}
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the declaration constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var declaration = ast.Declaration().MakeWithAttributes(
		comment,
		identifier,
		parameters,
	)
	return declaration
}

func Enumeration(arguments ...any) EnumerationLike {
	// Initialize the possible arguments.
	var parameter ParameterLike
	var identifiers = col.List[string]()

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ParameterLike:
			parameter = actual
		case string:
			identifiers.AppendValue(actual)
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the enumeration constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var enumeration = ast.Enumeration().MakeWithAttributes(
		parameter,
		identifiers,
	)
	return enumeration
}

func Function(arguments ...any) FunctionLike {
	// Initialize the possible arguments.
	var identifier string
	var parameters col.ListLike[ast.ParameterLike]
	var result ResultLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case ResultLike:
			result = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the function constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var function = ast.Function().MakeWithAttributes(
		identifier,
		parameters,
		result,
	)
	return function
}

func Functional(arguments ...any) FunctionalLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var parameters col.ListLike[ast.ParameterLike]
	var result ResultLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			declaration = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case ResultLike:
			result = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the functional constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var functional = ast.Functional().MakeWithAttributes(
		declaration,
		parameters,
		result,
	)
	return functional
}

func Header(arguments ...any) HeaderLike {
	// Initialize the possible arguments.
	var comment string
	var identifier string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			if len(comment) == 0 {
				comment = actual
			} else {
				identifier = actual
			}
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the header constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var header = ast.Header().MakeWithAttributes(
		comment,
		identifier,
	)
	return header
}

func Instance(arguments ...any) InstanceLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var attributes col.ListLike[ast.AttributeLike]
	var abstractions col.ListLike[ast.AbstractionLike]
	var methods col.ListLike[ast.MethodLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			declaration = actual
		case col.ListLike[ast.AttributeLike]:
			attributes = actual
		case col.ListLike[ast.AbstractionLike]:
			abstractions = actual
		case col.ListLike[ast.MethodLike]:
			methods = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the instance constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var instance = ast.Instance().MakeWithAttributes(
		declaration,
		attributes,
		abstractions,
		methods,
	)
	return instance
}

func Method(arguments ...any) MethodLike {
	// Initialize the possible arguments.
	var identifier string
	var parameters col.ListLike[ast.ParameterLike]
	var result ResultLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case ResultLike:
			result = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the method constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var method = ast.Method().MakeWithAttributes(
		identifier,
		parameters,
		result,
	)
	return method
}

func Model(arguments ...any) ModelLike {
	// Initialize the possible arguments.
	var notice NoticeLike
	var header HeaderLike
	var modules col.ListLike[ast.ModuleLike]
	var types col.ListLike[ast.TypeLike]
	var functionals col.ListLike[ast.FunctionalLike]
	var aspects col.ListLike[ast.AspectLike]
	var classes col.ListLike[ast.ClassLike]
	var instances col.ListLike[ast.InstanceLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case NoticeLike:
			notice = actual
		case HeaderLike:
			header = actual
		case col.ListLike[ast.ModuleLike]:
			modules = actual
		case col.ListLike[ast.TypeLike]:
			types = actual
		case col.ListLike[ast.FunctionalLike]:
			functionals = actual
		case col.ListLike[ast.AspectLike]:
			aspects = actual
		case col.ListLike[ast.ClassLike]:
			classes = actual
		case col.ListLike[ast.InstanceLike]:
			instances = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the model constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var model = ast.Model().MakeWithAttributes(
		notice,
		header,
		modules,
		types,
		functionals,
		aspects,
		classes,
		instances,
	)
	return model
}

func Module(arguments ...any) ModuleLike {
	// Initialize the possible arguments.
	var identifier string
	var text string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			if len(identifier) == 0 {
				identifier = actual
			} else {
				text = actual
			}
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the module constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var module = ast.Module().MakeWithAttributes(
		identifier,
		text,
	)
	return module
}

func Notice(arguments ...any) NoticeLike {
	// Initialize the possible arguments.
	var comment string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			comment = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the notice constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var notice = ast.Notice().MakeWithComment(
		comment,
	)
	return notice
}

func Parameter(arguments ...any) ParameterLike {
	// Initialize the possible arguments.
	var identifier string
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the parameter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var parameter = ast.Parameter().MakeWithAttributes(
		identifier,
		abstraction,
	)
	return parameter
}

func Prefix(arguments ...any) PrefixLike {
	// Initialize the possible arguments.
	var identifier string
	var type_ ast.PrefixType

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case ast.PrefixType:
			type_ = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the prefix constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var prefix = ast.Prefix().MakeWithAttributes(
		identifier,
		type_,
	)
	return prefix
}

func Result(arguments ...any) ResultLike {
	// Initialize the possible arguments.
	var abstraction AbstractionLike
	var parameters col.ListLike[ast.ParameterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case AbstractionLike:
			abstraction = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the result constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var class = ast.Result()
	var result ResultLike
	switch {
	case abstraction != nil:
		result = class.MakeWithAbstraction(abstraction)
	case parameters != nil:
		result = class.MakeWithParameters(parameters)
	default:
		panic("The constructor for a result requires an argument.")
	}
	return result
}

func Type(arguments ...any) TypeLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var abstraction AbstractionLike
	var enumeration EnumerationLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case DeclarationLike:
			declaration = actual
		case AbstractionLike:
			abstraction = actual
		case EnumerationLike:
			enumeration = actual
		default:
			var message = fmt.Sprintf(
				"Unknown argument type passed into the type constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var type_ = ast.Type().MakeWithAttributes(
		declaration,
		abstraction,
		enumeration,
	)
	return type_
}

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
