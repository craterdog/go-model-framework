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

package agent

import (
	fmt "fmt"
	cdc "github.com/craterdog/go-collection-framework/v4/cdcn"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	sts "strings"
	tim "time"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// Initialize the class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	// Define the instance attributes.
	class_ GeneratorClassLike
}

// Attributes

func (v *generator_) GetClass() GeneratorClassLike {
	return v.class_
}

// Public

func (v *generator_) CreateClassType(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(angleTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) CreateGenericType(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(arrayTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) CreateClassStructure(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(complexTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) CreateGenericStructure(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(catalogTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) GenerateClass(
	model ast.ModelLike,
	name string,
) (
	implementation string,
) {
	var classIterator = model.GetClasses().GetClasses().GetIterator()
	var instanceIterator = model.GetInstances().GetInstances().GetIterator()
	for classIterator.HasNext() && instanceIterator.HasNext() {
		var class = classIterator.GetNext()
		var className = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetName(),
			"ClassLike",
		))
		var instance = instanceIterator.GetNext()
		var instanceName = sts.ToLower(sts.TrimSuffix(
			instance.GetDeclaration().GetName(),
			"Like",
		))
		if className == name && instanceName == name {
			implementation = v.generateClass(model, class, instance)
			return implementation
		}
	}
	var message = fmt.Sprintf(
		"The following class does not exist in the model: %v",
		name,
	)
	panic(message)
}

// Private

func (v *generator_) expandCopyright(copyright string) string {
	var maximum = 78
	var length = len(copyright)
	if length > maximum {
		var message = fmt.Sprintf(
			"The copyright notice cannot be longer than 78 characters: %v",
			copyright,
		)
		panic(message)
	}
	if length == 0 {
		copyright = fmt.Sprintf(
			"Copyright (c) %v.  All Rights Reserved.",
			tim.Now().Year(),
		)
		length = len(copyright)
	}
	var padding = (maximum - length) / 2
	for range padding {
		copyright = " " + copyright + " "
	}
	if len(copyright) < maximum {
		copyright = " " + copyright
	}
	copyright = "." + copyright + "."
	return copyright
}

func (v *generator_) extractArguments(
	parameters ast.ParametersLike,
) ast.ArgumentsLike {
	var notation = cdc.Notation().Make()
	var additionalArguments = col.List[ast.AdditionalArgumentLike](notation).Make()

	// Extract the first argument.
	var parameter = parameters.GetParameter()
	var abstraction = ast.Abstraction().Make(nil, nil, parameter.GetName(), nil)
	var argument = ast.Argument().Make(abstraction)

	// Extract any additional arguments.
	var additionalParameters = parameters.GetAdditionalParameters()
	if additionalParameters != nil {
		var iterator = additionalParameters.GetIterator()
		for iterator.HasNext() {
			parameter = iterator.GetNext().GetParameter()
			abstraction = ast.Abstraction().Make(nil, nil, parameter.GetName(), nil)
			var additionalArgument = ast.AdditionalArgument().Make(
				ast.Argument().Make(abstraction),
			)
			additionalArguments.AppendValue(additionalArgument)
		}
	}
	var arguments = ast.Arguments().Make(argument, additionalArguments)
	return arguments
}

func (v *generator_) extractConcreteMappings(
	parameters ast.ParametersLike,
	arguments ast.ArgumentsLike,
) col.CatalogLike[string, ast.AbstractionLike] {
	// Create the mappings catalog.
	var notation = cdc.Notation().Make()
	var mappings = col.Catalog[string, ast.AbstractionLike](notation).Make()

	// Map the name of the first parameter to its concrete type.
	var parameter = parameters.GetParameter()
	var name = parameter.GetName()
	var argument = arguments.GetArgument()
	var concreteType = argument.GetAbstraction()
	mappings.SetValue(name, concreteType)

	// Map the name of the additional parameters to their concrete types.
	var parameterIterator = parameters.GetAdditionalParameters().GetIterator()
	var argumentIterator = arguments.GetAdditionalArguments().GetIterator()
	for parameterIterator.HasNext() {
		var additionalParameter = parameterIterator.GetNext()
		parameter = additionalParameter.GetParameter()
		name = parameter.GetName()
		var additionalArgument = argumentIterator.GetNext()
		argument = additionalArgument.GetArgument()
		concreteType = argument.GetAbstraction()
		mappings.SetValue(name, concreteType)
	}

	return mappings
}

func (v *generator_) extractConstructorAttributes(
	class ast.ClassLike,
	catalog col.CatalogLike[string, string],
) {
	var constructors = class.GetConstructors()
	if constructors == nil {
		return
	}
	var iterator = constructors.GetConstructors().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var methodName = constructor.GetName()
		if methodName == "Make" || sts.HasPrefix(methodName, "MakeWith") {
			var parameters = constructor.GetOptionalParameters()
			if parameters != nil {
				v.extractParameterAttributes(parameters, catalog)
			}
		}
	}
}

func (v *generator_) extractInstanceAttributes(
	instance ast.InstanceLike,
	catalog col.CatalogLike[string, string],
) {
	var attributeName string
	var attributeType string
	var formatter = Formatter().Make()
	var attributes = instance.GetAttributes()
	if attributes == nil {
		return
	}
	var iterator = attributes.GetAttributes().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var name = attribute.GetName()
		var abstraction ast.AbstractionLike
		switch {
		case sts.HasPrefix(name, "Is"):
			attributeName = sts.TrimPrefix(name, "Is")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Was"):
			attributeName = sts.TrimPrefix(name, "Was")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Are"):
			attributeName = sts.TrimPrefix(name, "Are")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Were"):
			attributeName = sts.TrimPrefix(name, "Were")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Has"):
			attributeName = sts.TrimPrefix(name, "Has")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Had"):
			attributeName = sts.TrimPrefix(name, "Had")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Have"):
			attributeName = sts.TrimPrefix(name, "Have")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Get"):
			attributeName = sts.TrimPrefix(name, "Get")
			abstraction = attribute.GetOptionalAbstraction()
		case sts.HasPrefix(name, "Set"):
			attributeName = sts.TrimPrefix(name, "Set")
			var parameter = attribute.GetOptionalParameter()
			abstraction = parameter.GetAbstraction()
		}
		attributeName = v.makePrivate(attributeName)
		attributeType = formatter.FormatAbstraction(abstraction)
		catalog.SetValue(attributeName, attributeType)
	}
}

func (v *generator_) extractParameterAttribute(
	parameter ast.ParameterLike,
	catalog col.CatalogLike[string, string],
) {
	var parameterName = parameter.GetName()
	parameterName = sts.TrimSuffix(parameterName, "_")
	var abstraction = parameter.GetAbstraction()
	var formatter = Formatter().Make()
	var parameterType = formatter.FormatAbstraction(abstraction)
	catalog.SetValue(parameterName, parameterType)
}

func (v *generator_) extractParameterAttributes(
	parameters ast.ParametersLike,
	catalog col.CatalogLike[string, string],
) {
	var parameter = parameters.GetParameter()
	v.extractParameterAttribute(parameter, catalog)
	var iterator = parameters.GetAdditionalParameters().GetIterator()
	for iterator.HasNext() {
		parameter = iterator.GetNext().GetParameter()
		v.extractParameterAttribute(parameter, catalog)
	}
}

func (v *generator_) extractTargetType(
	class ast.ClassLike,
	instance ast.InstanceLike,
) (
	targetType string,
) {
	var sequence = instance.GetAttributes().GetAttributes()
	if sequence.GetSize() == 1 {
		// This class has no instance attributes.
		var iterator = class.GetConstructors().GetConstructors().GetIterator()
		for iterator.HasNext() {
			var constructor = iterator.GetNext()
			var name = constructor.GetName()
			if name == "MakeFromValue" {
				var parameter = constructor.GetOptionalParameters().GetParameter()
				var abstraction = parameter.GetAbstraction()
				var formatter = Formatter().Make()
				targetType = formatter.FormatAbstraction(abstraction)
				break
			}
		}
	}
	return targetType
}

func (v *generator_) generateAbstractionMethods(
	targetType string,
	aspect ast.AspectLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) (implementation string) {
	// Generate the method implementations for the aspect.
	var iterator = aspect.GetMethods().GetIterator()
	for iterator.HasNext() {
		var aspectMethod = iterator.GetNext()
		var methodImplementation = v.generateMethodImplementation(
			targetType,
			aspectMethod,
			mappings,
		)
		implementation += methodImplementation
	}

	return implementation
}

func (v *generator_) generateAbstractions(
	targetType string,
	model ast.ModelLike,
	instance ast.InstanceLike,
) (
	implementation string,
) {
	// Check for no aspect abstractions.
	var abstractions = instance.GetOptionalAbstractions()
	if abstractions == nil {
		return implementation
	}

	// Generate the methods for each aspect abstraction.
	var formatter = Formatter().Make()
	var iterator = abstractions.GetAbstractions().GetIterator()
	for iterator.HasNext() {
		// Each aspect abstraction binds to its own concrete arguments.
		var abstraction = iterator.GetNext()
		var aspectName = formatter.FormatAbstraction(abstraction)
		var instanceAspect = instanceAspectTemplate_
		instanceAspect = sts.ReplaceAll(instanceAspect, "<AspectName>", aspectName)
		var methods string
		if abstraction.GetOptionalAlias() == nil {
			// We only know the method signatures for the local aspects.
			var mappings col.CatalogLike[string, ast.AbstractionLike]
			var aspect = v.retrieveAspect(model, abstraction.GetName())
			var declaration = aspect.GetDeclaration()
			if declaration.GetOptionalGenericParameters() != nil {
				var parameters = declaration.GetOptionalGenericParameters().GetParameters()
				var arguments = abstraction.GetOptionalGenericArguments().GetArguments()
				mappings = v.extractConcreteMappings(parameters, arguments)
			}
			methods = v.generateAbstractionMethods(targetType, aspect, mappings)
		}
		instanceAspect = sts.ReplaceAll(instanceAspect, "<Methods>", methods)
		implementation += instanceAspect
	}

	return implementation
}

func (v *generator_) generateAttributeAssignment(
	parameter ast.ParameterLike,
) (
	implementation string,
) {
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	implementation = attributeAssignmentTemplate_
	implementation = sts.ReplaceAll(implementation, "<AttributeName>", attributeName)
	implementation = sts.ReplaceAll(implementation, "<ParameterName>", parameterName)
	return implementation
}

func (v *generator_) generateAttributeAssignments(
	class ast.ClassLike,
	constructor ast.ConstructorLike,
) (
	implementation string,
) {
	var name = constructor.GetName()
	if sts.HasPrefix(name, "MakeFrom") {
		return implementation
	}
	var parameters = constructor.GetOptionalParameters()
	if parameters == nil {
		return implementation
	}
	var parameter = parameters.GetParameter()
	var assignment = v.generateAttributeAssignment(parameter)
	implementation += assignment
	var additionalParameters = parameters.GetAdditionalParameters()
	var iterator = additionalParameters.GetIterator()
	for iterator.HasNext() {
		var additionalParameter = iterator.GetNext()
		parameter = additionalParameter.GetParameter()
		assignment = v.generateAttributeAssignment(parameter)
		implementation += assignment
	}
	return implementation
}

func (v *generator_) generateAttributeMethods(
	targetType string,
	class ast.ClassLike,
	instance ast.InstanceLike,
) (
	implementation string,
) {
	var formatter = Formatter().Make()
	var instanceAttributes = instance.GetAttributes()
	if instanceAttributes == nil {
		return implementation
	}
	implementation = "\n// Attributes\n"
	var iterator = instanceAttributes.GetAttributes().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var methodName = attribute.GetName()
		var attributeName string
		var body string

		var parameter string
		var attributeParameter = attribute.GetOptionalParameter()
		var parameterName string
		if attributeParameter != nil {
			attributeName = sts.TrimPrefix(methodName, "Set")
			parameterName = attributeParameter.GetName()
			parameter = formatter.FormatParameter(attributeParameter)
			body = setterBodyTemplate_
		}

		var resultType string
		var abstraction = attribute.GetOptionalAbstraction()
		if abstraction != nil {
			switch {
			case sts.HasPrefix(methodName, "Get"):
				attributeName = sts.TrimPrefix(methodName, "Get")
			case sts.HasPrefix(methodName, "Is"):
				attributeName = sts.TrimPrefix(methodName, "Is")
			case sts.HasPrefix(methodName, "Was"):
				attributeName = sts.TrimPrefix(methodName, "Was")
			case sts.HasPrefix(methodName, "Has"):
				attributeName = sts.TrimPrefix(methodName, "Has")
			}
			resultType = " " + formatter.FormatAbstraction(abstraction)
			body = getterBodyTemplate_
			if len(targetType) > 0 {
				body = getterClassTemplate_
			}
		}

		attributeName = v.makePrivate(attributeName)
		body = sts.ReplaceAll(body, "<AttributeName>", attributeName)
		body = sts.ReplaceAll(body, "<ParameterName>", parameterName)
		var method = instanceMethodTemplate_
		if len(targetType) > 0 {
			method = typeMethodTemplate_
		}
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameter)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		implementation += method
	}
	return implementation
}

func (v *generator_) generateClass(
	model ast.ModelLike,
	class ast.ClassLike,
	instance ast.InstanceLike,
) (
	implementation string,
) {
	// Generate the class model template.
	implementation = classTemplate_
	var notice = v.generateNotice(model)
	implementation = sts.ReplaceAll(implementation, "<Notice>", notice)
	var header = v.generateHeader(model)
	implementation = sts.ReplaceAll(implementation, "<Header>", header)
	var classAccess = v.generateClassAccess(class)
	implementation = sts.ReplaceAll(implementation, "<Access>", classAccess)

	// Insert the class methods.
	var targetType = v.extractTargetType(class, instance)
	var classMethods = v.generateClassMethods(targetType, class)
	implementation = sts.ReplaceAll(implementation, "<Class>", classMethods)

	// Insert the instance methods.
	var instanceMethods = v.generateInstanceMethods(
		targetType,
		model,
		class,
		instance,
	)
	implementation = sts.ReplaceAll(implementation, "<Instance>", instanceMethods)

	// Insert the actual class name into the template.
	var classDeclaration = class.GetDeclaration()
	var className = classDeclaration.GetName()
	className = sts.TrimSuffix(className, "ClassLike")
	implementation = sts.ReplaceAll(implementation, "<ClassName>", className)
	implementation = sts.ReplaceAll(implementation, "<TargetName>", v.makePrivate(className))

	// Insert any generic parameters and arguments into the template.
	var parameters string
	var arguments string
	var genericParameters = classDeclaration.GetOptionalGenericParameters()
	if genericParameters != nil {
		var classParameters = genericParameters.GetParameters()
		var classArguments = v.extractArguments(classParameters)
		var formatter = Formatter().Make()
		parameters = "[" + formatter.FormatParameters(classParameters) + "]"
		arguments = "[" + formatter.FormatArguments(classArguments) + "]"
	}
	implementation = sts.ReplaceAll(implementation, "[<Parameters>]", parameters)
	implementation = sts.ReplaceAll(implementation, "[<Arguments>]", arguments)

	// Insert any imported modules into the template.
	var imports = v.generateImports(model, implementation)
	implementation = sts.ReplaceAll(implementation, "<Imports>", imports)

	// Return the generated class implementation.
	return implementation
}

func (v *generator_) generateClassAccess(
	class ast.ClassLike,
) (
	implementation string,
) {
	var declaration = class.GetDeclaration()
	var genericParameters = declaration.GetOptionalGenericParameters()
	var reference = classReferenceTemplate_
	var function = classFunctionTemplate_
	if genericParameters != nil {
		reference = genericReferenceTemplate_
		function = genericFunctionTemplate_
	}
	implementation = classAccessTemplate_
	implementation = sts.ReplaceAll(implementation, "<Reference>", reference)
	implementation = sts.ReplaceAll(implementation, "<Function>", function)
	return implementation
}

func (v *generator_) generateClassConstants(
	class ast.ClassLike,
) (
	implementation string,
) {
	var formatter = Formatter().Make()
	var classConstants = class.GetOptionalConstants()
	if classConstants == nil {
		return implementation
	}
	var iterator = classConstants.GetConstants().GetIterator()
	for iterator.HasNext() {
		var classConstant = iterator.GetNext()
		var constantName = classConstant.GetName()
		var constantAbstraction = classConstant.GetAbstraction()
		constantName = v.makePrivate(constantName)
		var constantType = formatter.FormatAbstraction(constantAbstraction)
		var constant = classConstantTemplate_
		constant = sts.ReplaceAll(constant, "<ConstantName>", constantName)
		constant = sts.ReplaceAll(constant, "<ConstantType>", constantType)
		implementation += constant
	}
	return implementation
}

func (v *generator_) generateClassMethods(
	targetType string,
	class ast.ClassLike,
) (
	implementation string,
) {
	implementation = classMethodsTemplate_
	var target = v.generateClassTarget(class)
	implementation = sts.ReplaceAll(implementation, "<Target>", target)
	var constantMethods = v.generateConstantMethods(class)
	implementation = sts.ReplaceAll(implementation, "<Constants>", constantMethods)
	var constructorMethods = v.generateConstructorMethods(targetType, class)
	implementation = sts.ReplaceAll(implementation, "<Constructors>", constructorMethods)
	var functionMethods = v.generateFunctionMethods(class)
	implementation = sts.ReplaceAll(implementation, "<Functions>", functionMethods)
	return implementation
}

func (v *generator_) generateClassTarget(
	class ast.ClassLike,
) (
	implementation string,
) {
	implementation = classTargetTemplate_
	var constants = v.generateClassConstants(class)
	implementation = sts.ReplaceAll(implementation, "<Constants>", constants)
	return implementation
}

func (v *generator_) generateConstantMethods(
	class ast.ClassLike,
) (
	implementation string,
) {
	var formatter = Formatter().Make()
	var classConstants = class.GetOptionalConstants()
	if classConstants == nil {
		return implementation
	}
	implementation = "\n// Constants\n"
	var iterator = classConstants.GetConstants().GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		var methodName = constant.GetName()
		var constantName = v.makePrivate(methodName)
		var abstraction = constant.GetAbstraction()
		var resultType = " " + formatter.FormatAbstraction(abstraction)
		var body = constantBodyTemplate_
		body = sts.ReplaceAll(body, "<ConstantName>", constantName)
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", "")
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		implementation += method
	}
	return implementation
}

func (v *generator_) generateConstructorMethods(
	targetType string,
	class ast.ClassLike,
) (
	implementation string,
) {
	var formatter = Formatter().Make()
	var classConstructors = class.GetConstructors()
	if classConstructors == nil {
		return implementation
	}
	implementation = "\n// Constructors\n"
	var iterator = classConstructors.GetConstructors().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var methodName = constructor.GetName()
		var ConstructorParameters = constructor.GetOptionalParameters()
		var parameters string
		if ConstructorParameters != nil {
			parameters = formatter.FormatParameters(ConstructorParameters)
		}
		var abstraction = constructor.GetAbstraction()
		var resultType = " " + formatter.FormatAbstraction(abstraction)
		var assignments = v.generateAttributeAssignments(class, constructor)
		var body = constructorBodyTemplate_
		if len(targetType) > 0 {
			if methodName == "MakeFromValue" {
				body = typeBodyTemplate_
				body = sts.ReplaceAll(body, "<TargetType>", targetType)
			} else {
				body = resultBodyTemplate_
			}
		}
		body = sts.ReplaceAll(body, "<Assignments>", assignments)
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		implementation += method
	}
	return implementation
}

func (v *generator_) generateFunctionMethods(
	class ast.ClassLike,
) (
	implementation string,
) {
	var formatter = Formatter().Make()
	var classFunctions = class.GetOptionalFunctions()
	if classFunctions == nil {
		return implementation
	}
	implementation = "\n// Functions\n"
	var iterator = classFunctions.GetFunctions().GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		var name = function.GetName()
		var functionParameters = function.GetOptionalParameters()
		var parameters string
		if functionParameters != nil {
			parameters = formatter.FormatParameters(functionParameters)
		}
		var result = function.GetResult()
		var resultType = " " + formatter.FormatResult(result)
		var body = functionBodyTemplate_
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", name)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		implementation += method
	}
	return implementation
}

func (v *generator_) generateHeader(
	model ast.ModelLike,
) (
	implementation string,
) {
	var name = model.GetHeader().GetName()
	var header = headerTemplate_
	header = sts.ReplaceAll(header, "<Name>", name)
	return header
}

func (v *generator_) generateImports(
	model ast.ModelLike,
	class string,
) (
	implementation string,
) {
	var imports = model.GetOptionalImports()
	if imports != nil || sts.Contains(class, "syn.") {
		var modules = v.generateModules(imports, class)
		implementation = importsTemplate_
		implementation = sts.ReplaceAll(implementation, "<Modules>", modules)
	}
	return implementation
}

func (v *generator_) generateInstanceAttributes(
	class ast.ClassLike,
	instance ast.InstanceLike,
) string {
	var attributes string
	var notation = cdc.Notation().Make()
	var catalog = col.Catalog[string, string](notation).Make()
	v.extractInstanceAttributes(instance, catalog)
	v.extractConstructorAttributes(class, catalog)
	if catalog.IsEmpty() {
		attributes = "\n\t// TBA - Add private instance attributes.\n"
		return attributes
	}
	var iterator = catalog.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var attributeName = association.GetKey()
		var attributeType = association.GetValue()
		var attribute = instanceAttributeTemplate_
		attribute = sts.ReplaceAll(attribute, "<AttributeName>", attributeName)
		attribute = sts.ReplaceAll(attribute, "<AttributeType>", attributeType)
		attributes += attribute
	}
	return attributes
}

func (v *generator_) generateInstanceMethods(
	targetType string,
	model ast.ModelLike,
	class ast.ClassLike,
	instance ast.InstanceLike,
) string {
	var instanceMethods = instanceMethodsTemplate_
	var attributes = v.generateAttributeMethods(targetType, class, instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Attributes>", attributes)
	var abstractions = v.generateAbstractions(targetType, model, instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Abstractions>", abstractions)
	var methods = v.generatePublicMethods(targetType, instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Methods>", methods)
	var target = v.generateInstanceTarget(targetType, class, instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Target>", target)
	return instanceMethods
}

func (v *generator_) generateInstanceTarget(
	targetType string,
	class ast.ClassLike,
	instance ast.InstanceLike,
) string {
	var target = instanceTargetTemplate_
	if len(targetType) > 0 {
		target = typeTargetTemplate_
	}
	target = sts.ReplaceAll(target, "<TargetType>", targetType)
	var attributes = v.generateInstanceAttributes(class, instance)
	target = sts.ReplaceAll(target, "<Attributes>", attributes)
	return target
}

func (v *generator_) generateMethodImplementation(
	targetType string,
	method ast.MethodLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) (
	implementation string,
) {
	// Choose the right method template.
	implementation = instanceMethodTemplate_
	if len(targetType) > 0 {
		implementation = typeMethodTemplate_
	}
	var methodName = method.GetName()
	implementation = sts.ReplaceAll(implementation, "<MethodName>", methodName)

	// Generate the right method body.
	var body = methodBodyTemplate_
	var methodResult = method.GetOptionalResult()
	if methodResult != nil {
		switch actual := methodResult.GetAny().(type) {
		case ast.AbstractionLike:
			body = resultBodyTemplate_
		case ast.ParameterizedLike:
			body = returnBodyTemplate_
		default:
			var message = fmt.Sprintf(
				"An unknown method result type was found: %T",
				actual,
			)
			panic(message)
		}
	}
	implementation = sts.ReplaceAll(implementation, "<Body>", body)

	// Generate the method parameters.
	var parameters string
	var formatter = Formatter().Make()
	var methodParameters = method.GetOptionalParameters()
	if methodParameters != nil {
		if mappings != nil && mappings.GetSize() > 0 {
			methodParameters = v.replaceParameterTypes(methodParameters, mappings)
		}
		parameters = formatter.FormatParameters(methodParameters)
	}
	implementation = sts.ReplaceAll(implementation, "<Parameters>", parameters)

	// Generate the method result type.
	var resultType string
	if methodResult != nil {
		if mappings != nil && mappings.GetSize() > 0 {
			methodResult = v.replaceResultType(methodResult, mappings)
		}
		resultType = " " + formatter.FormatResult(methodResult)
	}
	implementation = sts.ReplaceAll(implementation, "<ResultType>", resultType)

	return implementation
}

func (v *generator_) generateModules(
	imports ast.ImportsLike,
	class string,
) (
	implementation string,
) {
	if imports != nil {
		var iterator = imports.GetModules().GetModules().GetIterator()
		for iterator.HasNext() {
			var packageModule = iterator.GetNext()
			var name = packageModule.GetName()
			var path = packageModule.GetPath()
			if sts.Contains(class, name+".") {
				implementation += "\n\t" + name + " " + path
			}
		}
	}
	if sts.Contains(class, "syn.") {
		implementation += "\n\tfmt \"fmt\""
		implementation += "\n\tsyn \"sync\""
	}
	if len(implementation) > 0 {
		implementation += "\n"
	}
	return implementation
}

func (v *generator_) generateNotice(model ast.ModelLike) string {
	var notice = model.GetNotice().GetComment()
	return notice
}

func (v *generator_) generatePublicMethods(
	targetType string,
	instance ast.InstanceLike,
) string {
	var formatter = Formatter().Make()
	var publicMethods string
	var instanceMethods = instance.GetOptionalMethods()
	if instanceMethods == nil {
		return publicMethods
	}
	publicMethods = "\n// Public\n"
	var iterator = instanceMethods.GetMethods().GetIterator()
	for iterator.HasNext() {
		var publicMethod = iterator.GetNext()
		var methodName = publicMethod.GetName()
		var methodParameters = publicMethod.GetOptionalParameters()
		var parameters string
		if methodParameters != nil {
			parameters = formatter.FormatParameters(methodParameters)
		}
		var body = methodBodyTemplate_
		var result = publicMethod.GetOptionalResult()
		var resultType string
		if result != nil {
			switch actual := result.GetAny().(type) {
			case ast.AbstractionLike:
				body = resultBodyTemplate_
			case ast.ParameterizedLike:
				body = returnBodyTemplate_
			default:
				var message = fmt.Sprintf(
					"An unknown result type was found: %T",
					actual,
				)
				panic(message)
			}
			resultType = " " + formatter.FormatResult(result)
		}
		var method = instanceMethodTemplate_
		if len(targetType) > 0 {
			method = typeMethodTemplate_
		}
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		publicMethods += method
	}
	return publicMethods
}

func (v *generator_) makePrivate(name string) string {
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	return string(runes)
}

func (v *generator_) replaceAbstractionType(
	abstraction ast.AbstractionLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.AbstractionLike {
	var prefix = abstraction.GetOptionalPrefix()
	var alias = abstraction.GetOptionalAlias()
	var typeName = abstraction.GetName()
	var genericArguments = abstraction.GetOptionalGenericArguments()

	if prefix != nil {
		prefix = v.replacePrefixType(prefix, mappings)
	}

	if genericArguments != nil {
		var arguments = genericArguments.GetArguments()
		arguments = v.replaceArgumentTypes(arguments, mappings)
		genericArguments = ast.GenericArguments().Make(arguments)
	}

	if alias == nil {
		var concreteType = mappings.GetValue(typeName)
		if concreteType != nil {
			alias = concreteType.GetOptionalAlias()
			typeName = concreteType.GetName()
			genericArguments = concreteType.GetOptionalGenericArguments()
		}
	}

	abstraction = ast.Abstraction().Make(
		prefix,
		alias,
		typeName,
		genericArguments,
	)
	return abstraction
}

func (v *generator_) replaceArgumentType(
	argument ast.ArgumentLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.ArgumentLike {
	var abstraction = argument.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	argument = ast.Argument().Make(abstraction)
	return argument
}

func (v *generator_) replaceArgumentTypes(
	arguments ast.ArgumentsLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.ArgumentsLike {
	// Handle the non-generic case.
	if mappings == nil {
		return arguments
	}

	// Replace the first argument.
	var argument = arguments.GetArgument()
	argument = v.replaceArgumentType(argument, mappings)

	// Replace any additional arguments.
	var notation = cdc.Notation().Make()
	var additionalArguments = col.List[ast.AdditionalArgumentLike](notation).Make()
	var iterator = arguments.GetAdditionalArguments().GetIterator()
	for iterator.HasNext() {
		var additionalArgument = iterator.GetNext()
		var argument = additionalArgument.GetArgument()
		argument = v.replaceArgumentType(argument, mappings)
		additionalArgument = ast.AdditionalArgument().Make(argument)
		additionalArguments.AppendValue(additionalArgument)
	}

	// Construct a new sequence of arguments.
	arguments = ast.Arguments().Make(argument, additionalArguments)
	return arguments
}

func (v *generator_) replaceParameterType(
	parameter ast.ParameterLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.ParameterLike {
	var parameterName = parameter.GetName()
	var abstraction = parameter.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	parameter = ast.Parameter().Make(parameterName, abstraction)
	return parameter
}

func (v *generator_) replaceParameterTypes(
	parameters ast.ParametersLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.ParametersLike {
	// Replace the first parameter.
	var parameter = parameters.GetParameter()
	parameter = v.replaceParameterType(parameter, mappings)

	// Replace any additional parameters.
	var notation = cdc.Notation().Make()
	var additionalParameters = col.List[ast.AdditionalParameterLike](notation).Make()
	var iterator = parameters.GetAdditionalParameters().GetIterator()
	for iterator.HasNext() {
		var additionalParameter = iterator.GetNext()
		var parameter = additionalParameter.GetParameter()
		parameter = v.replaceParameterType(parameter, mappings)
		additionalParameter = ast.AdditionalParameter().Make(parameter)
		additionalParameters.AppendValue(additionalParameter)
	}

	// Construct a new sequence of parameters.
	parameters = ast.Parameters().Make(parameter, additionalParameters)
	return parameters
}

func (v *generator_) replacePrefixType(
	prefix ast.PrefixLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.PrefixLike {
	switch actual := prefix.GetAny().(type) {
	case ast.MapLike:
		var typeName = actual.GetName()
		var concreteType = mappings.GetValue(typeName)
		typeName = concreteType.GetName()
		var map_ = ast.Map().Make(typeName)
		prefix = ast.Prefix().Make(map_)
	default:
		// Ignore the rest.
	}
	return prefix
}

func (v *generator_) replaceResultType(
	result ast.ResultLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.ResultLike {
	// Handle the different kinds of results.
	switch actual := result.GetAny().(type) {
	case ast.AbstractionLike:
		var abstraction = actual
		abstraction = v.replaceAbstractionType(abstraction, mappings)
		result = ast.Result().Make(abstraction)
	case ast.ParameterizedLike:
		var parameterized = actual
		var parameters = parameterized.GetParameters()
		parameters = v.replaceParameterTypes(parameters, mappings)
		parameterized = ast.Parameterized().Make(parameters)
		result = ast.Result().Make(parameterized)
	}
	return result
}

func (v *generator_) retrieveAspect(
	model ast.ModelLike,
	name string,
) ast.AspectLike {
	var aspects = model.GetOptionalAspects()
	if aspects != nil {
		var iterator = aspects.GetAspects().GetIterator()
		for iterator.HasNext() {
			var aspect = iterator.GetNext()
			var declaration = aspect.GetDeclaration()
			if declaration.GetName() == name {
				return aspect
			}
		}
	}
	var message = fmt.Sprintf(
		"Missing the following aspect definition: %v",
		name,
	)
	panic(message)
}
