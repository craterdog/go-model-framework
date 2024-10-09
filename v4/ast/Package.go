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

// Class Definitions

/*
AbstractionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete abstraction-like class.
*/
type AbstractionClassLike interface {
	// Constructor Methods
	Make(
		optionalPrefix PrefixLike,
		name string,
		optionalSuffix SuffixLike,
		optionalArguments ArgumentsLike,
	) AbstractionLike
}

/*
AdditionalArgumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-argument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructor Methods
	Make(
		argument ArgumentLike,
	) AdditionalArgumentLike
}

/*
AdditionalConstraintClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-constraint-like class.
*/
type AdditionalConstraintClassLike interface {
	// Constructor Methods
	Make(
		constraint ConstraintLike,
	) AdditionalConstraintLike
}

/*
AdditionalValueClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete additional-value-like class.
*/
type AdditionalValueClassLike interface {
	// Constructor Methods
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
	// Constructor Methods
	Make(
		abstraction AbstractionLike,
	) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructor Methods
	Make(
		argument ArgumentLike,
		additionalArguments abs.Sequential[AdditionalArgumentLike],
	) ArgumentsLike
}

/*
ArrayClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike interface {
	// Constructor Methods
	Make() ArrayLike
}

/*
AspectDefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-definition-like class.
*/
type AspectDefinitionClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		aspectMethods abs.Sequential[AspectMethodLike],
	) AspectDefinitionLike
}

/*
AspectInterfaceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-interface-like class.
*/
type AspectInterfaceClassLike interface {
	// Constructor Methods
	Make(
		abstraction AbstractionLike,
	) AspectInterfaceLike
}

/*
AspectMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-method-like class.
*/
type AspectMethodClassLike interface {
	// Constructor Methods
	Make(
		method MethodLike,
	) AspectMethodLike
}

/*
AspectSectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-section-like class.
*/
type AspectSectionClassLike interface {
	// Constructor Methods
	Make(
		aspectDefinitions abs.Sequential[AspectDefinitionLike],
	) AspectSectionLike
}

/*
AspectSubsectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete aspect-subsection-like class.
*/
type AspectSubsectionClassLike interface {
	// Constructor Methods
	Make(
		aspectInterfaces abs.Sequential[AspectInterfaceLike],
	) AspectSubsectionLike
}

/*
AttributeMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-method-like class.
*/
type AttributeMethodClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) AttributeMethodLike
}

/*
AttributeSubsectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-subsection-like class.
*/
type AttributeSubsectionClassLike interface {
	// Constructor Methods
	Make(
		attributeMethods abs.Sequential[AttributeMethodLike],
	) AttributeSubsectionLike
}

/*
ChannelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete channel-like class.
*/
type ChannelClassLike interface {
	// Constructor Methods
	Make() ChannelLike
}

/*
ClassDefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-definition-like class.
*/
type ClassDefinitionClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		classMethods ClassMethodsLike,
	) ClassDefinitionLike
}

/*
ClassMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-methods-like class.
*/
type ClassMethodsClassLike interface {
	// Constructor Methods
	Make(
		constructorSubsection ConstructorSubsectionLike,
		optionalConstantSubsection ConstantSubsectionLike,
		optionalFunctionSubsection FunctionSubsectionLike,
	) ClassMethodsLike
}

/*
ClassSectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete class-section-like class.
*/
type ClassSectionClassLike interface {
	// Constructor Methods
	Make(
		classDefinitions abs.Sequential[ClassDefinitionLike],
	) ClassSectionLike
}

/*
ConstantMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-method-like class.
*/
type ConstantMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) ConstantMethodLike
}

/*
ConstantSubsectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constant-subsection-like class.
*/
type ConstantSubsectionClassLike interface {
	// Constructor Methods
	Make(
		constantMethods abs.Sequential[ConstantMethodLike],
	) ConstantSubsectionLike
}

/*
ConstraintClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constraint-like class.
*/
type ConstraintClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) ConstraintLike
}

/*
ConstraintsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constraints-like class.
*/
type ConstraintsClassLike interface {
	// Constructor Methods
	Make(
		constraint ConstraintLike,
		additionalConstraints abs.Sequential[AdditionalConstraintLike],
	) ConstraintsLike
}

/*
ConstructorMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-method-like class.
*/
type ConstructorMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		abstraction AbstractionLike,
	) ConstructorMethodLike
}

/*
ConstructorSubsectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete constructor-subsection-like class.
*/
type ConstructorSubsectionClassLike interface {
	// Constructor Methods
	Make(
		constructorMethods abs.Sequential[ConstructorMethodLike],
	) ConstructorSubsectionLike
}

/*
DeclarationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete declaration-like class.
*/
type DeclarationClassLike interface {
	// Constructor Methods
	Make(
		comment string,
		name string,
		optionalConstraints ConstraintsLike,
	) DeclarationLike
}

/*
EnumerationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete enumeration-like class.
*/
type EnumerationClassLike interface {
	// Constructor Methods
	Make(
		value ValueLike,
		additionalValues abs.Sequential[AdditionalValueLike],
	) EnumerationLike
}

/*
FunctionMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-method-like class.
*/
type FunctionMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		parameters abs.Sequential[ParameterLike],
		result ResultLike,
	) FunctionMethodLike
}

/*
FunctionSubsectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-subsection-like class.
*/
type FunctionSubsectionClassLike interface {
	// Constructor Methods
	Make(
		functionMethods abs.Sequential[FunctionMethodLike],
	) FunctionSubsectionLike
}

/*
FunctionalDefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-definition-like class.
*/
type FunctionalDefinitionClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		parameters abs.Sequential[ParameterLike],
		result ResultLike,
	) FunctionalDefinitionLike
}

/*
FunctionalSectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete functional-section-like class.
*/
type FunctionalSectionClassLike interface {
	// Constructor Methods
	Make(
		functionalDefinitions abs.Sequential[FunctionalDefinitionLike],
	) FunctionalSectionLike
}

/*
GetterMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete getter-method-like class.
*/
type GetterMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) GetterMethodLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructor Methods
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
	// Constructor Methods
	Make(
		modules abs.Sequential[ModuleLike],
	) ImportsLike
}

/*
InstanceDefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-definition-like class.
*/
type InstanceDefinitionClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		instanceMethods InstanceMethodsLike,
	) InstanceDefinitionLike
}

/*
InstanceMethodsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-methods-like class.
*/
type InstanceMethodsClassLike interface {
	// Constructor Methods
	Make(
		publicSubsection PublicSubsectionLike,
		optionalAttributeSubsection AttributeSubsectionLike,
		optionalAspectSubsection AspectSubsectionLike,
	) InstanceMethodsLike
}

/*
InstanceSectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete instance-section-like class.
*/
type InstanceSectionClassLike interface {
	// Constructor Methods
	Make(
		instanceDefinitions abs.Sequential[InstanceDefinitionLike],
	) InstanceSectionLike
}

/*
InterfaceDefinitionsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete interface-definitions-like class.
*/
type InterfaceDefinitionsClassLike interface {
	// Constructor Methods
	Make(
		classSection ClassSectionLike,
		instanceSection InstanceSectionLike,
		optionalAspectSection AspectSectionLike,
	) InterfaceDefinitionsLike
}

/*
MapClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete map-like class.
*/
type MapClassLike interface {
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
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
	// Constructor Methods
	Make(
		optionalTypeSection TypeSectionLike,
		optionalFunctionalSection FunctionalSectionLike,
	) PrimitiveDefinitionsLike
}

/*
PublicMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete public-method-like class.
*/
type PublicMethodClassLike interface {
	// Constructor Methods
	Make(
		method MethodLike,
	) PublicMethodLike
}

/*
PublicSubsectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete public-subsection-like class.
*/
type PublicSubsectionClassLike interface {
	// Constructor Methods
	Make(
		publicMethods abs.Sequential[PublicMethodLike],
	) PublicSubsectionLike
}

/*
ResultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete result-like class.
*/
type ResultClassLike interface {
	// Constructor Methods
	Make(
		any_ any,
	) ResultLike
}

/*
SetterMethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete setter-method-like class.
*/
type SetterMethodClassLike interface {
	// Constructor Methods
	Make(
		name string,
		parameter ParameterLike,
	) SetterMethodLike
}

/*
SuffixClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete suffix-like class.
*/
type SuffixClassLike interface {
	// Constructor Methods
	Make(
		name string,
	) SuffixLike
}

/*
TypeDefinitionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-definition-like class.
*/
type TypeDefinitionClassLike interface {
	// Constructor Methods
	Make(
		declaration DeclarationLike,
		abstraction AbstractionLike,
		optionalEnumeration EnumerationLike,
	) TypeDefinitionLike
}

/*
TypeSectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete type-section-like class.
*/
type TypeSectionClassLike interface {
	// Constructor Methods
	Make(
		typeDefinitions abs.Sequential[TypeDefinitionLike],
	) TypeSectionLike
}

/*
ValueClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete value-like class.
*/
type ValueClassLike interface {
	// Constructor Methods
	Make(
		name string,
		abstraction AbstractionLike,
	) ValueLike
}

// Instance Definitions

/*
AbstractionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete abstraction-like class.
*/
type AbstractionLike interface {
	// Public Methods
	GetClass() AbstractionClassLike

	// Attribute Methods
	GetOptionalPrefix() PrefixLike
	GetName() string
	GetOptionalSuffix() SuffixLike
	GetOptionalArguments() ArgumentsLike
}

/*
AdditionalArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-argument-like class.
*/
type AdditionalArgumentLike interface {
	// Public Methods
	GetClass() AdditionalArgumentClassLike

	// Attribute Methods
	GetArgument() ArgumentLike
}

/*
AdditionalConstraintLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-constraint-like class.
*/
type AdditionalConstraintLike interface {
	// Public Methods
	GetClass() AdditionalConstraintClassLike

	// Attribute Methods
	GetConstraint() ConstraintLike
}

/*
AdditionalValueLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete additional-value-like class.
*/
type AdditionalValueLike interface {
	// Public Methods
	GetClass() AdditionalValueClassLike

	// Attribute Methods
	GetName() string
}

/*
ArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Public Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetAbstraction() AbstractionLike
}

/*
ArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Public Methods
	GetClass() ArgumentsClassLike

	// Attribute Methods
	GetArgument() ArgumentLike
	GetAdditionalArguments() abs.Sequential[AdditionalArgumentLike]
}

/*
ArrayLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete array-like class.
*/
type ArrayLike interface {
	// Public Methods
	GetClass() ArrayClassLike
}

/*
AspectDefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-definition-like class.
*/
type AspectDefinitionLike interface {
	// Public Methods
	GetClass() AspectDefinitionClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetAspectMethods() abs.Sequential[AspectMethodLike]
}

/*
AspectInterfaceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-interface-like class.
*/
type AspectInterfaceLike interface {
	// Public Methods
	GetClass() AspectInterfaceClassLike

	// Attribute Methods
	GetAbstraction() AbstractionLike
}

/*
AspectMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-method-like class.
*/
type AspectMethodLike interface {
	// Public Methods
	GetClass() AspectMethodClassLike

	// Attribute Methods
	GetMethod() MethodLike
}

/*
AspectSectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-section-like class.
*/
type AspectSectionLike interface {
	// Public Methods
	GetClass() AspectSectionClassLike

	// Attribute Methods
	GetAspectDefinitions() abs.Sequential[AspectDefinitionLike]
}

/*
AspectSubsectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete aspect-subsection-like class.
*/
type AspectSubsectionLike interface {
	// Public Methods
	GetClass() AspectSubsectionClassLike

	// Attribute Methods
	GetAspectInterfaces() abs.Sequential[AspectInterfaceLike]
}

/*
AttributeMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-method-like class.
*/
type AttributeMethodLike interface {
	// Public Methods
	GetClass() AttributeMethodClassLike

	// Attribute Methods
	GetAny() any
}

/*
AttributeSubsectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-subsection-like class.
*/
type AttributeSubsectionLike interface {
	// Public Methods
	GetClass() AttributeSubsectionClassLike

	// Attribute Methods
	GetAttributeMethods() abs.Sequential[AttributeMethodLike]
}

/*
ChannelLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete channel-like class.
*/
type ChannelLike interface {
	// Public Methods
	GetClass() ChannelClassLike
}

/*
ClassDefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-definition-like class.
*/
type ClassDefinitionLike interface {
	// Public Methods
	GetClass() ClassDefinitionClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetClassMethods() ClassMethodsLike
}

/*
ClassMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-methods-like class.
*/
type ClassMethodsLike interface {
	// Public Methods
	GetClass() ClassMethodsClassLike

	// Attribute Methods
	GetConstructorSubsection() ConstructorSubsectionLike
	GetOptionalConstantSubsection() ConstantSubsectionLike
	GetOptionalFunctionSubsection() FunctionSubsectionLike
}

/*
ClassSectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete class-section-like class.
*/
type ClassSectionLike interface {
	// Public Methods
	GetClass() ClassSectionClassLike

	// Attribute Methods
	GetClassDefinitions() abs.Sequential[ClassDefinitionLike]
}

/*
ConstantMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-method-like class.
*/
type ConstantMethodLike interface {
	// Public Methods
	GetClass() ConstantMethodClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ConstantSubsectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constant-subsection-like class.
*/
type ConstantSubsectionLike interface {
	// Public Methods
	GetClass() ConstantSubsectionClassLike

	// Attribute Methods
	GetConstantMethods() abs.Sequential[ConstantMethodLike]
}

/*
ConstraintLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constraint-like class.
*/
type ConstraintLike interface {
	// Public Methods
	GetClass() ConstraintClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ConstraintsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constraints-like class.
*/
type ConstraintsLike interface {
	// Public Methods
	GetClass() ConstraintsClassLike

	// Attribute Methods
	GetConstraint() ConstraintLike
	GetAdditionalConstraints() abs.Sequential[AdditionalConstraintLike]
}

/*
ConstructorMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-method-like class.
*/
type ConstructorMethodLike interface {
	// Public Methods
	GetClass() ConstructorMethodClassLike

	// Attribute Methods
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetAbstraction() AbstractionLike
}

/*
ConstructorSubsectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete constructor-subsection-like class.
*/
type ConstructorSubsectionLike interface {
	// Public Methods
	GetClass() ConstructorSubsectionClassLike

	// Attribute Methods
	GetConstructorMethods() abs.Sequential[ConstructorMethodLike]
}

/*
DeclarationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete declaration-like class.
*/
type DeclarationLike interface {
	// Public Methods
	GetClass() DeclarationClassLike

	// Attribute Methods
	GetComment() string
	GetName() string
	GetOptionalConstraints() ConstraintsLike
}

/*
EnumerationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete enumeration-like class.
*/
type EnumerationLike interface {
	// Public Methods
	GetClass() EnumerationClassLike

	// Attribute Methods
	GetValue() ValueLike
	GetAdditionalValues() abs.Sequential[AdditionalValueLike]
}

/*
FunctionMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-method-like class.
*/
type FunctionMethodLike interface {
	// Public Methods
	GetClass() FunctionMethodClassLike

	// Attribute Methods
	GetName() string
	GetParameters() abs.Sequential[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionSubsectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-subsection-like class.
*/
type FunctionSubsectionLike interface {
	// Public Methods
	GetClass() FunctionSubsectionClassLike

	// Attribute Methods
	GetFunctionMethods() abs.Sequential[FunctionMethodLike]
}

/*
FunctionalDefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-definition-like class.
*/
type FunctionalDefinitionLike interface {
	// Public Methods
	GetClass() FunctionalDefinitionClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetParameters() abs.Sequential[ParameterLike]
	GetResult() ResultLike
}

/*
FunctionalSectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete functional-section-like class.
*/
type FunctionalSectionLike interface {
	// Public Methods
	GetClass() FunctionalSectionClassLike

	// Attribute Methods
	GetFunctionalDefinitions() abs.Sequential[FunctionalDefinitionLike]
}

/*
GetterMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete getter-method-like class.
*/
type GetterMethodLike interface {
	// Public Methods
	GetClass() GetterMethodClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
HeaderLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete header-like class.
*/
type HeaderLike interface {
	// Public Methods
	GetClass() HeaderClassLike

	// Attribute Methods
	GetComment() string
	GetName() string
}

/*
ImportsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete imports-like class.
*/
type ImportsLike interface {
	// Public Methods
	GetClass() ImportsClassLike

	// Attribute Methods
	GetModules() abs.Sequential[ModuleLike]
}

/*
InstanceDefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-definition-like class.
*/
type InstanceDefinitionLike interface {
	// Public Methods
	GetClass() InstanceDefinitionClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetInstanceMethods() InstanceMethodsLike
}

/*
InstanceMethodsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-methods-like class.
*/
type InstanceMethodsLike interface {
	// Public Methods
	GetClass() InstanceMethodsClassLike

	// Attribute Methods
	GetPublicSubsection() PublicSubsectionLike
	GetOptionalAttributeSubsection() AttributeSubsectionLike
	GetOptionalAspectSubsection() AspectSubsectionLike
}

/*
InstanceSectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete instance-section-like class.
*/
type InstanceSectionLike interface {
	// Public Methods
	GetClass() InstanceSectionClassLike

	// Attribute Methods
	GetInstanceDefinitions() abs.Sequential[InstanceDefinitionLike]
}

/*
InterfaceDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete interface-definitions-like class.
*/
type InterfaceDefinitionsLike interface {
	// Public Methods
	GetClass() InterfaceDefinitionsClassLike

	// Attribute Methods
	GetClassSection() ClassSectionLike
	GetInstanceSection() InstanceSectionLike
	GetOptionalAspectSection() AspectSectionLike
}

/*
MapLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete map-like class.
*/
type MapLike interface {
	// Public Methods
	GetClass() MapClassLike

	// Attribute Methods
	GetName() string
}

/*
MethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete method-like class.
*/
type MethodLike interface {
	// Public Methods
	GetClass() MethodClassLike

	// Attribute Methods
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
	// Public Methods
	GetClass() ModelClassLike

	// Attribute Methods
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
	// Public Methods
	GetClass() ModuleClassLike

	// Attribute Methods
	GetName() string
	GetPath() string
}

/*
ModuleDefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete module-definition-like class.
*/
type ModuleDefinitionLike interface {
	// Public Methods
	GetClass() ModuleDefinitionClassLike

	// Attribute Methods
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
	// Public Methods
	GetClass() NoneClassLike

	// Attribute Methods
	GetNewline() string
}

/*
NoticeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete notice-like class.
*/
type NoticeLike interface {
	// Public Methods
	GetClass() NoticeClassLike

	// Attribute Methods
	GetComment() string
}

/*
ParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Public Methods
	GetClass() ParameterClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}

/*
ParameterizedLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameterized-like class.
*/
type ParameterizedLike interface {
	// Public Methods
	GetClass() ParameterizedClassLike

	// Attribute Methods
	GetParameters() abs.Sequential[ParameterLike]
}

/*
PrefixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete prefix-like class.
*/
type PrefixLike interface {
	// Public Methods
	GetClass() PrefixClassLike

	// Attribute Methods
	GetAny() any
}

/*
PrimitiveDefinitionsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete primitive-definitions-like class.
*/
type PrimitiveDefinitionsLike interface {
	// Public Methods
	GetClass() PrimitiveDefinitionsClassLike

	// Attribute Methods
	GetOptionalTypeSection() TypeSectionLike
	GetOptionalFunctionalSection() FunctionalSectionLike
}

/*
PublicMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete public-method-like class.
*/
type PublicMethodLike interface {
	// Public Methods
	GetClass() PublicMethodClassLike

	// Attribute Methods
	GetMethod() MethodLike
}

/*
PublicSubsectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete public-subsection-like class.
*/
type PublicSubsectionLike interface {
	// Public Methods
	GetClass() PublicSubsectionClassLike

	// Attribute Methods
	GetPublicMethods() abs.Sequential[PublicMethodLike]
}

/*
ResultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Public Methods
	GetClass() ResultClassLike

	// Attribute Methods
	GetAny() any
}

/*
SetterMethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete setter-method-like class.
*/
type SetterMethodLike interface {
	// Public Methods
	GetClass() SetterMethodClassLike

	// Attribute Methods
	GetName() string
	GetParameter() ParameterLike
}

/*
SuffixLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete suffix-like class.
*/
type SuffixLike interface {
	// Public Methods
	GetClass() SuffixClassLike

	// Attribute Methods
	GetName() string
}

/*
TypeDefinitionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete type-definition-like class.
*/
type TypeDefinitionLike interface {
	// Public Methods
	GetClass() TypeDefinitionClassLike

	// Attribute Methods
	GetDeclaration() DeclarationLike
	GetAbstraction() AbstractionLike
	GetOptionalEnumeration() EnumerationLike
}

/*
TypeSectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete type-section-like class.
*/
type TypeSectionLike interface {
	// Public Methods
	GetClass() TypeSectionClassLike

	// Attribute Methods
	GetTypeDefinitions() abs.Sequential[TypeDefinitionLike]
}

/*
ValueLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete value-like class.
*/
type ValueLike interface {
	// Public Methods
	GetClass() ValueClassLike

	// Attribute Methods
	GetName() string
	GetAbstraction() AbstractionLike
}
