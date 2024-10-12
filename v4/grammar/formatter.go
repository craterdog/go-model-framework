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
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// Constructor Methods

func (c *formatterClass_) Make() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	instance.visitor_ = Visitor().Make(instance)
	return instance
}

// INSTANCE INTERFACE

// Methodical Methods

func (v *formatter_) ProcessComment(comment string) {
	v.appendString(comment)
}

func (v *formatter_) ProcessName(name string) {
	v.appendString(name)
}

func (v *formatter_) ProcessPath(path string) {
	v.appendString(path)
}

func (v *formatter_) ProcessSpace(space string) {
	v.appendString(space)
}

func (v *formatter_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	v.appendString(", ")
}

func (v *formatter_) PreprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	size uint,
) {
	v.appendString(", ")
}

func (v *formatter_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessArguments(arguments ast.ArgumentsLike) {
	v.appendString("[")
}

func (v *formatter_) PostprocessArguments(arguments ast.ArgumentsLike) {
	v.appendString("]")
}

func (v *formatter_) PreprocessArray(array ast.ArrayLike) {
	v.appendString("[")
}

func (v *formatter_) PostprocessArray(array ast.ArrayLike) {
	v.appendString("]")
}

func (v *formatter_) PreprocessAspectDefinition(
	aspectDefinition ast.AspectDefinitionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessAspectDefinitionSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" interface {")
		v.depth_++
	}
}

func (v *formatter_) PostprocessAspectDefinition(
	aspectDefinition ast.AspectDefinitionLike,
	index uint,
	size uint,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	size uint,
) {
}

func (v *formatter_) PreprocessAspectSection(aspectSection ast.AspectSectionLike) {
	v.appendNewline()
	v.appendString("// Aspect Definitions")
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectSubsection(aspectSubsection ast.AspectSubsectionLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Aspect Methods")
}

func (v *formatter_) PreprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAttributeSubsection(attributeSubsection ast.AttributeSubsectionLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Attribute Methods")
}

func (v *formatter_) PreprocessChannel(channel ast.ChannelLike) {
	v.appendString("chan ")
}

func (v *formatter_) PreprocessClassDefinition(
	classDefinition ast.ClassDefinitionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessClassDefinitionSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" interface {")
		v.depth_++
	}
}

func (v *formatter_) PostprocessClassDefinition(
	classDefinition ast.ClassDefinitionLike,
	index uint,
	size uint,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) PreprocessClassSection(classSection ast.ClassSectionLike) {
	v.appendNewline()
	v.appendString("// Class Definitions")
	v.appendNewline()
}

func (v *formatter_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstantMethodSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("() ")
	}
}

func (v *formatter_) PreprocessConstantSubsection(constantSubsection ast.ConstantSubsectionLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Constant Methods")
}

func (v *formatter_) ProcessConstraintSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessConstraints(constraints ast.ConstraintsLike) {
	v.appendString("[")
	v.depth_++
}

func (v *formatter_) PostprocessConstraints(constraints ast.ConstraintsLike) {
	v.depth_--
	v.appendString("]")
}

func (v *formatter_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstructorMethodSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	case 2:
		v.appendString(") ")
	}
}

func (v *formatter_) PreprocessConstructorSubsection(constructorSubsection ast.ConstructorSubsectionLike) {
	v.appendNewline()
	v.appendString("// Constructor Methods")
}

func (v *formatter_) ProcessDeclarationSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("type ")
	}
}

func (v *formatter_) PreprocessEnumeration(enumeration ast.EnumerationLike) {
	v.appendNewline()
	v.appendNewline()
	v.appendString("const (")
	v.depth_++
}

func (v *formatter_) PostprocessEnumeration(enumeration ast.EnumerationLike) {
	v.depth_--
	v.appendNewline()
	v.appendString(")")
}

func (v *formatter_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionMethodSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	case 2:
		v.appendString(")")
	}
}

func (v *formatter_) PreprocessFunctionSubsection(functionSubsection ast.FunctionSubsectionLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Function Methods")
}

func (v *formatter_) PreprocessFunctionalDefinition(
	functionalDefinition ast.FunctionalDefinitionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionalDefinitionSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" func(")
	case 2:
		v.appendString(")")
	}
}

func (v *formatter_) PostprocessFunctionalDefinition(
	functionalDefinition ast.FunctionalDefinitionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessFunctionalSection(functionalSection ast.FunctionalSectionLike) {
	v.appendNewline()
	v.appendString("// Functional Definitions")
	v.appendNewline()
}

func (v *formatter_) ProcessGetterMethodSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("() ")
	}
}

func (v *formatter_) ProcessHeaderSlot(slot uint) {
	v.appendString("package ")
}

func (v *formatter_) PreprocessImports(imports ast.ImportsLike) {
	v.appendNewline()
	v.appendNewline()
	v.appendString("import (")
	v.depth_++
}

func (v *formatter_) PostprocessImports(imports ast.ImportsLike) {
	v.depth_--
	v.appendNewline()
	v.appendString(")")
}

func (v *formatter_) PreprocessInstanceDefinition(
	instanceDefinition ast.InstanceDefinitionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessInstanceDefinitionSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" interface {")
		v.depth_++
	}
}

func (v *formatter_) PostprocessInstanceDefinition(
	instanceDefinition ast.InstanceDefinitionLike,
	index uint,
	size uint,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) PreprocessInstanceSection(instanceSection ast.InstanceSectionLike) {
	v.appendNewline()
	v.appendString("// Instance Definitions")
	v.appendNewline()
}

func (v *formatter_) PreprocessMap(map_ ast.MapLike) {
	v.appendString("map[")
}

func (v *formatter_) PostprocessMap(map_ ast.MapLike) {
	v.appendString("]")
}

func (v *formatter_) PreprocessMethod(
	method ast.MethodLike,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessMethodSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	case 2:
		v.appendString(")")
	}
}

func (v *formatter_) PreprocessModule(
	module ast.ModuleLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessModuleSlot(slot uint) {
	v.appendString(" ")
}

func (v *formatter_) PostprocessModuleDefinition(moduleDefinition_ ast.ModuleDefinitionLike) {
	v.appendNewline()
}

func (v *formatter_) PostprocessNotice(notice ast.NoticeLike) {
	v.appendNewline()
}

func (v *formatter_) PreprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
	if index == 1 {
		v.depth_++
	}
	v.appendNewline()
}

func (v *formatter_) ProcessParameterSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
	v.appendString(",")
	if index == size {
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) PreprocessParameterized(parameterized ast.ParameterizedLike) {
	v.appendString("(")
}

func (v *formatter_) PostprocessParameterized(parameterized ast.ParameterizedLike) {
	v.appendString(")")
}

func (v *formatter_) PreprocessPublicSubsection(publicSubsection ast.PublicSubsectionLike) {
	v.appendNewline()
	v.appendString("// Public Methods")
}

func (v *formatter_) PreprocessResult(result ast.ResultLike) {
	switch result.GetAny().(type) {
	case ast.NoneLike:
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessSetterMethodSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	}
}

func (v *formatter_) PostprocessSetterMethod(setterMethod ast.SetterMethodLike) {
	v.appendString(")")
}

func (v *formatter_) PreprocessSuffix(suffix ast.SuffixLike) {
	v.appendString(".")
}

func (v *formatter_) PreprocessTypeDefinition(
	typeDefinition ast.TypeDefinitionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessTypeDefinitionSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessTypeDefinition(
	typeDefinition ast.TypeDefinitionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessTypeSection(typeSection ast.TypeSectionLike) {
	v.appendNewline()
	v.appendString("// Type Definitions")
	v.appendNewline()
}

func (v *formatter_) PreprocessValue(value ast.ValueLike) {
	v.appendNewline()
}

func (v *formatter_) ProcessValueSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessValue(value ast.ValueLike) {
	v.appendString(" = iota")
}

// Public Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return v.getClass()
}

func (v *formatter_) FormatModel(
	model ast.ModelLike,
) string {
	var result_ string
	v.visitor_.VisitModel(model)
	result_ = v.getResult()
	return result_
}

// Private Methods

func (v *formatter_) getClass() *formatterClass_ {
	return v.class_
}

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// PRIVATE INTERFACE

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	class_   *formatterClass_
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
}
