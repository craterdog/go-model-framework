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
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-model-framework/v4/ast"
)

// CLASS INTERFACE

// Access Function

func Visitor() VisitorClassLike {
	return visitorClass
}

// Constructor Methods

func (c *visitorClass_) Make(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		class_:     c,
		processor_: processor,
	}
	return instance
}

// INSTANCE INTERFACE

// Public Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return v.getClass()
}

func (v *visitor_) VisitModel(
	model ast.ModelLike,
) {
	v.processor_.PreprocessModel(model)
	v.visitModel(model)
	v.processor_.PostprocessModel(model)
}

// Private Methods

func (v *visitor_) getClass() *visitorClass_ {
	return v.class_
}

func (v *visitor_) visitAbstraction(abstraction ast.AbstractionLike) {
	// Visit the optional prefix rule.
	var optionalPrefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(optionalPrefix) {
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
	if uti.IsDefined(optionalSuffix) {
		v.processor_.PreprocessSuffix(optionalSuffix)
		v.visitSuffix(optionalSuffix)
		v.processor_.PostprocessSuffix(optionalSuffix)
	}

	// Visit slot 3 between references.
	v.processor_.ProcessAbstractionSlot(3)

	// Visit the optional arguments rule.
	var optionalArguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
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

func (v *visitor_) visitArray(array ast.ArrayLike) {}

func (v *visitor_) visitAspectDefinition(aspectDefinition ast.AspectDefinitionLike) {
	// Visit the declaration rule.
	var declaration = aspectDefinition.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessAspectDefinitionSlot(1)

	// Visit each aspectMethod rule.
	var aspectMethodIndex uint
	var aspectMethods = aspectDefinition.GetAspectMethods().GetIterator()
	var aspectMethodsSize = uint(aspectMethods.GetSize())
	for aspectMethods.HasNext() {
		aspectMethodIndex++
		var aspectMethod = aspectMethods.GetNext()
		v.processor_.PreprocessAspectMethod(
			aspectMethod,
			aspectMethodIndex,
			aspectMethodsSize,
		)
		v.visitAspectMethod(aspectMethod)
		v.processor_.PostprocessAspectMethod(
			aspectMethod,
			aspectMethodIndex,
			aspectMethodsSize,
		)
	}
}

func (v *visitor_) visitAspectInterface(aspectInterface ast.AspectInterfaceLike) {
	// Visit the abstraction rule.
	var abstraction = aspectInterface.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitAspectMethod(aspectMethod ast.AspectMethodLike) {
	// Visit the method rule.
	var method = aspectMethod.GetMethod()
	v.processor_.PreprocessMethod(method)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(method)
}

func (v *visitor_) visitAspectSection(aspectSection ast.AspectSectionLike) {
	// Visit each aspectDefinition rule.
	var aspectDefinitionIndex uint
	var aspectDefinitions = aspectSection.GetAspectDefinitions().GetIterator()
	var aspectDefinitionsSize = uint(aspectDefinitions.GetSize())
	for aspectDefinitions.HasNext() {
		aspectDefinitionIndex++
		var aspectDefinition = aspectDefinitions.GetNext()
		v.processor_.PreprocessAspectDefinition(
			aspectDefinition,
			aspectDefinitionIndex,
			aspectDefinitionsSize,
		)
		v.visitAspectDefinition(aspectDefinition)
		v.processor_.PostprocessAspectDefinition(
			aspectDefinition,
			aspectDefinitionIndex,
			aspectDefinitionsSize,
		)
	}
}

func (v *visitor_) visitAspectSubsection(aspectSubsection ast.AspectSubsectionLike) {
	// Visit each interface rule.
	var interfaceIndex uint
	var interfaces = aspectSubsection.GetAspectInterfaces().GetIterator()
	var interfacesSize = uint(interfaces.GetSize())
	for interfaces.HasNext() {
		interfaceIndex++
		var aspectInterface = interfaces.GetNext()
		v.processor_.PreprocessAspectInterface(
			aspectInterface,
			interfaceIndex,
			interfacesSize,
		)
		v.visitAspectInterface(aspectInterface)
		v.processor_.PostprocessAspectInterface(
			aspectInterface,
			interfaceIndex,
			interfacesSize,
		)
	}
}

func (v *visitor_) visitAttributeMethod(attributeMethod ast.AttributeMethodLike) {
	// Visit the possible attributeMethod types.
	switch actual := attributeMethod.GetAny().(type) {
	case ast.GetterMethodLike:
		v.processor_.PreprocessGetterMethod(actual)
		v.visitGetterMethod(actual)
		v.processor_.PostprocessGetterMethod(actual)
	case ast.SetterMethodLike:
		v.processor_.PreprocessSetterMethod(actual)
		v.visitSetterMethod(actual)
		v.processor_.PostprocessSetterMethod(actual)
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitAttributeSubsection(attributeSubsection ast.AttributeSubsectionLike) {
	// Visit each attributeMethod rule.
	var attributeMethodIndex uint
	var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
	var attributeMethodsSize = uint(attributeMethods.GetSize())
	for attributeMethods.HasNext() {
		attributeMethodIndex++
		var attributeMethod = attributeMethods.GetNext()
		v.processor_.PreprocessAttributeMethod(
			attributeMethod,
			attributeMethodIndex,
			attributeMethodsSize,
		)
		v.visitAttributeMethod(attributeMethod)
		v.processor_.PostprocessAttributeMethod(
			attributeMethod,
			attributeMethodIndex,
			attributeMethodsSize,
		)
	}
}

func (v *visitor_) visitChannel(channel ast.ChannelLike) {
}

func (v *visitor_) visitClassDefinition(classDefinition ast.ClassDefinitionLike) {
	// Visit the declaration rule.
	var declaration = classDefinition.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessClassDefinitionSlot(1)

	// Visit the classMethods rule.
	var classMethods = classDefinition.GetClassMethods()
	v.processor_.PreprocessClassMethods(classMethods)
	v.visitClassMethods(classMethods)
	v.processor_.PostprocessClassMethods(classMethods)
}

func (v *visitor_) visitClassMethods(classMethods ast.ClassMethodsLike) {
	// Visit the constructorSubsection rule.
	var constructorSubsection = classMethods.GetConstructorSubsection()
	v.processor_.PreprocessConstructorSubsection(constructorSubsection)
	v.visitConstructorSubsection(constructorSubsection)
	v.processor_.PostprocessConstructorSubsection(constructorSubsection)

	// Visit slot 1 between references.
	v.processor_.ProcessClassMethodsSlot(1)

	// Visit the optional constantSubsection rule.
	var optionalConstantSubsection = classMethods.GetOptionalConstantSubsection()
	if uti.IsDefined(optionalConstantSubsection) {
		v.processor_.PreprocessConstantSubsection(optionalConstantSubsection)
		v.visitConstantSubsection(optionalConstantSubsection)
		v.processor_.PostprocessConstantSubsection(optionalConstantSubsection)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessClassMethodsSlot(2)

	// Visit the optional functionSubsection rule.
	var optionalFunctionSubsection = classMethods.GetOptionalFunctionSubsection()
	if uti.IsDefined(optionalFunctionSubsection) {
		v.processor_.PreprocessFunctionSubsection(optionalFunctionSubsection)
		v.visitFunctionSubsection(optionalFunctionSubsection)
		v.processor_.PostprocessFunctionSubsection(optionalFunctionSubsection)
	}
}

func (v *visitor_) visitClassSection(classSection ast.ClassSectionLike) {
	// Visit each classDefinition rule.
	var classDefinitionIndex uint
	var classDefinitions = classSection.GetClassDefinitions().GetIterator()
	var classDefinitionsSize = uint(classDefinitions.GetSize())
	for classDefinitions.HasNext() {
		classDefinitionIndex++
		var classDefinition = classDefinitions.GetNext()
		v.processor_.PreprocessClassDefinition(
			classDefinition,
			classDefinitionIndex,
			classDefinitionsSize,
		)
		v.visitClassDefinition(classDefinition)
		v.processor_.PostprocessClassDefinition(
			classDefinition,
			classDefinitionIndex,
			classDefinitionsSize,
		)
	}
}

func (v *visitor_) visitConstantMethod(constantMethod ast.ConstantMethodLike) {
	// Visit the name token.
	var name = constantMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstantMethodSlot(1)

	// Visit the abstraction rule.
	var abstraction = constantMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstantSubsection(constantSubsection ast.ConstantSubsectionLike) {
	// Visit each constantMethod rule.
	var constantMethodIndex uint
	var constantMethods = constantSubsection.GetConstantMethods().GetIterator()
	var constantMethodsSize = uint(constantMethods.GetSize())
	for constantMethods.HasNext() {
		constantMethodIndex++
		var constantMethod = constantMethods.GetNext()
		v.processor_.PreprocessConstantMethod(
			constantMethod,
			constantMethodIndex,
			constantMethodsSize,
		)
		v.visitConstantMethod(constantMethod)
		v.processor_.PostprocessConstantMethod(
			constantMethod,
			constantMethodIndex,
			constantMethodsSize,
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

func (v *visitor_) visitConstructorMethod(constructorMethod ast.ConstructorMethodLike) {
	// Visit the name token.
	var name = constructorMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessConstructorMethodSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = constructorMethod.GetParameters().GetIterator()
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
	v.processor_.ProcessConstructorMethodSlot(2)

	// Visit the abstraction rule.
	var abstraction = constructorMethod.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)
}

func (v *visitor_) visitConstructorSubsection(constructorSubsection ast.ConstructorSubsectionLike) {
	// Visit each constructorMethod rule.
	var constructorIndex uint
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	var constructorMethodsSize = uint(constructorMethods.GetSize())
	for constructorMethods.HasNext() {
		constructorIndex++
		var constructorMethod = constructorMethods.GetNext()
		v.processor_.PreprocessConstructorMethod(
			constructorMethod,
			constructorIndex,
			constructorMethodsSize,
		)
		v.visitConstructorMethod(constructorMethod)
		v.processor_.PostprocessConstructorMethod(
			constructorMethod,
			constructorIndex,
			constructorMethodsSize,
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
	if uti.IsDefined(optionalConstraints) {
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

func (v *visitor_) visitFunctionMethod(functionMethod ast.FunctionMethodLike) {
	// Visit the name token.
	var name = functionMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessFunctionMethodSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = functionMethod.GetParameters().GetIterator()
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
	v.processor_.ProcessFunctionMethodSlot(2)

	// Visit the result rule.
	var result = functionMethod.GetResult()
	v.processor_.PreprocessResult(result)
	v.visitResult(result)
	v.processor_.PostprocessResult(result)
}

func (v *visitor_) visitFunctionSubsection(functionSubsection ast.FunctionSubsectionLike) {
	// Visit each functionMethod rule.
	var functionMethodIndex uint
	var functionMethods = functionSubsection.GetFunctionMethods().GetIterator()
	var functionMethodsSize = uint(functionMethods.GetSize())
	for functionMethods.HasNext() {
		functionMethodIndex++
		var functionMethod = functionMethods.GetNext()
		v.processor_.PreprocessFunctionMethod(
			functionMethod,
			functionMethodIndex,
			functionMethodsSize,
		)
		v.visitFunctionMethod(functionMethod)
		v.processor_.PostprocessFunctionMethod(
			functionMethod,
			functionMethodIndex,
			functionMethodsSize,
		)
	}
}

func (v *visitor_) visitFunctionalDefinition(functionalDefinition ast.FunctionalDefinitionLike) {
	// Visit the declaration rule.
	var declaration = functionalDefinition.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessFunctionalDefinitionSlot(1)

	// Visit each parameter rule.
	var parameterIndex uint
	var parameters = functionalDefinition.GetParameters().GetIterator()
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
	v.processor_.ProcessFunctionalDefinitionSlot(2)

	// Visit the result rule.
	var result = functionalDefinition.GetResult()
	v.processor_.PreprocessResult(result)
	v.visitResult(result)
	v.processor_.PostprocessResult(result)
}

func (v *visitor_) visitFunctionalSection(functionalSection ast.FunctionalSectionLike) {
	// Visit each functionalDefinition rule.
	var functionalDefinitionIndex uint
	var functionalDefinitions = functionalSection.GetFunctionalDefinitions().GetIterator()
	var functionalDefinitionsSize = uint(functionalDefinitions.GetSize())
	for functionalDefinitions.HasNext() {
		functionalDefinitionIndex++
		var functionalDefinition = functionalDefinitions.GetNext()
		v.processor_.PreprocessFunctionalDefinition(
			functionalDefinition,
			functionalDefinitionIndex,
			functionalDefinitionsSize,
		)
		v.visitFunctionalDefinition(functionalDefinition)
		v.processor_.PostprocessFunctionalDefinition(
			functionalDefinition,
			functionalDefinitionIndex,
			functionalDefinitionsSize,
		)
	}
}

func (v *visitor_) visitGetterMethod(getterMethod ast.GetterMethodLike) {
	// Visit the name token.
	var name = getterMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessGetterMethodSlot(1)

	// Visit the single abstraction rule.
	var abstraction = getterMethod.GetAbstraction()
	if uti.IsDefined(abstraction) {
		v.processor_.PreprocessAbstraction(abstraction)
		v.visitAbstraction(abstraction)
		v.processor_.PostprocessAbstraction(abstraction)
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

func (v *visitor_) visitInstanceDefinition(instanceDefinition ast.InstanceDefinitionLike) {
	// Visit the declaration rule.
	var declaration = instanceDefinition.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessInstanceDefinitionSlot(1)

	// Visit the instanceMethods rule.
	var instanceMethods = instanceDefinition.GetInstanceMethods()
	v.processor_.PreprocessInstanceMethods(instanceMethods)
	v.visitInstanceMethods(instanceMethods)
	v.processor_.PostprocessInstanceMethods(instanceMethods)
}

func (v *visitor_) visitInstanceMethods(instanceMethods ast.InstanceMethodsLike) {
	// Visit the publicSubsection rule.
	var publicSubsection = instanceMethods.GetPublicSubsection()
	v.processor_.PreprocessPublicSubsection(publicSubsection)
	v.visitPublicSubsection(publicSubsection)
	v.processor_.PostprocessPublicSubsection(publicSubsection)

	// Visit slot 1 between references.
	v.processor_.ProcessInstanceMethodsSlot(1)

	// Visit the optional attributeSubsection rule.
	var optionalAttributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(optionalAttributeSubsection) {
		v.processor_.PreprocessAttributeSubsection(optionalAttributeSubsection)
		v.visitAttributeSubsection(optionalAttributeSubsection)
		v.processor_.PostprocessAttributeSubsection(optionalAttributeSubsection)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessInstanceMethodsSlot(2)

	// Visit the optional aspectSubsection rule.
	var optionalAspectSubsection = instanceMethods.GetOptionalAspectSubsection()
	if uti.IsDefined(optionalAspectSubsection) {
		v.processor_.PreprocessAspectSubsection(optionalAspectSubsection)
		v.visitAspectSubsection(optionalAspectSubsection)
		v.processor_.PostprocessAspectSubsection(optionalAspectSubsection)
	}
}

func (v *visitor_) visitInstanceSection(instanceSection ast.InstanceSectionLike) {
	// Visit each instanceDefinition rule.
	var instanceDefinitionIndex uint
	var instanceDefinitions = instanceSection.GetInstanceDefinitions().GetIterator()
	var instanceDefinitionsSize = uint(instanceDefinitions.GetSize())
	for instanceDefinitions.HasNext() {
		instanceDefinitionIndex++
		var instanceDefinition = instanceDefinitions.GetNext()
		v.processor_.PreprocessInstanceDefinition(
			instanceDefinition,
			instanceDefinitionIndex,
			instanceDefinitionsSize,
		)
		v.visitInstanceDefinition(instanceDefinition)
		v.processor_.PostprocessInstanceDefinition(
			instanceDefinition,
			instanceDefinitionIndex,
			instanceDefinitionsSize,
		)
	}
}

func (v *visitor_) visitInterfaceDefinitions(interfaceDefinitions ast.InterfaceDefinitionsLike) {
	// Visit the classSection rule.
	var classSection = interfaceDefinitions.GetClassSection()
	v.processor_.PreprocessClassSection(classSection)
	v.visitClassSection(classSection)
	v.processor_.PostprocessClassSection(classSection)

	// Visit slot 1 between references.
	v.processor_.ProcessInterfaceDefinitionsSlot(1)

	// Visit the instanceSection rule.
	var instanceSection = interfaceDefinitions.GetInstanceSection()
	v.processor_.PreprocessInstanceSection(instanceSection)
	v.visitInstanceSection(instanceSection)
	v.processor_.PostprocessInstanceSection(instanceSection)

	// Visit slot 2 between references.
	v.processor_.ProcessInterfaceDefinitionsSlot(2)

	// Visit the optional aspectSection rule.
	var optionalAspectSection = interfaceDefinitions.GetOptionalAspectSection()
	if uti.IsDefined(optionalAspectSection) {
		v.processor_.PreprocessAspectSection(optionalAspectSection)
		v.visitAspectSection(optionalAspectSection)
		v.processor_.PostprocessAspectSection(optionalAspectSection)
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
	if uti.IsDefined(optionalResult) {
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
	if uti.IsDefined(optionalImports) {
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
	// Visit the optional typeSection rule.
	var optionalTypeSection = primitiveDefinitions.GetOptionalTypeSection()
	if uti.IsDefined(optionalTypeSection) {
		v.processor_.PreprocessTypeSection(optionalTypeSection)
		v.visitTypeSection(optionalTypeSection)
		v.processor_.PostprocessTypeSection(optionalTypeSection)
	}

	// Visit slot 1 between references.
	v.processor_.ProcessPrimitiveDefinitionsSlot(1)

	// Visit the optional functionalSection rule.
	var optionalFunctionalSection = primitiveDefinitions.GetOptionalFunctionalSection()
	if uti.IsDefined(optionalFunctionalSection) {
		v.processor_.PreprocessFunctionalSection(optionalFunctionalSection)
		v.visitFunctionalSection(optionalFunctionalSection)
		v.processor_.PostprocessFunctionalSection(optionalFunctionalSection)
	}
}

func (v *visitor_) visitPublicMethod(publicMethod ast.PublicMethodLike) {
	// Visit the method rule.
	var method = publicMethod.GetMethod()
	v.processor_.PreprocessMethod(method)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(method)
}

func (v *visitor_) visitPublicSubsection(publicSubsection ast.PublicSubsectionLike) {
	// Visit each publicMethod rule.
	var publicMethodIndex uint
	var publicMethods = publicSubsection.GetPublicMethods().GetIterator()
	var publicMethodsSize = uint(publicMethods.GetSize())
	for publicMethods.HasNext() {
		publicMethodIndex++
		var publicMethod = publicMethods.GetNext()
		v.processor_.PreprocessPublicMethod(
			publicMethod,
			publicMethodIndex,
			publicMethodsSize,
		)
		v.visitPublicMethod(publicMethod)
		v.processor_.PostprocessPublicMethod(
			publicMethod,
			publicMethodIndex,
			publicMethodsSize,
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

func (v *visitor_) visitSetterMethod(setterMethod ast.SetterMethodLike) {
	// Visit the name token.
	var name = setterMethod.GetName()
	v.processor_.ProcessName(name)

	// Visit slot 1 between references.
	v.processor_.ProcessSetterMethodSlot(1)

	// Visit the parameter rule.
	var parameter = setterMethod.GetParameter()
	if uti.IsDefined(parameter) {
		v.processor_.PreprocessParameter(parameter, 1, 1)
		v.visitParameter(parameter)
		v.processor_.PostprocessParameter(parameter, 1, 1)
	}
}

func (v *visitor_) visitSuffix(suffix ast.SuffixLike) {
	// Visit the name token.
	var name = suffix.GetName()
	v.processor_.ProcessName(name)
}

func (v *visitor_) visitTypeDefinition(typeDefinition ast.TypeDefinitionLike) {
	// Visit the declaration rule.
	var declaration = typeDefinition.GetDeclaration()
	v.processor_.PreprocessDeclaration(declaration)
	v.visitDeclaration(declaration)
	v.processor_.PostprocessDeclaration(declaration)

	// Visit slot 1 between references.
	v.processor_.ProcessTypeDefinitionSlot(1)

	// Visit the abstraction rule.
	var abstraction = typeDefinition.GetAbstraction()
	v.processor_.PreprocessAbstraction(abstraction)
	v.visitAbstraction(abstraction)
	v.processor_.PostprocessAbstraction(abstraction)

	// Visit slot 2 between references.
	v.processor_.ProcessTypeDefinitionSlot(2)

	// Visit the optional enumeration rule.
	var optionalEnumeration = typeDefinition.GetOptionalEnumeration()
	if uti.IsDefined(optionalEnumeration) {
		v.processor_.PreprocessEnumeration(optionalEnumeration)
		v.visitEnumeration(optionalEnumeration)
		v.processor_.PostprocessEnumeration(optionalEnumeration)
	}
}

func (v *visitor_) visitTypeSection(typeSection ast.TypeSectionLike) {
	// Visit each typeDefinition rule.
	var typeDefinitionIndex uint
	var typeDefinitions = typeSection.GetTypeDefinitions().GetIterator()
	var typeDefinitionsSize = uint(typeDefinitions.GetSize())
	for typeDefinitions.HasNext() {
		typeDefinitionIndex++
		var typeDefinition = typeDefinitions.GetNext()
		v.processor_.PreprocessTypeDefinition(
			typeDefinition,
			typeDefinitionIndex,
			typeDefinitionsSize,
		)
		v.visitTypeDefinition(typeDefinition)
		v.processor_.PostprocessTypeDefinition(
			typeDefinition,
			typeDefinitionIndex,
			typeDefinitionsSize,
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

// PRIVATE INTERFACE

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	class_     *visitorClass_
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

var visitorClass = &visitorClass_{
	// Initialize the class constants.
}
