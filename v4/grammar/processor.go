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

package grammar

import (
	ast "github.com/craterdog/go-model-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var processorClass = &processorClass_{
	// Initialize the class constants.
}

// Function

func Processor() ProcessorClassLike {
	return processorClass
}

// CLASS METHODS

// Target

type processorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *processorClass_) Make() ProcessorLike {
	var processor = &processor_{
		// Initialize the instance attributes.
		class_: c,
	}
	return processor
}

// INSTANCE METHODS

// Target

type processor_ struct {
	// Define the instance attributes.
	class_ *processorClass_
}

// Public

func (v *processor_) GetClass() ProcessorClassLike {
	return v.class_
}

// Methodical

func (v *processor_) ProcessComment(comment string) {
}

func (v *processor_) ProcessName(name string) {
}

func (v *processor_) ProcessNewline(newline string) {
}

func (v *processor_) ProcessPath(path string) {
}

func (v *processor_) ProcessSpace(space string) {
}

func (v *processor_) PreprocessAbstraction(abstraction ast.AbstractionLike) {
}

func (v *processor_) ProcessAbstractionSlot(slot uint) {
}

func (v *processor_) PostprocessAbstraction(abstraction ast.AbstractionLike) {
}

func (v *processor_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAdditionalArgumentSlot(slot uint) {
}

func (v *processor_) PostprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAdditionalConstraintSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAdditionalValueSlot(slot uint) {
}

func (v *processor_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessArgument(argument ast.ArgumentLike) {
}

func (v *processor_) ProcessArgumentSlot(slot uint) {
}

func (v *processor_) PostprocessArgument(argument ast.ArgumentLike) {
}

func (v *processor_) PreprocessArguments(Arguments ast.ArgumentsLike) {
}

func (v *processor_) ProcessArgumentsSlot(slot uint) {
}

func (v *processor_) PostprocessArguments(Arguments ast.ArgumentsLike) {
}

func (v *processor_) PreprocessArray(array ast.ArrayLike) {
}

func (v *processor_) ProcessArraySlot(slot uint) {
}

func (v *processor_) PostprocessArray(array ast.ArrayLike) {
}

func (v *processor_) PreprocessAspect(
	aspect ast.AspectLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAspectSlot(slot uint) {
}

func (v *processor_) PostprocessAspect(
	aspect ast.AspectLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAspectDefinitions(aspectDefinitions ast.AspectDefinitionsLike) {
}

func (v *processor_) ProcessAspectDefinitionsSlot(slot uint) {
}

func (v *processor_) PostprocessAspectDefinitions(aspectDefinitions ast.AspectDefinitionsLike) {
}

func (v *processor_) PreprocessAspectInterfaces(aspectInterfaces ast.AspectInterfacesLike) {
}

func (v *processor_) ProcessAspectInterfacesSlot(slot uint) {
}

func (v *processor_) PostprocessAspectInterfaces(aspectInterfaces ast.AspectInterfacesLike) {
}

func (v *processor_) PreprocessAttribute(
	attribute ast.AttributeLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAttributeSlot(slot uint) {
}

func (v *processor_) PostprocessAttribute(
	attribute ast.AttributeLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAttributeMethods(attributeMethods ast.AttributeMethodsLike) {
}

func (v *processor_) ProcessAttributeMethodsSlot(slot uint) {
}

func (v *processor_) PostprocessAttributeMethods(attributeMethods ast.AttributeMethodsLike) {
}

func (v *processor_) PreprocessChannel(channel ast.ChannelLike) {
}

func (v *processor_) ProcessChannelSlot(slot uint) {
}

func (v *processor_) PostprocessChannel(channel ast.ChannelLike) {
}

func (v *processor_) PreprocessClass(
	class ast.ClassLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessClassSlot(slot uint) {
}

func (v *processor_) PostprocessClass(
	class ast.ClassLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessClassDefinitions(classDefinitions ast.ClassDefinitionsLike) {
}

func (v *processor_) ProcessClassDefinitionsSlot(slot uint) {
}

func (v *processor_) PostprocessClassDefinitions(classDefinitions ast.ClassDefinitionsLike) {
}

func (v *processor_) PreprocessClassMethods(classMethods ast.ClassMethodsLike) {
}

func (v *processor_) ProcessClassMethodsSlot(slot uint) {
}

func (v *processor_) PostprocessClassMethods(classMethods ast.ClassMethodsLike) {
}

func (v *processor_) PreprocessConstant(
	constant ast.ConstantLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessConstantSlot(slot uint) {
}

func (v *processor_) PostprocessConstant(
	constant ast.ConstantLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessConstantMethods(constantMethods ast.ConstantMethodsLike) {
}

func (v *processor_) ProcessConstantMethodsSlot(slot uint) {
}

func (v *processor_) PostprocessConstantMethods(constantMethods ast.ConstantMethodsLike) {
}

func (v *processor_) PreprocessConstraint(
	constraint ast.ConstraintLike,
) {
}

func (v *processor_) ProcessConstraintSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessConstraint(
	constraint ast.ConstraintLike,
) {
}

func (v *processor_) PreprocessConstraints(
	constraints ast.ConstraintsLike,
) {
}

func (v *processor_) ProcessConstraintsSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessConstraints(
	constraints ast.ConstraintsLike,
) {
}

func (v *processor_) PreprocessConstructor(
	constructor ast.ConstructorLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessConstructorSlot(slot uint) {
}

func (v *processor_) PostprocessConstructor(
	constructor ast.ConstructorLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessConstructorMethods(constructorMethods ast.ConstructorMethodsLike) {
}

func (v *processor_) ProcessConstructorMethodsSlot(slot uint) {
}

func (v *processor_) PostprocessConstructorMethods(constructorMethods ast.ConstructorMethodsLike) {
}

func (v *processor_) PreprocessDeclaration(declaration ast.DeclarationLike) {
}

func (v *processor_) ProcessDeclarationSlot(slot uint) {
}

func (v *processor_) PostprocessDeclaration(declaration ast.DeclarationLike) {
}

func (v *processor_) PreprocessEnumeration(enumeration ast.EnumerationLike) {
}

func (v *processor_) ProcessEnumerationSlot(slot uint) {
}

func (v *processor_) PostprocessEnumeration(enumeration ast.EnumerationLike) {
}

func (v *processor_) PreprocessFunction(
	function ast.FunctionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessFunctionSlot(slot uint) {
}

func (v *processor_) PostprocessFunction(
	function ast.FunctionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessFunctionMethods(functionMethods ast.FunctionMethodsLike) {
}

func (v *processor_) ProcessFunctionMethodsSlot(slot uint) {
}

func (v *processor_) PostprocessFunctionMethods(functionMethods ast.FunctionMethodsLike) {
}

func (v *processor_) PreprocessFunctional(
	functional ast.FunctionalLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessFunctionalSlot(slot uint) {
}

func (v *processor_) PostprocessFunctional(
	functional ast.FunctionalLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessFunctionalDefinitions(functionalDefinitions ast.FunctionalDefinitionsLike) {
}

func (v *processor_) ProcessFunctionalDefinitionsSlot(slot uint) {
}

func (v *processor_) PostprocessFunctionalDefinitions(functionalDefinitions ast.FunctionalDefinitionsLike) {
}

func (v *processor_) PreprocessHeader(header ast.HeaderLike) {
}

func (v *processor_) ProcessHeaderSlot(slot uint) {
}

func (v *processor_) PostprocessHeader(header ast.HeaderLike) {
}

func (v *processor_) PreprocessImports(imports ast.ImportsLike) {
}

func (v *processor_) ProcessImportsSlot(slot uint) {
}

func (v *processor_) PostprocessImports(imports ast.ImportsLike) {
}

func (v *processor_) PreprocessInstance(
	instance ast.InstanceLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessInstanceSlot(slot uint) {
}

func (v *processor_) PostprocessInstance(
	instance ast.InstanceLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessInstanceDefinitions(instanceDefinitions ast.InstanceDefinitionsLike) {
}

func (v *processor_) ProcessInstanceDefinitionsSlot(slot uint) {
}

func (v *processor_) PostprocessInstanceDefinitions(instanceDefinitions ast.InstanceDefinitionsLike) {
}

func (v *processor_) PreprocessInstanceMethods(instanceMethods ast.InstanceMethodsLike) {
}

func (v *processor_) ProcessInstanceMethodsSlot(slot uint) {
}

func (v *processor_) PostprocessInstanceMethods(instanceMethods ast.InstanceMethodsLike) {
}

func (v *processor_) PreprocessInterface(
	interface_ ast.InterfaceLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessInterfaceSlot(slot uint) {
}

func (v *processor_) PostprocessInterface(
	interface_ ast.InterfaceLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessInterfaceDefinitions(interfaceDefinitions ast.InterfaceDefinitionsLike) {
}

func (v *processor_) ProcessInterfaceDefinitionsSlot(slot uint) {
}

func (v *processor_) PostprocessInterfaceDefinitions(interfaceDefinitions ast.InterfaceDefinitionsLike) {
}

func (v *processor_) PreprocessMap(map_ ast.MapLike) {
}

func (v *processor_) ProcessMapSlot(slot uint) {
}

func (v *processor_) PostprocessMap(map_ ast.MapLike) {
}

func (v *processor_) PreprocessMethod(
	method ast.MethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessMethodSlot(slot uint) {
}

func (v *processor_) PostprocessMethod(
	method ast.MethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessModel(model ast.ModelLike) {
}

func (v *processor_) ProcessModelSlot(slot uint) {
}

func (v *processor_) PostprocessModel(model ast.ModelLike) {
}

func (v *processor_) PreprocessModule(
	module ast.ModuleLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessModuleSlot(slot uint) {
}

func (v *processor_) PostprocessModule(
	module ast.ModuleLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessModuleDefinition(moduleDefinition ast.ModuleDefinitionLike) {
}

func (v *processor_) ProcessModuleDefinitionSlot(slot uint) {
}

func (v *processor_) PostprocessModuleDefinition(moduleDefinition ast.ModuleDefinitionLike) {
}

func (v *processor_) PreprocessNone(none ast.NoneLike) {
}

func (v *processor_) ProcessNoneSlot(slot uint) {
}

func (v *processor_) PostprocessNone(none ast.NoneLike) {
}

func (v *processor_) PreprocessNotice(notice ast.NoticeLike) {
}

func (v *processor_) ProcessNoticeSlot(slot uint) {
}

func (v *processor_) PostprocessNotice(notice ast.NoticeLike) {
}

func (v *processor_) PreprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessParameterSlot(slot uint) {
}

func (v *processor_) PostprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessParameterized(parameterized ast.ParameterizedLike) {
}

func (v *processor_) ProcessParameterizedSlot(slot uint) {
}

func (v *processor_) PostprocessParameterized(parameterized ast.ParameterizedLike) {
}

func (v *processor_) PreprocessPrefix(prefix ast.PrefixLike) {
}

func (v *processor_) ProcessPrefixSlot(slot uint) {
}

func (v *processor_) PostprocessPrefix(prefix ast.PrefixLike) {
}

func (v *processor_) PreprocessPrimitiveDefinitions(primitiveDefinitions ast.PrimitiveDefinitionsLike) {
}

func (v *processor_) ProcessPrimitiveDefinitionsSlot(slot uint) {
}

func (v *processor_) PostprocessPrimitiveDefinitions(primitiveDefinitions ast.PrimitiveDefinitionsLike) {
}

func (v *processor_) PreprocessPublicMethods(publicMethods ast.PublicMethodsLike) {
}

func (v *processor_) ProcessPublicMethodsSlot(slot uint) {
}

func (v *processor_) PostprocessPublicMethods(publicMethods ast.PublicMethodsLike) {
}

func (v *processor_) PreprocessResult(result ast.ResultLike) {
}

func (v *processor_) ProcessResultSlot(slot uint) {
}

func (v *processor_) PostprocessResult(result ast.ResultLike) {
}

func (v *processor_) PreprocessSuffix(suffix ast.SuffixLike) {
}

func (v *processor_) ProcessSuffixSlot(slot uint) {
}

func (v *processor_) PostprocessSuffix(suffix ast.SuffixLike) {
}

func (v *processor_) PreprocessType(
	type_ ast.TypeLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessTypeSlot(slot uint) {
}

func (v *processor_) PostprocessType(
	type_ ast.TypeLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessTypeDefinitions(typeDefinitions ast.TypeDefinitionsLike) {
}

func (v *processor_) ProcessTypeDefinitionsSlot(slot uint) {
}

func (v *processor_) PostprocessTypeDefinitions(typeDefinitions ast.TypeDefinitionsLike) {
}

func (v *processor_) PreprocessValue(value ast.ValueLike) {
}

func (v *processor_) ProcessValueSlot(slot uint) {
}

func (v *processor_) PostprocessValue(value ast.ValueLike) {
}
