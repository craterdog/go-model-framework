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
	// Extract the first argument.
	var parameter = parameters.GetParameter()
	var abstraction = ast.Abstraction().Make(nil, nil, parameter.GetName(), nil)
	var argument = ast.Argument().Make(abstraction)

	// Extract any additional arguments.
	var notation = cdc.Notation().Make()
	var additionalArguments = col.List[ast.AdditionalArgumentLike](notation).Make()
	var additionalParameters = parameters.GetAdditionalParameters()
	var iterator = additionalParameters.GetIterator()
	for iterator.HasNext() {
		parameter = iterator.GetNext().GetParameter()
		abstraction = ast.Abstraction().Make(nil, nil, parameter.GetName(), nil)
		var additionalArgument = ast.AdditionalArgument().Make(
			ast.Argument().Make(abstraction),
		)
		additionalArguments.AppendValue(additionalArgument)
	}

	// Assemble the arguments.
	var arguments = ast.Arguments().Make(argument, additionalArguments)
	return arguments
}

func (v *generator_) extractAttributeNameAndType(
	attribute ast.AttributeLike,
) (
	attributeName string,
	attributeType string,
) {
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
	var formatter = Formatter().Make()
	attributeType = formatter.FormatAbstraction(abstraction)
	return attributeName, attributeType
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
	attributes col.CatalogLike[string, string],
) {
	var constructors = class.GetConstructors()
	var iterator = constructors.GetConstructors().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var methodName = constructor.GetName()
		var parameters = constructor.GetOptionalParameters()
		// Focus on constructors that are passed attributes as arguments.
		if parameters != nil &&
			(methodName == "Make" || sts.HasPrefix(methodName, "MakeWith")) {
			v.extractParameterAttributes(parameters, attributes)
		}
	}
}

func (v *generator_) extractInstanceAttributes(
	instance ast.InstanceLike,
	attributes col.CatalogLike[string, string],
) {
	var iterator = instance.GetAttributes().GetAttributes().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var attributeName, attributeType = v.extractAttributeNameAndType(attribute)
		attributes.SetValue(attributeName, attributeType)
	}
}

func (v *generator_) extractParameterAttribute(
	parameter ast.ParameterLike,
	attributes col.CatalogLike[string, string],
) {
	var parameterName = parameter.GetName()
	parameterName = sts.TrimSuffix(parameterName, "_")
	var abstraction = parameter.GetAbstraction()
	var formatter = Formatter().Make()
	var parameterType = formatter.FormatAbstraction(abstraction)
	attributes.SetValue(parameterName, parameterType)
}

func (v *generator_) extractParameterAttributes(
	parameters ast.ParametersLike,
	attributes col.CatalogLike[string, string],
) {
	var parameter = parameters.GetParameter()
	v.extractParameterAttribute(parameter, attributes)
	var iterator = parameters.GetAdditionalParameters().GetIterator()
	for iterator.HasNext() {
		parameter = iterator.GetNext().GetParameter()
		v.extractParameterAttribute(parameter, attributes)
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
		// This class has no instance attributes besides the GetClass() attribute.
		var iterator = class.GetConstructors().GetConstructors().GetIterator()
		for iterator.HasNext() {
			var constructor = iterator.GetNext()
			var name = constructor.GetName()
			if name == "MakeFromValue" {
				// We found a primitive value target type.
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
) (
	implementation string,
) {
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
	// Check to see if this instance interface includes aspect abstractions.
	var abstractions = instance.GetOptionalAbstractions()
	if abstractions == nil {
		return implementation
	}

	// Generate the methods for each aspect abstraction.
	var iterator = abstractions.GetAbstractions().GetIterator()
	for iterator.HasNext() {
		// Each aspect abstraction binds to its own concrete arguments.
		var abstraction = iterator.GetNext()
		var formatter = Formatter().Make()
		var aspectName = formatter.FormatAbstraction(abstraction)
		var instanceAspect = instanceAspectTemplate_
		instanceAspect = sts.ReplaceAll(instanceAspect, "<AspectName>", aspectName)
		var methods string
		if abstraction.GetOptionalAlias() == nil {
			// We will only know the method signatures for the local aspects.
			var mappings col.CatalogLike[string, ast.AbstractionLike]
			var aspect = v.retrieveAspect(model, abstraction.GetName())
			var declaration = aspect.GetDeclaration()
			var genericParameters = declaration.GetOptionalGenericParameters()
			var genericArguments = abstraction.GetOptionalGenericArguments()
			if genericParameters != nil && genericArguments != nil {
				var parameters = genericParameters.GetParameters()
				var arguments = genericArguments.GetArguments()
				mappings = v.extractConcreteMappings(parameters, arguments)
			}
			methods = v.generateAbstractionMethods(targetType, aspect, mappings)
		}
		instanceAspect = sts.ReplaceAll(instanceAspect, "<Methods>", methods)
		implementation += instanceAspect
	}

	return implementation
}

func (v *generator_) generateAttributeInitialization(
	parameter ast.ParameterLike,
) (
	implementation string,
) {
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	implementation = attributeInitializationTemplate_
	implementation = sts.ReplaceAll(implementation, "<AttributeName>", attributeName)
	implementation = sts.ReplaceAll(implementation, "<ParameterName>", parameterName)
	return implementation
}

func (v *generator_) generateAttributeInitializations(
	class ast.ClassLike,
	constructor ast.ConstructorLike,
) (
	implementation string,
) {
	// Ignore a constructor that doesn't take attributes as parameters.
	var name = constructor.GetName()
	if sts.HasPrefix(name, "MakeFrom") {
		return implementation
	}

	// Ignore a constructor that doesn't take any parameters.
	var parameters = constructor.GetOptionalParameters()
	if parameters == nil {
		return implementation
	}

	// Generate the first attribute initialization.
	var parameter = parameters.GetParameter()
	var initialization = v.generateAttributeInitialization(parameter)
	implementation += initialization

	// Generate any additional attribute initializations.
	var additionalParameters = parameters.GetAdditionalParameters()
	var iterator = additionalParameters.GetIterator()
	for iterator.HasNext() {
		var additionalParameter = iterator.GetNext()
		parameter = additionalParameter.GetParameter()
		initialization = v.generateAttributeInitialization(parameter)
		implementation += initialization
	}

	return implementation
}

func (v *generator_) generateAttributeCheck(
	parameter ast.ParameterLike,
) (
	implementation string,
) {
	// Ignore an optional attribute.
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	if sts.HasPrefix(attributeName, "optional") {
		return implementation
	}

	// Generate the attribute check code.
	implementation = attributeCheckTemplate_
	implementation = sts.ReplaceAll(implementation, "<AttributeName>", attributeName)
	implementation = sts.ReplaceAll(implementation, "<ParameterName>", parameterName)

	return implementation
}

func (v *generator_) generateAttributeChecks(
	class ast.ClassLike,
	constructor ast.ConstructorLike,
) (
	implementation string,
) {
	// Ignore a constructor that doesn't take attributes as parameters.
	var name = constructor.GetName()
	if sts.HasPrefix(name, "MakeFrom") {
		return implementation
	}

	// Ignore a constructor that doesn't take any parameters.
	var parameters = constructor.GetOptionalParameters()
	if parameters == nil {
		return implementation
	}

	// Generate the attribute check for the first parameter.
	var parameter = parameters.GetParameter()
	var check = v.generateAttributeCheck(parameter)
	implementation += check

	// Generate attribute checks for any additional parameters.
	var additionalParameters = parameters.GetAdditionalParameters()
	var iterator = additionalParameters.GetIterator()
	for iterator.HasNext() {
		var additionalParameter = iterator.GetNext()
		parameter = additionalParameter.GetParameter()
		check = v.generateAttributeCheck(parameter)
		implementation += check
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
	implementation = "\n// Attributes\n"

	// Generate each instance attribute method.
	var instanceAttributes = instance.GetAttributes()
	var iterator = instanceAttributes.GetAttributes().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()

		// Fill in the attribute method body template.
		var body string
		var parameter string
		var parameterName string
		var resultType string
		var attributeName, attributeType = v.extractAttributeNameAndType(attribute)
		var attributeParameter = attribute.GetOptionalParameter()
		var methodName = attribute.GetName()
		if sts.HasPrefix(methodName, "Set") {
			// This is a setter method.
			switch {
			case sts.HasPrefix(methodName, "SetOptional"):
				body = setterOptionalTemplate_
			default:
				body = setterClassTemplate_
			}
			var formatter = Formatter().Make()
			parameter = formatter.FormatParameter(attributeParameter)
			parameterName = attributeParameter.GetName()
		} else {
			// This is a getter method.
			body = getterClassTemplate_
			if len(targetType) > 0 {
				body = getterTypeTemplate_
			}
			resultType = " " + attributeType
		}
		body = sts.ReplaceAll(body, "<AttributeName>", attributeName)
		body = sts.ReplaceAll(body, "<ParameterName>", parameterName)

		// Generate the attribute method implementation.
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

	return implementation
}

func (v *generator_) generateClassAccess(
	class ast.ClassLike,
) (
	implementation string,
) {
	// Assume a non-generic class model.
	var reference = classReferenceTemplate_
	var function = classFunctionTemplate_
	var declaration = class.GetDeclaration()

	// Switch to a generic class model if necessary.
	var genericParameters = declaration.GetOptionalGenericParameters()
	if genericParameters != nil {
		reference = genericReferenceTemplate_
		function = genericFunctionTemplate_
	}

	// Generate the class access implementation.
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
	// Check to see if this class model includes class constants.
	var classConstants = class.GetOptionalConstants()
	if classConstants == nil {
		return implementation
	}

	// Generate the code for each private class constant declaration.
	var formatter = Formatter().Make()
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

	// Generate the class method target.
	var target = v.generateClassTarget(class)
	implementation = sts.ReplaceAll(implementation, "<Target>", target)

	// Generate the class constructor methods.
	var constructorMethods = v.generateConstructorMethods(targetType, class)
	implementation = sts.ReplaceAll(implementation, "<Constructors>", constructorMethods)

	// Generate the class constant access methods.
	var constantMethods = v.generateConstantMethods(class)
	implementation = sts.ReplaceAll(implementation, "<Constants>", constantMethods)

	// Generate the class function methods.
	var functionMethods = v.generateFunctionMethods(class)
	implementation = sts.ReplaceAll(implementation, "<Functions>", functionMethods)

	// Generate any private class methods.
	var privateMethods = v.generatePrivateMethods(targetType, class)
	implementation = sts.ReplaceAll(implementation, "<Private>", privateMethods)

	return implementation
}

func (v *generator_) generateClassTarget(
	class ast.ClassLike,
) (
	implementation string,
) {
	implementation = classTargetTemplate_

	// Generate the private class constant definitions.
	var constants = v.generateClassConstants(class)
	implementation = sts.ReplaceAll(implementation, "<Constants>", constants)

	return implementation
}

func (v *generator_) generateConstantMethods(
	class ast.ClassLike,
) (
	implementation string,
) {
	// Check to see if this class model includes class constants.
	var classConstants = class.GetOptionalConstants()
	if classConstants == nil {
		return implementation
	}

	// Generate the code for each class constant access method.
	implementation = "\n// Constants\n"
	var formatter = Formatter().Make()
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
	// Generate the code for each class constructor method.
	implementation = "\n// Constructors\n"
	var formatter = Formatter().Make()
	var classConstructors = class.GetConstructors()
	var iterator = classConstructors.GetConstructors().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var method = classMethodTemplate_

		// Generate the name of the class constructor.
		var methodName = constructor.GetName()
		method = sts.ReplaceAll(method, "<MethodName>", methodName)

		// Choose the appropriate class constructor method body.
		var body = constructorBodyTemplate_
		if len(targetType) > 0 {
			if methodName == "MakeFromValue" {
				body = typeBodyTemplate_
				body = sts.ReplaceAll(body, "<TargetType>", targetType)
			} else {
				body = resultBodyTemplate_
			}
		}

		// Generate the attribute value checks.
		var checks = v.generateAttributeChecks(class, constructor)
		body = sts.ReplaceAll(body, "<Checks>", checks)

		// Generate the attribute value initializations.
		var initializations = v.generateAttributeInitializations(class, constructor)
		body = sts.ReplaceAll(body, "<Initializations>", initializations)

		method = sts.ReplaceAll(method, "<Body>", body)

		// Generate any parameters for the class constructor.
		var constructorParameters = constructor.GetOptionalParameters()
		var parameters string
		if constructorParameters != nil {
			parameters = formatter.FormatParameters(constructorParameters)
		}
		method = sts.ReplaceAll(method, "<Parameters>", parameters)

		// Generate the class constructor method result type.
		var abstraction = constructor.GetAbstraction()
		var resultType = " " + formatter.FormatAbstraction(abstraction)
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
	// Check to see if this class model includes class functions.
	var classFunctions = class.GetOptionalFunctions()
	if classFunctions == nil {
		return implementation
	}

	// Generate the code for each class function method.
	implementation = "\n// Functions\n"
	var formatter = Formatter().Make()
	var iterator = classFunctions.GetFunctions().GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		var method = classMethodTemplate_

		// Generate the name of the class function.
		var name = function.GetName()
		method = sts.ReplaceAll(method, "<MethodName>", name)

		// Generate any parameters for the class function.
		var functionParameters = function.GetOptionalParameters()
		var parameters string
		if functionParameters != nil {
			parameters = formatter.FormatParameters(functionParameters)
		}
		method = sts.ReplaceAll(method, "<Parameters>", parameters)

		// Generate the body of the class function.
		var body = functionBodyTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)

		// Generate the result type for the class function.
		var result = function.GetResult()
		var resultType = " " + formatter.FormatResult(result)
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
	var header = headerTemplate_
	var name = model.GetHeader().GetName()
	header = sts.ReplaceAll(header, "<Name>", name)
	return header
}

func (v *generator_) generateImports(
	model ast.ModelLike,
	class string,
) (
	implementation string,
) {
	// Check to see if this class model includes module imports.
	var imports = model.GetOptionalImports()
	if imports != nil {
		var modules = imports.GetModules()
		implementation = v.generateModules(modules, class)
	}

	// Generate imports for specific modules that are referenced in the code.
	if sts.Contains(class, "syn.") {
		implementation += "\n\tfmt \"fmt\""
	}
	if sts.Contains(class, "ref.") {
		implementation += "\n\tref \"reflect\""
	}
	if sts.Contains(class, "syn.") {
		implementation += "\n\tsyn \"sync\""
	}

	// Generate an import statement with any imported modules.
	if len(implementation) > 0 {
		implementation += "\n"
		implementation = sts.ReplaceAll(
			importsTemplate_,
			"<Modules>",
			implementation,
		)
	}
	return implementation
}

func (v *generator_) generateInstanceAttributes(
	class ast.ClassLike,
	instance ast.InstanceLike,
) (
	implementation string,
) {
	// Create a catalog of attribute name-type mappings.
	var notation = cdc.Notation().Make()
	var attributes = col.Catalog[string, string](notation).Make()

	// Extract the attribute mappings from the class and instance definitions.
	v.extractInstanceAttributes(instance, attributes) // This must come first.
	v.extractConstructorAttributes(class, attributes)
	if attributes.IsEmpty() {
		implementation = "\n\t// TBA - Add private instance attributes.\n"
		return implementation
	}

	// Generate the instance attribute definitions for the class.
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var attribute = instanceAttributeTemplate_
		var attributeName = association.GetKey()
		attribute = sts.ReplaceAll(attribute, "<AttributeName>", attributeName)
		var attributeType = association.GetValue()
		attribute = sts.ReplaceAll(attribute, "<AttributeType>", attributeType)
		implementation += attribute
	}

	return implementation
}

func (v *generator_) generateInstanceMethods(
	targetType string,
	model ast.ModelLike,
	class ast.ClassLike,
	instance ast.InstanceLike,
) (
	implementation string,
) {
	implementation = instanceMethodsTemplate_

	// Generate the instance method target.
	var target = v.generateInstanceTarget(targetType, class, instance)
	implementation = sts.ReplaceAll(implementation, "<Target>", target)

	// Generate the instance attribute access methods for the class.
	var attributes = v.generateAttributeMethods(targetType, class, instance)
	implementation = sts.ReplaceAll(implementation, "<Attributes>", attributes)

	// Generate the instance abstraction methods for the class.
	var abstractions = v.generateAbstractions(targetType, model, instance)
	implementation = sts.ReplaceAll(implementation, "<Abstractions>", abstractions)

	// Generate the instance public methods for the class.
	var methods = v.generatePublicMethods(targetType, instance)
	implementation = sts.ReplaceAll(implementation, "<Methods>", methods)

	return implementation
}

func (v *generator_) generateInstanceTarget(
	targetType string,
	class ast.ClassLike,
	instance ast.InstanceLike,
) (
	implementation string,
) {
	// Generate the right instance target definition.
	if len(targetType) > 0 {
		implementation = typeTargetTemplate_
		implementation = sts.ReplaceAll(implementation, "<TargetType>", targetType)
	} else {
		implementation = instanceTargetTemplate_
		var attributes = v.generateInstanceAttributes(class, instance)
		implementation = sts.ReplaceAll(implementation, "<Attributes>", attributes)
	}
	return implementation
}

func (v *generator_) generateMethodImplementation(
	targetType string,
	method ast.MethodLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) (
	implementation string,
) {
	// Choose the right method template.
	if len(targetType) > 0 {
		implementation = typeMethodTemplate_
	} else {
		implementation = instanceMethodTemplate_
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
	modules ast.ModulesLike,
	class string,
) (
	implementation string,
) {
	var iterator = modules.GetModules().GetIterator()
	for iterator.HasNext() {
		var module = iterator.GetNext()
		var name = module.GetName()
		var path = module.GetPath()
		if sts.Contains(class, name+".") {
			implementation += "\n\t" + name + " " + path
		}
	}
	return implementation
}

func (v *generator_) generateNotice(model ast.ModelLike) string {
	var notice = model.GetNotice().GetComment()
	return notice
}

func (v *generator_) generatePrivateMethods(
	targetType string,
	class ast.ClassLike,
) (
	implementation string,
) {
	// Extended primitive type classes don't define private class methods.
	if len(targetType) > 0 {
		return implementation
	}

	// Check each class constructor method for attribute parameters.
	var constructors = class.GetConstructors()
	var constructorIterator = constructors.GetConstructors().GetIterator()
	for constructorIterator.HasNext() {
		var constructor = constructorIterator.GetNext()

		// Ignore class constructor methods that don't have attribute parameters.
		var name = constructor.GetName()
		var parameters = constructor.GetOptionalParameters()
		if parameters == nil || sts.HasPrefix(name, "MakeFrom") {
			continue
		}

		// Check the first parameter in the class constructor.
		var parameter = parameters.GetParameter()
		if !sts.HasPrefix(parameter.GetName(), "optional") {
			// Found a mandatory attribute parameter.
			implementation = privateMethodsTemplate_
			return implementation
		}

		// Check any additional parameters in the class constructor.
		var parameterIterator = parameters.GetAdditionalParameters().GetIterator()
		for parameterIterator.HasNext() {
			var additionalParameter = parameterIterator.GetNext()
			var parameter = additionalParameter.GetParameter()
			if !sts.HasPrefix(parameter.GetName(), "optional") {
				// Found a mandatory attribute parameter.
				implementation = privateMethodsTemplate_
				return implementation
			}
		}
	}

	// No mandatory attribute parameters were found.
	return implementation
}

func (v *generator_) generatePublicMethods(
	targetType string,
	instance ast.InstanceLike,
) (
	implementation string,
) {
	// Check to see if this instance interface includes public methods.
	var instanceMethods = instance.GetOptionalMethods()
	if instanceMethods == nil {
		return implementation
	}

	// Generate the code for each instance public method.
	implementation = "\n// Public\n"
	var formatter = Formatter().Make()
	var iterator = instanceMethods.GetMethods().GetIterator()
	for iterator.HasNext() {
		var publicMethod = iterator.GetNext()

		// Choose the appropriate method template.
		var method = instanceMethodTemplate_
		if len(targetType) > 0 {
			method = typeMethodTemplate_
		}

		// Generate the name of the public method.
		var methodName = publicMethod.GetName()
		method = sts.ReplaceAll(method, "<MethodName>", methodName)

		// Generate any parameters for the public method.
		var methodParameters = publicMethod.GetOptionalParameters()
		var parameters string
		if methodParameters != nil {
			parameters = formatter.FormatParameters(methodParameters)
		}
		method = sts.ReplaceAll(method, "<Parameters>", parameters)

		// Generate the body of the public method.
		var body = methodBodyTemplate_
		var result = publicMethod.GetOptionalResult()
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
		}
		method = sts.ReplaceAll(method, "<Body>", body)

		// Generate the result type for the public method.
		var resultType string
		if result != nil {
			resultType = " " + formatter.FormatResult(result)
		}
		method = sts.ReplaceAll(method, "<ResultType>", resultType)

		implementation += method
	}
	return implementation
}

func (v *generator_) makePrivate(name string) string {
	// All names beginning with a lowercase character are private.
	runes := []rune(name)
	runes[0] = uni.ToLower(runes[0])
	return string(runes)
}

func (v *generator_) replaceAbstractionType(
	abstraction ast.AbstractionLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.AbstractionLike {
	// Replace the generic type in a prefix with the concrete type.
	var prefix = abstraction.GetOptionalPrefix()
	if prefix != nil {
		prefix = v.replacePrefixType(prefix, mappings)
	}

	// Replace the generic types in a sequence of arguments with concrete types.
	var genericArguments = abstraction.GetOptionalGenericArguments()
	if genericArguments != nil {
		var arguments = genericArguments.GetArguments()
		arguments = v.replaceArgumentTypes(arguments, mappings)
		genericArguments = ast.GenericArguments().Make(arguments)
	}

	// Replace a non-aliased generic type with its concrete type.
	var typeName = abstraction.GetName()
	var alias = abstraction.GetOptionalAlias()
	if alias == nil {
		var concreteType = mappings.GetValue(typeName)
		if concreteType != nil {
			alias = concreteType.GetOptionalAlias()
			typeName = concreteType.GetName()
			genericArguments = concreteType.GetOptionalGenericArguments()
		}
	}

	// Recreate the abstraction using its updated types.
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
	// Ignore the non-generic case.
	if mappings == nil {
		return arguments
	}

	// Replace the generic type of the first argument with its concrete type.
	var argument = arguments.GetArgument()
	argument = v.replaceArgumentType(argument, mappings)

	// Replace the generic types of any additional arguments with concrete types.
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

	// Construct the updated sequence of arguments.
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
	// Ignore the non-generic case.
	if mappings == nil {
		return parameters
	}

	// Replace the generic type of the first parameter with its concrete type.
	var parameter = parameters.GetParameter()
	parameter = v.replaceParameterType(parameter, mappings)

	// Replace the generic types of any additional parameters with concrete types.
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

	// Construct the updated sequence of parameters.
	parameters = ast.Parameters().Make(parameter, additionalParameters)
	return parameters
}

func (v *generator_) replacePrefixType(
	prefix ast.PrefixLike,
	mappings col.CatalogLike[string, ast.AbstractionLike],
) ast.PrefixLike {
	switch actual := prefix.GetAny().(type) {
	case ast.MapLike:
		// eg. map[K]V -> map[string]int
		var typeName = actual.GetName()
		var concreteType = mappings.GetValue(typeName)
		typeName = concreteType.GetName()
		var map_ = ast.Map().Make(typeName)
		prefix = ast.Prefix().Make(map_)
	default:
		// Ignore the rest since they don't contain any generic types.
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
	default:
		var message = fmt.Sprintf(
			"An unknown result type was found: %T",
			actual,
		)
		panic(message)
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
