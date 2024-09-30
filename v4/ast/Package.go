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
Package "ast" provides the abstract syntax tree (AST) classes for this module.
Each AST class manages the attributes associated with the rule definition found
in the syntax grammar with the same rule name as the class.

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
package ast

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
AbstractionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructor
	Make(
		optionalPrefix PrefixLike,
		name string,
		optionalSuffix SuffixLike,
		optionalGenericArguments GenericArgumentsLike,
	) AbstractionLike
}

/*
AdditionalArgumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-argument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructor
	Make(
		argument ArgumentLike,
	) AdditionalArgumentLike
}

/*
AdditionalValueClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-value-like class.
*/
type AdditionalValueClassLike interface {
	// Constructor
	Make(
		name string,
	) AdditionalValueLike
}

/*
ArgumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor
	Make(
		abstraction AbstractionLike,
	) ArgumentLike
}

/*
ArrayClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike interface {
	// Constructor
	Make() ArrayLike
}

/*
AspectClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-like class.
*/
type AspectClassLike interface {
	// Constructor
	Make(
		declaration DeclarationLike,
		methods abs.Sequential[MethodLike],
	) AspectLike
}

/*
AspectDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-definitions-like class.
*/
type AspectDefinitionsClassLike interface {
	// Constructor
	Make(
		aspects abs.Sequential[AspectLike],
	) AspectDefinitionsLike
}

/*
AspectInterfacesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-interfaces-like class.
*/
type AspectInterfacesClassLike interface {
	// Constructor
	Make(
		interfaces abs.Sequential[InterfaceLike],
	) AspectInterfacesLike
}

/*
AttributeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructor
	Make(
		name string,
		optionalParameter ParameterLike,
		optionalAbstraction AbstractionLike,
	) AttributeLike
}

/*
AttributeMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-methods-like class.
*/
type AttributeMethodsClassLike interface {
	// Constructor
	Make(
		attributes abs.Sequential[AttributeLike],
	) AttributeMethodsLike
}

/*
ChannelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete channel-like class.
*/
type ChannelClassLike interface {
	// Constructor
	Make() ChannelLike
}

/*
ClassClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-like class.
*/
type ClassClassLike interface {
	// Constructor
	Make(
		declaration DeclarationLike,
		classMethods ClassMethodsLike,
	) ClassLike
}

/*
ClassDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-definitions-like class.
*/
type ClassDefinitionsClassLike interface {
	// Constructor
	Make(
		classes abs.Sequential[ClassLike],
	) ClassDefinitionsLike
}

/*
ClassMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-methods-like class.
*/
type ClassMethodsClassLike interface {
	// Constructor
	Make(
		constructorMethods ConstructorMethodsLike,
		optionalConstantMethods ConstantMethodsLike,
		optionalFunctionMethods FunctionMethodsLike,
	) ClassMethodsLike
}

/*
ConstantClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructor
	Make(
		name string,
		abstraction AbstractionLike,
	) ConstantLike
}

/*
ConstantMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-methods-like class.
*/
type ConstantMethodsClassLike interface {
	// Constructor
	Make(
		constants abs.Sequential[ConstantLike],
	) ConstantMethodsLike
}

/*
ConstructorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-like class.
*/
type ConstructorClassLike interface {
	// Constructor
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		abstraction AbstractionLike,
	) ConstructorLike
}

/*
ConstructorMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-methods-like class.
*/
type ConstructorMethodsClassLike interface {
	// Constructor
	Make(
		constructors abs.Sequential[ConstructorLike],
	) ConstructorMethodsLike
}

/*
DeclarationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructor
	Make(
		comment string,
		name string,
		optionalGenericParameters GenericParametersLike,
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructor
	Make(
		value ValueLike,
		additionalValues abs.Sequential[AdditionalValueLike],
	) EnumerationLike
}

/*
FunctionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructor
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		result ResultLike,
	) FunctionLike
}

/*
FunctionMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-methods-like class.
*/
type FunctionMethodsClassLike interface {
	// Constructor
	Make(
		functions abs.Sequential[FunctionLike],
	) FunctionMethodsLike
}

/*
FunctionalClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-like class.
*/
type FunctionalClassLike interface {
	// Constructor
	Make(
		declaration DeclarationLike,
		parameters abs.Sequential[ParameterLike],
		result ResultLike,
	) FunctionalLike
}

/*
FunctionalDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-definitions-like class.
*/
type FunctionalDefinitionsClassLike interface {
	// Constructor
	Make(
		functionals abs.Sequential[FunctionalLike],
	) FunctionalDefinitionsLike
}

/*
GenericArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete generic-arguments-like class.
*/
type GenericArgumentsClassLike interface {
	// Constructor
	Make(
		argument ArgumentLike,
		additionalArguments abs.Sequential[AdditionalArgumentLike],
	) GenericArgumentsLike
}

/*
GenericParametersClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete generic-parameters-like class.
*/
type GenericParametersClassLike interface {
	// Constructor
	Make(
		parameters abs.Sequential[ParameterLike],
	) GenericParametersLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructor
	Make(
		comment string,
		name string,
	) HeaderLike
}

/*
ImportsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete imports-like class.
*/
type ImportsClassLike interface {
	// Constructor
	Make(
		modules abs.Sequential[ModuleLike],
	) ImportsLike
}

/*
InstanceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-like class.
*/
type InstanceClassLike interface {
	// Constructor
	Make(
		declaration DeclarationLike,
		instanceMethods InstanceMethodsLike,
	) InstanceLike
}

/*
InstanceDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-definitions-like class.
*/
type InstanceDefinitionsClassLike interface {
	// Constructor
	Make(
		instances abs.Sequential[InstanceLike],
	) InstanceDefinitionsLike
}

/*
InstanceMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-methods-like class.
*/
type InstanceMethodsClassLike interface {
	// Constructor
	Make(
		publicMethods PublicMethodsLike,
		optionalAttributeMethods AttributeMethodsLike,
		optionalAspectInterfaces AspectInterfacesLike,
	) InstanceMethodsLike
}

/*
InterfaceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete interface-like class.
*/
type InterfaceClassLike interface {
	// Constructor
	Make(
		name string,
		optionalSuffix SuffixLike,
		optionalGenericArguments GenericArgumentsLike,
	) InterfaceLike
}

/*
InterfaceDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete interface-definitions-like class.
*/
type InterfaceDefinitionsClassLike interface {
	// Constructor
	Make(
		classDefinitions ClassDefinitionsLike,
		instanceDefinitions InstanceDefinitionsLike,
		optionalAspectDefinitions AspectDefinitionsLike,
	) InterfaceDefinitionsLike
}

/*
MapClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete map-like class.
*/
type MapClassLike interface {
	// Constructor
	Make(
		name string,
	) MapLike
}

/*
MethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete method-like class.
*/
type MethodClassLike interface {
	// Constructor
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		optionalResult ResultLike,
	) MethodLike
}

/*
ModelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete model-like class.
*/
type ModelClassLike interface {
	// Constructor
	Make(
		moduleDefinition ModuleDefinitionLike,
		primitiveDefinitions PrimitiveDefinitionsLike,
		interfaceDefinitions InterfaceDefinitionsLike,
	) ModelLike
}

/*
ModuleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete module-like class.
*/
type ModuleClassLike interface {
	// Constructor
	Make(
		name string,
		path string,
	) ModuleLike
}

/*
ModuleDefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete module-definition-like class.
*/
type ModuleDefinitionClassLike interface {
	// Constructor
	Make(
		notice NoticeLike,
		header HeaderLike,
		optionalImports ImportsLike,
	) ModuleDefinitionLike
}

/*
NoneClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete none-like class.
*/
type NoneClassLike interface {
	// Constructor
	Make(
		newline string,
	) NoneLike
}

/*
NoticeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete notice-like class.
*/
type NoticeClassLike interface {
	// Constructor
	Make(
		comment string,
	) NoticeLike
}

/*
ParameterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructor
	Make(
		name string,
		abstraction AbstractionLike,
	) ParameterLike
}

/*
ParameterizedClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameterized-like class.
*/
type ParameterizedClassLike interface {
	// Constructor
	Make(
		parameters abs.Sequential[ParameterLike],
	) ParameterizedLike
}

/*
PrefixClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete prefix-like class.
*/
type PrefixClassLike interface {
	// Constructor
	Make(
		any_ any,
	) PrefixLike
}

/*
PrimitiveDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete primitive-definitions-like class.
*/
type PrimitiveDefinitionsClassLike interface {
	// Constructor
	Make(
		optionalTypeDefinitions TypeDefinitionsLike,
		optionalFunctionalDefinitions FunctionalDefinitionsLike,
	) PrimitiveDefinitionsLike
}

/*
PublicMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete public-methods-like class.
*/
type PublicMethodsClassLike interface {
	// Constructor
	Make(
		methods abs.Sequential[MethodLike],
	) PublicMethodsLike
}

/*
ResultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete result-like class.
*/
type ResultClassLike interface {
	// Constructor
	Make(
		any_ any,
	) ResultLike
}

/*
SuffixClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete suffix-like class.
*/
type SuffixClassLike interface {
	// Constructor
	Make(
		name string,
	) SuffixLike
}

/*
TypeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-like class.
*/
type TypeClassLike interface {
	// Constructor
	Make(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		optionalEnumeration EnumerationLike,
	) TypeLike
}

/*
TypeDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-definitions-like class.
*/
type TypeDefinitionsClassLike interface {
	// Constructor
	Make(
		types abs.Sequential[TypeLike],
	) TypeDefinitionsLike
}

/*
ValueClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete value-like class.
*/
type ValueClassLike interface {
	// Constructor
	Make(
		name string,
		abstraction AbstractionLike,
	) ValueLike
}

// Instances

/*
AbstractionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Public
	GetClass() AbstractionClassLike

	// Attribute
	GetOptionalPrefix() PrefixLike
	GetName() string
	GetOptionalSuffix() SuffixLike
	GetOptionalGenericArguments() GenericArgumentsLike
}

/*
AdditionalArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-argument-like class.
*/
type AdditionalArgumentLike interface {
	// Public
	GetClass() AdditionalArgumentClassLike

	// Attribute
	GetArgument() ArgumentLike
}

/*
AdditionalValueLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-value-like class.
*/
type AdditionalValueLike interface {
	// Public
	GetClass() AdditionalValueClassLike

	// Attribute
	GetName() string
}

/*
ArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Public
	GetClass() ArgumentClassLike

	// Attribute
	GetAbstraction() AbstractionLike
}

/*
ArrayLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete array-like class.
*/
type ArrayLike interface {
	// Public
	GetClass() ArrayClassLike
}

/*
AspectLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-like class.
*/
type AspectLike interface {
	// Public
	GetClass() AspectClassLike

	// Attribute
	GetDeclaration() DeclarationLike
	GetMethods() abs.Sequential[MethodLike]
}

/*
AspectDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-definitions-like class.
*/
type AspectDefinitionsLike interface {
	// Public
	GetClass() AspectDefinitionsClassLike

	// Attribute
	GetAspects() abs.Sequential[AspectLike]
}

/*
AspectInterfacesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-interfaces-like class.
*/
type AspectInterfacesLike interface {
	// Public
	GetClass() AspectInterfacesClassLike

	// Attribute
	GetInterfaces() abs.Sequential[InterfaceLike]
}

/*
AttributeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-like class.
*/
type AttributeLike interface {
	// Public
	GetClass() AttributeClassLike

	// Attribute
	GetName() string
	GetOptionalParameter() ParameterLike
	GetOptionalAbstraction() AbstractionLike
}

/*
AttributeMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-methods-like class.
*/
type AttributeMethodsLike interface {
	// Public
	GetClass() AttributeMethodsClassLike

	// Attribute
	GetAttributes() abs.Sequential[AttributeLike]
}

/*
ChannelLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete channel-like class.
*/
type ChannelLike interface {
	// Public
	GetClass() ChannelClassLike
}

/*
ClassLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-like class.
*/
type ClassLike interface {
	// Public
	GetClass() ClassClassLike

	// Attribute
	GetDeclaration() DeclarationLike
	GetClassMethods() ClassMethodsLike
}

/*
ClassDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-definitions-like class.
*/
type ClassDefinitionsLike interface {
	// Public
	GetClass() ClassDefinitionsClassLike

	// Attribute
	GetClasses() abs.Sequential[ClassLike]
}

/*
ClassMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-methods-like class.
*/
type ClassMethodsLike interface {
	// Public
	GetClass() ClassMethodsClassLike

	// Attribute
	GetConstructorMethods() ConstructorMethodsLike
	GetOptionalConstantMethods() ConstantMethodsLike
	GetOptionalFunctionMethods() FunctionMethodsLike
}

/*
ConstantLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-like class.
*/
type ConstantLike interface {
	// Public
	GetClass() ConstantClassLike

	// Attribute
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ConstantMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-methods-like class.
*/
type ConstantMethodsLike interface {
	// Public
	GetClass() ConstantMethodsClassLike

	// Attribute
	GetConstants() abs.Sequential[ConstantLike]
}

/*
ConstructorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-like class.
*/
type ConstructorLike interface {
	// Public
	GetClass() ConstructorClassLike

	// Attribute
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetAbstraction() AbstractionLike
}

/*
ConstructorMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-methods-like class.
*/
type ConstructorMethodsLike interface {
	// Public
	GetClass() ConstructorMethodsClassLike

	// Attribute
	GetConstructors() abs.Sequential[ConstructorLike]
}

/*
DeclarationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Public
	GetClass() DeclarationClassLike

	// Attribute
	GetComment() string
	GetName() string
	GetOptionalGenericParameters() GenericParametersLike
}

/*
EnumerationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Public
	GetClass() EnumerationClassLike

	// Attribute
	GetValue() ValueLike
	GetAdditionalValues() abs.Sequential[AdditionalValueLike]
}

/*
FunctionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Public
	GetClass() FunctionClassLike

	// Attribute
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-methods-like class.
*/
type FunctionMethodsLike interface {
	// Public
	GetClass() FunctionMethodsClassLike

	// Attribute
	GetFunctions() abs.Sequential[FunctionLike]
}

/*
FunctionalLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-like class.
*/
type FunctionalLike interface {
	// Public
	GetClass() FunctionalClassLike

	// Attribute
	GetDeclaration() DeclarationLike
	GetParameters() abs.Sequential[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionalDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-definitions-like class.
*/
type FunctionalDefinitionsLike interface {
	// Public
	GetClass() FunctionalDefinitionsClassLike

	// Attribute
	GetFunctionals() abs.Sequential[FunctionalLike]
}

/*
GenericArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete generic-arguments-like class.
*/
type GenericArgumentsLike interface {
	// Public
	GetClass() GenericArgumentsClassLike

	// Attribute
	GetArgument() ArgumentLike
	GetAdditionalArguments() abs.Sequential[AdditionalArgumentLike]
}

/*
GenericParametersLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete generic-parameters-like class.
*/
type GenericParametersLike interface {
	// Public
	GetClass() GenericParametersClassLike

	// Attribute
	GetParameters() abs.Sequential[ParameterLike]
}

/*
HeaderLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete header-like class.
*/
type HeaderLike interface {
	// Public
	GetClass() HeaderClassLike

	// Attribute
	GetComment() string
	GetName() string
}

/*
ImportsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete imports-like class.
*/
type ImportsLike interface {
	// Public
	GetClass() ImportsClassLike

	// Attribute
	GetModules() abs.Sequential[ModuleLike]
}

/*
InstanceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-like class.
*/
type InstanceLike interface {
	// Public
	GetClass() InstanceClassLike

	// Attribute
	GetDeclaration() DeclarationLike
	GetInstanceMethods() InstanceMethodsLike
}

/*
InstanceDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-definitions-like class.
*/
type InstanceDefinitionsLike interface {
	// Public
	GetClass() InstanceDefinitionsClassLike

	// Attribute
	GetInstances() abs.Sequential[InstanceLike]
}

/*
InstanceMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-methods-like class.
*/
type InstanceMethodsLike interface {
	// Public
	GetClass() InstanceMethodsClassLike

	// Attribute
	GetPublicMethods() PublicMethodsLike
	GetOptionalAttributeMethods() AttributeMethodsLike
	GetOptionalAspectInterfaces() AspectInterfacesLike
}

/*
InterfaceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete interface-like class.
*/
type InterfaceLike interface {
	// Public
	GetClass() InterfaceClassLike

	// Attribute
	GetName() string
	GetOptionalSuffix() SuffixLike
	GetOptionalGenericArguments() GenericArgumentsLike
}

/*
InterfaceDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete interface-definitions-like class.
*/
type InterfaceDefinitionsLike interface {
	// Public
	GetClass() InterfaceDefinitionsClassLike

	// Attribute
	GetClassDefinitions() ClassDefinitionsLike
	GetInstanceDefinitions() InstanceDefinitionsLike
	GetOptionalAspectDefinitions() AspectDefinitionsLike
}

/*
MapLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete map-like class.
*/
type MapLike interface {
	// Public
	GetClass() MapClassLike

	// Attribute
	GetName() string
}

/*
MethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete method-like class.
*/
type MethodLike interface {
	// Public
	GetClass() MethodClassLike

	// Attribute
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetOptionalResult() ResultLike
}

/*
ModelLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete model-like class.
*/
type ModelLike interface {
	// Public
	GetClass() ModelClassLike

	// Attribute
	GetModuleDefinition() ModuleDefinitionLike
	GetPrimitiveDefinitions() PrimitiveDefinitionsLike
	GetInterfaceDefinitions() InterfaceDefinitionsLike
}

/*
ModuleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete module-like class.
*/
type ModuleLike interface {
	// Public
	GetClass() ModuleClassLike

	// Attribute
	GetName() string
	GetPath() string
}

/*
ModuleDefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete module-definition-like class.
*/
type ModuleDefinitionLike interface {
	// Public
	GetClass() ModuleDefinitionClassLike

	// Attribute
	GetNotice() NoticeLike
	GetHeader() HeaderLike
	GetOptionalImports() ImportsLike
}

/*
NoneLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete none-like class.
*/
type NoneLike interface {
	// Public
	GetClass() NoneClassLike

	// Attribute
	GetNewline() string
}

/*
NoticeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete notice-like class.
*/
type NoticeLike interface {
	// Public
	GetClass() NoticeClassLike

	// Attribute
	GetComment() string
}

/*
ParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Public
	GetClass() ParameterClassLike

	// Attribute
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ParameterizedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameterized-like class.
*/
type ParameterizedLike interface {
	// Public
	GetClass() ParameterizedClassLike

	// Attribute
	GetParameters() abs.Sequential[ParameterLike]
}

/*
PrefixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Public
	GetClass() PrefixClassLike

	// Attribute
	GetAny() any
}

/*
PrimitiveDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete primitive-definitions-like class.
*/
type PrimitiveDefinitionsLike interface {
	// Public
	GetClass() PrimitiveDefinitionsClassLike

	// Attribute
	GetOptionalTypeDefinitions() TypeDefinitionsLike
	GetOptionalFunctionalDefinitions() FunctionalDefinitionsLike
}

/*
PublicMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete public-methods-like class.
*/
type PublicMethodsLike interface {
	// Public
	GetClass() PublicMethodsClassLike

	// Attribute
	GetMethods() abs.Sequential[MethodLike]
}

/*
ResultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Public
	GetClass() ResultClassLike

	// Attribute
	GetAny() any
}

/*
SuffixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete suffix-like class.
*/
type SuffixLike interface {
	// Public
	GetClass() SuffixClassLike

	// Attribute
	GetName() string
}

/*
TypeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete type-like class.
*/
type TypeLike interface {
	// Public
	GetClass() TypeClassLike

	// Attribute
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetOptionalEnumeration() EnumerationLike
}

/*
TypeDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete type-definitions-like class.
*/
type TypeDefinitionsLike interface {
	// Public
	GetClass() TypeDefinitionsClassLike

	// Attribute
	GetTypes() abs.Sequential[TypeLike]
}

/*
ValueLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete value-like class.
*/
type ValueLike interface {
	// Public
	GetClass() ValueClassLike

	// Attribute
	GetName() string
	GetAbstraction() AbstractionLike
}
