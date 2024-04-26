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
	gcf "github.com/craterdog/go-collection-framework/v4"
	col "github.com/craterdog/go-collection-framework/v4/collection"
	gcm "github.com/craterdog/go-model-framework/v4/gcmn"
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
	class_        ValidatorClassLike
	modules_      col.CatalogLike[string, gcm.ModuleLike]
	types_        col.CatalogLike[string, gcm.TypeLike]
	functionals_  col.CatalogLike[string, gcm.FunctionalLike]
	aspects_      col.CatalogLike[string, gcm.AspectLike]
	classes_      col.CatalogLike[string, gcm.ClassLike]
	instances_    col.CatalogLike[string, gcm.InstanceLike]
	abstractions_ col.CatalogLike[string, gcm.AbstractionLike]
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateModel(model gcm.ModelLike) {
	// Initialize the catalogs.
	v.modules_ = gcf.Catalog[string, gcm.ModuleLike]()
	v.types_ = gcf.Catalog[string, gcm.TypeLike]()
	v.functionals_ = gcf.Catalog[string, gcm.FunctionalLike]()
	v.aspects_ = gcf.Catalog[string, gcm.AspectLike]()
	v.classes_ = gcf.Catalog[string, gcm.ClassLike]()
	v.instances_ = gcf.Catalog[string, gcm.InstanceLike]()
	v.abstractions_ = gcf.Catalog[string, gcm.AbstractionLike]()

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

func (v *validator_) extractAspects(model gcm.ModelLike) {
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

func (v *validator_) extractClasses(model gcm.ModelLike) {
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

func (v *validator_) extractFunctionals(model gcm.ModelLike) {
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

func (v *validator_) extractInstances(model gcm.ModelLike) {
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

func (v *validator_) extractModules(model gcm.ModelLike) {
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

func (v *validator_) extractTypes(model gcm.ModelLike) {
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

func (v *validator_) validateAbstraction(abstraction gcm.AbstractionLike) {
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

func (v *validator_) validateAbstractions(abstractions col.ListLike[gcm.AbstractionLike]) {
	var iterator = abstractions.GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.validateAbstraction(abstraction)
	}
}

func (v *validator_) validateArguments(arguments col.ListLike[gcm.AbstractionLike]) {
	var iterator = arguments.GetIterator()
	for iterator.HasNext() {
		var argument = iterator.GetNext()
		v.validateAbstraction(argument)
	}
}

func (v *validator_) validateAspect(aspect gcm.AspectLike) {
	var declaration = aspect.GetDeclaration()
	v.validateDeclaration(declaration)
	var methods = aspect.GetMethods()
	v.validateMethods(methods)
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

func (v *validator_) validateAspects(model gcm.ModelLike) {
	v.aspects_.SortValues()
	var iterator = v.aspects_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var aspect = association.GetValue()
		v.validateAspect(aspect)
	}
	var aspects = model.GetAspects()
	if aspects != nil {
		aspects.RemoveAll()
		aspects.AppendValues(v.aspects_.GetValues(v.aspects_.GetKeys()))
	}
}

func (v *validator_) validateAttribute(attribute gcm.AttributeLike) {
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

func (v *validator_) validateAttributes(attributes col.ListLike[gcm.AttributeLike]) {
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		v.validateAttribute(attribute)
	}
}

func (v *validator_) validateBoolean(abstraction gcm.AbstractionLike) {
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

func (v *validator_) validateClass(class gcm.ClassLike) {
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

func (v *validator_) validateClasses(model gcm.ModelLike) {
	v.classes_.SortValues()
	var iterator = v.classes_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var class = association.GetValue()
		v.validateClass(class)
	}
	var classes = model.GetClasses()
	if classes != nil {
		classes.RemoveAll()
		classes.AppendValues(v.classes_.GetValues(v.classes_.GetKeys()))
	}
}

func (v *validator_) validateConstant(constant gcm.ConstantLike) {
	var abstraction = constant.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstants(constants col.ListLike[gcm.ConstantLike]) {
	var iterator = constants.GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.validateConstant(constant)
	}
}

func (v *validator_) validateConstructor(constructor gcm.ConstructorLike) {
	var parameters = constructor.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var abstraction = constructor.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstructors(constructors col.ListLike[gcm.ConstructorLike]) {
	var iterator = constructors.GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.validateConstructor(constructor)
	}
}

func (v *validator_) validateDeclaration(declaration gcm.DeclarationLike) {
	var parameters = declaration.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
}

func (v *validator_) validateEnumeration(enumeration gcm.EnumerationLike) {
	var parameter = enumeration.GetParameter()
	v.validateParameter(parameter)
}

func (v *validator_) validateFunction(function gcm.FunctionLike) {
	var parameters = function.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = function.GetResult()
	v.validateResult(result)
}

func (v *validator_) validateFunctional(functional gcm.FunctionalLike) {
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

func (v *validator_) validateFunctionals(model gcm.ModelLike) {
	v.functionals_.SortValues()
	var iterator = v.functionals_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var functional = association.GetValue()
		v.validateFunctional(functional)
	}
	var functionals = model.GetFunctionals()
	if functionals != nil {
		functionals.RemoveAll()
		functionals.AppendValues(v.functionals_.GetValues(v.functionals_.GetKeys()))
	}
}

func (v *validator_) validateFunctions(functions col.ListLike[gcm.FunctionLike]) {
	var iterator = functions.GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.validateFunction(function)
	}
}

func (v *validator_) validateGetClassMethod(class string) {
	var instance = v.instances_.GetValue(class)
	var attributes = instance.GetAttributes()
	if attributes != nil {
		var iterator = attributes.GetIterator()
		for iterator.HasNext() {
			var attribute = iterator.GetNext()
			var identifier = attribute.GetIdentifier()
			if identifier == "GetClass" {
				var abstraction = attribute.GetAbstraction()
				if class+"classlike" == sts.ToLower(abstraction.GetIdentifier()) {
					return
				}
			}
		}
	}
	fmt.Printf(
		"The following class is missing a GetClass() instance method: %v\n",
		sts.TrimSuffix(instance.GetDeclaration().GetIdentifier(), "Like"),
	)
}

func (v *validator_) validateInstance(instance gcm.InstanceLike) {
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

func (v *validator_) validateInstances(model gcm.ModelLike) {
	v.instances_.SortValues()
	var iterator = v.instances_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var instance = association.GetValue()
		v.validateInstance(instance)
	}
	var instances = model.GetInstances()
	if instances != nil {
		instances.RemoveAll()
		instances.AppendValues(v.instances_.GetValues(v.instances_.GetKeys()))
	}
}

func (v *validator_) validateMethod(method gcm.MethodLike) {
	var parameters = method.GetParameters()
	if parameters != nil {
		v.validateParameters(parameters)
	}
	var result = method.GetResult()
	if result != nil {
		v.validateResult(result)
	}
}

func (v *validator_) validateMethods(methods col.ListLike[gcm.MethodLike]) {
	var iterator = methods.GetIterator()
	for iterator.HasNext() {
		var method = iterator.GetNext()
		v.validateMethod(method)
	}
}

func (v *validator_) validateModule(module gcm.ModuleLike) {
	var identifier = module.GetIdentifier()
	if len(identifier) != 3 {
		var message = fmt.Sprintf(
			"The length of the identifier for an imported module must be 3: %v",
			identifier,
		)
		panic(message)
	}
}

func (v *validator_) validateModules(model gcm.ModelLike) {
	v.modules_.SortValues()
	var iterator = v.modules_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var module = association.GetValue()
		v.validateModule(module)
	}
	var modules = model.GetModules()
	if modules != nil {
		modules.RemoveAll()
		modules.AppendValues(v.modules_.GetValues(v.modules_.GetKeys()))
	}
}

func (v *validator_) validatePairings(model gcm.ModelLike) {
	// Make sure each class interface has an associated instance interface.
	var classes = gcf.List[string](v.classes_.GetKeys())
	var instances = gcf.List[string](v.instances_.GetKeys())
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

func (v *validator_) validateParameter(parameter gcm.ParameterLike) {
	var abstraction = parameter.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateParameters(parameters col.ListLike[gcm.ParameterLike]) {
	var iterator = parameters.GetIterator()
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		v.validateParameter(parameter)
	}
}

func (v *validator_) validatePrefix(prefix gcm.PrefixLike) {
	if prefix.GetType() == gcm.AliasPrefix {
		var identifier = prefix.GetIdentifier()
		if v.modules_.GetValue(identifier) == nil {
			var message = fmt.Sprintf(
				"Unknown module alias: %v",
				identifier,
			)
			panic(message)
		}
	}
}

func (v *validator_) validateResult(result gcm.ResultLike) {
	if result != nil {
		var abstraction = result.GetAbstraction()
		if abstraction != nil {
			v.validateAbstraction(abstraction)
		} else {
			var parameters = result.GetParameters()
			v.validateParameters(parameters)
		}
	}
}

func (v *validator_) validateType(type_ gcm.TypeLike) {
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

func (v *validator_) validateTypes(model gcm.ModelLike) {
	v.types_.SortValues()
	var iterator = v.types_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var type_ = association.GetValue()
		v.validateType(type_)
	}
	var types = model.GetTypes()
	if types != nil {
		types.RemoveAll()
		types.AppendValues(v.types_.GetValues(v.types_.GetKeys()))
	}
}
