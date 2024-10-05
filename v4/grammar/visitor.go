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
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	ast "github.com/craterdog/go-model-framework/v4/ast"
)

// CLASS ACCESS

// Reference

var visitorClass = &visitorClass_{
	// Initialize the class constants.
}

// Function

func Visitor() VisitorClassLike {
	return visitorClass
}

// CLASS METHODS

// Target

type visitorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *visitorClass_) Make(processor Methodical) VisitorLike {
	return &visitor_{
		// Initialize the instance attributes.
		class_:     c,
		processor_: processor,
	}
}

// INSTANCE METHODS

// Target

type visitor_ struct {
	// Define the instance attributes.
	class_     *visitorClass_
	processor_ Methodical
}

// Public

func (v *visitor_) GetClass() VisitorClassLike {
	return v.class_
}

func (v *visitor_) VisitModel(model ast.ModelLike) {
	// Visit the model syntax.
	v.processor_.PreprocessModel(model)
	v.visitModel(model)
	v.processor_.PostprocessModel(model)
}

// Private

func (v *visitor_) visitAbstraction(abstraction ast.AbstractionLike) {
	// Visit the optional prefix rule.
	var optionalPrefix = abstraction.GetOptionalPrefix()
	if col.IsDefined(optionalPrefix) {
		v.processor_.PreprocessPrefix(optionalPrefix)
		v.visitPrefix(optionalPrefix)
		v.processor_.PostprocessPrefix(optionalPrefix)
	}

	// Visit slot 1 between references.
	v.processor_.ProcessAbstractionSlot(1)

	// Visit the name token.
	var name = abstraction.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 2 between references.
	v.processor_.ProcessAbstractionSlot(2)

	// Visit the optional suffix rule.
	var optionalSuffix = abstraction.GetOptionalSuffix()
	if col.IsDefined(optionalSuffix) {
		v.processor_.PreprocessSuffix(optionalSuffix)
		v.visitSuffix(optionalSuffix)
		v.processor_.PostprocessSuffix(optionalSuffix)
	}

	// Visit slot 3 between references.
	v.processor_.ProcessAbstractionSlot(3)

	// Visit the optional arguments rule.
	var optionalArguments = abstraction.GetOptionalArguments()
	if col.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(optionalArguments)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(optionalArguments)
	}
}

func (v *visitor_) visitAdditionalArgument(additionalArgument ast.AdditionalArgumentLike) {
	// Visit the argument rule.
	var argument = additionalArgument.GetArgument()
	v.processor_.PreprocessArgument(argument)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(argument)
}

func (v *visitor_) visitAdditionalConstraint(additionalConstraint ast.AdditionalConstraintLike) {
	// Visit the constraint rule.
	var constraint = additionalConstraint.GetConstraint()
	v.processor_.PreprocessConstraint(constraint)
	v.visitConstraint(constraint)
	v.processor_.PostprocessConstraint(constraint)
}

func (v *visitor_) visitAdditionalValue(additionalValue ast.AdditionalValueLike) {
	// Visit the name token.
	var name = additionalValue.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitArgument(argument ast.ArgumentLike) {
	// Visit the abstraction rule.
	var abstraction = argument.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitArray(array ast.ArrayLike) {}

func (v *visitor_) visitAspect(aspect ast.AspectLike) {
	// Visit the declaration rule.
	var declaration = aspect.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessAspectSlot(1)

	// Visit each method rule.
	var methodIndex uint
	var methods = aspect.GetMethods().GetIterator()
	var methodsSize = uint(methods.GetSize())
	for methods.HasNext() {
		methodIndex++
		var method = methods.GetNext()
		v.processor_.PreprocessMethod(
			method,
			methodIndex,
			methodsSize,
		)
		v.visitMethod(method)
		v.processor_.PostprocessMethod(
			method,
			methodIndex,
			methodsSize,
		)
	}
}

func (v *visitor_) visitAspectDefinitions(aspectDefinitions ast.AspectDefinitionsLike) {
	// Visit each aspect rule.
	var aspectIndex uint
	var aspects = aspectDefinitions.GetAspects().GetIterator()
	var aspectsSize = uint(aspects.GetSize())
	for aspects.HasNext() {
		aspectIndex++
		var aspect = aspects.GetNext()
		v.processor_.PreprocessAspect(
			aspect,
			aspectIndex,
			aspectsSize,
		)
		v.visitAspect(aspect)
		v.processor_.PostprocessAspect(
			aspect,
			aspectIndex,
			aspectsSize,
		)
	}
}

func (v *visitor_) visitAspectInterfaces(aspectInterfaces ast.AspectInterfacesLike) {
	// Visit each interface rule.
	var interfaceIndex uint
	var interfaces = aspectInterfaces.GetInterfaces().GetIterator()
	var interfacesSize = uint(interfaces.GetSize())
	for interfaces.HasNext() {
		interfaceIndex++
		var interface_ = interfaces.GetNext()
		v.processor_.PreprocessInterface(
			interface_,
			interfaceIndex,
			interfacesSize,
		)
		v.visitInterface(interface_)
		v.processor_.PostprocessInterface(
			interface_,
			interfaceIndex,
			interfacesSize,
		)
	}
}

func (v *visitor_) visitAttribute(attribute ast.AttributeLike) {
	// Visit the name token.
	var name = attribute.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessAttributeSlot(1)

	// Visit the parameter rule.
	var optionalParameter = attribute.GetOptionalParameter()
	if col.IsDefined(optionalParameter) {
		v.processor_.PreprocessParameter(optionalParameter, 1, 1)
		v.visitParameter(optionalParameter)
		v.processor_.PostprocessParameter(optionalParameter, 1, 1)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessAttributeSlot(2)

	// Visit the optional abstraction rule.
	var optionalAbstraction = attribute.GetOptionalAbstraction()
	if col.IsDefined(optionalAbstraction) {
		v.processor_.PreprocessAbstraction(optionalAbstraction)
		v.visitAbstraction(optionalAbstraction)
		v.processor_.PostprocessAbstraction(optionalAbstraction)
	}
}

func (v *visitor_) visitAttributeMethods(attributeMethods ast.AttributeMethodsLike) {
	// Visit each attribute rule.
	var attributeIndex uint
	var attributes = attributeMethods.GetAttributes().GetIterator()
	var attributesSize = uint(attributes.GetSize())
	for attributes.HasNext() {
		attributeIndex++
		var attribute = attributes.GetNext()
		v.processor_.PreprocessAttribute(
			attribute,
			attributeIndex,
			attributesSize,
		)
		v.visitAttribute(attribute)
		v.processor_.PostprocessAttribute(
			attribute,
			attributeIndex,
			attributesSize,
		)
	}
}

func (v *visitor_) visitChannel(channel ast.ChannelLike) {}

func (v *visitor_) visitClass(class ast.ClassLike) {
	// Visit the declaration rule.
	var declaration = class.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessClassSlot(1)

	// Visit the classMethods rule.
	var classMethods = class.GetClassMethods()
	v.processor_.PreprocessClassMethods(classMethods)
	v.visitClassMethods(classMethods)
	v.processor_.PostprocessClassMethods(classMethods)
}

func (v *visitor_) visitClassDefinitions(classDefinitions ast.ClassDefinitionsLike) {
	// Visit each class rule.
	var classIndex uint
	var classes = classDefinitions.GetClasses().GetIterator()
	var classesSize = uint(classes.GetSize())
	for classes.HasNext() {
		classIndex++
		var class = classes.GetNext()
		v.processor_.PreprocessClass(
			class,
			classIndex,
			classesSize,
		)
		v.visitClass(class)
		v.processor_.PostprocessClass(
			class,
			classIndex,
			classesSize,
		)
	}
}

func (v *visitor_) visitClassMethods(classMethods ast.ClassMethodsLike) {
	// Visit the constructorMethods rule.
	var constructorMethods = classMethods.GetConstructorMethods()
	v.processor_.PreprocessConstructorMethods(constructorMethods)
	v.visitConstructorMethods(constructorMethods)
	v.processor_.PostprocessConstructorMethods(constructorMethods)

	// Visit slot 1 between references.
	v.processor_.ProcessClassMethodsSlot(1)

	// Visit the optional constantMethods rule.
	var optionalConstantMethods = classMethods.GetOptionalConstantMethods()
	if col.IsDefined(optionalConstantMethods) {
		v.processor_.PreprocessConstantMethods(optionalConstantMethods)
		v.visitConstantMethods(optionalConstantMethods)
		v.processor_.PostprocessConstantMethods(optionalConstantMethods)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessClassMethodsSlot(2)

	// Visit the optional functionMethods rule.
	var optionalFunctionMethods = classMethods.GetOptionalFunctionMethods()
	if col.IsDefined(optionalFunctionMethods) {
		v.processor_.PreprocessFunctionMethods(optionalFunctionMethods)
		v.visitFunctionMethods(optionalFunctionMethods)
		v.processor_.PostprocessFunctionMethods(optionalFunctionMethods)
	}
}

func (v *visitor_) visitConstant(constant ast.ConstantLike) {
	// Visit the name token.
	var name = constant.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstantSlot(1)

	// Visit the abstraction rule.
	var abstraction = constant.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstantMethods(constantMethods ast.ConstantMethodsLike) {
	// Visit each constant rule.
	var constantIndex uint
	var constants = constantMethods.GetConstants().GetIterator()
	var constantsSize = uint(constants.GetSize())
	for constants.HasNext() {
		constantIndex++
		var constant = constants.GetNext()
		v.processor_.PreprocessConstant(
			constant,
			constantIndex,
			constantsSize,
		)
		v.visitConstant(constant)
		v.processor_.PostprocessConstant(
			constant,
			constantIndex,
			constantsSize,
		)
	}
}

func (v *visitor_) visitConstructor(constructor ast.ConstructorLike) {
	// Visit the name token.
	var name = constructor.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstructorSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = constructor.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessConstructorSlot(2)

	// Visit the abstraction rule.
	var abstraction = constructor.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstructorMethods(constructorMethods ast.ConstructorMethodsLike) {
	// Visit each constructor rule.
	var constructorIndex uint
	var constructors = constructorMethods.GetConstructors().GetIterator()
	var constructorsSize = uint(constructors.GetSize())
	for constructors.HasNext() {
		constructorIndex++
		var constructor = constructors.GetNext()
		v.processor_.PreprocessConstructor(
			constructor,
			constructorIndex,
			constructorsSize,
		)
		v.visitConstructor(constructor)
		v.processor_.PostprocessConstructor(
			constructor,
			constructorIndex,
			constructorsSize,
		)
	}
}

func (v *visitor_) visitDeclaration(declaration ast.DeclarationLike) {
	// Visit the comment token.
	var comment = declaration.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit slot 1 between references.
	v.processor_.ProcessDeclarationSlot(1)

	// Visit the name token.
	var name = declaration.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 2 between references.
	v.processor_.ProcessDeclarationSlot(2)

	// Visit the optional constraints rule.
	var optionalConstraints = declaration.GetOptionalConstraints()
	if col.IsDefined(optionalConstraints) {
		v.processor_.PreprocessConstraints(optionalConstraints)
		v.visitConstraints(optionalConstraints)
		v.processor_.PostprocessConstraints(optionalConstraints)
	}
}

func (v *visitor_) visitEnumeration(enumeration ast.EnumerationLike) {
	// Visit the value rule.
	var value = enumeration.GetValue()
	v.processor_.PreprocessValue(value)
	v.visitValue(value)
	v.processor_.PostprocessValue(value)

	// Visit slot 1 between references.
	v.processor_.ProcessEnumerationSlot(1)

	// Visit each additionalValue rule.
	var additionalValueIndex uint
	var additionalValues = enumeration.GetAdditionalValues().GetIterator()
	var additionalValuesSize = uint(additionalValues.GetSize())
	for additionalValues.HasNext() {
		additionalValueIndex++
		var additionalValue = additionalValues.GetNext()
		v.processor_.PreprocessAdditionalValue(
			additionalValue,
			additionalValueIndex,
			additionalValuesSize,
		)
		v.visitAdditionalValue(additionalValue)
		v.processor_.PostprocessAdditionalValue(
			additionalValue,
			additionalValueIndex,
			additionalValuesSize,
		)
	}
}

func (v *visitor_) visitFunction(function ast.FunctionLike) {
	// Visit the name token.
	var name = function.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessFunctionSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = function.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessFunctionSlot(2)

	// Visit the result rule.
	var result = function.GetResult()
	v.processor_.PreprocessResult(result)
	v.visitResult(result)
	v.processor_.PostprocessResult(result)
}

func (v *visitor_) visitFunctionMethods(functionMethods ast.FunctionMethodsLike) {
	// Visit each function rule.
	var functionIndex uint
	var functions = functionMethods.GetFunctions().GetIterator()
	var functionsSize = uint(functions.GetSize())
	for functions.HasNext() {
		functionIndex++
		var function = functions.GetNext()
		v.processor_.PreprocessFunction(
			function,
			functionIndex,
			functionsSize,
		)
		v.visitFunction(function)
		v.processor_.PostprocessFunction(
			function,
			functionIndex,
			functionsSize,
		)
	}
}

func (v *visitor_) visitFunctional(functional ast.FunctionalLike) {
	// Visit the declaration rule.
	var declaration = functional.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessFunctionalSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = functional.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessFunctionalSlot(2)

	// Visit the result rule.
	var result = functional.GetResult()
	v.processor_.PreprocessResult(result)
	v.visitResult(result)
	v.processor_.PostprocessResult(result)
}

func (v *visitor_) visitFunctionalDefinitions(functionalDefinitions ast.FunctionalDefinitionsLike) {
	// Visit each functional rule.
	var functionalIndex uint
	var functionals = functionalDefinitions.GetFunctionals().GetIterator()
	var functionalsSize = uint(functionals.GetSize())
	for functionals.HasNext() {
		functionalIndex++
		var functional = functionals.GetNext()
		v.processor_.PreprocessFunctional(
			functional,
			functionalIndex,
			functionalsSize,
		)
		v.visitFunctional(functional)
		v.processor_.PostprocessFunctional(
			functional,
			functionalIndex,
			functionalsSize,
		)
	}
}

func (v *visitor_) visitArguments(arguments ast.ArgumentsLike) {
	// Visit the argument rule.
	var argument = arguments.GetArgument()
	v.processor_.PreprocessArgument(argument)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(argument)

	// Visit slot 1 between references.
	v.processor_.ProcessArgumentsSlot(1)

	// Visit each additionalArgument rule.
	var additionalArgumentIndex uint
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	var additionalArgumentsSize = uint(additionalArguments.GetSize())
	for additionalArguments.HasNext() {
		additionalArgumentIndex++
		var additionalArgument = additionalArguments.GetNext()
		v.processor_.PreprocessAdditionalArgument(
			additionalArgument,
			additionalArgumentIndex,
			additionalArgumentsSize,
		)
		v.visitAdditionalArgument(additionalArgument)
		v.processor_.PostprocessAdditionalArgument(
			additionalArgument,
			additionalArgumentIndex,
			additionalArgumentsSize,
		)
	}
}

func (v *visitor_) visitConstraint(constraint ast.ConstraintLike) {
	// Visit the name token.
	var name = constraint.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstraintSlot(1)

	// Visit the abstraction rule.
	var abstraction = constraint.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstraints(constraints ast.ConstraintsLike) {
	// Visit the constraint rule.
	var constraint = constraints.GetConstraint()
	v.processor_.PreprocessConstraint(constraint)
	v.visitConstraint(constraint)
	v.processor_.PostprocessConstraint(constraint)

	// Visit slot 1 between references.
	v.processor_.ProcessConstraintsSlot(1)

	// Visit each additionalConstraint rule.
	var additionalConstraintIndex uint
	var additionalConstraints = constraints.GetAdditionalConstraints().GetIterator()
	var additionalConstraintsSize = uint(additionalConstraints.GetSize())
	for additionalConstraints.HasNext() {
		additionalConstraintIndex++
		var additionalConstraint = additionalConstraints.GetNext()
		v.processor_.PreprocessAdditionalConstraint(
			additionalConstraint,
			additionalConstraintIndex,
			additionalConstraintsSize,
		)
		v.visitAdditionalConstraint(additionalConstraint)
		v.processor_.PostprocessAdditionalConstraint(
			additionalConstraint,
			additionalConstraintIndex,
			additionalConstraintsSize,
		)
	}
}

func (v *visitor_) visitHeader(header ast.HeaderLike) {
	// Visit the comment token.
	var comment = header.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit slot 1 between references.
	v.processor_.ProcessHeaderSlot(1)

	// Visit the name token.
	var name = header.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitImports(imports ast.ImportsLike) {
	// Visit each module rule.
	var moduleIndex uint
	var modules = imports.GetModules().GetIterator()
	var modulesSize = uint(modules.GetSize())
	for modules.HasNext() {
		moduleIndex++
		var module = modules.GetNext()
		v.processor_.PreprocessModule(
			module,
			moduleIndex,
			modulesSize,
		)
		v.visitModule(module)
		v.processor_.PostprocessModule(
			module,
			moduleIndex,
			modulesSize,
		)
	}
}

func (v *visitor_) visitInstance(instance ast.InstanceLike) {
	// Visit the declaration rule.
	var declaration = instance.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessInstanceSlot(1)

	// Visit the instanceMethods rule.
	var instanceMethods = instance.GetInstanceMethods()
	v.processor_.PreprocessInstanceMethods(instanceMethods)
	v.visitInstanceMethods(instanceMethods)
	v.processor_.PostprocessInstanceMethods(instanceMethods)
}

func (v *visitor_) visitInstanceDefinitions(instanceDefinitions ast.InstanceDefinitionsLike) {
	// Visit each instance rule.
	var instanceIndex uint
	var instances = instanceDefinitions.GetInstances().GetIterator()
	var instancesSize = uint(instances.GetSize())
	for instances.HasNext() {
		instanceIndex++
		var instance = instances.GetNext()
		v.processor_.PreprocessInstance(
			instance,
			instanceIndex,
			instancesSize,
		)
		v.visitInstance(instance)
		v.processor_.PostprocessInstance(
			instance,
			instanceIndex,
			instancesSize,
		)
	}
}

func (v *visitor_) visitInstanceMethods(instanceMethods ast.InstanceMethodsLike) {
	// Visit the publicMethods rule.
	var publicMethods = instanceMethods.GetPublicMethods()
	v.processor_.PreprocessPublicMethods(publicMethods)
	v.visitPublicMethods(publicMethods)
	v.processor_.PostprocessPublicMethods(publicMethods)

	// Visit slot 1 between references.
	v.processor_.ProcessInstanceMethodsSlot(1)

	// Visit the optional attributeMethods rule.
	var optionalAttributeMethods = instanceMethods.GetOptionalAttributeMethods()
	if col.IsDefined(optionalAttributeMethods) {
		v.processor_.PreprocessAttributeMethods(optionalAttributeMethods)
		v.visitAttributeMethods(optionalAttributeMethods)
		v.processor_.PostprocessAttributeMethods(optionalAttributeMethods)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessInstanceMethodsSlot(2)

	// Visit the optional aspectInterfaces rule.
	var optionalAspectInterfaces = instanceMethods.GetOptionalAspectInterfaces()
	if col.IsDefined(optionalAspectInterfaces) {
		v.processor_.PreprocessAspectInterfaces(optionalAspectInterfaces)
		v.visitAspectInterfaces(optionalAspectInterfaces)
		v.processor_.PostprocessAspectInterfaces(optionalAspectInterfaces)
	}
}

func (v *visitor_) visitInterface(interface_ ast.InterfaceLike) {
	// Visit the abstraction rule.
	var abstraction = interface_.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)

	// Visit slot 1 between references.
	v.processor_.ProcessInterfaceSlot(1)
}

func (v *visitor_) visitInterfaceDefinitions(interfaceDefinitions ast.InterfaceDefinitionsLike) {
	// Visit the classDefinitions rule.
	var classDefinitions = interfaceDefinitions.GetClassDefinitions()
	v.processor_.PreprocessClassDefinitions(classDefinitions)
	v.visitClassDefinitions(classDefinitions)
	v.processor_.PostprocessClassDefinitions(classDefinitions)

	// Visit slot 1 between references.
	v.processor_.ProcessInterfaceDefinitionsSlot(1)

	// Visit the instanceDefinitions rule.
	var instanceDefinitions = interfaceDefinitions.GetInstanceDefinitions()
	v.processor_.PreprocessInstanceDefinitions(instanceDefinitions)
	v.visitInstanceDefinitions(instanceDefinitions)
	v.processor_.PostprocessInstanceDefinitions(instanceDefinitions)

	// Visit slot 2 between references.
	v.processor_.ProcessInterfaceDefinitionsSlot(2)

	// Visit the optional aspectDefinitions rule.
	var optionalAspectDefinitions = interfaceDefinitions.GetOptionalAspectDefinitions()
	if col.IsDefined(optionalAspectDefinitions) {
		v.processor_.PreprocessAspectDefinitions(optionalAspectDefinitions)
		v.visitAspectDefinitions(optionalAspectDefinitions)
		v.processor_.PostprocessAspectDefinitions(optionalAspectDefinitions)
	}
}

func (v *visitor_) visitMap(map_ ast.MapLike) {
	// Visit the name token.
	var name = map_.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitMethod(method ast.MethodLike) {
	// Visit the name token.
	var name = method.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessMethodSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = method.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessMethodSlot(2)

	// Visit the optional result rule.
	var optionalResult = method.GetOptionalResult()
	if col.IsDefined(optionalResult) {
		v.processor_.PreprocessResult(optionalResult)
		v.visitResult(optionalResult)
		v.processor_.PostprocessResult(optionalResult)
	}
}

func (v *visitor_) visitModel(model ast.ModelLike) {
	// Visit the moduleDefinition rule.
	var moduleDefinition = model.GetModuleDefinition()
	v.processor_.PreprocessModuleDefinition(moduleDefinition)
	v.visitModuleDefinition(moduleDefinition)
	v.processor_.PostprocessModuleDefinition(moduleDefinition)

	// Visit slot 1 between references.
	v.processor_.ProcessModelSlot(1)

	// Visit the primitiveDefinitions rule.
	var primitiveDefinitions = model.GetPrimitiveDefinitions()
	v.processor_.PreprocessPrimitiveDefinitions(primitiveDefinitions)
	v.visitPrimitiveDefinitions(primitiveDefinitions)
	v.processor_.PostprocessPrimitiveDefinitions(primitiveDefinitions)

	// Visit slot 2 between references.
	v.processor_.ProcessModelSlot(2)

	// Visit the interfaceDefinitions rule.
	var interfaceDefinitions = model.GetInterfaceDefinitions()
	v.processor_.PreprocessInterfaceDefinitions(interfaceDefinitions)
	v.visitInterfaceDefinitions(interfaceDefinitions)
	v.processor_.PostprocessInterfaceDefinitions(interfaceDefinitions)
}

func (v *visitor_) visitModule(module ast.ModuleLike) {
	// Visit the name token.
	var name = module.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessModuleSlot(1)

	// Visit the path token.
	var path = module.GetPath()
	v.processor_.ProcessPath(path)
}

func (v *visitor_) visitModuleDefinition(moduleDefinition ast.ModuleDefinitionLike) {
	// Visit the notice rule.
	var notice = moduleDefinition.GetNotice()
	v.processor_.PreprocessNotice(notice)
	v.visitNotice(notice)
	v.processor_.PostprocessNotice(notice)

	// Visit slot 1 between references.
	v.processor_.ProcessModuleDefinitionSlot(1)

	// Visit the header rule.
	var header = moduleDefinition.GetHeader()
	v.processor_.PreprocessHeader(header)
	v.visitHeader(header)
	v.processor_.PostprocessHeader(header)

	// Visit slot 2 between references.
	v.processor_.ProcessModuleDefinitionSlot(2)

	// Visit the optional imports rule.
	var optionalImports = moduleDefinition.GetOptionalImports()
	if col.IsDefined(optionalImports) {
		v.processor_.PreprocessImports(optionalImports)
		v.visitImports(optionalImports)
		v.processor_.PostprocessImports(optionalImports)
	}
}

func (v *visitor_) visitNone(none ast.NoneLike) {
	// Visit the newline token.
	var newline = none.GetNewline()
	v.processor_.ProcessNewline(newline)
}

func (v *visitor_) visitNotice(notice ast.NoticeLike) {
	// Visit the comment token.
	var comment = notice.GetComment()
	v.processor_.ProcessComment(comment)
}

func (v *visitor_) visitParameter(parameter ast.ParameterLike) {
	// Visit the name token.
	var name = parameter.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessParameterSlot(1)

	// Visit the abstraction rule.
	var abstraction = parameter.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitParameterized(parameterized ast.ParameterizedLike) {
	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = parameterized.GetParameters().GetIterator()
	var parametersSize = uint(parameters.GetSize())
	for parameters.HasNext() {
		parameterIndex++
		var parameter = parameters.GetNext()
		v.processor_.PreprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(
			parameter,
			parameterIndex,
			parametersSize,
		)
	}
}

func (v *visitor_) visitPrefix(prefix ast.PrefixLike) {
	// Visit the possible prefix types.
	switch actual := prefix.GetAny().(type) {
	case ast.ArrayLike:
		v.processor_.PreprocessArray(actual)
		v.visitArray(actual)
		v.processor_.PostprocessArray(actual)
	case ast.MapLike:
		v.processor_.PreprocessMap(actual)
		v.visitMap(actual)
		v.processor_.PostprocessMap(actual)
	case ast.ChannelLike:
		v.processor_.PreprocessChannel(actual)
		v.visitChannel(actual)
		v.processor_.PostprocessChannel(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitPrimitiveDefinitions(primitiveDefinitions ast.PrimitiveDefinitionsLike) {
	// Visit the optional typeDefinitions rule.
	var optionalTypeDefinitions = primitiveDefinitions.GetOptionalTypeDefinitions()
	if col.IsDefined(optionalTypeDefinitions) {
		v.processor_.PreprocessTypeDefinitions(optionalTypeDefinitions)
		v.visitTypeDefinitions(optionalTypeDefinitions)
		v.processor_.PostprocessTypeDefinitions(optionalTypeDefinitions)
	}

	// Visit slot 1 between references.
	v.processor_.ProcessPrimitiveDefinitionsSlot(1)

	// Visit the optional functionalDefinitions rule.
	var optionalFunctionalDefinitions = primitiveDefinitions.GetOptionalFunctionalDefinitions()
	if col.IsDefined(optionalFunctionalDefinitions) {
		v.processor_.PreprocessFunctionalDefinitions(optionalFunctionalDefinitions)
		v.visitFunctionalDefinitions(optionalFunctionalDefinitions)
		v.processor_.PostprocessFunctionalDefinitions(optionalFunctionalDefinitions)
	}
}

func (v *visitor_) visitPublicMethods(publicMethods ast.PublicMethodsLike) {
	// Visit each method rule.
	var methodIndex uint
	var methods = publicMethods.GetMethods().GetIterator()
	var methodsSize = uint(methods.GetSize())
	for methods.HasNext() {
		methodIndex++
		var method = methods.GetNext()
		v.processor_.PreprocessMethod(
			method,
			methodIndex,
			methodsSize,
		)
		v.visitMethod(method)
		v.processor_.PostprocessMethod(
			method,
			methodIndex,
			methodsSize,
		)
	}
}

func (v *visitor_) visitResult(result ast.ResultLike) {
	// Visit the possible result types.
	switch actual := result.GetAny().(type) {
	case ast.NoneLike:
		v.processor_.PreprocessNone(actual)
		v.visitNone(actual)
		v.processor_.PostprocessNone(actual)
	case ast.AbstractionLike:
		v.processor_.PreprocessAbstraction(actual)
		v.visitAbstraction(actual)
		v.processor_.PostprocessAbstraction(actual)
	case ast.ParameterizedLike:
		v.processor_.PreprocessParameterized(actual)
		v.visitParameterized(actual)
		v.processor_.PostprocessParameterized(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitSuffix(suffix ast.SuffixLike) {
	// Visit the name token.
	var name = suffix.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitType(type_ ast.TypeLike) {
	// Visit the declaration rule.
	var declaration = type_.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessTypeSlot(1)

	// Visit the abstraction rule.
	var abstraction = type_.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)

	// Visit slot 2 between references.
	v.processor_.ProcessTypeSlot(2)

	// Visit the optional enumeration rule.
	var optionalEnumeration = type_.GetOptionalEnumeration()
	if col.IsDefined(optionalEnumeration) {
		v.processor_.PreprocessEnumeration(optionalEnumeration)
		v.visitEnumeration(optionalEnumeration)
		v.processor_.PostprocessEnumeration(optionalEnumeration)
	}
}

func (v *visitor_) visitTypeDefinitions(typeDefinitions ast.TypeDefinitionsLike) {
	// Visit each type rule.
	var typeIndex uint
	var types = typeDefinitions.GetTypes().GetIterator()
	var typesSize = uint(types.GetSize())
	for types.HasNext() {
		typeIndex++
		var type_ = types.GetNext()
		v.processor_.PreprocessType(
			type_,
			typeIndex,
			typesSize,
		)
		v.visitType(type_)
		v.processor_.PostprocessType(
			type_,
			typeIndex,
			typesSize,
		)
	}
}

func (v *visitor_) visitValue(value ast.ValueLike) {
	// Visit the name token.
	var name = value.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessValueSlot(1)

	// Visit the abstraction rule.
	var abstraction = value.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}
