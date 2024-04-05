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
	sts "strings"
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// This class does not initialize any private class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// This class does not define any private class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	modules_      col.CatalogLike[string, ModuleLike]
	types_        col.CatalogLike[string, TypeLike]
	functionals_  col.CatalogLike[string, FunctionalLike]
	aspects_      col.CatalogLike[string, AspectLike]
	classes_      col.CatalogLike[string, ClassLike]
	instances_    col.CatalogLike[string, InstanceLike]
	abstractions_ col.CatalogLike[string, AbstractionLike]
}

// Public

func (v *validator_) ValidateModel(model ModelLike) {
	// Initialize the catalogs.
	v.modules_ = col.Catalog[string, ModuleLike]().Make()
	v.types_ = col.Catalog[string, TypeLike]().Make()
	v.functionals_ = col.Catalog[string, FunctionalLike]().Make()
	v.aspects_ = col.Catalog[string, AspectLike]().Make()
	v.classes_ = col.Catalog[string, ClassLike]().Make()
	v.instances_ = col.Catalog[string, InstanceLike]().Make()
	v.abstractions_ = col.Catalog[string, AbstractionLike]().Make()

	// Extract the catalogs.
	v.extractModules(model)
	v.extractTypes(model)
	v.extractFunctionals(model)
	v.extractAspects(model)
	v.extractClasses(model)
	v.extractInstances(model)
	v.validateClasses()

	// Validate the catalogs.
	v.validateModules()
	v.validateClasses()
	v.validateInstances()
	v.validatePairings()
	v.validateAspects()
	v.validateTypes()
	v.validateFunctionals()
}

// Private

func (v *validator_) extractAspects(model ModelLike) {
	var aspects = model.GetAspects()
	if aspects == nil {
		return
	}
	var iterator = aspects.GetIterator()
	for iterator.HasNext() {
		var aspect = iterator.GetNext()
		var identifier = sts.ToLower(aspect.GetDeclaration().GetIdentifier())
		v.aspects_.SetValue(identifier, aspect)
	}
}

func (v *validator_) extractClasses(model ModelLike) {
	var classes = model.GetClasses()
	if classes == nil {
		return
	}
	var iterator = classes.GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var identifier = class.GetDeclaration().GetIdentifier()
		identifier = sts.TrimSuffix(identifier, "ClassLike")
		identifier = sts.ToLower(identifier)
		v.classes_.SetValue(identifier, class)
	}
}

func (v *validator_) extractFunctionals(model ModelLike) {
	var functionals = model.GetFunctionals()
	if functionals == nil {
		return
	}
	var iterator = functionals.GetIterator()
	for iterator.HasNext() {
		var functional = iterator.GetNext()
		var identifier = sts.ToLower(functional.GetDeclaration().GetIdentifier())
		v.functionals_.SetValue(identifier, functional)
	}
}

func (v *validator_) extractInstances(model ModelLike) {
	var instances = model.GetInstances()
	if instances == nil {
		return
	}
	var iterator = instances.GetIterator()
	for iterator.HasNext() {
		var instance = iterator.GetNext()
		var identifier = instance.GetDeclaration().GetIdentifier()
		identifier = sts.TrimSuffix(identifier, "Like")
		identifier = sts.ToLower(identifier)
		v.instances_.SetValue(identifier, instance)
	}
}

func (v *validator_) extractModules(model ModelLike) {
	var modules = model.GetModules()
	if modules == nil {
		return
	}
	var iterator = modules.GetIterator()
	for iterator.HasNext() {
		var module = iterator.GetNext()
		var identifier = sts.ToLower(module.GetIdentifier())
		v.modules_.SetValue(identifier, module)
	}
}

func (v *validator_) extractTypes(model ModelLike) {
	var types = model.GetTypes()
	if types == nil {
		return
	}
	var iterator = types.GetIterator()
	for iterator.HasNext() {
		var type_ = iterator.GetNext()
		var identifier = sts.ToLower(type_.GetDeclaration().GetIdentifier())
		v.types_.SetValue(identifier, type_)
	}
}

func (v *validator_) validateAbstraction(abstraction AbstractionLike) {
	var prefix = abstraction.GetPrefix()
	if prefix != nil {
		v.validatePrefix(prefix)
	}
	var identifier = abstraction.GetIdentifier()
	v.abstractions_.SetValue(identifier, abstraction)
	var arguments = abstraction.GetArguments()
	if arguments != nil {
		v.validateArguments(arguments)
	}
}

func (v *validator_) validateAbstractions(abstractions col.ListLike[AbstractionLike]) {
	var iterator = abstractions.GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.validateAbstraction(abstraction)
	}
}

func (v *validator_) validateArguments(arguments col.ListLike[AbstractionLike]) {
	var iterator = arguments.GetIterator()
	for iterator.HasNext() {
		var argument = iterator.GetNext()
		v.validateAbstraction(argument)
	}
}

func (v *validator_) validateAspect(aspect AspectLike) {
	var declaration = aspect.GetDeclaration()
	v.validateDeclaration(declaration)
	var methods = aspect.GetMethods()
	if methods != nil {
		v.validateMethods(methods)
	}
	var identifier = declaration.GetIdentifier()
	var abstraction = v.abstractions_.GetValue(identifier)
	if abstraction == nil {
		var message = fmt.Sprintf(
			"The following aspect is never used in this package: %v",
			identifier,
		)
		panic(message)
	}
}

func (v *validator_) validateAspects() {
	v.aspects_.SortValues()
	var iterator = v.aspects_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var aspect = association.GetValue()
		v.validateAspect(aspect)
	}
}

func (v *validator_) validateAttribute(attribute AttributeLike) {
	var identifier = attribute.GetIdentifier()
	var parameter = attribute.GetParameter()
	var abstraction = attribute.GetAbstraction()
	switch {
	case sts.HasPrefix(identifier, "Get"):
		v.validateAbstraction(abstraction)
	case sts.HasPrefix(identifier, "Set"):
		v.validateParameter(parameter)
	case sts.HasPrefix(identifier, "Is"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(identifier, "Are"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(identifier, "Was"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(identifier, "Were"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(identifier, "Has"):
		v.validateBoolean(abstraction)
	case sts.HasPrefix(identifier, "Had"):
		v.validateBoolean(abstraction)
	default:
		var message = fmt.Sprintf(
			"Found an illegal attribute method name: %v",
			identifier,
		)
		panic(message)
	}
}

func (v *validator_) validateAttributes(attributes col.ListLike[AttributeLike]) {
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		v.validateAttribute(attribute)
	}
}

func (v *validator_) validateBoolean(abstraction AbstractionLike) {
	var prefix = abstraction.GetPrefix()
	if prefix != nil {
		panic("A boolean type cannot have a prefix.")
	}
	var identifier = abstraction.GetIdentifier()
	if identifier != "bool" {
		panic("A question attribute must have a boolean type.")
	}
	var arguments = abstraction.GetArguments()
	if arguments != nil {
		panic("A boolean type cannot be a generic type.")
	}
}

func (v *validator_) validateClass(class ClassLike) {
	var declaration = class.GetDeclaration()
	v.validateDeclaration(declaration)
	var constants = class.GetConstants()
	if constants != nil {
		v.validateConstants(constants)
	}
	var constructors = class.GetConstructors()
	if constructors != nil {
		v.validateConstructors(constructors)
	}
	var functions = class.GetFunctions()
	if functions != nil {
		v.validateFunctions(functions)
	}
}

func (v *validator_) validateClasses() {
	v.classes_.SortValues()
	var iterator = v.classes_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var class = association.GetValue()
		v.validateClass(class)
	}
}

func (v *validator_) validateConstant(constant ConstantLike) {
	var abstraction = constant.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstants(constants col.ListLike[ConstantLike]) {
	var iterator = constants.GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.validateConstant(constant)
	}
}

func (v *validator_) validateConstructor(constructor ConstructorLike) {
	var parameters = constructor.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var abstraction = constructor.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstructors(constructors col.ListLike[ConstructorLike]) {
	var iterator = constructors.GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.validateConstructor(constructor)
	}
}

func (v *validator_) validateDeclaration(declaration DeclarationLike) {
	var parameters = declaration.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
}

func (v *validator_) validateEnumeration(enumeration EnumerationLike) {
	var parameter = enumeration.GetParameter()
	v.validateParameter(parameter)
}

func (v *validator_) validateFunction(function FunctionLike) {
	var parameters = function.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = function.GetResult()
	v.validateResult(result)
}

func (v *validator_) validateFunctional(functional FunctionalLike) {
	var declaration = functional.GetDeclaration()
	v.validateDeclaration(declaration)
	var parameters = functional.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = functional.GetResult()
	v.validateResult(result)
	var identifier = declaration.GetIdentifier()
	var abstraction = v.abstractions_.GetValue(identifier)
	if abstraction == nil {
		var message = fmt.Sprintf(
			"The following functional is never used in this package: %v",
			identifier,
		)
		panic(message)
	}
}

func (v *validator_) validateFunctionals() {
	v.functionals_.SortValues()
	var iterator = v.functionals_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var functional = association.GetValue()
		v.validateFunctional(functional)
	}
}

func (v *validator_) validateFunctions(functions col.ListLike[FunctionLike]) {
	var iterator = functions.GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.validateFunction(function)
	}
}

func (v *validator_) validateInstance(instance InstanceLike) {
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

func (v *validator_) validateInstances() {
	v.instances_.SortValues()
	var iterator = v.instances_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var instance = association.GetValue()
		v.validateInstance(instance)
	}
}

func (v *validator_) validateMethod(method MethodLike) {
	var parameters = method.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = method.GetResult()
	if result != nil {
		v.validateResult(result)
	}
}

func (v *validator_) validateMethods(methods col.ListLike[MethodLike]) {
	var iterator = methods.GetIterator()
	for iterator.HasNext() {
		var method = iterator.GetNext()
		v.validateMethod(method)
	}
}

func (v *validator_) validateModule(module ModuleLike) {
	var identifier = module.GetIdentifier()
	if len(identifier) != 3 {
		var message = fmt.Sprintf(
			"The length of the identifier for an imported module must be 3: %v",
			identifier,
		)
		panic(message)
	}
}

func (v *validator_) validateModules() {
	var iterator = v.modules_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var module = association.GetValue()
		v.validateModule(module)
	}
}

func (v *validator_) validatePairings() {
	// Make sure each class interface has an associated instance interface.
	var classes = col.List[string]().MakeFromSequence(v.classes_.GetKeys())
	var instances = col.List[string]().MakeFromSequence(v.instances_.GetKeys())
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
	}
}

func (v *validator_) validateParameter(parameter ParameterLike) {
	var abstraction = parameter.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateParameters(parameters col.ListLike[ParameterLike]) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		v.validateParameter(parameter)
	}
}

func (v *validator_) validatePrefix(prefix PrefixLike) {
	if prefix == nil || prefix.GetType() != AliasPrefix {
		return
	}
	var identifier = prefix.GetIdentifier()
	if v.modules_.GetValue(identifier) == nil {
		var message = fmt.Sprintf(
			"Unknown module alias: %v",
			identifier,
		)
		panic(message)
	}
}

func (v *validator_) validateResult(result ResultLike) {
	var abstraction = result.GetAbstraction()
	if abstraction != nil {
		v.validateAbstraction(abstraction)
	} else {
		var parameters = result.GetParameters()
		v.validateParameters(parameters)
	}
}

func (v *validator_) validateType(type_ TypeLike) {
	var declaration = type_.GetDeclaration()
	v.validateDeclaration(declaration)
	var abstraction = type_.GetAbstraction()
	v.validateAbstraction(abstraction)
	var enumeration = type_.GetEnumeration()
	if enumeration != nil {
		v.validateEnumeration(enumeration)
	}
	var identifier = declaration.GetIdentifier()
	abstraction = v.abstractions_.GetValue(identifier)
	if abstraction == nil {
		var message = fmt.Sprintf(
			"The following type is never used in this package: %v",
			identifier,
		)
		panic(message)
	}
}

func (v *validator_) validateTypes() {
	v.types_.SortValues()
	var iterator = v.types_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var type_ = association.GetValue()
		v.validateType(type_)
	}
}
