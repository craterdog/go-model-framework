/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

/*
Package "models" provides the ability to generate Go class files based on a
Go Package.go file that follows the format shown in the following code template:
  - https://github.com/craterdog/go-model-framework/blob/main/models/Package.go

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package models

import (
	col "github.com/craterdog/go-collection-framework/v3"
)

// Types

/*
PrefixType is a constrained type representing a prefix type for the Go language.
*/
type PrefixType uint8

const (
	ErrorPrefix PrefixType = iota
	AliasPrefix
	ArrayPrefix
	ChannelPrefix
	MapPrefix
)

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	CommentToken
	DelimiterToken
	EOFToken
	EOLToken
	IdentifierToken
	NoteToken
	SpaceToken
	TextToken
)

// Classes

/*
AbstractionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructors
	MakeWithAttributes(
		prefix PrefixLike,
		identifier string,
		arguments col.ListLike[AbstractionLike],
	) AbstractionLike
}

/*
AspectClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-like class.
*/
type AspectClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		methods col.ListLike[MethodLike],
	) AspectLike
}

/*
AttributeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameter ParameterLike,
		abstraction AbstractionLike,
	) AttributeLike
}

/*
ClassClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-like class.
*/
type ClassClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		constants col.ListLike[ConstantLike],
		constructors col.ListLike[ConstructorLike],
		functions col.ListLike[FunctionLike],
	) ClassLike
}

/*
ConstantClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ConstantLike
}

/*
ConstructorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-like class.
*/
type ConstructorClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters col.ListLike[ParameterLike],
		abstraction AbstractionLike,
	) ConstructorLike
}

/*
DeclarationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
		parameters col.ListLike[ParameterLike],
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructors
	MakeWithAttributes(
		parameter ParameterLike,
		identifiers col.ListLike[string],
	) EnumerationLike
}

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
FunctionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters col.ListLike[ParameterLike],
		result ResultLike,
	) FunctionLike
}

/*
FunctionalClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-like class.
*/
type FunctionalClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		parameters col.ListLike[ParameterLike],
		result ResultLike,
	) FunctionalLike
}

/*
GeneratorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete generator-like class.
*/
type GeneratorClassLike interface {
	// Constructors
	Make() GeneratorLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
	) HeaderLike
}

/*
InstanceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-like class.
*/
type InstanceClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		attributes col.ListLike[AttributeLike],
		abstractions col.ListLike[AbstractionLike],
		methods col.ListLike[MethodLike],
	) InstanceLike
}

/*
MethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete method-like class.
*/
type MethodClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		parameters col.ListLike[ParameterLike],
		result ResultLike,
	) MethodLike
}

/*
ModelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete model-like class.
*/
type ModelClassLike interface {
	// Constructors
	MakeWithAttributes(
		notice NoticeLike,
		header HeaderLike,
		modules col.ListLike[ModuleLike],
		types col.ListLike[TypeLike],
		functionals col.ListLike[FunctionalLike],
		aspects col.ListLike[AspectLike],
		classes col.ListLike[ClassLike],
		instances col.ListLike[InstanceLike],
	) ModelLike
}

/*
ModuleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete module-like class.
*/
type ModuleClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		text string,
	) ModuleLike
}

/*
NoticeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete notice-like class.
*/
type NoticeClassLike interface {
	// Constructors
	MakeWithComment(comment string) NoticeLike
}

/*
ParameterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ParameterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
PrefixClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete prefix-like class.
*/
type PrefixClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		type_ PrefixType,
	) PrefixLike
}

/*
ResultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete result-like class.
*/
type ResultClassLike interface {
	// Constructors
	MakeWithAbstraction(abstraction AbstractionLike) ResultLike
	MakeWithParameters(parameters col.ListLike[ParameterLike]) ResultLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

MatchToken() a list of strings representing any matches found in the specified
text of the specified token type using the regular expression defined for that
token type.  If the regular expression contains submatch patterns the matching
substrings are returned as additional values in the list.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	MatchToken(
		type_ TokenType,
		text string,
	) col.ListLike[string]
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	MakeWithAttributes(
		line int,
		position int,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
TypeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-like class.
*/
type TypeClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		enumeration EnumerationLike,
	) TypeLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

// Instances

/*
AbstractionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Attributes
	GetPrefix() PrefixLike
	GetIdentifier() string
	GetArguments() col.ListLike[AbstractionLike]
}

/*
AspectLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-like class.
*/
type AspectLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetMethods() col.ListLike[MethodLike]
}

/*
AttributeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-like class.
*/
type AttributeLike interface {
	// Attributes
	GetIdentifier() string
	GetParameter() ParameterLike
	GetAbstraction() AbstractionLike
}

/*
ClassLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-like class.
*/
type ClassLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetConstants() col.ListLike[ConstantLike]
	GetConstructors() col.ListLike[ConstructorLike]
	GetFunctions() col.ListLike[FunctionLike]
}

/*
ConstantLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-like class.
*/
type ConstantLike interface {
	// Attributes
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
ConstructorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-like class.
*/
type ConstructorLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
	GetAbstraction() AbstractionLike
}

/*
DeclarationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
}

/*
EnumerationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Attributes
	GetParameter() ParameterLike
	GetIdentifiers() col.ListLike[string]
}

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Methods
	FormatAbstraction(abstraction AbstractionLike) string
	FormatArguments(arguments col.ListLike[AbstractionLike]) string
	FormatGenerics(parameters col.ListLike[ParameterLike]) string
	FormatMethod(method MethodLike) string
	FormatModel(model ModelLike) string
	FormatParameter(parameter ParameterLike) string
	FormatParameterNames(parameters col.ListLike[ParameterLike]) string
	FormatParameters(parameters col.ListLike[ParameterLike]) string
	FormatResult(result ResultLike) string
}

/*
FunctionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionalLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-like class.
*/
type FunctionalLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetParameters() col.ListLike[ParameterLike]
	GetResult() ResultLike
}

/*
GeneratorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete generator-like class.
*/
type GeneratorLike interface {
	// Methods
	CreateModel(
		directory string,
		name string,
		copyright string,
	)
	GeneratePackage(directory string)
}

/*
HeaderLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete header-like class.
*/
type HeaderLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
}

/*
InstanceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-like class.
*/
type InstanceLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetAttributes() col.ListLike[AttributeLike]
	GetAbstractions() col.ListLike[AbstractionLike]
	GetMethods() col.ListLike[MethodLike]
}

/*
MethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete method-like class.
*/
type MethodLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
	GetResult() ResultLike
}

/*
ModelLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete model-like class.
*/
type ModelLike interface {
	// Attributes
	GetNotice() NoticeLike
	GetHeader() HeaderLike
	GetModules() col.ListLike[ModuleLike]
	GetTypes() col.ListLike[TypeLike]
	GetFunctionals() col.ListLike[FunctionalLike]
	GetAspects() col.ListLike[AspectLike]
	GetClasses() col.ListLike[ClassLike]
	GetInstances() col.ListLike[InstanceLike]
}

/*
ModuleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete module-like class.
*/
type ModuleLike interface {
	// Attributes
	GetIdentifier() string
	GetText() string
}

/*
NoticeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete notice-like class.
*/
type NoticeLike interface {
	// Attributes
	GetComment() string
}

/*
ParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Attributes
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) ModelLike
}

/*
PrefixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Attributes
	GetType() PrefixType
	GetIdentifier() string
}

/*
ResultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetParameters() col.ListLike[ParameterLike]
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
TypeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete type-like class.
*/
type TypeLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetEnumeration() EnumerationLike
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Methods
	ValidateModel(model ModelLike)
}
