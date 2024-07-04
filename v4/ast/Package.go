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
Package "ast" provides the ability to generate Go class files based on a
Go Package.go file that follows the format shown in the following code template:
  - https://github.com/craterdog/go-model-framework/blob/main/models/Package.go

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-model-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	col "github.com/craterdog/go-collection-framework/v4/collection"
)

// Classes

/*
AbstractionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructors
	Make(
		prefix PrefixLike,
		name string,
		genericArguments GenericArgumentsLike,
	) AbstractionLike
}

/*
AbstractionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstractions-like class.
*/
type AbstractionsClassLike interface {
	// Constructors
	Make(
		note string,
		abstractions col.Sequential[AbstractionLike],
	) AbstractionsLike
}

/*
AdditionalArgumentClassLike is a class interface that defines the complete set
of class constants, constructors and functions that must be supported by each
concrete additional-argument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructors
	Make(argument ArgumentLike) AdditionalArgumentLike
}

/*
AdditionalParameterClassLike is a class interface that defines the complete set
of class constants, constructors and functions that must be supported by each
concrete additional-parameter-like class.
*/
type AdditionalParameterClassLike interface {
	// Constructors
	Make(parameter ParameterLike) AdditionalParameterLike
}

/*
AdditionalValueClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-value-like class.
*/
type AdditionalValueClassLike interface {
	// Constructors
	Make(name string) AdditionalValueLike
}

/*
AliasClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete alias-like class.
*/
type AliasClassLike interface {
	// Constructors
	Make(name string) AliasLike
}

/*
ArgumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructors
	Make(abstraction AbstractionLike) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructors
	Make(
		argument ArgumentLike,
		additionalArguments col.Sequential[AdditionalArgumentLike],
	) ArgumentsLike
}

/*
ArrayClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike interface {
	// Constructors
	Make() ArrayLike
}

/*
AspectClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-like class.
*/
type AspectClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		methods MethodsLike,
	) AspectLike
}

/*
AspectsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspects-like class.
*/
type AspectsClassLike interface {
	// Constructors
	Make(
		note string,
		aspects col.Sequential[AspectLike],
	) AspectsLike
}

/*
AttributeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructors
	Make(
		name string,
		parameter ParameterLike,
		abstraction AbstractionLike,
	) AttributeLike
}

/*
AttributesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attributes-like class.
*/
type AttributesClassLike interface {
	// Constructors
	Make(
		note string,
		attributes col.Sequential[AttributeLike],
	) AttributesLike
}

/*
ChannelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete channel-like class.
*/
type ChannelClassLike interface {
	// Constructors
	Make() ChannelLike
}

/*
ClassClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-like class.
*/
type ClassClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		constructors ConstructorsLike,
		constants ConstantsLike,
		functions FunctionsLike,
	) ClassLike
}

/*
ClassesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete classes-like class.
*/
type ClassesClassLike interface {
	// Constructors
	Make(
		note string,
		classes col.Sequential[ClassLike],
	) ClassesLike
}

/*
ConstantClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-like class.
*/
type ConstantClassLike interface {
	// Constructors
	Make(
		name string,
		abstraction AbstractionLike,
	) ConstantLike
}

/*
ConstantsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constants-like class.
*/
type ConstantsClassLike interface {
	// Constructors
	Make(
		note string,
		constants col.Sequential[ConstantLike],
	) ConstantsLike
}

/*
ConstructorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-like class.
*/
type ConstructorClassLike interface {
	// Constructors
	Make(
		name string,
		parameters ParametersLike,
		abstraction AbstractionLike,
	) ConstructorLike
}

/*
ConstructorsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructors-like class.
*/
type ConstructorsClassLike interface {
	// Constructors
	Make(
		note string,
		constructors col.Sequential[ConstructorLike],
	) ConstructorsLike
}

/*
DeclarationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructors
	Make(
		comment string,
		name string,
		genericParameters GenericParametersLike,
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructors
	Make(values ValuesLike) EnumerationLike
}

/*
FunctionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructors
	Make(
		name string,
		parameters ParametersLike,
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
	Make(
		declaration DeclarationLike,
		parameters ParametersLike,
		result ResultLike,
	) FunctionalLike
}

/*
FunctionalsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functionals-like class.
*/
type FunctionalsClassLike interface {
	// Constructors
	Make(
		note string,
		functionals col.Sequential[FunctionalLike],
	) FunctionalsLike
}

/*
FunctionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functions-like class.
*/
type FunctionsClassLike interface {
	// Constructors
	Make(
		note string,
		functions col.Sequential[FunctionLike],
	) FunctionsLike
}

/*
GenericArgumentsClassLike is a class interface that defines the complete set
of class constants, constructors and functions that must be supported by each
concrete genericarguments-like class.
*/
type GenericArgumentsClassLike interface {
	// Constructors
	Make(arguments ArgumentsLike) GenericArgumentsLike
}

/*
GenericParametersClassLike is a class interface that defines the complete set
of class constants, constructors and functions that must be supported by each
concrete genericparameters-like class.
*/
type GenericParametersClassLike interface {
	// Constructors
	Make(parameters ParametersLike) GenericParametersLike
}

/*
HeaderClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	Make(
		comment string,
		name string,
	) HeaderLike
}

/*
ImportsClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
imports-like class.
*/
type ImportsClassLike interface {
	// Constructors
	Make(modules ModulesLike) ImportsLike
}

/*
InstanceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-like class.
*/
type InstanceClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		attributes AttributesLike,
		abstractions AbstractionsLike,
		methods MethodsLike,
	) InstanceLike
}

/*
InstancesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instances-like class.
*/
type InstancesClassLike interface {
	// Constructors
	Make(
		note string,
		instances col.Sequential[InstanceLike],
	) InstancesLike
}

/*
MapClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
map-like class.
*/
type MapClassLike interface {
	// Constructors
	Make(name string) MapLike
}

/*
MethodClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
method-like class.
*/
type MethodClassLike interface {
	// Constructors
	Make(
		name string,
		parameters ParametersLike,
		result ResultLike,
	) MethodLike
}

/*
MethodsClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
methods-like class.
*/
type MethodsClassLike interface {
	// Constructors
	Make(
		note string,
		methods col.Sequential[MethodLike],
	) MethodsLike
}

/*
ModelClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
model-like class.
*/
type ModelClassLike interface {
	// Constructors
	Make(
		notice NoticeLike,
		header HeaderLike,
		imports ImportsLike,
		types TypesLike,
		functionals FunctionalsLike,
		classes ClassesLike,
		instances InstancesLike,
		aspects AspectsLike,
	) ModelLike
}

/*
ModuleClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
module-like class.
*/
type ModuleClassLike interface {
	// Constructors
	Make(
		name string,
		path string,
	) ModuleLike
}

/*
ModulesClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
modules-like class.
*/
type ModulesClassLike interface {
	// Constructors
	Make(modules col.Sequential[ModuleLike]) ModulesLike
}

/*
NoticeClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
notice-like class.
*/
type NoticeClassLike interface {
	// Constructors
	Make(comment string) NoticeLike
}

/*
ParameterClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
parameter-like class.
*/
type ParameterClassLike interface {
	// Constructors
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
	// Constructors
	Make(parameters ParametersLike) ParameterizedLike
}

/*
ParametersClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
parameters-like class.
*/
type ParametersClassLike interface {
	// Constructors
	Make(
		parameter ParameterLike,
		additionalParameters col.Sequential[AdditionalParameterLike],
	) ParametersLike
}

/*
PrefixClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
prefix-like class.
*/
type PrefixClassLike interface {
	// Constructors
	Make(any any) PrefixLike
}

/*
ResultClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
result-like class.
*/
type ResultClassLike interface {
	// Constructors
	Make(any any) ResultLike
}

/*
TypeClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
type-like class.
*/
type TypeClassLike interface {
	// Constructors
	Make(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		enumeration EnumerationLike,
	) TypeLike
}

/*
TypesClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
types-like class.
*/
type TypesClassLike interface {
	// Constructors
	Make(
		note string,
		types col.Sequential[TypeLike],
	) TypesLike
}

/*
ValueClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
value-like class.
*/
type ValueClassLike interface {
	// Constructors
	Make(
		name string,
		abstraction AbstractionLike,
	) ValueLike
}

/*
ValuesClassLike is a class interface that defines the complete set of class
constants, constructors and functions that must be supported by each concrete
values-like class.
*/
type ValuesClassLike interface {
	// Constructors
	Make(
		value ValueLike,
		additionalValues col.Sequential[AdditionalValueLike],
	) ValuesLike
}

// Instances

/*
AbstractionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Attributes
	GetClass() AbstractionClassLike
	GetPrefix() PrefixLike
	GetName() string
	GetGenericArguments() GenericArgumentsLike
}

/*
AbstractionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstractions-like class.
*/
type AbstractionsLike interface {
	// Attributes
	GetClass() AbstractionsClassLike
	GetNote() string
	GetAbstractions() col.Sequential[AbstractionLike]
}

/*
AdditionalArgumentLike is an instance interface that defines the complete set
of instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-argument-like class.
*/
type AdditionalArgumentLike interface {
	// Attributes
	GetClass() AdditionalArgumentClassLike
	GetArgument() ArgumentLike
}

/*
AdditionalParameterLike is an instance interface that defines the complete set
of instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-parameter-like class.
*/
type AdditionalParameterLike interface {
	// Attributes
	GetClass() AdditionalParameterClassLike
	GetParameter() ParameterLike
}

/*
AdditionalValueLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-value-like class.
*/
type AdditionalValueLike interface {
	// Attributes
	GetClass() AdditionalValueClassLike
	GetName() string
}

/*
AliasLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete alias-like class.
*/
type AliasLike interface {
	// Attributes
	GetClass() AliasClassLike
	GetName() string
}

/*
ArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Attributes
	GetClass() ArgumentClassLike
	GetAbstraction() AbstractionLike
}

/*
ArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Attributes
	GetClass() ArgumentsClassLike
	GetArgument() ArgumentLike
	GetAdditionalArguments() col.Sequential[AdditionalArgumentLike]
}

/*
ArrayLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete array-like class.
*/
type ArrayLike interface {
	// Attributes
	GetClass() ArrayClassLike
}

/*
AspectLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete aspect-like class.
*/
type AspectLike interface {
	// Attributes
	GetClass() AspectClassLike
	GetDeclaration() DeclarationLike
	GetMethods() MethodsLike
}

/*
AspectsLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete aspects-like class.
*/
type AspectsLike interface {
	// Attributes
	GetClass() AspectsClassLike
	GetNote() string
	GetAspects() col.Sequential[AspectLike]
}

/*
AttributeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-like class.
*/
type AttributeLike interface {
	// Attributes
	GetClass() AttributeClassLike
	GetName() string
	GetParameter() ParameterLike
	GetAbstraction() AbstractionLike
}

/*
AttributesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attributes-like class.
*/
type AttributesLike interface {
	// Attributes
	GetClass() AttributesClassLike
	GetNote() string
	GetAttributes() col.Sequential[AttributeLike]
}

/*
ChannelLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete channel-like class.
*/
type ChannelLike interface {
	// Attributes
	GetClass() ChannelClassLike
}

/*
ClassLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete class-like class.
*/
type ClassLike interface {
	// Attributes
	GetClass() ClassClassLike
	GetDeclaration() DeclarationLike
	GetConstructors() ConstructorsLike
	GetConstants() ConstantsLike
	GetFunctions() FunctionsLike
}

/*
ClassesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete classes-like class.
*/
type ClassesLike interface {
	// Attributes
	GetClass() ClassesClassLike
	GetNote() string
	GetClasses() col.Sequential[ClassLike]
}

/*
ConstantLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-like class.
*/
type ConstantLike interface {
	// Attributes
	GetClass() ConstantClassLike
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ConstantsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constants-like class.
*/
type ConstantsLike interface {
	// Attributes
	GetClass() ConstantsClassLike
	GetNote() string
	GetConstants() col.Sequential[ConstantLike]
}

/*
ConstructorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-like class.
*/
type ConstructorLike interface {
	// Attributes
	GetClass() ConstructorClassLike
	GetName() string
	GetParameters() ParametersLike
	GetAbstraction() AbstractionLike
}

/*
ConstructorsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructors-like class.
*/
type ConstructorsLike interface {
	// Attributes
	GetClass() ConstructorsClassLike
	GetNote() string
	GetConstructors() col.Sequential[ConstructorLike]
}

/*
DeclarationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Attributes
	GetClass() DeclarationClassLike
	GetComment() string
	GetName() string
	GetGenericParameters() GenericParametersLike
}

/*
EnumerationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Attributes
	GetClass() EnumerationClassLike
	GetValues() ValuesLike
}

/*
FunctionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Attributes
	GetClass() FunctionClassLike
	GetName() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
FunctionalLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-like class.
*/
type FunctionalLike interface {
	// Attributes
	GetClass() FunctionalClassLike
	GetDeclaration() DeclarationLike
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
FunctionalsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functionals-like class.
*/
type FunctionalsLike interface {
	// Attributes
	GetClass() FunctionalsClassLike
	GetNote() string
	GetFunctionals() col.Sequential[FunctionalLike]
}

/*
FunctionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by
each instance of a concrete functions-like class.
*/
type FunctionsLike interface {
	// Attributes
	GetClass() FunctionsClassLike
	GetNote() string
	GetFunctions() col.Sequential[FunctionLike]
}

/*
GenericArgumentsLike is an instance interface that defines the complete set
of instance attributes, abstractions and methods that must be supported by
each instance of a concrete genericarguments-like class.
*/
type GenericArgumentsLike interface {
	// Attributes
	GetClass() GenericArgumentsClassLike
	GetArguments() ArgumentsLike
}

/*
GenericParametersLike is an instance interface that defines the complete set
of instance attributes, abstractions and methods that must be supported by
each instance of a concrete genericparameters-like class.
*/
type GenericParametersLike interface {
	// Attributes
	GetClass() GenericParametersClassLike
	GetParameters() ParametersLike
}

/*
HeaderLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete header-like class.
*/
type HeaderLike interface {
	// Attributes
	GetClass() HeaderClassLike
	GetComment() string
	GetName() string
}

/*
ImportsLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete imports-like class.
*/
type ImportsLike interface {
	// Attributes
	GetClass() ImportsClassLike
	GetModules() ModulesLike
}

/*
InstanceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-like class.
*/
type InstanceLike interface {
	// Attributes
	GetClass() InstanceClassLike
	GetDeclaration() DeclarationLike
	GetAttributes() AttributesLike
	GetAbstractions() AbstractionsLike
	GetMethods() MethodsLike
}

/*
InstancesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instances-like class.
*/
type InstancesLike interface {
	// Attributes
	GetClass() InstancesClassLike
	GetNote() string
	GetInstances() col.Sequential[InstanceLike]
}

/*
MapLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete map-like class.
*/
type MapLike interface {
	// Attributes
	GetClass() MapClassLike
	GetName() string
}

/*
MethodLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete method-like class.
*/
type MethodLike interface {
	// Attributes
	GetClass() MethodClassLike
	GetName() string
	GetParameters() ParametersLike
	GetResult() ResultLike
}

/*
MethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by
each instance of a concrete methods-like class.
*/
type MethodsLike interface {
	// Attributes
	GetClass() MethodsClassLike
	GetNote() string
	GetMethods() col.Sequential[MethodLike]
}

/*
ModelLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete model-like class.
*/
type ModelLike interface {
	// Attributes
	GetClass() ModelClassLike
	GetNotice() NoticeLike
	GetHeader() HeaderLike
	GetImports() ImportsLike
	GetTypes() TypesLike
	GetFunctionals() FunctionalsLike
	GetClasses() ClassesLike
	GetInstances() InstancesLike
	GetAspects() AspectsLike
}

/*
ModuleLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete module-like class.
*/
type ModuleLike interface {
	// Attributes
	GetClass() ModuleClassLike
	GetName() string
	GetPath() string
}

/*
ModulesLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete modules-like class.
*/
type ModulesLike interface {
	// Attributes
	GetClass() ModulesClassLike
	GetModules() col.Sequential[ModuleLike]
}

/*
NoticeLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete notice-like class.
*/
type NoticeLike interface {
	// Attributes
	GetClass() NoticeClassLike
	GetComment() string
}

/*
ParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by
each instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Attributes
	GetClass() ParameterClassLike
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ParameterizedLike is an instance interface that defines the complete set
of instance attributes, abstractions and methods that must be supported by
each instance of a concrete parameterized-like class.
*/
type ParameterizedLike interface {
	// Attributes
	GetClass() ParameterizedClassLike
	GetParameters() ParametersLike
}

/*
ParametersLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Attributes
	GetClass() ParametersClassLike
	GetParameter() ParameterLike
	GetAdditionalParameters() col.Sequential[AdditionalParameterLike]
}

/*
PrefixLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Attributes
	GetClass() PrefixClassLike
	GetAny() any
}

/*
ResultLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete result-like class.
*/
type ResultLike interface {
	// Attributes
	GetClass() ResultClassLike
	GetAny() any
}

/*
TypeLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete type-like class.
*/
type TypeLike interface {
	// Attributes
	GetClass() TypeClassLike
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetEnumeration() EnumerationLike
}

/*
TypesLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete types-like class.
*/
type TypesLike interface {
	// Attributes
	GetClass() TypesClassLike
	GetNote() string
	GetTypes() col.Sequential[TypeLike]
}

/*
ValueLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete value-like class.
*/
type ValueLike interface {
	// Attributes
	GetClass() ValueClassLike
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ValuesLike is an instance interface that defines the complete set of instance
attributes, abstractions and methods that must be supported by each instance
of a concrete values-like class.
*/
type ValuesLike interface {
	// Attributes
	GetClass() ValuesClassLike
	GetValue() ValueLike
	GetAdditionalValues() col.Sequential[AdditionalValueLike]
}
