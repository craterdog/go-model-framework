/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See http://opensource.org/licenses/MIT)                        .
................................................................................
*/

package models

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
	osx "os"
	sts "strings"
	tim "time"
	uni "unicode"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// This class does not initialize any class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// This class does not define any private class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		// This class does not initialize any private instance attributes.
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	// This class does not define any private instance attributes.
}

// Public

func (v *generator_) CreateModel(directory string, name string, copyright string) {
	// Center and insert the copyright notice into the model template.
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
	var template = sts.ReplaceAll(modelTemplate_, "<Copyright>", copyright)
	template = sts.ReplaceAll(template, "<packagename>", name)
	var bytes = []byte(template[1:]) // Remove leading "\n".

	// Save the new model template.
	v.createDirectory(directory)
	var modelFile = directory + "Package.go"
	fmt.Printf(
		"The model file %q does not exist, creating a template for it.\n",
		modelFile,
	)
	var err = osx.WriteFile(modelFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) GeneratePackage(directory string) {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var model = v.parseModel(directory)
	if model == nil {
		return
	}
	v.generateModel(directory, model)
	v.generateClasses(directory, model)
}

// Private

func (v *generator_) createDirectory(directory string) {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) extractConstructorAttributes(
	class ClassLike,
	catalog col.CatalogLike[string, string],
) {
	var constructors = class.GetConstructors()
	if constructors != nil {
		var iterator = constructors.GetIterator()
		for iterator.HasNext() {
			var constructor = iterator.GetNext()
			var methodName = constructor.GetIdentifier()
			if sts.HasPrefix(methodName, "MakeWith") {
				var parameters = constructor.GetParameters()
				v.extractParameterAttributes(parameters, catalog)
			}
		}
	}
}

func (v *generator_) extractInstanceAttributes(
	instance InstanceLike,
	catalog col.CatalogLike[string, string],
) {
	var attributeName string
	var attributeType string
	var formatter = Formatter().Make()
	var attributes = instance.GetAttributes()
	if attributes != nil {
		var iterator = attributes.GetIterator()
		for iterator.HasNext() {
			var attribute = iterator.GetNext()
			var identifier = attribute.GetIdentifier()
			var abstraction AbstractionLike
			switch {
			case sts.HasPrefix(identifier, "Get"):
				attributeName = sts.TrimPrefix(identifier, "Get")
				abstraction = attribute.GetAbstraction()
			case sts.HasPrefix(identifier, "Is"):
				attributeName = sts.TrimPrefix(identifier, "Is")
				abstraction = attribute.GetAbstraction()
			case sts.HasPrefix(identifier, "Was"):
				attributeName = sts.TrimPrefix(identifier, "Was")
				abstraction = attribute.GetAbstraction()
			case sts.HasPrefix(identifier, "Has"):
				attributeName = sts.TrimPrefix(identifier, "Has")
				abstraction = attribute.GetAbstraction()
			default:
				if attributeName == v.makePrivate(sts.TrimPrefix(identifier, "Set")) {
					// This attribute was already added.
					continue
				}
				attributeName = sts.TrimPrefix(identifier, "Set")
				var parameter = attribute.GetParameter()
				abstraction = parameter.GetAbstraction()
			}
			attributeName = v.makePrivate(attributeName)
			attributeType = formatter.FormatAbstraction(abstraction)
			catalog.SetValue(attributeName, attributeType)
		}
	}
}

func (v *generator_) extractParameterAttributes(
	parameters col.ListLike[ParameterLike],
	catalog col.CatalogLike[string, string],
) {
	var formatter = Formatter().Make()
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var attributeName = attribute.GetIdentifier()
		attributeName = sts.TrimSuffix(attributeName, "_")
		var abstraction = attribute.GetAbstraction()
		var attributeType = formatter.FormatAbstraction(abstraction)
		catalog.SetValue(attributeName, attributeType)
	}
}

func (v *generator_) generateAbstractionMethods(
	aspect AspectLike,
	abstraction AbstractionLike,
) string {
	var formatter = Formatter().Make()
	var aspectDeclaration = aspect.GetDeclaration()
	var genericTypes = aspectDeclaration.GetParameters()
	var concreteTypes = abstraction.GetArguments()
	var abstractionMethods string
	var aspectMethods = aspect.GetMethods()
	if aspectMethods == nil {
		return abstractionMethods
	}
	var iterator = aspectMethods.GetIterator()
	for iterator.HasNext() {
		var aspectMethod = iterator.GetNext()
		var methodName = aspectMethod.GetIdentifier()
		var methodParameters = aspectMethod.GetParameters()
		var parameters string
		if methodParameters != nil {
			if genericTypes != nil {
				// Replace the generic type names from the aspect definition
				// with the actual types defined in the instance interface.
				methodParameters = v.replaceParameterTypes(
					genericTypes,
					concreteTypes,
					methodParameters,
				)
			}
			parameters = formatter.FormatParameters(methodParameters)
		}
		var resultType string
		var body = methodBodyTemplate_
		var methodResult = aspectMethod.GetResult()
		if methodResult != nil {
			if genericTypes != nil {
				// Replace the generic type names from the aspect definition
				// with the actual types defined in the instance interface.
				methodResult = v.replaceResultTypes(
					genericTypes,
					concreteTypes,
					methodResult,
				)
			}
			resultType = " " + formatter.FormatResult(methodResult)
			if methodResult.GetAbstraction() != nil {
				body = resultBodyTemplate_
			} else {
				body = returnBodyTemplate_
			}
		}
		var method = instanceMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		abstractionMethods += method + "\n"
	}
	return abstractionMethods
}

func (v *generator_) generateAbstractions(
	model ModelLike,
	instance InstanceLike,
) string {
	var formatter = Formatter().Make()
	var result string
	var abstractions = instance.GetAbstractions()
	if abstractions == nil {
		return result
	}
	var iterator = abstractions.GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		var prefix = abstraction.GetPrefix()
		var identifier = abstraction.GetIdentifier()
		var aspectName = formatter.FormatAbstraction(abstraction)
		var methods string
		if prefix == nil {
			// We only know the method signatures for the local aspects.
			var aspect = v.retrieveAspect(model, identifier)
			methods = v.generateAbstractionMethods(aspect, abstraction)
		}
		var instanceAspect = instanceAspectTemplate_
		instanceAspect = sts.ReplaceAll(instanceAspect, "<AspectName>", aspectName)
		instanceAspect = sts.ReplaceAll(instanceAspect, "<Methods>", methods)
		result += instanceAspect
	}
	return result
}

func (v *generator_) generateAttributeAssignments(
	class ClassLike,
	constructor ConstructorLike,
) string {
	var assignments string
	var identifier = constructor.GetIdentifier()
	if !sts.HasPrefix(identifier, "MakeWith") {
		return assignments
	}
	var parameters = constructor.GetParameters()
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var parameterName = parameter.GetIdentifier()
		var attributeName = sts.TrimSuffix(parameterName, "_")
		var assignment = attributeAssignmentTemplate_
		assignment = sts.ReplaceAll(assignment, "<AttributeName>", attributeName)
		assignment = sts.ReplaceAll(assignment, "<ParameterName>", parameterName)
		assignments += assignment
	}
	assignments += "\n\t"
	return assignments
}

func (v *generator_) generateAttributeMethods(instance InstanceLike) string {
	var formatter = Formatter().Make()
	var methods string
	var instanceAttributes = instance.GetAttributes()
	if instanceAttributes == nil {
		return methods
	}
	var iterator = instanceAttributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var methodName = attribute.GetIdentifier()
		var attributeName string
		var body string

		var parameter string
		var attributeParameter = attribute.GetParameter()
		var parameterName string
		if attributeParameter != nil {
			attributeName = sts.TrimPrefix(methodName, "Set")
			parameterName = attributeParameter.GetIdentifier()
			parameter = formatter.FormatParameter(attributeParameter)
			body = setterBodyTemplate_
		}

		var resultType string
		var abstraction = attribute.GetAbstraction()
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
		}

		attributeName = v.makePrivate(attributeName)
		body = sts.ReplaceAll(body, "<AttributeName>", attributeName)
		body = sts.ReplaceAll(body, "<ParameterName>", parameterName)
		var method = instanceMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameter)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateClass(
	directory string,
	model ModelLike,
	class ClassLike,
	instance InstanceLike,
) {
	var template = classTemplate_

	var notice = model.GetNotice().GetComment()
	template = sts.ReplaceAll(template, "<Notice>", notice)

	var header = v.generateHeader(model)
	template = sts.ReplaceAll(template, "<Header>", header)

	var classAccess = v.generateClassAccess(class)
	template = sts.ReplaceAll(template, "<Access>", classAccess)

	var classMethods = v.generateClassMethods(class)
	template = sts.ReplaceAll(template, "<Class>", classMethods)

	var instanceMethods = v.generateInstanceMethods(
		model,
		class,
		instance,
	)
	template = sts.ReplaceAll(template, "<Instance>", instanceMethods)

	var classDeclaration = class.GetDeclaration()
	var classIdentifier = classDeclaration.GetIdentifier()
	var className = sts.TrimSuffix(classIdentifier, "ClassLike")
	template = sts.ReplaceAll(template, "<ClassName>", className)
	template = sts.ReplaceAll(template, "<TargetName>", v.makePrivate(className))

	var parameters string
	var arguments string
	var classParameters = classDeclaration.GetParameters()
	if classParameters != nil {
		var formatter = Formatter().Make()
		parameters = "[" + formatter.FormatGenerics(classParameters) + "]"
		arguments = "[" + formatter.FormatParameterNames(classParameters) + "]"
	}
	template = sts.ReplaceAll(template, "[<Parameters>]", parameters)
	template = sts.ReplaceAll(template, "[<Arguments>]", arguments)

	var imports = v.generateImports(model, template)
	template = sts.ReplaceAll(template, "<Imports>", imports)

	var fileName = sts.ToLower(className)
	var classFile = directory + fileName + ".go"
	v.outputClass(classFile, template)
}

func (v *generator_) generateClassAccess(class ClassLike) string {
	var declaration = class.GetDeclaration()
	var parameters = declaration.GetParameters()
	var reference = classReferenceTemplate_
	var function = classFunctionTemplate_
	if parameters != nil {
		reference = genericReferenceTemplate_
		function = genericFunctionTemplate_
	}
	var access = classAccessTemplate_
	access = sts.ReplaceAll(access, "<Reference>", reference)
	access = sts.ReplaceAll(access, "<Function>", function)
	return access + "\n"
}

func (v *generator_) generateClassConstants(class ClassLike) string {
	var formatter = Formatter().Make()
	var constants string
	var classConstants = class.GetConstants()
	if classConstants == nil {
		constants = "\n\t// This class has no private constants.\n"
		return constants
	}
	var iterator = classConstants.GetIterator()
	for iterator.HasNext() {
		var classConstant = iterator.GetNext()
		var constantIdentifier = classConstant.GetIdentifier()
		var constantAbstraction = classConstant.GetAbstraction()
		var constantName = v.makePrivate(constantIdentifier)
		var constantType = formatter.FormatAbstraction(constantAbstraction)
		var constant = classConstantTemplate_
		constant = sts.ReplaceAll(constant, "<ConstantName>", constantName)
		constant = sts.ReplaceAll(constant, "<ConstantType>", constantType)
		constants += constant
	}
	constants += "\n"
	return constants
}

func (v *generator_) generateClasses(directory string, model ModelLike) {
	var classes = model.GetClasses()
	var instances = model.GetInstances()
	if classes == nil || instances == nil {
		return
	}
	var classIterator = classes.GetIterator()
	var instanceIterator = instances.GetIterator()
	for classIterator.HasNext() {
		var class = classIterator.GetNext()
		var instance = instanceIterator.GetNext()
		v.generateClass(directory, model, class, instance)
	}
}

func (v *generator_) generateClassMethods(class ClassLike) string {
	var methods = classMethodsTemplate_
	var target = v.generateClassTarget(class)
	methods = sts.ReplaceAll(methods, "<Target>", target)
	var constantMethods = v.generateConstantMethods(class)
	methods = sts.ReplaceAll(methods, "<Constants>", constantMethods)
	var constructorMethods = v.generateConstructorMethods(class)
	methods = sts.ReplaceAll(methods, "<Constructors>", constructorMethods)
	var functionMethods = v.generateFunctionMethods(class)
	methods = sts.ReplaceAll(methods, "<Functions>", functionMethods)
	return methods
}

func (v *generator_) generateClassTarget(class ClassLike) string {
	var target = classTargetTemplate_
	var constants = v.generateClassConstants(class)
	target = sts.ReplaceAll(target, "<Constants>", constants) + "\n"
	return target
}

func (v *generator_) generateConstantMethods(class ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var classConstants = class.GetConstants()
	if classConstants == nil {
		return methods
	}
	var iterator = classConstants.GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		var methodName = constant.GetIdentifier()
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
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateConstructorMethods(class ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var classConstructors = class.GetConstructors()
	if classConstructors == nil {
		return methods
	}
	var iterator = classConstructors.GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var methodName = constructor.GetIdentifier()
		var ConstructorParameters = constructor.GetParameters()
		var parameters string
		if ConstructorParameters != nil {
			parameters = formatter.FormatParameters(ConstructorParameters)
		}
		var abstraction = constructor.GetAbstraction()
		var resultType = " " + formatter.FormatAbstraction(abstraction)
		var assignments = v.generateAttributeAssignments(class, constructor)
		var body = constructorBodyTemplate_
		body = sts.ReplaceAll(body, "<Assignments>", assignments)
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateFunctionMethods(class ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var classFunctions = class.GetFunctions()
	if classFunctions == nil {
		return methods
	}
	var iterator = classFunctions.GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		var identifier = function.GetIdentifier()
		var functionParameters = function.GetParameters()
		var parameters string
		if functionParameters != nil {
			parameters = formatter.FormatParameters(functionParameters)
		}
		var result = function.GetResult()
		var resultType = " " + formatter.FormatResult(result)
		var body = functionBodyTemplate_
		var method = classMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", identifier)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method + "\n"
	}
	return methods
}

func (v *generator_) generateHeader(model ModelLike) string {
	var packageName = model.GetHeader().GetIdentifier()
	var header = headerTemplate_
	header = sts.ReplaceAll(header, "<PackageName>", packageName) + "\n"
	return header
}

func (v *generator_) generateImports(model ModelLike, class string) string {
	var modules string
	var packageModules = model.GetModules()
	if packageModules != nil {
		var iterator = packageModules.GetIterator()
		for iterator.HasNext() {
			var packageModule = iterator.GetNext()
			var identifier = packageModule.GetIdentifier()
			var text = packageModule.GetText()
			if sts.Contains(class, identifier+".") {
				modules += "\n\t" + identifier + " " + text
			}
		}
	}
	if sts.Contains(class, "syn.") {
		modules += "\n\tfmt \"fmt\""
		modules += "\n\tsyn \"sync\""
	}
	if len(modules) > 0 {
		modules += "\n"
	}
	var imports = importsTemplate_
	imports = sts.ReplaceAll(imports, "<Modules>", modules) + "\n"
	return imports
}

func (v *generator_) generateInstanceAttributes(
	class ClassLike,
	instance InstanceLike,
) string {
	var attributes string
	var catalog = col.Catalog[string, string]().Make()
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
	attributes += "\n"
	return attributes
}

func (v *generator_) generateInstanceMethods(
	model ModelLike,
	class ClassLike,
	instance InstanceLike,
) string {
	var instanceMethods = instanceMethodsTemplate_
	var target = v.generateInstanceTarget(class, instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Target>", target)
	var attributes = v.generateAttributeMethods(instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Attributes>", attributes)
	var abstractions = v.generateAbstractions(model, instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Abstractions>", abstractions)
	var methods = v.generatePublicMethods(instance)
	instanceMethods = sts.ReplaceAll(instanceMethods, "<Methods>", methods)
	return instanceMethods
}

func (v *generator_) generateInstanceTarget(
	class ClassLike,
	instance InstanceLike,
) string {
	var target = instanceTargetTemplate_
	var attributes = v.generateInstanceAttributes(class, instance)
	target = sts.ReplaceAll(target, "<Attributes>", attributes) + "\n"
	return target
}

func (v *generator_) generateModel(directory string, model ModelLike) {
	var formatter = Formatter().Make()
	var source = formatter.FormatModel(model)
	var bytes = []byte(source)
	var modelFile = directory + "Package.go"
	var err = osx.WriteFile(modelFile, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) generatePublicMethods(instance InstanceLike) string {
	var formatter = Formatter().Make()
	var publicMethods string
	var instanceMethods = instance.GetMethods()
	if instanceMethods == nil {
		return publicMethods
	}
	var iterator = instanceMethods.GetIterator()
	for iterator.HasNext() {
		var publicMethod = iterator.GetNext()
		var methodName = publicMethod.GetIdentifier()
		var methodParameters = publicMethod.GetParameters()
		var parameters string
		if methodParameters != nil {
			parameters = formatter.FormatParameters(methodParameters)
		}
		var body = methodBodyTemplate_
		var result = publicMethod.GetResult()
		var resultType string
		if result != nil {
			if result.GetAbstraction() != nil {
				body = resultBodyTemplate_
			} else {
				body = returnBodyTemplate_
			}
			resultType = " " + formatter.FormatResult(result)
		}
		var method = instanceMethodTemplate_
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		publicMethods += method + "\n"
	}
	return publicMethods
}

func (v *generator_) makePrivate(identifier string) string {
	runes := []rune(identifier)
	runes[0] = uni.ToLower(runes[0])
	return string(runes)
}

func (v *generator_) outputClass(classFile, class string) {
	var _, err = osx.ReadFile(classFile)
	if err == nil {
		// Don't overwrite an existing class file.
		fmt.Printf(
			"The class file %q already exists, leaving it alone.\n",
			classFile,
		)
		return
	}
	err = osx.WriteFile(classFile, []byte(class), 0644)
	if err != nil {
		panic(err)
	}
}

func (v *generator_) parseModel(directory string) ModelLike {
	var modelFile = directory + "Package.go"
	var bytes, err = osx.ReadFile(modelFile)
	if err != nil {
		var message = fmt.Sprintf(
			"The specified directory is missing a model file: %v",
			modelFile,
		)
		panic(message)
	}
	var source = string(bytes)
	var parser = Parser().Make()
	var model = parser.ParseSource(source)
	var validator = Validator().Make()
	validator.ValidateModel(model)
	return model
}

func (v *generator_) replaceGenericType(
	genericTypes col.ListLike[ParameterLike],
	concreteTypes col.ListLike[AbstractionLike],
	abstraction AbstractionLike,
) AbstractionLike {
	var formatter = Formatter().Make()
	var prefix = abstraction.GetPrefix()
	var identifier = abstraction.GetIdentifier()
	var arguments = abstraction.GetArguments()
	var genericIterator = genericTypes.GetIterator()
	var concreteIterator = concreteTypes.GetIterator()
	for genericIterator.HasNext() {
		var genericType = genericIterator.GetNext()
		var genericName = genericType.GetIdentifier()
		var concreteType = concreteIterator.GetNext()
		if identifier == genericName {
			identifier = formatter.FormatAbstraction(concreteType)
			break
		}
	}
	if arguments != nil {
		var argumentIterator = arguments.GetIterator()
		arguments = col.List[AbstractionLike]().Make()
		for argumentIterator.HasNext() {
			var argument = argumentIterator.GetNext()
			argument = v.replaceGenericType(
				genericTypes,
				concreteTypes,
				argument,
			)
			arguments.AppendValue(argument)
		}
	}
	abstraction = Abstraction().MakeWithAttributes(prefix, identifier, arguments)
	return abstraction
}

func (v *generator_) replaceParameterTypes(
	genericTypes col.ListLike[ParameterLike],
	concreteTypes col.ListLike[AbstractionLike],
	methodParameters col.ListLike[ParameterLike],
) col.ListLike[ParameterLike] {
	var parameterIterator = methodParameters.GetIterator()
	methodParameters = col.List[ParameterLike]().Make()
	for parameterIterator.HasNext() {
		var methodParameter = parameterIterator.GetNext()
		var parameterName = methodParameter.GetIdentifier()
		var parameterType = methodParameter.GetAbstraction()
		parameterType = v.replaceGenericType(
			genericTypes,
			concreteTypes,
			parameterType,
		)
		methodParameter = Parameter().MakeWithAttributes(parameterName, parameterType)
		methodParameters.AppendValue(methodParameter)
	}
	return methodParameters
}

func (v *generator_) replaceResultTypes(
	genericTypes col.ListLike[ParameterLike],
	concreteTypes col.ListLike[AbstractionLike],
	methodResult ResultLike,
) ResultLike {
	var resultAbstraction = methodResult.GetAbstraction()
	if resultAbstraction != nil {
		resultAbstraction = v.replaceGenericType(
			genericTypes,
			concreteTypes,
			resultAbstraction,
		)
		methodResult = Result().MakeWithAbstraction(resultAbstraction)
	} else {
		var resultParameters = methodResult.GetParameters()
		resultParameters = v.replaceParameterTypes(
			genericTypes,
			concreteTypes,
			resultParameters,
		)
		methodResult = Result().MakeWithParameters(resultParameters)
	}
	return methodResult
}

func (v *generator_) retrieveAspect(
	model ModelLike,
	identifier string,
) AspectLike {
	var aspects = model.GetAspects()
	if aspects != nil {
		var iterator = aspects.GetIterator()
		for iterator.HasNext() {
			var aspect = iterator.GetNext()
			var declaration = aspect.GetDeclaration()
			if declaration.GetIdentifier() == identifier {
				return aspect
			}
		}
	}
	var message = fmt.Sprintf(
		"Missing the following aspect definition: %v",
		identifier,
	)
	panic(message)
}
