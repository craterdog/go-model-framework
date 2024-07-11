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
	col "github.com/craterdog/go-collection-framework/v4/collection"
	age "github.com/craterdog/go-model-framework/v4/agent"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	ref "reflect"
)

// TYPE ALIASES

// AST

type (
	AbstractionLike         = ast.AbstractionLike
	AbstractionsLike        = ast.AbstractionsLike
	AdditionalArgumentLike  = ast.AdditionalArgumentLike
	AdditionalParameterLike = ast.AdditionalParameterLike
	AdditionalValueLike     = ast.AdditionalValueLike
	AliasLike               = ast.AliasLike
	ArgumentLike            = ast.ArgumentLike
	ArgumentsLike           = ast.ArgumentsLike
	ArrayLike               = ast.ArrayLike
	AspectLike              = ast.AspectLike
	AspectsLike             = ast.AspectsLike
	AttributeLike           = ast.AttributeLike
	AttributesLike          = ast.AttributesLike
	ChannelLike             = ast.ChannelLike
	ClassLike               = ast.ClassLike
	ClassesLike             = ast.ClassesLike
	ConstantLike            = ast.ConstantLike
	ConstantsLike           = ast.ConstantsLike
	ConstructorLike         = ast.ConstructorLike
	ConstructorsLike        = ast.ConstructorsLike
	DeclarationLike         = ast.DeclarationLike
	EnumerationLike         = ast.EnumerationLike
	FunctionLike            = ast.FunctionLike
	FunctionalLike          = ast.FunctionalLike
	FunctionalsLike         = ast.FunctionalsLike
	FunctionsLike           = ast.FunctionsLike
	GenericArgumentsLike    = ast.GenericArgumentsLike
	GenericParametersLike   = ast.GenericParametersLike
	HeaderLike              = ast.HeaderLike
	ImportsLike             = ast.ImportsLike
	InstanceLike            = ast.InstanceLike
	InstancesLike           = ast.InstancesLike
	MapLike                 = ast.MapLike
	MethodLike              = ast.MethodLike
	MethodsLike             = ast.MethodsLike
	ModelLike               = ast.ModelLike
	ModuleLike              = ast.ModuleLike
	ModulesLike             = ast.ModulesLike
	NoticeLike              = ast.NoticeLike
	ParameterizedLike       = ast.ParameterizedLike
	ParameterLike           = ast.ParameterLike
	ParametersLike          = ast.ParametersLike
	PrefixLike              = ast.PrefixLike
	ResultLike              = ast.ResultLike
	TypeLike                = ast.TypeLike
	TypesLike               = ast.TypesLike
	ValueLike               = ast.ValueLike
	ValuesLike              = ast.ValuesLike
)

// Agents

type (
	FormatterLike = age.FormatterLike
	GeneratorLike = age.GeneratorLike
	ParserLike    = age.ParserLike
	ValidatorLike = age.ValidatorLike
)

// UNIVERSAL CONSTRUCTORS

// AST

func Abstraction(args ...any) AbstractionLike {
	// Initialize the possible arguments.
	var prefix PrefixLike
	var alias AliasLike
	var name string
	var genericArguments GenericArgumentsLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case PrefixLike:
			prefix = actual
		case AliasLike:
			alias = actual
		case string:
			name = actual
		case GenericArgumentsLike:
			genericArguments = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the abstraction constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var abstraction = ast.Abstraction().Make(
		prefix,
		alias,
		name,
		genericArguments,
	)
	return abstraction
}

func Abstractions(args ...any) AbstractionsLike {
	// Initialize the possible arguments.
	var note = "// Abstractions"
	var sequence col.Sequential[AbstractionLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[AbstractionLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[AbstractionLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the abstractions constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var abstractions = ast.Abstractions().Make(
		note,
		sequence,
	)
	return abstractions
}

func AdditionalArgument(args ...any) AdditionalArgumentLike {
	// Initialize the possible arguments.
	var argument ArgumentLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ArgumentLike:
			argument = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the additional argument constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var additionalArgument = ast.AdditionalArgument().Make(argument)
	return additionalArgument
}

func AdditionalParameter(args ...any) AdditionalParameterLike {
	// Initialize the possible arguments.
	var parameter ParameterLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ParameterLike:
			parameter = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the additional parameter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var additionalParameter = ast.AdditionalParameter().Make(parameter)
	return additionalParameter
}

func AdditionalValue(args ...any) AdditionalValueLike {
	// Initialize the possible arguments.
	var name string

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the additional value constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var additionalValue = ast.AdditionalValue().Make(name)
	return additionalValue
}

func Alias(args ...any) AliasLike {
	// Initialize the possible arguments.
	var name string

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the alias constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var alias = ast.Alias().Make(name)
	return alias
}

func Argument(args ...any) ArgumentLike {
	// Initialize the possible arguments.
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the argument constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var argument = ast.Argument().Make(abstraction)
	return argument
}

func Arguments(args ...any) ArgumentsLike {
	// Initialize the possible arguments.
	var argument ArgumentLike
	var additionalArguments col.Sequential[AdditionalArgumentLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ArgumentLike:
			argument = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[AdditionalArgumentLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				additionalArguments = arg.(col.Sequential[AdditionalArgumentLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the additional arguments constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var arguments = ast.Arguments().Make(
		argument,
		additionalArguments,
	)
	return arguments
}

func Array(args ...any) ArrayLike {
	// Initialize the possible arguments.

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the array constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var array = ast.Array().Make()
	return array
}

func Aspect(args ...any) AspectLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var methods col.Sequential[MethodLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case DeclarationLike:
			declaration = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[MethodLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				methods = arg.(col.Sequential[MethodLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the aspect constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var aspect = ast.Aspect().Make(
		declaration,
		methods,
	)
	return aspect
}

func Aspects(args ...any) AspectsLike {
	// Initialize the possible arguments.
	var note = "// Aspects"
	var sequence col.Sequential[AspectLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[AspectLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[AspectLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the aspects constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var aspects = ast.Aspects().Make(
		note,
		sequence,
	)
	return aspects
}

func Attribute(args ...any) AttributeLike {
	// Initialize the possible arguments.
	var name string
	var parameter ParameterLike
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		case ParameterLike:
			parameter = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the attribute constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var attribute = ast.Attribute().Make(
		name,
		parameter,
		abstraction,
	)
	return attribute
}

func Attributes(args ...any) AttributesLike {
	// Initialize the possible arguments.
	var note = "// Attributes"
	var sequence col.Sequential[AttributeLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[AttributeLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[AttributeLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the attributes constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var attributes = ast.Attributes().Make(
		note,
		sequence,
	)
	return attributes
}

func Channel(args ...any) ChannelLike {
	// Initialize the possible arguments.

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the channel constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var channel = ast.Channel().Make()
	return channel
}

func Class(args ...any) ClassLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var constructors ConstructorsLike
	var constants ConstantsLike
	var functions FunctionsLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case DeclarationLike:
			declaration = actual
		case ConstructorsLike:
			constructors = actual
		case ConstantsLike:
			constants = actual
		case FunctionsLike:
			functions = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the class constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var class = ast.Class().Make(
		declaration,
		constructors,
		constants,
		functions,
	)
	return class
}

func Classes(args ...any) ClassesLike {
	// Initialize the possible arguments.
	var note = "// Classes"
	var sequence col.Sequential[ClassLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[ClassLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[ClassLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the classes constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var classes = ast.Classes().Make(
		note,
		sequence,
	)
	return classes
}

func Constant(args ...any) ConstantLike {
	// Initialize the possible arguments.
	var name string
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the constant constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constant = ast.Constant().Make(
		name,
		abstraction,
	)
	return constant
}

func Constants(args ...any) ConstantsLike {
	// Initialize the possible arguments.
	var note = "// Constants"
	var sequence col.Sequential[ConstantLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[ConstantLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[ConstantLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the constants constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var constants = ast.Constants().Make(
		note,
		sequence,
	)
	return constants
}

func Constructor(args ...any) ConstructorLike {
	// Initialize the possible arguments.
	var name string
	var parameters ParametersLike
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		case ParametersLike:
			parameters = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the constructor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constructor = ast.Constructor().Make(
		name,
		parameters,
		abstraction,
	)
	return constructor
}

func Constructors(args ...any) ConstructorsLike {
	// Initialize the possible arguments.
	var note = "// Constructors"
	var sequence col.Sequential[ConstructorLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[ConstructorLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[ConstructorLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the constructors constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var constructors = ast.Constructors().Make(
		note,
		sequence,
	)
	return constructors
}

func Declaration(args ...any) DeclarationLike {
	// Initialize the possible arguments.
	var comment string
	var name string
	var genericParameters GenericParametersLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			if len(comment) == 0 {
				comment = actual
			} else {
				name = actual
			}
		case GenericParametersLike:
			genericParameters = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the declaration constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var declaration = ast.Declaration().Make(
		comment,
		name,
		genericParameters,
	)
	return declaration
}

func Enumeration(args ...any) EnumerationLike {
	// Initialize the possible arguments.
	var values ValuesLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ValuesLike:
			values = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the enumeration constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var enumeration = ast.Enumeration().Make(
		values,
	)
	return enumeration
}

func Function(args ...any) FunctionLike {
	// Initialize the possible arguments.
	var name string
	var parameters ParametersLike
	var result ResultLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		case ParametersLike:
			parameters = actual
		case ResultLike:
			result = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the function constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var function = ast.Function().Make(
		name,
		parameters,
		result,
	)
	return function
}

func Functional(args ...any) FunctionalLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var parameters ParametersLike
	var result ResultLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case DeclarationLike:
			declaration = actual
		case ParametersLike:
			parameters = actual
		case ResultLike:
			result = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the functional constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var functional = ast.Functional().Make(
		declaration,
		parameters,
		result,
	)
	return functional
}

func Functionals(args ...any) FunctionalsLike {
	// Initialize the possible arguments.
	var note = "// Functionals"
	var sequence col.Sequential[FunctionalLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[FunctionalLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[FunctionalLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the functionals constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var functionals = ast.Functionals().Make(
		note,
		sequence,
	)
	return functionals
}

func Functions(args ...any) FunctionsLike {
	// Initialize the possible arguments.
	var note = "// Functions"
	var sequence col.Sequential[FunctionLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[FunctionLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[FunctionLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the functions constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var functions = ast.Functions().Make(
		note,
		sequence,
	)
	return functions
}

func GenericArguments(args ...any) GenericArgumentsLike {
	// Initialize the possible arguments.
	var arguments ArgumentsLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ArgumentsLike:
			arguments = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the generic arguments constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var genericArguments = ast.GenericArguments().Make(arguments)
	return genericArguments
}

func GenericParameters(args ...any) GenericParametersLike {
	// Initialize the possible arguments.
	var parameters ParametersLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ParametersLike:
			parameters = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the generic parameters constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var genericParameters = ast.GenericParameters().Make(parameters)
	return genericParameters
}

func Header(args ...any) HeaderLike {
	// Initialize the possible arguments.
	var comment string
	var name string

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			if len(comment) == 0 {
				comment = actual
			} else {
				name = actual
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the header constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var header = ast.Header().Make(
		comment,
		name,
	)
	return header
}

func Imports(args ...any) ImportsLike {
	// Initialize the possible arguments.
	var modules ModulesLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ModulesLike:
			modules = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the imports constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var imports = ast.Imports().Make(modules)
	return imports
}

func Instance(args ...any) InstanceLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var attributes AttributesLike
	var abstractions AbstractionsLike
	var methods MethodsLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case DeclarationLike:
			declaration = actual
		case AttributesLike:
			attributes = actual
		case AbstractionsLike:
			abstractions = actual
		case MethodsLike:
			methods = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the instance constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var instance = ast.Instance().Make(
		declaration,
		attributes,
		abstractions,
		methods,
	)
	return instance
}

func Instances(args ...any) InstancesLike {
	// Initialize the possible arguments.
	var note = "// Instances"
	var sequence col.Sequential[InstanceLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[InstanceLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[InstanceLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the instances constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var instances = ast.Instances().Make(
		note,
		sequence,
	)
	return instances
}

func Map(args ...any) MapLike {
	// Initialize the possible arguments.
	var name string

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the map constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var map_ = ast.Map().Make(name)
	return map_
}

func Method(args ...any) MethodLike {
	// Initialize the possible arguments.
	var name string
	var parameters ParametersLike
	var result ResultLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		case ParametersLike:
			parameters = actual
		case ResultLike:
			result = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the method constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var method = ast.Method().Make(
		name,
		parameters,
		result,
	)
	return method
}

func Methods(args ...any) MethodsLike {
	// Initialize the possible arguments.
	var note = "// Methods"
	var sequence col.Sequential[MethodLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[MethodLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[MethodLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the methods constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var methods = ast.Methods().Make(
		note,
		sequence,
	)
	return methods
}

func Model(args ...any) ModelLike {
	// Initialize the possible arguments.
	var notice NoticeLike
	var header HeaderLike
	var imports ImportsLike
	var types TypesLike
	var functionals FunctionalsLike
	var classes ClassesLike
	var instances InstancesLike
	var aspects AspectsLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case NoticeLike:
			notice = actual
		case HeaderLike:
			header = actual
		case ImportsLike:
			imports = actual
		case TypesLike:
			types = actual
		case FunctionalsLike:
			functionals = actual
		case ClassesLike:
			classes = actual
		case InstancesLike:
			instances = actual
		case AspectsLike:
			aspects = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the model constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var model = ast.Model().Make(
		notice,
		header,
		imports,
		types,
		functionals,
		classes,
		instances,
		aspects,
	)
	return model
}

func Module(args ...any) ModuleLike {
	// Initialize the possible arguments.
	var name string
	var text string

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			if len(name) == 0 {
				name = actual
			} else {
				text = actual
			}
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the module constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var module = ast.Module().Make(
		name,
		text,
	)
	return module
}

func Modules(args ...any) ModulesLike {
	// Initialize the possible arguments.
	var sequence col.Sequential[ModuleLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[ModuleLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[ModuleLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the modules constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var modules = ast.Modules().Make(sequence)
	return modules
}

func Notice(args ...any) NoticeLike {
	// Initialize the possible arguments.
	var comment string

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			comment = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the notice constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var notice = ast.Notice().Make(
		comment,
	)
	return notice
}

func Parameter(args ...any) ParameterLike {
	// Initialize the possible arguments.
	var name string
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the parameter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var parameter = ast.Parameter().Make(
		name,
		abstraction,
	)
	return parameter
}

func Parameterized(args ...any) ParameterizedLike {
	// Initialize the possible arguments.
	var parameters ParametersLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ParametersLike:
			parameters = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the imports constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var imports = ast.Parameterized().Make(parameters)
	return imports
}

func Parameters(args ...any) ParametersLike {
	// Initialize the possible arguments.
	var parameter ParameterLike
	var additionalParameters col.Sequential[AdditionalParameterLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ParameterLike:
			parameter = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[AdditionalParameterLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				additionalParameters = arg.(col.Sequential[AdditionalParameterLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the parameters constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var parameters = ast.Parameters().Make(
		parameter,
		additionalParameters,
	)
	return parameters
}

func Prefix(args ...any) PrefixLike {
	// Initialize the possible arguments.
	var array ast.ArrayLike
	var map_ ast.MapLike
	var channel ast.ChannelLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ast.ArrayLike:
			array = actual
		case ast.MapLike:
			map_ = actual
		case ast.ChannelLike:
			channel = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the prefix constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var prefix ast.PrefixLike
	switch {
	case array != nil:
		prefix = ast.Prefix().Make(array)
	case map_ != nil:
		prefix = ast.Prefix().Make(map_)
	case channel != nil:
		prefix = ast.Prefix().Make(channel)
	default:
		panic("The constructor for a result requires an argument.")
	}
	return prefix
}

func Result(args ...any) ResultLike {
	// Initialize the possible arguments.
	var abstraction AbstractionLike
	var parameterized ParameterizedLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case AbstractionLike:
			abstraction = actual
		case ParameterizedLike:
			parameterized = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the result constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var result ast.ResultLike
	switch {
	case abstraction != nil:
		result = ast.Result().Make(abstraction)
	case parameterized != nil:
		result = ast.Result().Make(parameterized)
	default:
		panic("The constructor for a result requires an argument.")
	}
	return result
}

func Type(args ...any) TypeLike {
	// Initialize the possible arguments.
	var declaration DeclarationLike
	var abstraction AbstractionLike
	var enumeration EnumerationLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case DeclarationLike:
			declaration = actual
		case AbstractionLike:
			abstraction = actual
		case EnumerationLike:
			enumeration = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the type constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var type_ = ast.Type().Make(
		declaration,
		abstraction,
		enumeration,
	)
	return type_
}

func Types(args ...any) TypesLike {
	// Initialize the possible arguments.
	var note = "// Types"
	var sequence col.Sequential[TypeLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			note = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[TypeLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				sequence = arg.(col.Sequential[TypeLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the types constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var types = ast.Types().Make(
		note,
		sequence,
	)
	return types
}

func Value(args ...any) ValueLike {
	// Initialize the possible arguments.
	var name string
	var abstraction AbstractionLike

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case string:
			name = actual
		case AbstractionLike:
			abstraction = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the value constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var value = ast.Value().Make(
		name,
		abstraction,
	)
	return value
}

func Values(args ...any) ValuesLike {
	// Initialize the possible arguments.
	var value ValueLike
	var additionalValues col.Sequential[AdditionalValueLike]

	// Process the actual arguments.
	for _, arg := range args {
		switch actual := arg.(type) {
		case ValueLike:
			value = actual
		default:
			// Unfortunately generic types must be handled reflectively.
			var sequenceType = ref.TypeOf((*col.Sequential[AdditionalValueLike])(nil)).Elem()
			var reflectedType = ref.TypeOf(arg)
			switch {
			case reflectedType.Implements(sequenceType):
				additionalValues = arg.(col.Sequential[AdditionalValueLike])
			default:
				var message = fmt.Sprintf(
					"An unknown argument type was passed into the additional values constructor: %T\n",
					actual,
				)
				panic(message)
			}
		}
	}

	// Call the constructor.
	var values = ast.Values().Make(
		value,
		additionalValues,
	)
	return values
}

// Agents

func Formatter(args ...any) FormatterLike {
	if len(args) > 0 {
		panic("The formatter constructor does not take any arguments.")
	}
	var formatter = age.Formatter().Make()
	return formatter
}

func Generator(args ...any) GeneratorLike {
	if len(args) > 0 {
		panic("The generator constructor does not take any arguments.")
	}
	var generator = age.Generator().Make()
	return generator
}

func Parser(args ...any) ParserLike {
	if len(args) > 0 {
		panic("The parser constructor does not take any arguments.")
	}
	var parser = age.Parser().Make()
	return parser
}

func Validator(args ...any) ValidatorLike {
	if len(args) > 0 {
		panic("The validator constructor does not take any arguments.")
	}
	var validator = age.Validator().Make()
	return validator
}
