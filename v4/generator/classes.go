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

package generator

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func Classes() ClassesClassLike {
	return classesClass
}

// Constructor Methods

func (c *classesClass_) Make() ClassesLike {
	var instance = &classes_{
		// Initialize the instance attributes.
		class_: c,
	}
	return instance
}

// INSTANCE INTERFACE

// Public Methods

func (v *classes_) GetClass() ClassesClassLike {
	return v.getClass()
}

func (v *classes_) GenerateModelClasses(
	model ast.ModelLike,
) abs.CatalogLike[string, string] {
	var result_ = col.Catalog[string, string]()
	var interfaceDefinitions = model.GetInterfaceDefinitions()
	var classSection = interfaceDefinitions.GetClassSection()
	var classDefinitions = classSection.GetClassDefinitions().GetIterator()
	var instanceSection = interfaceDefinitions.GetInstanceSection()
	var instanceDefinitions = instanceSection.GetInstanceDefinitions().GetIterator()
	for classDefinitions.HasNext() && instanceDefinitions.HasNext() {
		var classDefinition = classDefinitions.GetNext()
		var className = v.extractClassName(classDefinition)
		var instanceDefinition = instanceDefinitions.GetNext()
		var implementation = v.generateClass(
			model,
			classDefinition,
			instanceDefinition,
		)
		result_.SetValue(className, implementation)
	}
	return result_
}

// Private Methods

func (v *classes_) getClass() *classesClass_ {
	return v.class_
}

func (v *classes_) analyzeClassConstants(
	classDefinition ast.ClassDefinitionLike,
) {
	v.constants_ = col.Catalog[string, string]()
	var classMethods = classDefinition.GetClassMethods()
	var constantSubsection = classMethods.GetOptionalConstantSubsection()
	if uti.IsDefined(constantSubsection) {
		var constantMethods = constantSubsection.GetConstantMethods().GetIterator()
		for constantMethods.HasNext() {
			var constantMethod = constantMethods.GetNext()
			var constantName = constantMethod.GetName()
			var constantType = v.extractType(constantMethod.GetAbstraction())
			v.constants_.SetValue(constantName, constantType)
		}
	}
}

func (v *classes_) analyzeClassDefinition(
	classDefinition ast.ClassDefinitionLike,
	instanceDefinition ast.InstanceDefinitionLike,
) {
	v.analyzeClassGenerics(classDefinition)
	v.analyzeClassConstants(classDefinition)
	v.analyzePublicAttributes(instanceDefinition)
	v.analyzePrivateAttributes(classDefinition)
}

func (v *classes_) analyzeClassGenerics(
	classDefinition ast.ClassDefinitionLike,
) {
	v.isGeneric_ = false
	var declaration = classDefinition.GetDeclaration()
	var constraints = declaration.GetOptionalConstraints()
	if uti.IsDefined(constraints) {
		v.isGeneric_ = true
	}
}

func (v *classes_) analyzePrivateAttributes(
	classDefinition ast.ClassDefinitionLike,
) {
	var classMethods = classDefinition.GetClassMethods()
	var constructorSubsection = classMethods.GetConstructorSubsection()
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		var name = constructorMethod.GetName()
		// Focus only on constructors that are passed attributes as arguments.
		if name == "Make" || sts.HasPrefix(name, "MakeWith") {
			var parameters = constructorMethod.GetParameters().GetIterator()
			for parameters.HasNext() {
				var parameter = parameters.GetNext()
				var attributeName = sts.TrimSuffix(parameter.GetName(), "_")
				var attributeType = v.extractType(parameter.GetAbstraction())
				v.attributes_.SetValue(attributeName, attributeType)
			}
		}
	}
}

func (v *classes_) analyzePublicAttributes(
	instanceDefinition ast.InstanceDefinitionLike,
) {
	v.attributes_ = col.Catalog[string, string]()
	var instanceMethods = instanceDefinition.GetInstanceMethods()
	var attributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(attributeSubsection) {
		var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
		for attributeMethods.HasNext() {
			var attributeMethod = attributeMethods.GetNext()
			var attributeName string
			var abstraction ast.AbstractionLike
			switch actual := attributeMethod.GetAny().(type) {
			case ast.GetterMethodLike:
				attributeName = v.extractAttributeName(actual.GetName())
				abstraction = actual.GetAbstraction()
			case ast.SetterMethodLike:
				attributeName = v.extractAttributeName(actual.GetName())
				abstraction = actual.GetParameter().GetAbstraction()
			}
			var attributeType = v.extractType(abstraction)
			v.attributes_.SetValue(attributeName, attributeType)
		}
	}
}

func (v *classes_) extractAttributeName(accessorName string) string {
	var attributeName string
	switch {
	case sts.HasPrefix(accessorName, "Get"):
		attributeName = sts.TrimPrefix(accessorName, "Get")
	case sts.HasPrefix(accessorName, "Set"):
		attributeName = sts.TrimPrefix(accessorName, "Set")
	case sts.HasPrefix(accessorName, "Is"):
		attributeName = sts.TrimPrefix(accessorName, "Is")
	case sts.HasPrefix(accessorName, "Was"):
		attributeName = sts.TrimPrefix(accessorName, "Was")
	case sts.HasPrefix(accessorName, "Are"):
		attributeName = sts.TrimPrefix(accessorName, "Are")
	case sts.HasPrefix(accessorName, "Were"):
		attributeName = sts.TrimPrefix(accessorName, "Were")
	case sts.HasPrefix(accessorName, "Has"):
		attributeName = sts.TrimPrefix(accessorName, "Has")
	case sts.HasPrefix(accessorName, "Had"):
		attributeName = sts.TrimPrefix(accessorName, "Had")
	case sts.HasPrefix(accessorName, "Have"):
		attributeName = sts.TrimPrefix(accessorName, "Have")
	default:
		var message = fmt.Sprintf(
			"An unknown accessor name was found: %q",
			accessorName,
		)
		panic(message)
	}
	attributeName = uti.MakeLowerCase(attributeName)
	return attributeName
}

func (v *classes_) extractClassName(
	classDefinition ast.ClassDefinitionLike,
) string {
	var className = classDefinition.GetDeclaration().GetName()
	className = sts.TrimSuffix(className, "ClassLike")
	className = uti.MakeLowerCase(className)
	return className
}

func (v *classes_) extractConcreteMappings(
	constraints ast.ConstraintsLike,
	arguments ast.ArgumentsLike,
) abs.CatalogLike[string, ast.AbstractionLike] {
	// Create the mappings catalog.
	var mappings = col.Catalog[string, ast.AbstractionLike]()
	if uti.IsUndefined(constraints) || uti.IsUndefined(arguments) {
		return mappings
	}

	// Map the name of the first constraint to its concrete type.
	var constraint = constraints.GetConstraint()
	var constraintName = constraint.GetName()
	var argument = arguments.GetArgument()
	var concreteType = argument.GetAbstraction()
	mappings.SetValue(constraintName, concreteType)

	// Map the name of the additional constraints to their concrete types.
	var additionalConstraints = constraints.GetAdditionalConstraints().GetIterator()
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	for additionalConstraints.HasNext() {
		constraint = additionalConstraints.GetNext().GetConstraint()
		constraintName = constraint.GetName()
		argument = additionalArguments.GetNext().GetArgument()
		concreteType = argument.GetAbstraction()
		mappings.SetValue(constraintName, concreteType)
	}

	return mappings
}

func (v *classes_) extractNotice(model ast.ModelLike) string {
	var definition = model.GetModuleDefinition()
	var notice = definition.GetNotice().GetComment()
	return notice
}

func (v *classes_) extractPackageName(model ast.ModelLike) string {
	var definition = model.GetModuleDefinition()
	var header = definition.GetHeader()
	var packageName = header.GetName()
	return packageName
}

func (v *classes_) extractType(abstraction ast.AbstractionLike) string {
	var abstractType string
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		switch actual := prefix.GetAny().(type) {
		case ast.ArrayLike:
			abstractType = "[]"
		case ast.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		case ast.ChannelLike:
			abstractType = "chan "
		}
	}
	var name = abstraction.GetName()
	abstractType += name
	var suffix = abstraction.GetOptionalSuffix()
	if uti.IsDefined(suffix) {
		abstractType += "." + suffix.GetName()
	}
	var arguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(arguments) {
		var argument = v.extractType(arguments.GetArgument().GetAbstraction())
		abstractType += "[" + argument
		var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
		for additionalArguments.HasNext() {
			var additionalArgument = additionalArguments.GetNext().GetArgument()
			argument = v.extractType(additionalArgument.GetAbstraction())
			abstractType += ", " + argument
		}
		abstractType += "]"
	}
	return abstractType
}

func (v *classes_) generateAccessFunction() (
	implementation string,
) {
	implementation = v.getClass().accessFunction_
	var function = v.getClass().classFunction_
	if v.isGeneric_ {
		function = v.getClass().genericFunction_
	}
	implementation = uti.ReplaceAll(implementation, "function", function)
	return implementation
}

func (v *classes_) generateArguments(
	classDefinition ast.ClassDefinitionLike,
) (
	arguments string,
) {
	if v.isGeneric_ {
		arguments = "["
		var classDeclaration = classDefinition.GetDeclaration()
		var optionalConstraints = classDeclaration.GetOptionalConstraints()
		var constraint = optionalConstraints.GetConstraint()
		var argument = constraint.GetName()
		arguments += argument
		var additionalConstraints = optionalConstraints.GetAdditionalConstraints().GetIterator()
		for additionalConstraints.HasNext() {
			constraint = additionalConstraints.GetNext().GetConstraint()
			argument = constraint.GetName()
			arguments += ", " + argument
		}
		arguments += "]"
	}
	return arguments
}

func (v *classes_) generateAspectInterface(
	aspectType ast.AbstractionLike,
	aspectSection ast.AspectSectionLike,
) (
	implementation string,
) {
	var methods string
	if uti.IsDefined(aspectSection) {
		var aspectDefinitions = aspectSection.GetAspectDefinitions().GetIterator()
		for aspectDefinitions.HasNext() {
			var aspectDefinition = aspectDefinitions.GetNext()
			var declaration = aspectDefinition.GetDeclaration()
			var constraints = declaration.GetOptionalConstraints()
			var arguments = aspectType.GetOptionalArguments()
			if uti.IsUndefined(aspectType.GetOptionalSuffix()) &&
				declaration.GetName() == aspectType.GetName() {
				var mappings = v.extractConcreteMappings(constraints, arguments)
				methods = v.generateAspectMethods(
					aspectType,
					aspectDefinition,
					mappings,
				)
			}
		}
	}
	implementation = v.getClass().aspectInterface_
	implementation = uti.ReplaceAll(
		implementation,
		"aspectType",
		v.extractType(aspectType),
	)
	implementation = uti.ReplaceAll(
		implementation,
		"methods",
		methods,
	)
	return implementation
}

func (v *classes_) generateAspectInterfaces(
	aspectSection ast.AspectSectionLike,
	aspectSubsection ast.AspectSubsectionLike,
) (
	implementation string,
) {
	if uti.IsDefined(aspectSubsection) {
		var aspectInterfaces = aspectSubsection.GetAspectInterfaces().GetIterator()
		for aspectInterfaces.HasNext() {
			var aspectType = aspectInterfaces.GetNext().GetAbstraction()
			implementation += v.generateAspectInterface(aspectType, aspectSection)
		}
	}
	return implementation
}

func (v *classes_) generateAspectMethod(
	aspectType ast.AbstractionLike,
	aspectMethod ast.AspectMethodLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) (
	implementation string,
) {
	var method = aspectMethod.GetMethod()
	var methodName = method.GetName()
	var methodParameters = method.GetParameters()
	var methodResult = method.GetOptionalResult()
	if mappings.GetSize() > 0 {
		methodParameters = v.replaceParameterTypes(method.GetParameters(), mappings)
		if uti.IsDefined(methodResult) {
			methodResult = v.replaceResultType(method.GetOptionalResult(), mappings)
		}
	}
	var parameters = v.generateParameters(methodParameters)
	var resultType = v.generateResult(methodResult)
	implementation = v.getClass().instanceMethod_
	if uti.IsDefined(resultType) {
		implementation = v.getClass().instanceFunction_
		implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	}
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	return implementation
}

func (v *classes_) generateAspectMethods(
	aspectType ast.AbstractionLike,
	aspectDefinition ast.AspectDefinitionLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) (
	implementation string,
) {
	var aspectMethods = aspectDefinition.GetAspectMethods().GetIterator()
	for aspectMethods.HasNext() {
		var aspectMethod = aspectMethods.GetNext()
		implementation += v.generateAspectMethod(
			aspectType,
			aspectMethod,
			mappings,
		)
	}
	return implementation
}

func (v *classes_) generateAttributeCheck(parameter ast.ParameterLike) (
	implementation string,
) {
	var parameterName = parameter.GetName()
	var attributeName = sts.TrimSuffix(parameterName, "_")
	// Ignore optional attributes.
	if !sts.HasPrefix(attributeName, "optional") {
		var template = v.getClass().attributeCheck_
		template = uti.ReplaceAll(template, "attributeName", attributeName)
		implementation += template
	}
	return implementation
}

func (v *classes_) generateAttributeChecks(
	parameters abs.Sequential[ast.ParameterLike],
) (
	implementation string,
) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var attributeCheck = v.generateAttributeCheck(parameter)
		implementation += attributeCheck
	}
	return implementation
}

func (v *classes_) generateAttributeDeclarations() (
	implementation string,
) {
	var attributes = v.attributes_.GetIterator()
	for attributes.HasNext() {
		var attribute = attributes.GetNext()
		var attributeName = attribute.GetKey()
		var attributeType = attribute.GetValue()
		var declaration = v.getClass().attributeDeclaration_
		declaration = uti.ReplaceAll(declaration, "attributeName", attributeName)
		declaration = uti.ReplaceAll(declaration, "attributeType", attributeType)
		implementation += declaration
	}
	return implementation
}

func (v *classes_) generateAttributeInitializations(
	parameters abs.Sequential[ast.ParameterLike],
) (
	implementation string,
) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var parameterName = parameter.GetName()
		var attributeName = sts.TrimSuffix(parameterName, "_")
		if uti.IsDefined(v.attributes_.GetValue(attributeName)) {
			var template = v.getClass().attributeInitialization_
			template = uti.ReplaceAll(template, "attributeName", attributeName)
			implementation += template
		}
	}
	return implementation
}

func (v *classes_) generateAttributeMethods(
	instanceDefinition ast.InstanceDefinitionLike,
) (
	implementation string,
) {
	var instanceMethods = instanceDefinition.GetInstanceMethods()
	var attributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(attributeSubsection) {
		var methods string
		var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
		for attributeMethods.HasNext() {
			var method string
			var attributeMethod = attributeMethods.GetNext()
			switch actual := attributeMethod.GetAny().(type) {
			case ast.GetterMethodLike:
				method = v.generateGetterMethod(actual)
			case ast.SetterMethodLike:
				method = v.generateSetterMethod(actual)
			}
			methods += method
		}
		implementation = v.getClass().attributeMethods_
		implementation = uti.ReplaceAll(implementation, "methods", methods)
	}
	return implementation
}

func (v *classes_) generateClass(
	model ast.ModelLike,
	classDefinition ast.ClassDefinitionLike,
	instanceDefinition ast.InstanceDefinitionLike,
) (
	implementation string,
) {
	// Analyze the class.
	v.analyzeClassDefinition(classDefinition, instanceDefinition)

	// Start with the class template.
	implementation = v.getClass().classTemplate_
	var notice = v.extractNotice(model)
	implementation = uti.ReplaceAll(implementation, "notice", notice)

	// Add in the package declaration.
	var packageDeclaration = v.generatePackageDeclaration(model)
	implementation = uti.ReplaceAll(
		implementation,
		"packageDeclaration",
		packageDeclaration,
	)

	// Add in the class access function.
	var accessFunction = v.generateAccessFunction()
	implementation = uti.ReplaceAll(
		implementation,
		"accessFunction",
		accessFunction,
	)

	// Add in the class constructor methods.
	var constructorMethods = v.generateConstructorMethods(classDefinition)
	implementation = uti.ReplaceAll(
		implementation,
		"constructorMethods",
		constructorMethods,
	)

	// Add in the class constant methods.
	var constantMethods = v.generateConstantMethods(classDefinition)
	implementation = uti.ReplaceAll(
		implementation,
		"constantMethods",
		constantMethods,
	)

	// Add in the class function methods.
	var functionMethods = v.generateFunctionMethods(classDefinition)
	implementation = uti.ReplaceAll(
		implementation,
		"functionMethods",
		functionMethods,
	)

	// Add in the instance attribute methods.
	var attributeMethods = v.generateAttributeMethods(instanceDefinition)
	implementation = uti.ReplaceAll(
		implementation,
		"attributeMethods",
		attributeMethods,
	)

	// Add in the instance aspect interfaces.
	var interfaceDefinitions = model.GetInterfaceDefinitions()
	var aspectSection = interfaceDefinitions.GetOptionalAspectSection()
	var instanceMethods = instanceDefinition.GetInstanceMethods()
	var aspectSubsection = instanceMethods.GetOptionalAspectSubsection()
	var aspectInterfaces = v.generateAspectInterfaces(
		aspectSection,
		aspectSubsection,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"aspectInterfaces",
		aspectInterfaces,
	)

	// Add in the instance public methods.
	var publicMethods = v.generatePublicMethods(instanceDefinition)
	implementation = uti.ReplaceAll(
		implementation,
		"publicMethods",
		publicMethods,
	)

	// Add in the instance private methods.
	var privateMethods = v.generatePrivateMethods(instanceDefinition)
	implementation = uti.ReplaceAll(
		implementation,
		"privateMethods",
		privateMethods,
	)

	// Add in the private instance structure.
	var instanceStructure = v.generateInstanceStructure()
	implementation = uti.ReplaceAll(
		implementation,
		"instanceStructure",
		instanceStructure,
	)

	// Add in the private class structure.
	var classStructure = v.generateClassStructure()
	implementation = uti.ReplaceAll(
		implementation,
		"classStructure",
		classStructure,
	)

	// Add in the private class reference.
	var classReference = v.generateClassReference()
	implementation = uti.ReplaceAll(
		implementation,
		"classReference",
		classReference,
	)

	// Set the classname.
	var className = v.extractClassName(classDefinition)
	implementation = uti.ReplaceAll(
		implementation,
		"className",
		className,
	)

	// Insert generics if necessary.
	var constraints string
	var arguments string
	if v.isGeneric_ {
		constraints = v.generateConstraints(classDefinition)
		arguments = v.generateArguments(classDefinition)
	}
	implementation = uti.ReplaceAll(
		implementation,
		"constraints",
		constraints,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"arguments",
		arguments,
	)

	// Insert any imported modules (this must be done last).
	var moduleImports = v.generateImports(model, implementation)
	implementation = uti.ReplaceAll(
		implementation,
		"moduleImports",
		moduleImports,
	)

	return implementation
}

func (v *classes_) generateClassReference() (
	implementation string,
) {
	implementation = v.getClass().classReference_
	if v.isGeneric_ {
		implementation = v.getClass().genericReference_
	}
	var constantInitializations = v.generateConstantInitializations()
	implementation = uti.ReplaceAll(
		implementation,
		"constantInitializations",
		constantInitializations,
	)
	return implementation
}

func (v *classes_) generateClassStructure() (
	implementation string,
) {
	implementation = v.getClass().classStructure_
	var constantDeclarations = v.generateConstantDeclarations()
	implementation = uti.ReplaceAll(
		implementation,
		"constantDeclarations",
		constantDeclarations,
	)
	return implementation
}

func (v *classes_) generateConstantDeclarations() (
	implementation string,
) {
	var constants = v.constants_.GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var constantType = constant.GetValue()
		var declaration = v.getClass().constantDeclaration_
		declaration = uti.ReplaceAll(declaration, "constantName", constantName)
		declaration = uti.ReplaceAll(declaration, "constantType", constantType)
		implementation += declaration
	}
	return implementation
}

func (v *classes_) generateConstantInitializations() (
	implementation string,
) {
	var constants = v.constants_.GetIterator()
	for constants.HasNext() {
		var constant = constants.GetNext()
		var constantName = constant.GetKey()
		var initialization = v.getClass().constantInitialization_
		initialization = uti.ReplaceAll(initialization, "constantName", constantName)
		implementation += initialization
	}
	return implementation
}

func (v *classes_) generateConstantMethod(constantMethod ast.ConstantMethodLike) (
	implementation string,
) {
	var methodName = constantMethod.GetName()
	var resultType = v.extractType(constantMethod.GetAbstraction())
	implementation = v.getClass().constantMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	return implementation
}

func (v *classes_) generateConstantMethods(
	classDefinition ast.ClassDefinitionLike,
) (
	implementation string,
) {
	var classMethods = classDefinition.GetClassMethods()
	var constantSubsection = classMethods.GetOptionalConstantSubsection()
	if uti.IsDefined(constantSubsection) {
		var methods string
		var constantMethods = constantSubsection.GetConstantMethods().GetIterator()
		for constantMethods.HasNext() {
			var constantMethod = constantMethods.GetNext()
			methods += v.generateConstantMethod(constantMethod)
		}
		implementation = v.getClass().constantMethods_
		implementation = uti.ReplaceAll(implementation, "methods", methods)
	}
	return implementation
}

func (v *classes_) generateConstraints(
	classDefinition ast.ClassDefinitionLike,
) (
	constraints string,
) {
	if v.isGeneric_ {
		constraints = "["
		var classDeclaration = classDefinition.GetDeclaration()
		var optionalConstraints = classDeclaration.GetOptionalConstraints()
		var constraint = optionalConstraints.GetConstraint()
		var constraintName = constraint.GetName()
		var constraintType = v.extractType(constraint.GetAbstraction())
		constraints += constraintName + " " + constraintType
		var additionalConstraints = optionalConstraints.GetAdditionalConstraints().GetIterator()
		for additionalConstraints.HasNext() {
			constraint = additionalConstraints.GetNext().GetConstraint()
			constraintName = constraint.GetName()
			constraintType = v.extractType(constraint.GetAbstraction())
			constraints += ", " + constraintName + " " + constraintType
		}
		constraints += "]"
	}
	return constraints
}

func (v *classes_) generateConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
) (
	implementation string,
) {
	var methodName = constructorMethod.GetName()
	var constructorParameters = constructorMethod.GetParameters()
	var parameters = v.generateParameters(constructorParameters)
	var resultType = v.extractType(constructorMethod.GetAbstraction())
	var attributeChecks = v.generateAttributeChecks(constructorParameters)
	var attributeInitializations = v.generateAttributeInitializations(
		constructorParameters,
	)
	implementation = v.getClass().constructorMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	implementation = uti.ReplaceAll(implementation, "attributeChecks", attributeChecks)
	implementation = uti.ReplaceAll(implementation, "attributeInitializations", attributeInitializations)
	return implementation
}

func (v *classes_) generateConstructorMethods(
	classDefinition ast.ClassDefinitionLike,
) (
	implementation string,
) {
	var methods string
	var classMethods = classDefinition.GetClassMethods()
	var constructorSubsection = classMethods.GetConstructorSubsection()
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		methods += v.generateConstructorMethod(constructorMethod)
	}
	implementation = v.getClass().constructorMethods_
	implementation = uti.ReplaceAll(implementation, "methods", methods)
	return implementation
}

func (v *classes_) generateFunctionMethod(functionMethod ast.FunctionMethodLike) (
	implementation string,
) {
	var methodName = functionMethod.GetName()
	var parameters = v.generateParameters(functionMethod.GetParameters())
	var resultType = v.generateResult(functionMethod.GetResult())
	implementation = v.getClass().functionMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	return implementation
}

func (v *classes_) generateFunctionMethods(
	classDefinition ast.ClassDefinitionLike,
) (
	implementation string,
) {
	var classMethods = classDefinition.GetClassMethods()
	var functionSubsection = classMethods.GetOptionalFunctionSubsection()
	if uti.IsDefined(functionSubsection) {
		var methods string
		var functionMethods = functionSubsection.GetFunctionMethods().GetIterator()
		for functionMethods.HasNext() {
			var functionMethod = functionMethods.GetNext()
			methods += v.generateFunctionMethod(functionMethod)
		}
		implementation = v.getClass().functionMethods_
		implementation = uti.ReplaceAll(implementation, "methods", methods)
	}
	return implementation
}

func (v *classes_) generateGetterMethod(getterMethod ast.GetterMethodLike) (
	implementation string,
) {
	var methodName = getterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var attributeType = v.extractType(getterMethod.GetAbstraction())
	implementation = v.getClass().getterMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "attributeName", attributeName)
	implementation = uti.ReplaceAll(implementation, "attributeType", attributeType)
	return implementation
}

func (v *classes_) generateImports(
	model ast.ModelLike,
	class string,
) (
	implementation string,
) {
	var modules = v.generateModules(model, class)
	if uti.IsDefined(modules) {
		implementation = v.getClass().moduleImports_
		implementation = uti.ReplaceAll(implementation, "modules", modules)
	}
	return implementation
}

func (v *classes_) generateInstanceStructure() (
	implementation string,
) {
	implementation = v.getClass().instanceStructure_
	var attributeDeclarations = v.generateAttributeDeclarations()
	implementation = uti.ReplaceAll(
		implementation,
		"attributeDeclarations",
		attributeDeclarations,
	)
	return implementation
}

func (v *classes_) generateModules(
	model ast.ModelLike,
	class string,
) (
	implementation string,
) {
	var moduleDefinition = model.GetModuleDefinition()
	var imports = moduleDefinition.GetOptionalImports()
	if uti.IsDefined(imports) {
		var modules = imports.GetModules().GetIterator()
		for modules.HasNext() {
			var module = modules.GetNext()
			var moduleName = module.GetName()
			var prefix = moduleName + "."
			var modulePath = module.GetPath()
			if sts.Contains(class, prefix) && !sts.Contains(implementation, prefix) {
				var alias = v.getClass().moduleAlias_
				alias = uti.ReplaceAll(alias, "moduleName", moduleName)
				alias = uti.ReplaceAll(alias, "modulePath", modulePath)
				implementation += alias
			}
		}
	}
	if sts.Contains(class, "fmt.") && !sts.Contains(implementation, "fmt.") {
		var alias = v.getClass().moduleAlias_
		alias = uti.ReplaceAll(alias, "moduleName", "fmt")
		alias = uti.ReplaceAll(alias, "modulePath", "\"fmt\"")
		implementation += alias
	}
	if sts.Contains(class, "uti.") && !sts.Contains(implementation, "uti.") {
		var alias = v.getClass().moduleAlias_
		alias = uti.ReplaceAll(alias, "moduleName", "uti")
		alias = uti.ReplaceAll(alias, "modulePath", "\"github.com/craterdog/go-missing-utilities/v2\"")
		implementation += alias
	}
	if sts.Contains(class, "col.") && !sts.Contains(implementation, "col.") {
		var alias = v.getClass().moduleAlias_
		alias = uti.ReplaceAll(alias, "moduleName", "col")
		alias = uti.ReplaceAll(alias, "modulePath", "\"github.com/craterdog/go-collection-framework/v4\"")
		implementation += alias
	}
	if sts.Contains(class, "abs.") && !sts.Contains(implementation, "abs.") {
		var alias = v.getClass().moduleAlias_
		alias = uti.ReplaceAll(alias, "moduleName", "abs")
		alias = uti.ReplaceAll(alias, "modulePath", "\"github.com/craterdog/go-collection-framework/v4/collection\"")
		implementation += alias
	}
	if sts.Contains(class, "syn.") && !sts.Contains(implementation, "syn.") {
		var alias = v.getClass().moduleAlias_
		alias = uti.ReplaceAll(alias, "moduleName", "syn")
		alias = uti.ReplaceAll(alias, "modulePath", "\"sync\"")
		implementation += alias
	}
	if uti.IsDefined(implementation) {
		implementation += "\n"
	}
	return implementation
}

func (v *classes_) generatePackageDeclaration(model ast.ModelLike) (
	implementation string,
) {
	var packageName = v.extractPackageName(model)
	implementation = v.getClass().packageDeclaration_
	implementation = uti.ReplaceAll(implementation, "packageName", packageName)
	return implementation
}

func (v *classes_) generateParameters(
	parameters abs.Sequential[ast.ParameterLike],
) (
	implementation string,
) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(parameter.GetAbstraction())
		var template = v.getClass().methodParameter_
		template = uti.ReplaceAll(template, "parameterName", parameterName)
		template = uti.ReplaceAll(template, "parameterType", parameterType)
		implementation += template
	}
	if uti.IsDefined(implementation) {
		implementation += "\n"
	}
	return implementation
}

func (v *classes_) generatePublicMethod(publicMethod ast.PublicMethodLike) (
	implementation string,
) {
	var method = publicMethod.GetMethod()
	var methodName = method.GetName()
	var parameters = v.generateParameters(method.GetParameters())
	var resultType = v.generateResult(method.GetOptionalResult())
	implementation = v.getClass().instanceMethod_
	if uti.IsDefined(resultType) {
		implementation = v.getClass().instanceFunction_
		implementation = uti.ReplaceAll(implementation, "resultType", resultType)
	}
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "parameters", parameters)
	return implementation
}

func (v *classes_) generatePublicMethods(
	instanceDefinition ast.InstanceDefinitionLike,
) (
	implementation string,
) {
	var methods string
	var instanceMethods = instanceDefinition.GetInstanceMethods()
	var publicSubsection = instanceMethods.GetPublicSubsection()
	var publicMethods = publicSubsection.GetPublicMethods().GetIterator()
	for publicMethods.HasNext() {
		var publicMethod = publicMethods.GetNext()
		if publicMethod.GetMethod().GetName() != "GetClass" {
			methods += v.generatePublicMethod(publicMethod)
		}
	}
	implementation = v.getClass().publicMethods_
	implementation = uti.ReplaceAll(implementation, "methods", methods)
	return implementation
}

func (v *classes_) generatePrivateMethods(
	instance ast.InstanceDefinitionLike,
) (
	implementation string,
) {
	implementation = v.getClass().privateMethods_
	return implementation
}

func (v *classes_) generateResult(
	result ast.ResultLike,
) (
	implementation string,
) {
	if uti.IsDefined(result) {
		switch actual := result.GetAny().(type) {
		case ast.AbstractionLike:
			implementation = v.extractType(actual)
		case ast.ParameterizedLike:
			implementation = "(" + v.generateParameters(actual.GetParameters()) + "\n)"
		}
	}
	return implementation
}

func (v *classes_) generateSetterMethod(setterMethod ast.SetterMethodLike) (
	implementation string,
) {
	var methodName = setterMethod.GetName()
	var attributeName = v.extractAttributeName(methodName)
	var parameter = setterMethod.GetParameter()
	var attributeType = v.extractType(parameter.GetAbstraction())
	var attributeCheck = v.generateAttributeCheck(parameter)
	implementation = v.getClass().setterMethod_
	implementation = uti.ReplaceAll(implementation, "methodName", methodName)
	implementation = uti.ReplaceAll(implementation, "attributeName", attributeName)
	implementation = uti.ReplaceAll(implementation, "attributeType", attributeType)
	implementation = uti.ReplaceAll(implementation, "attributeCheck", attributeCheck)
	return implementation
}

func (v *classes_) replaceAbstractionType(
	abstraction ast.AbstractionLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) ast.AbstractionLike {
	// Replace the generic type in a prefix with the concrete type.
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		prefix = v.replacePrefixType(prefix, mappings)
	}

	// Replace the generic types in a sequence of arguments with concrete types.
	var arguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(arguments) {
		arguments = v.replaceArgumentTypes(arguments, mappings)
	}

	// Replace a non-suffixed generic type with its concrete type.
	var typeName = abstraction.GetName()
	var suffix = abstraction.GetOptionalSuffix()
	if uti.IsUndefined(suffix) {
		var concreteType = mappings.GetValue(typeName)
		if uti.IsDefined(concreteType) {
			suffix = concreteType.GetOptionalSuffix()
			typeName = concreteType.GetName()
			arguments = concreteType.GetOptionalArguments()
		}
	}

	// Recreate the abstraction using its updated types.
	abstraction = ast.Abstraction().Make(
		prefix,
		typeName,
		suffix,
		arguments,
	)

	return abstraction
}

func (v *classes_) replaceArgumentType(
	argument ast.ArgumentLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) ast.ArgumentLike {
	var abstraction = argument.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	argument = ast.Argument().Make(abstraction)
	return argument
}

func (v *classes_) replaceArgumentTypes(
	arguments ast.ArgumentsLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) ast.ArgumentsLike {
	// Replace the generic type of the first argument with its concrete type.
	var argument = arguments.GetArgument()
	argument = v.replaceArgumentType(argument, mappings)

	// Replace the generic types of any additional arguments with concrete types.
	var additionalArguments = col.List[ast.AdditionalArgumentLike]()
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

func (v *classes_) replaceParameterizedTypes(
	parameterized ast.ParameterizedLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) ast.ParameterizedLike {
	var parameters = parameterized.GetParameters()
	var replacedParameters = v.replaceParameterTypes(parameters, mappings)
	parameterized = ast.Parameterized().Make(replacedParameters)
	return parameterized
}

func (v *classes_) replaceParameterType(
	parameter ast.ParameterLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) ast.ParameterLike {
	var parameterName = parameter.GetName()
	var abstraction = parameter.GetAbstraction()
	abstraction = v.replaceAbstractionType(abstraction, mappings)
	parameter = ast.Parameter().Make(parameterName, abstraction)
	return parameter
}

func (v *classes_) replaceParameterTypes(
	parameters abs.Sequential[ast.ParameterLike],
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) abs.Sequential[ast.ParameterLike] {
	var replacedParameters = col.List[ast.ParameterLike]()
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		parameter = v.replaceParameterType(parameter, mappings)
		replacedParameters.AppendValue(parameter)
	}
	return replacedParameters
}

func (v *classes_) replacePrefixType(
	prefix ast.PrefixLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
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

func (v *classes_) replaceResultType(
	result ast.ResultLike,
	mappings abs.CatalogLike[string, ast.AbstractionLike],
) ast.ResultLike {
	if uti.IsUndefined(result) {
		return result
	}
	switch actual := result.GetAny().(type) {
	case ast.NoneLike:
	case ast.AbstractionLike:
		var abstraction = actual
		abstraction = v.replaceAbstractionType(abstraction, mappings)
		result = ast.Result().Make(abstraction)
	case ast.ParameterizedLike:
		var parameterized = actual
		parameterized = v.replaceParameterizedTypes(parameterized, mappings)
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

// PRIVATE INTERFACE

// Instance Structure

type classes_ struct {
	// Declare the instance attributes.
	class_      *classesClass_
	isGeneric_  bool
	constants_  abs.CatalogLike[string, string]
	attributes_ abs.CatalogLike[string, string]
}

// Class Structure

type classesClass_ struct {
	// Declare the class constants.
	classTemplate_           string
	packageDeclaration_      string
	moduleImports_           string
	moduleAlias_             string
	accessFunction_          string
	classFunction_           string
	genericFunction_         string
	methodParameter_         string
	constructorMethods_      string
	constructorMethod_       string
	attributeCheck_          string
	attributeInitialization_ string
	constantMethods_         string
	constantMethod_          string
	functionMethods_         string
	functionMethod_          string
	attributeMethods_        string
	getterMethod_            string
	setterMethod_            string
	aspectInterface_         string
	publicMethods_           string
	privateMethods_          string
	instanceMethod_          string
	instanceFunction_        string
	instanceStructure_       string
	attributeDeclaration_    string
	classStructure_          string
	constantDeclaration_     string
	constantInitialization_  string
	classReference_          string
	genericReference_        string
}

// Class Reference

var classesClass = &classesClass_{
	// Initialize the class constants.
	classTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// CLASS INTERFACE<AccessFunction><ConstructorMethods><ConstantMethods><FunctionMethods>

// INSTANCE INTERFACE<AttributeMethods><AspectInterfaces><PublicMethods><PrivateMethods>

// PRIVATE INTERFACE<InstanceStructure><ClassStructure><ClassReference>
`,

	packageDeclaration_: `
package <~packageName>`,

	moduleImports_: `

import (<Modules>)`,

	moduleAlias_: `
	<~moduleName> <modulePath>`,

	accessFunction_: `

// Access Function

func <~ClassName><Constraints>() <~ClassName>ClassLike<Arguments> {
	return <~className>Reference<Arguments>()
}`,

	methodParameter_: `
	<parameterName_> <ParameterType>,`,

	constructorMethods_: `

// Constructor Methods<Methods>`,

	constructorMethod_: `

func (c *<~className>Class_<Arguments>) <MethodName>(<Parameters>) <~ClassName>Like<Arguments> {<AttributeChecks>
	var instance = &<~className>_<Arguments>{
		// Initialize the instance attributes.
		class_: c,<AttributeInitializations>
	}
	return instance
}`,

	attributeCheck_: `
	if uti.IsUndefined(<attributeName_>) {
		panic("The \"<~attributeName>\" attribute is required by this class.")
	}`,

	attributeInitialization_: `
		<~attributeName>_: <attributeName_>,`,

	constantMethods_: `

// Constant Methods<Methods>`,

	constantMethod_: `

func (c *<~className>Class_<Arguments>) <~MethodName>() <ResultType> {
	return c.<~methodName>_
}`,

	functionMethods_: `

// Function Methods<Methods>`,

	functionMethod_: `

func (c *<~className>Class_<Arguments>) <~MethodName>(<Parameters>) <ResultType> {
	var result_ <ResultType>
	// TBD - Add the function implementation.
	return result_
}`,

	attributeMethods_: `

// Attribute Methods<Methods>`,

	getterMethod_: `

func (v *<~className>_<Arguments>) <~MethodName>() <AttributeType> {
	return v.<~attributeName>_
}`,

	setterMethod_: `

func (v *<~className>_<Arguments>) <~MethodName>(
	<attributeName_> <AttributeType>,
) {<AttributeCheck>
	v.<~attributeName>_ = <attributeName_>
}`,

	aspectInterface_: `

// <AspectType> Methods<Methods>`,

	publicMethods_: `

// Public Methods

func (v *<~className>_<Arguments>) GetClass() <~ClassName>ClassLike<Arguments> {
	return v.getClass()
}<Methods>`,

	privateMethods_: `

// Private Methods

func (v *<~className>_<Arguments>) getClass() *<~className>Class_<Arguments> {
	return v.class_
}`,

	instanceMethod_: `

func (v *<~className>_<Arguments>) <~MethodName>(<Parameters>) {
	// TBD - Add the method implementation.
}`,

	instanceFunction_: `

func (v *<~className>_<Arguments>) <~MethodName>(<Parameters>) <ResultType> {
	var result_ <ResultType>
	// TBD - Add the method implementation.
	return result_
}`,

	instanceStructure_: `

// Instance Structure

type <~className>_<Constraints> struct {
	// Declare the instance attributes.
	class_ *<~className>Class_<Arguments><AttributeDeclarations>
}`,

	attributeDeclaration_: `
	<~attributeName>_ <AttributeType>`,

	classStructure_: `

// Class Structure

type <~className>Class_<Constraints> struct {
	// Declare the class constants.<ConstantDeclarations>
}`,

	constantDeclaration_: `
	<~constantName>_ <ConstantType>`,

	constantInitialization_: `
	// <~constantName>_: constantValue,`,

	classReference_: `

// Class Reference

func <~className>Reference() *<~className>Class_ {
	return <~className>Reference_
}

var <~className>Reference_ = &<~className>Class_{
	// Initialize the class constants.<ConstantInitializations>
}`,

	genericReference_: `

// Class Reference

var <~className>Reference_ = map[string]any{}
var <~className>Mutex_ syn.Mutex

func <~className>Reference<Constraints>() *<~className>Class_<Arguments> {
	// Generate the name of the bound class type.
	var class *<className>Class_<Arguments>
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	<className>Mutex_.Lock()
	var value = <className>Reference_[name]
	switch actual := value.(type) {
	case *<className>Class_<Arguments>:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &<className>Class_<Arguments>{
			// Initialize the class constants.<ConstantInitializations>
		}
		<className>Reference_[name] = class
	}
	<className>Mutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}`,
}
