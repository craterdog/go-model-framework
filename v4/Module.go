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
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	age "github.com/craterdog/go-model-framework/v4/gcmn/agent"
	ast "github.com/craterdog/go-model-framework/v4/gcmn/ast"
)

// Agents

func Formatter(arguments ...any) age.FormatterLike {
	if len(arguments) > 0 {
		panic("The formatter constructor does not take any arguments.")
	}
	var formatter = age.Formatter().Make()
	return formatter
}

func Generator(arguments ...any) age.GeneratorLike {
	if len(arguments) > 0 {
		panic("The generator constructor does not take any arguments.")
	}
	var generator = age.Generator().Make()
	return generator
}

func Parser(arguments ...any) age.ParserLike {
	if len(arguments) > 0 {
		panic("The parser constructor does not take any arguments.")
	}
	var parser = age.Parser().Make()
	return parser
}

func Validator(arguments ...any) age.ValidatorLike {
	if len(arguments) > 0 {
		panic("The validator constructor does not take any arguments.")
	}
	var validator = age.Validator().Make()
	return validator
}

// AST

func Abstraction(arguments ...any) ast.AbstractionLike {
	// Initialize the possible arguments.
	var prefix ast.PrefixLike
	var identifier string
	var arguments_ col.ListLike[ast.AbstractionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.PrefixLike:
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

func Aspect(arguments ...any) ast.AspectLike {
	// Initialize the possible arguments.
	var declaration ast.DeclarationLike
	var methods col.ListLike[ast.MethodLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.DeclarationLike:
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

func Attribute(arguments ...any) ast.AttributeLike {
	// Initialize the possible arguments.
	var identifier string
	var parameter ast.ParameterLike
	var abstraction ast.AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case ast.ParameterLike:
			parameter = actual
		case ast.AbstractionLike:
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

func Class(arguments ...any) ast.ClassLike {
	// Initialize the possible arguments.
	var declaration ast.DeclarationLike
	var constants col.ListLike[ast.ConstantLike]
	var constructors col.ListLike[ast.ConstructorLike]
	var functions col.ListLike[ast.FunctionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.DeclarationLike:
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

func Constant(arguments ...any) ast.ConstantLike {
	// Initialize the possible arguments.
	var identifier string
	var abstraction ast.AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case ast.AbstractionLike:
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

func Constructor(arguments ...any) ast.ConstructorLike {
	// Initialize the possible arguments.
	var identifier string
	var parameters col.ListLike[ast.ParameterLike]
	var abstraction ast.AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case ast.AbstractionLike:
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

func Declaration(arguments ...any) ast.DeclarationLike {
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

func Enumeration(arguments ...any) ast.EnumerationLike {
	// Initialize the possible arguments.
	var parameter ast.ParameterLike
	var notation = cdc.Notation().Make()
	var identifiers = col.List[string](notation).Make()

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.ParameterLike:
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

func Function(arguments ...any) ast.FunctionLike {
	// Initialize the possible arguments.
	var identifier string
	var parameters col.ListLike[ast.ParameterLike]
	var result ast.ResultLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case ast.ResultLike:
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

func Functional(arguments ...any) ast.FunctionalLike {
	// Initialize the possible arguments.
	var declaration ast.DeclarationLike
	var parameters col.ListLike[ast.ParameterLike]
	var result ast.ResultLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.DeclarationLike:
			declaration = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case ast.ResultLike:
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

func Header(arguments ...any) ast.HeaderLike {
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

func Instance(arguments ...any) ast.InstanceLike {
	// Initialize the possible arguments.
	var declaration ast.DeclarationLike
	var attributes col.ListLike[ast.AttributeLike]
	var abstractions col.ListLike[ast.AbstractionLike]
	var methods col.ListLike[ast.MethodLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.DeclarationLike:
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

func Method(arguments ...any) ast.MethodLike {
	// Initialize the possible arguments.
	var identifier string
	var parameters col.ListLike[ast.ParameterLike]
	var result ast.ResultLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case col.ListLike[ast.ParameterLike]:
			parameters = actual
		case ast.ResultLike:
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

func Model(arguments ...any) ast.ModelLike {
	// Initialize the possible arguments.
	var notice ast.NoticeLike
	var header ast.HeaderLike
	var modules col.ListLike[ast.ModuleLike]
	var types col.ListLike[ast.TypeLike]
	var functionals col.ListLike[ast.FunctionalLike]
	var aspects col.ListLike[ast.AspectLike]
	var classes col.ListLike[ast.ClassLike]
	var instances col.ListLike[ast.InstanceLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.NoticeLike:
			notice = actual
		case ast.HeaderLike:
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

func Module(arguments ...any) ast.ModuleLike {
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

func Notice(arguments ...any) ast.NoticeLike {
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

func Parameter(arguments ...any) ast.ParameterLike {
	// Initialize the possible arguments.
	var identifier string
	var abstraction ast.AbstractionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			identifier = actual
		case ast.AbstractionLike:
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

func Prefix(arguments ...any) ast.PrefixLike {
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

func Result(arguments ...any) ast.ResultLike {
	// Initialize the possible arguments.
	var abstraction ast.AbstractionLike
	var parameters col.ListLike[ast.ParameterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.AbstractionLike:
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
	var result ast.ResultLike
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

func Type(arguments ...any) ast.TypeLike {
	// Initialize the possible arguments.
	var declaration ast.DeclarationLike
	var abstraction ast.AbstractionLike
	var enumeration ast.EnumerationLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ast.DeclarationLike:
			declaration = actual
		case ast.AbstractionLike:
			abstraction = actual
		case ast.EnumerationLike:
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
