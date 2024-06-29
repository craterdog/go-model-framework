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
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Initialize class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// Define class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define instance attributes.
	class_        ValidatorClassLike
	imports_      col.CatalogLike[string, ast.ModuleLike]
	types_        col.CatalogLike[string, ast.TypeLike]
	functionals_  col.CatalogLike[string, ast.FunctionalLike]
	aspects_      col.CatalogLike[string, ast.AspectLike]
	classes_      col.CatalogLike[string, ast.ClassLike]
	instances_    col.CatalogLike[string, ast.InstanceLike]
	abstractions_ col.CatalogLike[string, ast.AbstractionLike]
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateModel(model ast.ModelLike) {
	// Initialize the catalogs.
	var notation = cdc.Notation().Make()
	v.imports_ = col.Catalog[string, ast.ModuleLike](notation).Make()
	v.types_ = col.Catalog[string, ast.TypeLike](notation).Make()
	v.functionals_ = col.Catalog[string, ast.FunctionalLike](notation).Make()
	v.aspects_ = col.Catalog[string, ast.AspectLike](notation).Make()
	v.classes_ = col.Catalog[string, ast.ClassLike](notation).Make()
	v.instances_ = col.Catalog[string, ast.InstanceLike](notation).Make()
	v.abstractions_ = col.Catalog[string, ast.AbstractionLike](notation).Make()

	// Extract the catalogs.
	v.extractModules(model)
	v.extractTypes(model)
	v.extractFunctionals(model)
	v.extractAspects(model)
	v.extractClasses(model)
	v.extractInstances(model)

	// Validate the catalogs (note, the order matters).
	v.validateModules(model)
	v.validateClasses(model)
	v.validateInstances(model)
	v.validatePairings(model)
	v.validateAspects(model)
	v.validateTypes(model)
	v.validateFunctionals(model)
}

// Private

func (v *validator_) extractAspects(model ast.ModelLike) {
	var aspects = model.GetAspects()
	if aspects == nil {
		return
	}
	var iterator = aspects.GetAspectIterator()
	for iterator.HasNext() {
		var aspect = iterator.GetNext()
		var name = sts.ToLower(aspect.GetDeclaration().GetName())
		v.aspects_.SetValue(name, aspect)
	}
}

func (v *validator_) extractClasses(model ast.ModelLike) {
	var classes = model.GetClasses()
	if classes == nil {
		return
	}
	var iterator = classes.GetClassIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var name = class.GetDeclaration().GetName()
		name = sts.TrimSuffix(name, "ClassLike")
		name = sts.ToLower(name)
		v.classes_.SetValue(name, class)
	}
}

func (v *validator_) extractFunctionals(model ast.ModelLike) {
	var functionals = model.GetFunctionals()
	if functionals == nil {
		return
	}
	var iterator = functionals.GetFunctionalIterator()
	for iterator.HasNext() {
		var functional = iterator.GetNext()
		var name = sts.ToLower(functional.GetDeclaration().GetName())
		v.functionals_.SetValue(name, functional)
	}
}

func (v *validator_) extractInstances(model ast.ModelLike) {
	var instances = model.GetInstances()
	if instances == nil {
		return
	}
	var iterator = instances.GetInstanceIterator()
	for iterator.HasNext() {
		var instance = iterator.GetNext()
		var name = instance.GetDeclaration().GetName()
		name = sts.TrimSuffix(name, "Like")
		name = sts.ToLower(name)
		v.instances_.SetValue(name, instance)
	}
}

func (v *validator_) extractModules(model ast.ModelLike) {
	var imports = model.GetImports()
	if imports == nil {
		return
	}
	var iterator = imports.GetModules().GetModuleIterator()
	for iterator.HasNext() {
		var module = iterator.GetNext()
		var path = module.GetPath()
		v.imports_.SetValue(path, module)
	}
}

func (v *validator_) extractTypes(model ast.ModelLike) {
	var types = model.GetTypes()
	if types == nil {
		return
	}
	var iterator = types.GetTypeIterator()
	for iterator.HasNext() {
		var type_ = iterator.GetNext()
		var name = sts.ToLower(type_.GetDeclaration().GetName())
		v.types_.SetValue(name, type_)
	}
}

func (v *validator_) validateAbstraction(abstraction ast.AbstractionLike) {
	var prefix = abstraction.GetPrefix()
	if prefix != nil {
		v.validatePrefix(prefix)
	}
	var name = abstraction.GetName()
	v.abstractions_.SetValue(name, abstraction)
	var arguments = abstraction.GetGenericArguments().GetArguments()
	if arguments != nil {
		v.validateArguments(arguments)
	}
}

func (v *validator_) validateAbstractions(abstractions ast.AbstractionsLike) {
	var iterator = abstractions.GetAbstractionIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.validateAbstraction(abstraction)
	}
}

func (v *validator_) validateArguments(arguments ast.ArgumentsLike) {
	var abstraction = arguments.GetArgument().GetAbstraction()
	v.validateAbstraction(abstraction)
	var iterator = arguments.GetAdditionalArguments().GetAdditionalArgumentIterator()
	for iterator.HasNext() {
		abstraction = iterator.GetNext().GetArgument().GetAbstraction()
		v.validateAbstraction(abstraction)
	}
}

func (v *validator_) validateAspect(aspect ast.AspectLike) {
	var declaration = aspect.GetDeclaration()
	v.validateDeclaration(declaration)
	var methods = aspect.GetMethods()
	v.validateMethods(methods)
	var name = declaration.GetName()
	var abstraction = v.abstractions_.GetValue(name)
	if abstraction == nil {
		var message = fmt.Sprintf(
			"The following aspect is never used in this package: %v",
			name,
		)
		panic(message)
	}
}

func (v *validator_) validateAspects(model ast.ModelLike) {
	var iterator = v.aspects_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var aspect = association.GetValue()
		v.validateAspect(aspect)
	}
}

func (v *validator_) validateAttribute(attribute ast.AttributeLike) {
	var name = attribute.GetName()
	var parameter = attribute.GetParameter()
	var abstraction = attribute.GetAbstraction()
	switch {
	case sts.HasPrefix(name, "Get"):
		v.validateAbstraction(abstraction)
	case sts.HasPrefix(name, "Set"):
		v.validateParameter(parameter)
	case sts.HasPrefix(name, "Is"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(name, "Are"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(name, "Was"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(name, "Were"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(name, "Has"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(name, "Had"):
		v.validateBoolean(abstraction)
	default:
		var message = fmt.Sprintf(
			"Found an illegal attribute method name: %v",
			name,
		)
		panic(message)
	}
}

func (v *validator_) validateAttributes(attributes ast.AttributesLike) {
	var iterator = attributes.GetAttributeIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		v.validateAttribute(attribute)
	}
}

func (v *validator_) validateBoolean(abstraction ast.AbstractionLike) {
	var prefix = abstraction.GetPrefix()
	if prefix != nil {
		panic("A boolean type cannot have a prefix.")
	}
	var name = abstraction.GetName()
	if name != "bool" {
		panic("A question attribute must have a boolean type.")
	}
	var arguments = abstraction.GetGenericArguments()
	if arguments != nil {
		panic("A boolean type cannot be a generic type.")
	}
}

func (v *validator_) validateClass(class ast.ClassLike) {
	var declaration = class.GetDeclaration()
	v.validateDeclaration(declaration)
	var constructors = class.GetConstructors()
	if constructors != nil {
		v.validateConstructors(constructors)
	}
	var constants = class.GetConstants()
	if constants != nil {
		v.validateConstants(constants)
	}
	var functions = class.GetFunctions()
	if functions != nil {
		v.validateFunctions(functions)
	}
}

func (v *validator_) validateClasses(model ast.ModelLike) {
	var iterator = v.classes_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var class = association.GetValue()
		v.validateClass(class)
	}
}

func (v *validator_) validateConstant(constant ast.ConstantLike) {
	var abstraction = constant.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstants(constants ast.ConstantsLike) {
	var iterator = constants.GetConstantIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.validateConstant(constant)
	}
}

func (v *validator_) validateConstructor(constructor ast.ConstructorLike) {
	var parameters = constructor.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var abstraction = constructor.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstructors(constructors ast.ConstructorsLike) {
	var iterator = constructors.GetConstructorIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.validateConstructor(constructor)
	}
}

func (v *validator_) validateDeclaration(declaration ast.DeclarationLike) {
	var parameters = declaration.GetGenericParameters()
	if parameters != nil {
		v.validateParameters(parameters.GetParameters())
	}
}

func (v *validator_) validateEnumeration(enumeration ast.EnumerationLike) {
	var values = enumeration.GetValues()
	var value = values.GetValue()
	var abstraction = value.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateFunction(function ast.FunctionLike) {
	var parameters = function.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = function.GetResult()
	v.validateResult(result)
}

func (v *validator_) validateFunctional(functional ast.FunctionalLike) {
	var declaration = functional.GetDeclaration()
	v.validateDeclaration(declaration)
	var parameters = functional.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = functional.GetResult()
	v.validateResult(result)
	var name = declaration.GetName()
	var abstraction = v.abstractions_.GetValue(name)
	if abstraction == nil {
		var message = fmt.Sprintf(
			"The following functional is never used in this package: %v",
			name,
		)
		panic(message)
	}
}

func (v *validator_) validateFunctionals(model ast.ModelLike) {
	var iterator = v.functionals_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var functional = association.GetValue()
		v.validateFunctional(functional)
	}
}

func (v *validator_) validateFunctions(functions ast.FunctionsLike) {
	var iterator = functions.GetFunctionIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.validateFunction(function)
	}
}

func (v *validator_) validateGetClassMethod(class string) {
	var instance = v.instances_.GetValue(class)
	var attributes = instance.GetAttributes()
	if attributes != nil {
		var iterator = attributes.GetAttributeIterator()
		for iterator.HasNext() {
			var attribute = iterator.GetNext()
			var name = attribute.GetName()
			if name == "GetClass" {
				var abstraction = attribute.GetAbstraction()
				if class+"classlike" == sts.ToLower(abstraction.GetName()) {
					return
				}
			}
		}
	}
	fmt.Printf(
		"The following class is missing a GetClass() instance method: %v\n",
		sts.TrimSuffix(instance.GetDeclaration().GetName(), "Like"),
	)
}

func (v *validator_) validateInstance(instance ast.InstanceLike) {
	var declaration = instance.GetDeclaration()
	v.validateDeclaration(declaration)
	var attributes = instance.GetAttributes()
	if attributes != nil {
		v.validateAttributes(attributes)
	}
	var abstractions = instance.GetAbstractions()
	if abstractions != nil {
		v.validateAbstractions(abstractions)
	}
	var methods = instance.GetMethods()
	if methods != nil {
		v.validateMethods(methods)
	}
}

func (v *validator_) validateInstances(model ast.ModelLike) {
	var iterator = v.instances_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var instance = association.GetValue()
		v.validateInstance(instance)
	}
}

func (v *validator_) validateMethod(method ast.MethodLike) {
	var parameters = method.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = method.GetResult()
	if result != nil {
		v.validateResult(result)
	}
}

func (v *validator_) validateMethods(methods ast.MethodsLike) {
	var iterator = methods.GetMethodIterator()
	for iterator.HasNext() {
		var method = iterator.GetNext()
		v.validateMethod(method)
	}
}

func (v *validator_) validateModule(module ast.ModuleLike) {
	var name = module.GetName()
	if len(name) != 3 {
		var message = fmt.Sprintf(
			"The length of the name for an imported module must be 3: %v",
			name,
		)
		panic(message)
	}
}

func (v *validator_) validateModules(model ast.ModelLike) {
	var iterator = v.imports_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var module = association.GetValue()
		v.validateModule(module)
	}
}

func (v *validator_) validatePairings(model ast.ModelLike) {
	// Make sure each class interface has an associated instance interface.
	var notation = cdc.Notation().Make()
	var classes = col.List[string](notation).MakeFromSequence(v.classes_.GetKeys())
	var instances = col.List[string](notation).MakeFromSequence(v.instances_.GetKeys())
	if classes.GetSize() != instances.GetSize() {
		var message = fmt.Sprintf(
			"Mismatched class and instance interfaces:\n%v\n%v\n",
			classes,
			instances,
		)
		panic(message)
	}
	var classIterator = classes.GetIterator()
	var instanceIterator = instances.GetIterator()
	for classIterator.HasNext() {
		var class = classIterator.GetNext()
		var instance = instanceIterator.GetNext()
		if class != instance {
			var message = fmt.Sprintf(
				"Mismatched class and instance interfaces:\n%v\n%v\n",
				class,
				instance,
			)
			panic(message)
		}
		v.validateGetClassMethod(class)
	}
}

func (v *validator_) validateParameter(parameter ast.ParameterLike) {
	var abstraction = parameter.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateParameters(parameters ast.ParametersLike) {
	var parameter = parameters.GetParameter()
	v.validateParameter(parameter)
	var additionalParameters = parameters.GetAdditionalParameters()
	if additionalParameters != nil {
		var iterator = additionalParameters.GetAdditionalParameterIterator()
		for iterator.HasNext() {
			var parameter = iterator.GetNext()
			v.validateParameter(parameter.GetParameter())
		}
	}
}

func (v *validator_) validatePrefix(prefix ast.PrefixLike) {
	switch actual := prefix.GetAny().(type) {
	case ast.AliasLike:
		var name = actual.GetName()
		var iterator = v.imports_.GetIterator()
		for iterator.HasNext() {
			var association = iterator.GetNext()
			var module = association.GetValue()
			if module.GetName() == name {
				// Found a matching alias.
				return
			}
		}
		var message = fmt.Sprintf(
			"Unknown module alias name: %v",
			name,
		)
		panic(message)
	default:
		// Ignore the other prefix types.
	}
}

func (v *validator_) validateResult(result ast.ResultLike) {
	switch actual := result.GetAny().(type) {
	case ast.AbstractionLike:
		v.validateAbstraction(actual)
	case ast.ParameterizedLike:
		var parameters = actual.GetParameters()
		v.validateParameters(parameters)
	default:
		var message = fmt.Sprintf(
			"Found an unknown result type: %T",
			actual,
		)
		panic(message)
	}
}

func (v *validator_) validateType(type_ ast.TypeLike) {
	var declaration = type_.GetDeclaration()
	v.validateDeclaration(declaration)
	var abstraction = type_.GetAbstraction()
	v.validateAbstraction(abstraction)
	var enumeration = type_.GetEnumeration()
	if enumeration != nil {
		v.validateEnumeration(enumeration)
	}
	var name = declaration.GetName()
	abstraction = v.abstractions_.GetValue(name)
	if abstraction == nil {
		var message = fmt.Sprintf(
			"The following type is never used in this package: %v",
			name,
		)
		panic(message)
	}
}

func (v *validator_) validateTypes(model ast.ModelLike) {
	var iterator = v.types_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var type_ = association.GetValue()
		v.validateType(type_)
	}
}
