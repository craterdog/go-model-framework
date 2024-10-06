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

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	var formatter = &formatter_{
		// Initialize the instance attributes.
		class_: c,

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	formatter.visitor_ = Visitor().Make(formatter)
	return formatter
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_   *formatterClass_
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Define the inherited aspects.
	Methodical
}

// Public

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) FormatModel(model ast.ModelLike) string {
	v.visitor_.VisitModel(model)
	return v.getResult()
}

// Methodical

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

func (v *formatter_) PreprocessArray(array ast.ArrayLike) {
	v.appendString("[")
}

func (v *formatter_) PostprocessArray(array ast.ArrayLike) {
	v.appendString("]")
}

func (v *formatter_) PreprocessAspect(
	aspect ast.AspectLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessAspectSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" interface {")
		v.depth_++
	}
}

func (v *formatter_) PostprocessAspect(
	aspect ast.AspectLike,
	index uint,
	size uint,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectDefinitions(aspectDefinitions ast.AspectDefinitionsLike) {
	v.appendNewline()
	v.appendString("// Aspects")
	v.appendNewline()
}

func (v *formatter_) PreprocessAspectInterfaces(aspectInterfaces ast.AspectInterfacesLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Aspect")
}

func (v *formatter_) PreprocessAccessor(
	accessor ast.AccessorLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessGetterSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("() ")
	}
}

func (v *formatter_) ProcessSetterSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	}
}

func (v *formatter_) PostprocessSetter(setter ast.SetterLike) {
	v.appendString(")")
}

func (v *formatter_) PreprocessAttributeMethods(attributeMethods ast.AttributeMethodsLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Attribute")
}

func (v *formatter_) PreprocessChannel(channel ast.ChannelLike) {
	v.appendString("chan ")
}

func (v *formatter_) ProcessConstraintSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessClass(
	class ast.ClassLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessClassSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" interface {")
		v.depth_++
	}
}

func (v *formatter_) PostprocessClass(
	class ast.ClassLike,
	index uint,
	size uint,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) PreprocessClassDefinitions(classDefinitions ast.ClassDefinitionsLike) {
	v.appendNewline()
	v.appendString("// Classes")
	v.appendNewline()
}

func (v *formatter_) PreprocessConstant(
	constant ast.ConstantLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstantSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("() ")
	}
}

func (v *formatter_) PreprocessConstantMethods(constantMethods ast.ConstantMethodsLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Constant")
}

func (v *formatter_) PreprocessConstructor(
	constructor ast.ConstructorLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessConstructorSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	case 2:
		v.appendString(") ")
	}
}

func (v *formatter_) PreprocessConstructorMethods(constructorMethods ast.ConstructorMethodsLike) {
	v.appendNewline()
	v.appendString("// Constructor")
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

func (v *formatter_) PreprocessFunction(
	function ast.FunctionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	case 2:
		v.appendString(")")
	}
}

func (v *formatter_) PreprocessFunctionMethods(functionMethods ast.FunctionMethodsLike) {
	v.appendString("\n")
	v.appendNewline()
	v.appendString("// Function")
}

func (v *formatter_) PreprocessFunctional(
	functional ast.FunctionalLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionalSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" func(")
	case 2:
		v.appendString(")")
	}
}

func (v *formatter_) PostprocessFunctional(
	functional ast.FunctionalLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessFunctionalDefinitions(functionalDefinitions ast.FunctionalDefinitionsLike) {
	v.appendNewline()
	v.appendString("// Functionals")
	v.appendNewline()
}

func (v *formatter_) PreprocessArguments(arguments ast.ArgumentsLike) {
	v.appendString("[")
}

func (v *formatter_) PostprocessArguments(arguments ast.ArgumentsLike) {
	v.appendString("]")
}

func (v *formatter_) PreprocessConstraints(constraints ast.ConstraintsLike) {
	v.appendString("[")
	v.depth_++
}

func (v *formatter_) PostprocessConstraints(constraints ast.ConstraintsLike) {
	v.depth_--
	v.appendString("]")
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

func (v *formatter_) PreprocessInstance(
	instance ast.InstanceLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessInstanceSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" interface {")
		v.depth_++
	}
}

func (v *formatter_) PostprocessInstance(
	instance ast.InstanceLike,
	index uint,
	size uint,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) PreprocessInstanceDefinitions(instanceDefinitions ast.InstanceDefinitionsLike) {
	v.appendNewline()
	v.appendString("// Instances")
	v.appendNewline()
}

func (v *formatter_) PreprocessInterface(
	interface_ ast.InterfaceLike,
	index uint,
	size uint,
) {
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
	index uint,
	size uint,
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

func (v *formatter_) ProcessParameterizedSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString("(")
	case 2:
		v.appendString(")")
	}
}

func (v *formatter_) PreprocessPublicMethods(publicMethods ast.PublicMethodsLike) {
	v.appendNewline()
	v.appendString("// Public")
}

func (v *formatter_) PreprocessResult(result ast.ResultLike) {
	switch result.GetAny().(type) {
	case ast.NoneLike:
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessSuffix(suffix ast.SuffixLike) {
	v.appendString(".")
}

func (v *formatter_) PreprocessType(
	type_ ast.TypeLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessTypeSlot(slot uint) {
	switch slot {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PostprocessType(
	type_ ast.TypeLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessTypeDefinitions(typeDefinitions ast.TypeDefinitionsLike) {
	v.appendNewline()
	v.appendString("// Types")
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

// Private

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
