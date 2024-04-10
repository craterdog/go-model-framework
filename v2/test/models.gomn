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

Additional implementations of the concrete classes provided by this package can
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
PrefixType is a constrained type representing a prefix type.
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
AbstractionClassLike defines the set of class constants, constructors and
functions that must be supported by all abstraction-class-like classes.
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
AspectClassLike defines the set of class constants, constructors and
functions that must be supported by all aspect-class-like classes.
*/
type AspectClassLike interface {
	// Constructors
	MakeWithAttributes(
		declaration DeclarationLike,
		methods col.ListLike[MethodLike],
	) AspectLike
}

/*
AttributeClassLike defines the set of class constants, constructors and
functions that must be supported by all attribute-class-like classes.
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
ClassClassLike defines the set of class constants, constructors and
functions that must be supported by all class-class-like classes.
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
ConstantClassLike defines the set of class constants, constructors and
functions that must be supported by all constant-class-like classes.
*/
type ConstantClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ConstantLike
}

/*
ConstructorClassLike defines the set of class constants, constructors and
functions that must be supported by all constructor-class-like classes.
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
DeclarationClassLike defines the set of class constants, constructors and
functions that must be supported by all declaration-class-like classes.
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
EnumerationClassLike defines the set of class constants, constructors and
functions that must be supported by all enumeration-class-like classes.
*/
type EnumerationClassLike interface {
	// Constructors
	MakeWithAttributes(
		parameter ParameterLike,
		identifiers col.ListLike[string],
	) EnumerationLike
}

/*
FormatterClassLike defines the set of class constants, constructors and
functions that must be supported by all formatter-class-like classes.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
FunctionClassLike defines the set of class constants, constructors and
functions that must be supported by all function-class-like classes.
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
FunctionalClassLike defines the set of class constants, constructors and
functions that must be supported by all functional-class-like classes.
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
GeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all generator-class-like classes.
*/
type GeneratorClassLike interface {
	// Constructors
	Make() GeneratorLike
}

/*
HeaderClassLike defines the set of class constants, constructors and functions
that must be supported by all header-class-like classes.
*/
type HeaderClassLike interface {
	// Constructors
	MakeWithAttributes(
		comment string,
		identifier string,
	) HeaderLike
}

/*
InstanceClassLike defines the set of class constants, constructors and functions
that must be supported by all instance-class-like classes.
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
MethodClassLike defines the set of class constants, constructors and functions
that must be supported by all method-class-like classes.
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
ModelClassLike defines the set of class constants, constructors and functions
that must be supported by all package-class-like classes.
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
ModuleClassLike defines the set of class constants, constructors and
functions that must be supported by all module-class-like classes.
*/
type ModuleClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		text string,
	) ModuleLike
}

/*
NoticeClassLike defines the set of class constants, constructors and
functions that must be supported by all notice-class-like classes.
*/
type NoticeClassLike interface {
	// Constructors
	MakeWithComment(comment string) NoticeLike
}

/*
ParameterClassLike defines the set of class constants, constructors and
functions that must be supported by all parameter-class-like classes.
*/
type ParameterClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		abstraction AbstractionLike,
	) ParameterLike
}

/*
ParserClassLike defines the set of class constants, constructors and functions
that must be supported by all parser-class-like classes.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
PrefixClassLike defines the set of class constants, constructors and
functions that must be supported by all prefix-class-like classes.
*/
type PrefixClassLike interface {
	// Constructors
	MakeWithAttributes(
		identifier string,
		type_ PrefixType,
	) PrefixLike
}

/*
ResultClassLike defines the set of class constants, constructors and functions
that must be supported by all result-class-like classes.
*/
type ResultClassLike interface {
	// Constructors
	MakeWithAbstraction(abstraction AbstractionLike) ResultLike
	MakeWithParameters(parameters col.ListLike[ParameterLike]) ResultLike
}

/*
ScannerClassLike is a class interface that defines the set of class
constants, constructors and functions that must be supported by each
scanner-class-like concrete class.  The following functions are supported:

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
TokenClassLike is a class interface that defines the set of class
constants, constructors and functions that must be supported by each
token-class-like concrete class.
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
TypeClassLike defines the set of class constants, constructors and
functions that must be supported by all type-class-like classes.
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
ValidatorClassLike defines the set of class constants, constructors and
functions that must be supported by all validator-class-like classes.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

// Instances

/*
AbstractionLike defines the set of abstractions and methods that must be
supported by all abstraction-like instances.
*/
type AbstractionLike interface {
	// Attributes
	GetPrefix() PrefixLike
	GetIdentifier() string
	GetArguments() col.ListLike[AbstractionLike]
}

/*
AspectLike defines the set of abstractions and methods that must be supported by
all aspect-like instances.
*/
type AspectLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetMethods() col.ListLike[MethodLike]
}

/*
AttributeLike defines the set of abstractions and methods that must be supported
by all attribute-like instances.
*/
type AttributeLike interface {
	// Attributes
	GetIdentifier() string
	GetParameter() ParameterLike
	GetAbstraction() AbstractionLike
}

/*
ClassLike defines the set of abstractions and methods that must be supported by
all class-like instances.
*/
type ClassLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetConstants() col.ListLike[ConstantLike]
	GetConstructors() col.ListLike[ConstructorLike]
	GetFunctions() col.ListLike[FunctionLike]
}

/*
ConstantLike defines the set of abstractions and methods that must be supported
by all constant-like instances.
*/
type ConstantLike interface {
	// Attributes
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
ConstructorLike defines the set of abstractions and methods that must be
supported by all constructor-like instances.
*/
type ConstructorLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
	GetAbstraction() AbstractionLike
}

/*
DeclarationLike defines the set of abstractions and methods that must be
supported by all declaration-like instances.
*/
type DeclarationLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
}

/*
EnumerationLike defines the set of abstractions and methods that must be
supported by all enumeration-like instances.
*/
type EnumerationLike interface {
	// Attributes
	GetParameter() ParameterLike
	GetIdentifiers() col.ListLike[string]
}

/*
FormatterLike defines the set of abstractions and methods that must be
supported by all formatter-like instances.
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
FunctionLike defines the set of abstractions and methods that must be supported
by all function-like instances.
*/
type FunctionLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionalLike defines the set of abstractions and methods that must be
supported by all functional-like instances.
*/
type FunctionalLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetParameters() col.ListLike[ParameterLike]
	GetResult() ResultLike
}

/*
GeneratorLike defines the set of abstractions and methods that must be
supported by all generator-like instances.
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
HeaderLike defines the set of abstractions and methods that must be supported by
all header-like instances.
*/
type HeaderLike interface {
	// Attributes
	GetComment() string
	GetIdentifier() string
}

/*
InstanceLike defines the set of abstractions and methods that must be supported
by all instance-like instances.
*/
type InstanceLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetAttributes() col.ListLike[AttributeLike]
	GetAbstractions() col.ListLike[AbstractionLike]
	GetMethods() col.ListLike[MethodLike]
}

/*
MethodLike defines the set of abstractions and methods that must be supported by
all method-like instances.
*/
type MethodLike interface {
	// Attributes
	GetIdentifier() string
	GetParameters() col.ListLike[ParameterLike]
	GetResult() ResultLike
}

/*
ModelLike defines the set of abstractions and methods that must be supported by
all package-like instances.
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
ModuleLike defines the set of abstractions and methods that must be
supported by all module-like instances.
*/
type ModuleLike interface {
	// Attributes
	GetIdentifier() string
	GetText() string
}

/*
NoticeLike defines the set of abstractions and methods that must be supported
by all notice-like instances.
*/
type NoticeLike interface {
	// Attributes
	GetComment() string
}

/*
ParameterLike defines the set of abstractions and methods that must be supported
by all parameter-like instances.
*/
type ParameterLike interface {
	// Attributes
	GetIdentifier() string
	GetAbstraction() AbstractionLike
}

/*
ParserLike defines the set of abstractions and methods that must be supported by
all parser-like instances.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) ModelLike
}

/*
PrefixLike defines the set of abstractions and methods that must be
supported by all prefix-like instances.
*/
type PrefixLike interface {
	// Attributes
	GetType() PrefixType
	GetIdentifier() string
}

/*
ResultLike defines the set of abstractions and methods that must be supported by
all result-like instances.
*/
type ResultLike interface {
	// Attributes
	GetAbstraction() AbstractionLike
	GetParameters() col.ListLike[ParameterLike]
}

/*
ScannerLike defines the set of abstractions and methods that must be supported
by all scanner-like instances.
*/
type ScannerLike interface {
}

/*
TokenLike is an instance interface that defines the complete set of
abstractions and methods that must be supported by each instance of a
token-like concrete class.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
TypeLike defines the set of abstractions and methods that must be
supported by all type-like instances.
*/
type TypeLike interface {
	// Attributes
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetEnumeration() EnumerationLike
}

/*
ValidatorLike defines the set of abstractions and methods that must be
supported by all validator-like instances.
*/
type ValidatorLike interface {
	// Methods
	ValidateModel(model ModelLike)
}
