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
) string {
	var classIterator = model.GetClasses().GetClassIterator()
	var instanceIterator = model.GetInstances().GetInstanceIterator()
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
			var source = v.generateClass(model, class, instance)
			return source
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

func (v *generator_) extractConstructorAttributes(
	class ast.ClassLike,
	catalog col.CatalogLike[string, string],
) {
	var constructors = class.GetConstructors()
	if constructors == nil {
		return
	}
	var iterator = constructors.GetConstructorIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var methodName = constructor.GetName()
		if sts.HasPrefix(methodName, "MakeWith") {
			var parameters = constructor.GetParameters()
			v.extractParameterAttributes(parameters, catalog)
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
	var iterator = attributes.GetAttributeIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var name = attribute.GetName()
		var abstraction ast.AbstractionLike
		switch {
		case sts.HasPrefix(name, "Get"):
			attributeName = sts.TrimPrefix(name, "Get")
			abstraction = attribute.GetAbstraction()
		case sts.HasPrefix(name, "Is"):
			attributeName = sts.TrimPrefix(name, "Is")
			abstraction = attribute.GetAbstraction()
		case sts.HasPrefix(name, "Was"):
			attributeName = sts.TrimPrefix(name, "Was")
			abstraction = attribute.GetAbstraction()
		case sts.HasPrefix(name, "Has"):
			attributeName = sts.TrimPrefix(name, "Has")
			abstraction = attribute.GetAbstraction()
		case sts.HasPrefix(name, "Set"):
			attributeName = sts.TrimPrefix(name, "Set")
			var parameter = attribute.GetParameter()
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
	var iterator = parameters.GetAdditionalParameters().GetAdditionalParameterIterator()
	for iterator.HasNext() {
		parameter = iterator.GetNext().GetParameter()
		v.extractParameterAttribute(parameter, catalog)
	}
}

func (v *generator_) generateAbstractionMethods(
	targetType string,
	aspect ast.AspectLike,
	abstraction ast.AbstractionLike,
) string {
	var formatter = Formatter().Make()
	var aspectDeclaration = aspect.GetDeclaration()
	var genericTypes = aspectDeclaration.GetGenericParameters()
	var concreteTypes = abstraction.GetGenericArguments()
	var abstractionMethods string
	var aspectMethods = aspect.GetMethods()
	if aspectMethods == nil {
		return abstractionMethods
	}
	var iterator = aspectMethods.GetMethodIterator()
	for iterator.HasNext() {
		var aspectMethod = iterator.GetNext()
		var methodName = aspectMethod.GetName()
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
			switch actual := methodResult.GetAny().(type) {
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
		var method = instanceMethodTemplate_
		if len(targetType) > 0 {
			method = typeMethodTemplate_
		}
		method = sts.ReplaceAll(method, "<Body>", body)
		method = sts.ReplaceAll(method, "<MethodName>", methodName)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		abstractionMethods += method
	}
	return abstractionMethods
}

func (v *generator_) generateAbstractions(
	targetType string,
	model ast.ModelLike,
	instance ast.InstanceLike,
) string {
	var formatter = Formatter().Make()
	var aspects string
	var abstractions = instance.GetAbstractions()
	if abstractions == nil {
		return aspects
	}
	var iterator = abstractions.GetAbstractionIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		var prefix = abstraction.GetPrefix()
		var name = abstraction.GetName()
		var aspectName = formatter.FormatAbstraction(abstraction)
		var methods string
		if prefix == nil {
			// We only know the method signatures for the local aspects.
			var aspect = v.retrieveAspect(model, name)
			methods = v.generateAbstractionMethods(targetType, aspect, abstraction)
		}
		var instanceAspect = instanceAspectTemplate_
		instanceAspect = sts.ReplaceAll(instanceAspect, "<AspectName>", aspectName)
		instanceAspect = sts.ReplaceAll(instanceAspect, "<Methods>", methods)
		aspects += instanceAspect
	}
	return aspects
}

func (v *generator_) generateAttributeAssignment(
	parameter ast.ParameterLike,
) string {
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	var assignment = attributeAssignmentTemplate_
	assignment = sts.ReplaceAll(assignment, "<AttributeName>", attributeName)
	assignment = sts.ReplaceAll(assignment, "<ParameterName>", parameterName)
	return assignment
}

func (v *generator_) generateAttributeAssignments(
	class ast.ClassLike,
	constructor ast.ConstructorLike,
) string {
	var assignments string
	var name = constructor.GetName()
	if sts.HasPrefix(name, "MakeFrom") {
		return assignments
	}
	var parameters = constructor.GetParameters()
	if parameters == nil {
		return assignments
	}
	var parameter = parameters.GetParameter()
	var assignment = v.generateAttributeAssignment(parameter)
	assignments += assignment
	var additionalParameters = parameters.GetAdditionalParameters()
	var iterator = additionalParameters.GetAdditionalParameterIterator()
	for iterator.HasNext() {
		var additionalParameter = iterator.GetNext()
		parameter = additionalParameter.GetParameter()
		assignment = v.generateAttributeAssignment(parameter)
		assignments += assignment
	}
	return assignments
}

func (v *generator_) generateAttributeMethods(
	targetType string,
	class ast.ClassLike,
	instance ast.InstanceLike,
) string {
	var formatter = Formatter().Make()
	var methods string
	var instanceAttributes = instance.GetAttributes()
	if instanceAttributes == nil {
		return methods
	}
	methods = "\n// Attributes\n"
	var iterator = instanceAttributes.GetAttributeIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		var methodName = attribute.GetName()
		var attributeName string
		var body string

		var parameter string
		var attributeParameter = attribute.GetParameter()
		var parameterName string
		if attributeParameter != nil {
			attributeName = sts.TrimPrefix(methodName, "Set")
			parameterName = attributeParameter.GetName()
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
		methods += method
	}
	return methods
}

func (v *generator_) extractArguments(
	parameters ast.ParametersLike,
) ast.ArgumentsLike {
	var notation = cdc.Notation().Make()
	var list = col.List[ast.AdditionalArgumentLike](notation).Make()
	var parameter = parameters.GetParameter()
	var iterator = parameters.GetAdditionalParameters().GetAdditionalParameterIterator()
	var abstraction = ast.Abstraction().Make(nil, parameter.GetName(), nil)
	var argument = ast.Argument().Make(abstraction)
	for iterator.HasNext() {
		parameter = iterator.GetNext().GetParameter()
		abstraction = ast.Abstraction().Make(nil, parameter.GetName(), nil)
		var additionalArgument = ast.AdditionalArgument().Make(
			ast.Argument().Make(abstraction),
		)
		list.AppendValue(additionalArgument)
	}
	var additionalArguments = ast.AdditionalArguments().Make(list.GetIterator())
	var arguments = ast.Arguments().Make(argument, additionalArguments)
	return arguments
}

func (v *generator_) generateClass(
	model ast.ModelLike,
	class ast.ClassLike,
	instance ast.InstanceLike,
) string {
	var targetType string
	var attributeIterator = instance.GetAttributes().GetAttributeIterator()
	attributeIterator.ToEnd()
	if attributeIterator.GetSlot() == 1 {
		var iterator = class.GetConstructors().GetConstructorIterator()
		for iterator.HasNext() {
			var constructor = iterator.GetNext()
			var name = constructor.GetName()
			if name == "MakeFromValue" {
				var parameter = constructor.GetParameters().GetParameter()
				var abstraction = parameter.GetAbstraction()
				var formatter = Formatter().Make()
				targetType = formatter.FormatAbstraction(abstraction)
				break
			}
		}
	}

	var template = classTemplate_

	var notice = v.generateNotice(model)
	template = sts.ReplaceAll(template, "<Notice>", notice)

	var header = v.generateHeader(model)
	template = sts.ReplaceAll(template, "<Header>", header)

	var classAccess = v.generateClassAccess(class)
	template = sts.ReplaceAll(template, "<Access>", classAccess)

	var classMethods = v.generateClassMethods(targetType, class)
	template = sts.ReplaceAll(template, "<Class>", classMethods)

	var instanceMethods = v.generateInstanceMethods(
		targetType,
		model,
		class,
		instance,
	)
	template = sts.ReplaceAll(template, "<Instance>", instanceMethods)

	var classDeclaration = class.GetDeclaration()
	var className = classDeclaration.GetName()
	className = sts.TrimSuffix(className, "ClassLike")
	template = sts.ReplaceAll(template, "<ClassName>", className)
	template = sts.ReplaceAll(template, "<TargetName>", v.makePrivate(className))

	var parameters string
	var arguments string
	var genericParameters = classDeclaration.GetGenericParameters()
	if genericParameters != nil {
		var classParameters = genericParameters.GetParameters()
		var classArguments = v.extractArguments(classParameters)
		var formatter = Formatter().Make()
		parameters = "[" + formatter.FormatParameters(classParameters) + "]"
		arguments = "[" + formatter.FormatArguments(classArguments) + "]"
	}
	template = sts.ReplaceAll(template, "[<Parameters>]", parameters)
	template = sts.ReplaceAll(template, "[<Arguments>]", arguments)

	var imports = v.generateImports(model, template)
	template = sts.ReplaceAll(template, "<Imports>", imports)
	return template
}

func (v *generator_) generateClassAccess(class ast.ClassLike) string {
	var declaration = class.GetDeclaration()
	var genericParameters = declaration.GetGenericParameters()
	var reference = classReferenceTemplate_
	var function = classFunctionTemplate_
	if genericParameters != nil {
		reference = genericReferenceTemplate_
		function = genericFunctionTemplate_
	}
	var access = classAccessTemplate_
	access = sts.ReplaceAll(access, "<Reference>", reference)
	access = sts.ReplaceAll(access, "<Function>", function)
	return access
}

func (v *generator_) generateClassConstants(class ast.ClassLike) string {
	var formatter = Formatter().Make()
	var constants string
	var classConstants = class.GetConstants()
	if classConstants == nil {
		return constants
	}
	var iterator = classConstants.GetConstantIterator()
	for iterator.HasNext() {
		var classConstant = iterator.GetNext()
		var constantName = classConstant.GetName()
		var constantAbstraction = classConstant.GetAbstraction()
		constantName = v.makePrivate(constantName)
		var constantType = formatter.FormatAbstraction(constantAbstraction)
		var constant = classConstantTemplate_
		constant = sts.ReplaceAll(constant, "<ConstantName>", constantName)
		constant = sts.ReplaceAll(constant, "<ConstantType>", constantType)
		constants += constant
	}
	return constants
}

func (v *generator_) generateClassMethods(
	targetType string,
	class ast.ClassLike,
) string {
	var methods = classMethodsTemplate_
	var target = v.generateClassTarget(class)
	methods = sts.ReplaceAll(methods, "<Target>", target)
	var constantMethods = v.generateConstantMethods(class)
	methods = sts.ReplaceAll(methods, "<Constants>", constantMethods)
	var constructorMethods = v.generateConstructorMethods(targetType, class)
	methods = sts.ReplaceAll(methods, "<Constructors>", constructorMethods)
	var functionMethods = v.generateFunctionMethods(class)
	methods = sts.ReplaceAll(methods, "<Functions>", functionMethods)
	return methods
}

func (v *generator_) generateClassTarget(class ast.ClassLike) string {
	var target = classTargetTemplate_
	var constants = v.generateClassConstants(class)
	target = sts.ReplaceAll(target, "<Constants>", constants)
	return target
}

func (v *generator_) generateConstantMethods(class ast.ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var classConstants = class.GetConstants()
	if classConstants == nil {
		return methods
	}
	methods = "\n// Constants\n"
	var iterator = classConstants.GetConstantIterator()
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
		methods += method
	}
	return methods
}

func (v *generator_) generateConstructorMethods(
	targetType string,
	class ast.ClassLike,
) string {
	var formatter = Formatter().Make()
	var methods string
	var classConstructors = class.GetConstructors()
	if classConstructors == nil {
		return methods
	}
	methods = "\n// Constructors\n"
	var iterator = classConstructors.GetConstructorIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		var methodName = constructor.GetName()
		var ConstructorParameters = constructor.GetParameters()
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
		methods += method
	}
	return methods
}

func (v *generator_) generateFunctionMethods(class ast.ClassLike) string {
	var formatter = Formatter().Make()
	var methods string
	var classFunctions = class.GetFunctions()
	if classFunctions == nil {
		return methods
	}
	methods = "\n// Functions\n"
	var iterator = classFunctions.GetFunctionIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		var name = function.GetName()
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
		method = sts.ReplaceAll(method, "<MethodName>", name)
		method = sts.ReplaceAll(method, "<Parameters>", parameters)
		method = sts.ReplaceAll(method, "<ResultType>", resultType)
		methods += method
	}
	return methods
}

func (v *generator_) generateHeader(model ast.ModelLike) string {
	var name = model.GetHeader().GetName()
	var header = headerTemplate_
	header = sts.ReplaceAll(header, "<Name>", name)
	return header
}

func (v *generator_) generateImports(model ast.ModelLike, class string) string {
	var modules string
	if model.GetImports() != nil {
		var iterator = model.GetImports().GetModules().GetModuleIterator()
		for iterator.HasNext() {
			var packageModule = iterator.GetNext()
			var name = packageModule.GetName()
			var path = packageModule.GetPath()
			if sts.Contains(class, name+".") {
				modules += "\n\t" + name + " " + path
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
	imports = sts.ReplaceAll(imports, "<Modules>", modules)
	return imports
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
		target = sts.ReplaceAll(target, "<TargetType>", targetType)
	}
	var attributes = v.generateInstanceAttributes(class, instance)
	target = sts.ReplaceAll(target, "<Attributes>", attributes)
	return target
}

func (v *generator_) generateNotice(model ast.ModelLike) string {
	var notice = model.GetNotice().GetComment() + "\n"
	return notice
}

func (v *generator_) generatePublicMethods(
	targetType string,
	instance ast.InstanceLike,
) string {
	var formatter = Formatter().Make()
	var publicMethods string
	var instanceMethods = instance.GetMethods()
	if instanceMethods == nil {
		return publicMethods
	}
	publicMethods = "\n// Public\n"
	var iterator = instanceMethods.GetMethodIterator()
	for iterator.HasNext() {
		var publicMethod = iterator.GetNext()
		var methodName = publicMethod.GetName()
		var methodParameters = publicMethod.GetParameters()
		var parameters string
		if methodParameters != nil {
			parameters = formatter.FormatParameters(methodParameters)
		}
		var body = methodBodyTemplate_
		var result = publicMethod.GetResult()
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

func (v *generator_) lookupConcreteType(
	genericTypeName string,
	genericParameters ast.GenericParametersLike,
	concreteArguments ast.GenericArgumentsLike,
) (
	concreteType ast.AbstractionLike,
) {
	var parameters = genericParameters.GetParameters()
	var genericParameter = parameters.GetParameter()
	var arguments = concreteArguments.GetArguments()
	var concreteArgument = arguments.GetArgument()
	if genericTypeName == genericParameter.GetName() {
		concreteType = concreteArgument.GetAbstraction()
	} else {
		var additionalParameters = parameters.GetAdditionalParameters()
		var additionalArguments = arguments.GetAdditionalArguments()
		if additionalParameters != nil && additionalArguments != nil {
			var parameterIterator = additionalParameters.GetAdditionalParameterIterator()
			var argumentIterator = additionalArguments.GetAdditionalArgumentIterator()
			for parameterIterator.HasNext() {
				genericParameter = parameterIterator.GetNext().GetParameter()
				concreteArgument = argumentIterator.GetNext().GetArgument()
				if genericTypeName == genericParameter.GetName() {
					concreteType = concreteArgument.GetAbstraction()
					break
				}
			}
		}
	}
	if concreteType == nil {
		var message = fmt.Sprintf(
			"An unknown generic type name was passed: %q",
			genericTypeName,
		)
		panic(message)
	}
	return concreteType
}

func (v *generator_) replaceParameterTypes(
	genericTypes ast.GenericParametersLike,
	concreteTypes ast.GenericArgumentsLike,
	methodParameters ast.ParametersLike,
) ast.ParametersLike {
	var parameter = methodParameters.GetParameter()
	var parameterName = parameter.GetName()
	var parameterType = v.lookupConcreteType(
		parameterName,
		genericTypes,
		concreteTypes,
	)
	parameter = ast.Parameter().Make(parameterName, parameterType)

	var additionalParameterIterator = methodParameters.GetAdditionalParameters().GetAdditionalParameterIterator()
	var notation = cdc.Notation().Make()
	var list = col.List[ast.AdditionalParameterLike](notation).Make()
	for additionalParameterIterator.HasNext() {
		parameter = additionalParameterIterator.GetNext().GetParameter()
		parameterName = parameter.GetName()
		parameterType = v.lookupConcreteType(
			parameterName,
			genericTypes,
			concreteTypes,
		)
		parameter = ast.Parameter().Make(parameterName, parameterType)
		var additionalParameter = ast.AdditionalParameter().Make(parameter)
		list.AppendValue(additionalParameter)
	}
	var additionalParameters = ast.AdditionalParameters().Make(list.GetIterator())
	var parameters = ast.Parameters().Make(parameter, additionalParameters)
	return parameters
}

func (v *generator_) replaceResultTypes(
	genericTypes ast.GenericParametersLike,
	concreteTypes ast.GenericArgumentsLike,
	methodResult ast.ResultLike,
) ast.ResultLike {
	switch actual := methodResult.GetAny().(type) {
	case ast.AbstractionLike:
		var parameterName = actual.GetName()
		var resultAbstraction = v.lookupConcreteType(
			parameterName,
			genericTypes,
			concreteTypes,
		)
		methodResult = ast.Result().Make(resultAbstraction)
	case ast.ParameterizedLike:
		var resultParameters = actual.GetParameters()
		resultParameters = v.replaceParameterTypes(
			genericTypes,
			concreteTypes,
			resultParameters,
		)
		methodResult = ast.Result().Make(resultParameters)
	default:
		var message = fmt.Sprintf(
			"An unknown result type was passed: %T",
			actual,
		)
		panic(message)
	}
	return methodResult
}

func (v *generator_) retrieveAspect(
	model ast.ModelLike,
	name string,
) ast.AspectLike {
	var aspects = model.GetAspects()
	if aspects != nil {
		var iterator = aspects.GetAspectIterator()
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
