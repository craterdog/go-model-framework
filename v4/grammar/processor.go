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

// CLASS INTERFACE

// Access Function

func Processor() ProcessorClassLike {
	return processorReference()
}

// Constructor Methods

func (c *processorClass_) Make() ProcessorLike {
	var instance = &processor_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Methodical Methods

func (v *processor_) ProcessComment(
	comment string,
) {
}

func (v *processor_) ProcessName(
	name string,
) {
}

func (v *processor_) ProcessNewline(
	newline string,
) {
}

func (v *processor_) ProcessPath(
	path string,
) {
}

func (v *processor_) ProcessSpace(
	space string,
) {
}

func (v *processor_) PreprocessAbstraction(
	abstraction ast.AbstractionLike,
) {
}

func (v *processor_) ProcessAbstractionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAbstraction(
	abstraction ast.AbstractionLike,
) {
}

func (v *processor_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
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

func (v *processor_) ProcessAdditionalValueSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessArgument(
	argument ast.ArgumentLike,
) {
}

func (v *processor_) ProcessArgumentSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessArgument(
	argument ast.ArgumentLike,
) {
}

func (v *processor_) PreprocessArguments(
	arguments ast.ArgumentsLike,
) {
}

func (v *processor_) ProcessArgumentsSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessArguments(
	arguments ast.ArgumentsLike,
) {
}

func (v *processor_) PreprocessArray(
	array ast.ArrayLike,
) {
}

func (v *processor_) ProcessArraySlot(
	slot uint,
) {
}

func (v *processor_) PostprocessArray(
	array ast.ArrayLike,
) {
}

func (v *processor_) PreprocessAspectDefinition(
	aspectDefinition ast.AspectDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAspectDefinitionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAspectDefinition(
	aspectDefinition ast.AspectDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAspectInterfaceSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAspectMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAspectSection(
	aspectSection ast.AspectSectionLike,
) {
}

func (v *processor_) ProcessAspectSectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAspectSection(
	aspectSection ast.AspectSectionLike,
) {
}

func (v *processor_) PreprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
) {
}

func (v *processor_) ProcessAspectSubsectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
) {
}

func (v *processor_) PreprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessAttributeMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
) {
}

func (v *processor_) ProcessAttributeSubsectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
) {
}

func (v *processor_) PreprocessChannel(
	channel ast.ChannelLike,
) {
}

func (v *processor_) ProcessChannelSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessChannel(
	channel ast.ChannelLike,
) {
}

func (v *processor_) PreprocessClassDefinition(
	classDefinition ast.ClassDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessClassDefinitionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessClassDefinition(
	classDefinition ast.ClassDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessClassMethods(
	classMethods ast.ClassMethodsLike,
) {
}

func (v *processor_) ProcessClassMethodsSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessClassMethods(
	classMethods ast.ClassMethodsLike,
) {
}

func (v *processor_) PreprocessClassSection(
	classSection ast.ClassSectionLike,
) {
}

func (v *processor_) ProcessClassSectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessClassSection(
	classSection ast.ClassSectionLike,
) {
}

func (v *processor_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessConstantMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
) {
}

func (v *processor_) ProcessConstantSubsectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
) {
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

func (v *processor_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessConstructorMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
) {
}

func (v *processor_) ProcessConstructorSubsectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
) {
}

func (v *processor_) PreprocessDeclaration(
	declaration ast.DeclarationLike,
) {
}

func (v *processor_) ProcessDeclarationSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessDeclaration(
	declaration ast.DeclarationLike,
) {
}

func (v *processor_) PreprocessEnumeration(
	enumeration ast.EnumerationLike,
) {
}

func (v *processor_) ProcessEnumerationSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessEnumeration(
	enumeration ast.EnumerationLike,
) {
}

func (v *processor_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessFunctionMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
) {
}

func (v *processor_) ProcessFunctionSubsectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
) {
}

func (v *processor_) PreprocessFunctionalDefinition(
	functionalDefinition ast.FunctionalDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessFunctionalDefinitionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessFunctionalDefinition(
	functionalDefinition ast.FunctionalDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
) {
}

func (v *processor_) ProcessFunctionalSectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
) {
}

func (v *processor_) PreprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
) {
}

func (v *processor_) ProcessGetterMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
) {
}

func (v *processor_) PreprocessHeader(
	header ast.HeaderLike,
) {
}

func (v *processor_) ProcessHeaderSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessHeader(
	header ast.HeaderLike,
) {
}

func (v *processor_) PreprocessImports(
	imports ast.ImportsLike,
) {
}

func (v *processor_) ProcessImportsSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessImports(
	imports ast.ImportsLike,
) {
}

func (v *processor_) PreprocessInstanceDefinition(
	instanceDefinition ast.InstanceDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessInstanceDefinitionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessInstanceDefinition(
	instanceDefinition ast.InstanceDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
) {
}

func (v *processor_) ProcessInstanceMethodsSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
) {
}

func (v *processor_) PreprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
) {
}

func (v *processor_) ProcessInstanceSectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
) {
}

func (v *processor_) PreprocessInterfaceDefinitions(
	interfaceDefinitions ast.InterfaceDefinitionsLike,
) {
}

func (v *processor_) ProcessInterfaceDefinitionsSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessInterfaceDefinitions(
	interfaceDefinitions ast.InterfaceDefinitionsLike,
) {
}

func (v *processor_) PreprocessMap(
	map_ ast.MapLike,
) {
}

func (v *processor_) ProcessMapSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessMap(
	map_ ast.MapLike,
) {
}

func (v *processor_) PreprocessMethod(
	method ast.MethodLike,
) {
}

func (v *processor_) ProcessMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessMethod(
	method ast.MethodLike,
) {
}

func (v *processor_) PreprocessModel(
	model ast.ModelLike,
) {
}

func (v *processor_) ProcessModelSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessModel(
	model ast.ModelLike,
) {
}

func (v *processor_) PreprocessModule(
	module ast.ModuleLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessModuleSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessModule(
	module ast.ModuleLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessModuleDefinition(
	moduleDefinition ast.ModuleDefinitionLike,
) {
}

func (v *processor_) ProcessModuleDefinitionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessModuleDefinition(
	moduleDefinition ast.ModuleDefinitionLike,
) {
}

func (v *processor_) PreprocessNone(
	none ast.NoneLike,
) {
}

func (v *processor_) ProcessNoneSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessNone(
	none ast.NoneLike,
) {
}

func (v *processor_) PreprocessNotice(
	notice ast.NoticeLike,
) {
}

func (v *processor_) ProcessNoticeSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessNotice(
	notice ast.NoticeLike,
) {
}

func (v *processor_) PreprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessParameterSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessParameterized(
	parameterized ast.ParameterizedLike,
) {
}

func (v *processor_) ProcessParameterizedSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessParameterized(
	parameterized ast.ParameterizedLike,
) {
}

func (v *processor_) PreprocessPrefix(
	prefix ast.PrefixLike,
) {
}

func (v *processor_) ProcessPrefixSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessPrefix(
	prefix ast.PrefixLike,
) {
}

func (v *processor_) PreprocessPrimitiveDefinitions(
	primitiveDefinitions ast.PrimitiveDefinitionsLike,
) {
}

func (v *processor_) ProcessPrimitiveDefinitionsSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessPrimitiveDefinitions(
	primitiveDefinitions ast.PrimitiveDefinitionsLike,
) {
}

func (v *processor_) PreprocessPublicMethod(
	publicMethod ast.PublicMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessPublicMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessPublicMethod(
	publicMethod ast.PublicMethodLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessPublicSubsection(
	publicSubsection ast.PublicSubsectionLike,
) {
}

func (v *processor_) ProcessPublicSubsectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessPublicSubsection(
	publicSubsection ast.PublicSubsectionLike,
) {
}

func (v *processor_) PreprocessResult(
	result ast.ResultLike,
) {
}

func (v *processor_) ProcessResultSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessResult(
	result ast.ResultLike,
) {
}

func (v *processor_) PreprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
) {
}

func (v *processor_) ProcessSetterMethodSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
) {
}

func (v *processor_) PreprocessSuffix(
	suffix ast.SuffixLike,
) {
}

func (v *processor_) ProcessSuffixSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessSuffix(
	suffix ast.SuffixLike,
) {
}

func (v *processor_) PreprocessTypeDefinition(
	typeDefinition ast.TypeDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) ProcessTypeDefinitionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessTypeDefinition(
	typeDefinition ast.TypeDefinitionLike,
	index uint,
	size uint,
) {
}

func (v *processor_) PreprocessTypeSection(
	typeSection ast.TypeSectionLike,
) {
}

func (v *processor_) ProcessTypeSectionSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessTypeSection(
	typeSection ast.TypeSectionLike,
) {
}

func (v *processor_) PreprocessValue(
	value ast.ValueLike,
) {
}

func (v *processor_) ProcessValueSlot(
	slot uint,
) {
}

func (v *processor_) PostprocessValue(
	value ast.ValueLike,
) {
}

// Public Methods

func (v *processor_) GetClass() ProcessorClassLike {
	return v.getClass()
}

// Private Methods

func (v *processor_) getClass() *processorClass_ {
	return processorReference()
}

// PRIVATE INTERFACE

// Instance Structure

type processor_ struct {
	// Declare the instance attributes.
}

// Class Structure

type processorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func processorReference() *processorClass_ {
	return processorReference_
}

var processorReference_ = &processorClass_{
	// Initialize the class constants.
}
