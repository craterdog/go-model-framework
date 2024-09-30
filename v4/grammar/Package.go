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
Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-model-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-model-framework/v4/ast"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	NameToken
	NewlineToken
	PathToken
	SpaceToken
)

// Classes

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor
	Make() FormatterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor
	Make() ParserLike
}

/*
ProcessorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor
	Make() ProcessorLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
	) ScannerLike

	// Function
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructor
	Make(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor
	Make() ValidatorLike
}

/*
VisitorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor
	Make(
		processor Methodical,
	) VisitorLike
}

// Instances

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Public
	GetClass() FormatterClassLike
	FormatModel(
		model ast.ModelLike,
	) string

	// Aspect
	Methodical
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Public
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.ModelLike
}

/*
ProcessorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Public
	GetClass() ProcessorClassLike

	// Aspect
	Methodical
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Public
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Public
	GetClass() TokenClassLike

	// Attribute
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Public
	GetClass() ValidatorClassLike
	ValidateModel(
		model ast.ModelLike,
	)

	// Aspect
	Methodical
}

/*
VisitorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Public
	GetClass() VisitorClassLike
	VisitModel(
		model ast.ModelLike,
	)
}

// Aspects

/*
Methodical defines the set of method signatures that must be supported
by all methodical processors.
*/
type Methodical interface {
	ProcessComment(
		comment string,
	)
	ProcessName(
		name string,
	)
	ProcessNewline(
		newline string,
	)
	ProcessPath(
		path string,
	)
	ProcessSpace(
		space string,
	)
	PreprocessAbstraction(
		abstraction ast.AbstractionLike,
	)
	ProcessAbstractionSlot(
		slot uint,
	)
	PostprocessAbstraction(
		abstraction ast.AbstractionLike,
	)
	PreprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		size uint,
	)
	ProcessAdditionalArgumentSlot(
		slot uint,
	)
	PostprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		size uint,
	)
	PreprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		size uint,
	)
	ProcessAdditionalValueSlot(
		slot uint,
	)
	PostprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		size uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
	)
	ProcessArgumentSlot(
		slot uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
	)
	PreprocessArray(
		array ast.ArrayLike,
	)
	ProcessArraySlot(
		slot uint,
	)
	PostprocessArray(
		array ast.ArrayLike,
	)
	PreprocessAspect(
		aspect ast.AspectLike,
		index uint,
		size uint,
	)
	ProcessAspectSlot(
		slot uint,
	)
	PostprocessAspect(
		aspect ast.AspectLike,
		index uint,
		size uint,
	)
	PreprocessAspectDefinitions(
		aspectDefinitions ast.AspectDefinitionsLike,
	)
	ProcessAspectDefinitionsSlot(
		slot uint,
	)
	PostprocessAspectDefinitions(
		aspectDefinitions ast.AspectDefinitionsLike,
	)
	PreprocessAspectInterfaces(
		aspectInterfaces ast.AspectInterfacesLike,
	)
	ProcessAspectInterfacesSlot(
		slot uint,
	)
	PostprocessAspectInterfaces(
		aspectInterfaces ast.AspectInterfacesLike,
	)
	PreprocessAttribute(
		attribute ast.AttributeLike,
		index uint,
		size uint,
	)
	ProcessAttributeSlot(
		slot uint,
	)
	PostprocessAttribute(
		attribute ast.AttributeLike,
		index uint,
		size uint,
	)
	PreprocessAttributeMethods(
		attributeMethods ast.AttributeMethodsLike,
	)
	ProcessAttributeMethodsSlot(
		slot uint,
	)
	PostprocessAttributeMethods(
		attributeMethods ast.AttributeMethodsLike,
	)
	PreprocessChannel(
		channel ast.ChannelLike,
	)
	ProcessChannelSlot(
		slot uint,
	)
	PostprocessChannel(
		channel ast.ChannelLike,
	)
	PreprocessClass(
		class ast.ClassLike,
		index uint,
		size uint,
	)
	ProcessClassSlot(
		slot uint,
	)
	PostprocessClass(
		class ast.ClassLike,
		index uint,
		size uint,
	)
	PreprocessClassDefinitions(
		classDefinitions ast.ClassDefinitionsLike,
	)
	ProcessClassDefinitionsSlot(
		slot uint,
	)
	PostprocessClassDefinitions(
		classDefinitions ast.ClassDefinitionsLike,
	)
	PreprocessClassMethods(
		classMethods ast.ClassMethodsLike,
	)
	ProcessClassMethodsSlot(
		slot uint,
	)
	PostprocessClassMethods(
		classMethods ast.ClassMethodsLike,
	)
	PreprocessConstant(
		constant ast.ConstantLike,
		index uint,
		size uint,
	)
	ProcessConstantSlot(
		slot uint,
	)
	PostprocessConstant(
		constant ast.ConstantLike,
		index uint,
		size uint,
	)
	PreprocessConstantMethods(
		constantMethods ast.ConstantMethodsLike,
	)
	ProcessConstantMethodsSlot(
		slot uint,
	)
	PostprocessConstantMethods(
		constantMethods ast.ConstantMethodsLike,
	)
	PreprocessConstructor(
		constructor ast.ConstructorLike,
		index uint,
		size uint,
	)
	ProcessConstructorSlot(
		slot uint,
	)
	PostprocessConstructor(
		constructor ast.ConstructorLike,
		index uint,
		size uint,
	)
	PreprocessConstructorMethods(
		constructorMethods ast.ConstructorMethodsLike,
	)
	ProcessConstructorMethodsSlot(
		slot uint,
	)
	PostprocessConstructorMethods(
		constructorMethods ast.ConstructorMethodsLike,
	)
	PreprocessDeclaration(
		declaration ast.DeclarationLike,
	)
	ProcessDeclarationSlot(
		slot uint,
	)
	PostprocessDeclaration(
		declaration ast.DeclarationLike,
	)
	PreprocessEnumeration(
		enumeration ast.EnumerationLike,
	)
	ProcessEnumerationSlot(
		slot uint,
	)
	PostprocessEnumeration(
		enumeration ast.EnumerationLike,
	)
	PreprocessFunction(
		function ast.FunctionLike,
		index uint,
		size uint,
	)
	ProcessFunctionSlot(
		slot uint,
	)
	PostprocessFunction(
		function ast.FunctionLike,
		index uint,
		size uint,
	)
	PreprocessFunctionMethods(
		functionMethods ast.FunctionMethodsLike,
	)
	ProcessFunctionMethodsSlot(
		slot uint,
	)
	PostprocessFunctionMethods(
		functionMethods ast.FunctionMethodsLike,
	)
	PreprocessFunctional(
		functional ast.FunctionalLike,
		index uint,
		size uint,
	)
	ProcessFunctionalSlot(
		slot uint,
	)
	PostprocessFunctional(
		functional ast.FunctionalLike,
		index uint,
		size uint,
	)
	PreprocessFunctionalDefinitions(
		functionalDefinitions ast.FunctionalDefinitionsLike,
	)
	ProcessFunctionalDefinitionsSlot(
		slot uint,
	)
	PostprocessFunctionalDefinitions(
		functionalDefinitions ast.FunctionalDefinitionsLike,
	)
	PreprocessGenericArguments(
		genericArguments ast.GenericArgumentsLike,
	)
	ProcessGenericArgumentsSlot(
		slot uint,
	)
	PostprocessGenericArguments(
		genericArguments ast.GenericArgumentsLike,
	)
	PreprocessGenericParameters(
		genericParameters ast.GenericParametersLike,
	)
	ProcessGenericParametersSlot(
		slot uint,
	)
	PostprocessGenericParameters(
		genericParameters ast.GenericParametersLike,
	)
	PreprocessHeader(
		header ast.HeaderLike,
	)
	ProcessHeaderSlot(
		slot uint,
	)
	PostprocessHeader(
		header ast.HeaderLike,
	)
	PreprocessImports(
		imports ast.ImportsLike,
	)
	ProcessImportsSlot(
		slot uint,
	)
	PostprocessImports(
		imports ast.ImportsLike,
	)
	PreprocessInstance(
		instance ast.InstanceLike,
		index uint,
		size uint,
	)
	ProcessInstanceSlot(
		slot uint,
	)
	PostprocessInstance(
		instance ast.InstanceLike,
		index uint,
		size uint,
	)
	PreprocessInstanceDefinitions(
		instanceDefinitions ast.InstanceDefinitionsLike,
	)
	ProcessInstanceDefinitionsSlot(
		slot uint,
	)
	PostprocessInstanceDefinitions(
		instanceDefinitions ast.InstanceDefinitionsLike,
	)
	PreprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
	)
	ProcessInstanceMethodsSlot(
		slot uint,
	)
	PostprocessInstanceMethods(
		instanceMethods ast.InstanceMethodsLike,
	)
	PreprocessInterface(
		interface_ ast.InterfaceLike,
		index uint,
		size uint,
	)
	ProcessInterfaceSlot(
		slot uint,
	)
	PostprocessInterface(
		interface_ ast.InterfaceLike,
		index uint,
		size uint,
	)
	PreprocessInterfaceDefinitions(
		interfaceDefinitions ast.InterfaceDefinitionsLike,
	)
	ProcessInterfaceDefinitionsSlot(
		slot uint,
	)
	PostprocessInterfaceDefinitions(
		interfaceDefinitions ast.InterfaceDefinitionsLike,
	)
	PreprocessMap(
		map_ ast.MapLike,
	)
	ProcessMapSlot(
		slot uint,
	)
	PostprocessMap(
		map_ ast.MapLike,
	)
	PreprocessMethod(
		method ast.MethodLike,
		index uint,
		size uint,
	)
	ProcessMethodSlot(
		slot uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
		index uint,
		size uint,
	)
	PreprocessModel(
		model ast.ModelLike,
	)
	ProcessModelSlot(
		slot uint,
	)
	PostprocessModel(
		model ast.ModelLike,
	)
	PreprocessModule(
		module ast.ModuleLike,
		index uint,
		size uint,
	)
	ProcessModuleSlot(
		slot uint,
	)
	PostprocessModule(
		module ast.ModuleLike,
		index uint,
		size uint,
	)
	PreprocessModuleDefinition(
		moduleDefinition ast.ModuleDefinitionLike,
	)
	ProcessModuleDefinitionSlot(
		slot uint,
	)
	PostprocessModuleDefinition(
		moduleDefinition ast.ModuleDefinitionLike,
	)
	PreprocessNone(
		none ast.NoneLike,
	)
	ProcessNoneSlot(
		slot uint,
	)
	PostprocessNone(
		none ast.NoneLike,
	)
	PreprocessNotice(
		notice ast.NoticeLike,
	)
	ProcessNoticeSlot(
		slot uint,
	)
	PostprocessNotice(
		notice ast.NoticeLike,
	)
	PreprocessParameter(
		parameter ast.ParameterLike,
		index uint,
		size uint,
	)
	ProcessParameterSlot(
		slot uint,
	)
	PostprocessParameter(
		parameter ast.ParameterLike,
		index uint,
		size uint,
	)
	PreprocessParameterized(
		parameterized ast.ParameterizedLike,
	)
	ProcessParameterizedSlot(
		slot uint,
	)
	PostprocessParameterized(
		parameterized ast.ParameterizedLike,
	)
	PreprocessPrefix(
		prefix ast.PrefixLike,
	)
	ProcessPrefixSlot(
		slot uint,
	)
	PostprocessPrefix(
		prefix ast.PrefixLike,
	)
	PreprocessPrimitiveDefinitions(
		primitiveDefinitions ast.PrimitiveDefinitionsLike,
	)
	ProcessPrimitiveDefinitionsSlot(
		slot uint,
	)
	PostprocessPrimitiveDefinitions(
		primitiveDefinitions ast.PrimitiveDefinitionsLike,
	)
	PreprocessPublicMethods(
		publicMethods ast.PublicMethodsLike,
	)
	ProcessPublicMethodsSlot(
		slot uint,
	)
	PostprocessPublicMethods(
		publicMethods ast.PublicMethodsLike,
	)
	PreprocessResult(
		result ast.ResultLike,
	)
	ProcessResultSlot(
		slot uint,
	)
	PostprocessResult(
		result ast.ResultLike,
	)
	PreprocessSuffix(
		suffix ast.SuffixLike,
	)
	ProcessSuffixSlot(
		slot uint,
	)
	PostprocessSuffix(
		suffix ast.SuffixLike,
	)
	PreprocessType(
		type_ ast.TypeLike,
		index uint,
		size uint,
	)
	ProcessTypeSlot(
		slot uint,
	)
	PostprocessType(
		type_ ast.TypeLike,
		index uint,
		size uint,
	)
	PreprocessTypeDefinitions(
		typeDefinitions ast.TypeDefinitionsLike,
	)
	ProcessTypeDefinitionsSlot(
		slot uint,
	)
	PostprocessTypeDefinitions(
		typeDefinitions ast.TypeDefinitionsLike,
	)
	PreprocessValue(
		value ast.ValueLike,
	)
	ProcessValueSlot(
		slot uint,
	)
	PostprocessValue(
		value ast.ValueLike,
	)
}
