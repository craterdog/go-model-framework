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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	sts "strings"
)

// CLASS ACCESS

// Reference

var validatorClass = &validatorClass_{
	// Initialize the class constants.
}

// Function

func Validator() ValidatorClassLike {
	return validatorClass
}

// CLASS METHODS

// Target

type validatorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *validatorClass_) Make() ValidatorLike {
	return &validator_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type validator_ struct {
	// Define the instance attributes.
	class_        ValidatorClassLike
	imports_      abs.CatalogLike[string, ast.ModuleLike]
	types_        abs.CatalogLike[string, ast.TypeLike]
	functionals_  abs.CatalogLike[string, ast.FunctionalLike]
	aspects_      abs.CatalogLike[string, ast.AspectLike]
	classes_      abs.CatalogLike[string, ast.ClassLike]
	instances_    abs.CatalogLike[string, ast.InstanceLike]
	abstractions_ abs.CatalogLike[string, ast.AbstractionLike]
}

// Attributes

func (v *validator_) GetClass() ValidatorClassLike {
	return v.class_
}

// Public

func (v *validator_) ValidateModel(model ast.ModelLike) {
	// Initialize the catalogs.
	v.imports_ = col.Catalog[string, ast.ModuleLike]()
	v.types_ = col.Catalog[string, ast.TypeLike]()
	v.functionals_ = col.Catalog[string, ast.FunctionalLike]()
	v.classes_ = col.Catalog[string, ast.ClassLike]()
	v.instances_ = col.Catalog[string, ast.InstanceLike]()
	v.aspects_ = col.Catalog[string, ast.AspectLike]()
	v.abstractions_ = col.Catalog[string, ast.AbstractionLike]()

	// Extract the catalogs.
	v.extractModules(model)
	v.extractTypes(model)
	v.extractFunctionals(model)
	v.extractClasses(model)
	v.extractInstances(model)
	v.extractAspects(model)

	// Validate the catalogs (note, the order matters).
	v.validateModules(model)
	v.validateClasses(model)
	v.validateInstances(model)
	v.validateAspects(model)
	v.validateTypes(model)
	v.validateFunctionals(model)
}

// Private

func (v *validator_) extractAspects(model ast.ModelLike) {
	var aspects = model.GetOptionalAspects()
	if col.IsDefined(aspects) {
		var iterator = aspects.GetAspects().GetIterator()
		for iterator.HasNext() {
			var aspect = iterator.GetNext()
			var name = aspect.GetDeclaration().GetName()
			v.aspects_.SetValue(name, aspect)
		}
	}
}

func (v *validator_) extractClasses(model ast.ModelLike) {
	var classes = model.GetClasses()
	if col.IsDefined(classes) {
		var iterator = classes.GetClasses().GetIterator()
		for iterator.HasNext() {
			var class = iterator.GetNext()
			var name = class.GetDeclaration().GetName()
			name = sts.TrimSuffix(name, "ClassLike")
			v.classes_.SetValue(name, class)
		}
	}
}

func (v *validator_) extractFunctionals(model ast.ModelLike) {
	var functionals = model.GetOptionalFunctionals()
	if col.IsDefined(functionals) {
		var iterator = functionals.GetFunctionals().GetIterator()
		for iterator.HasNext() {
			var functional = iterator.GetNext()
			var name = functional.GetDeclaration().GetName()
			v.functionals_.SetValue(name, functional)
		}
	}
}

func (v *validator_) extractInstances(model ast.ModelLike) {
	var instances = model.GetInstances()
	if col.IsDefined(instances) {
		var iterator = instances.GetInstances().GetIterator()
		for iterator.HasNext() {
			var instance = iterator.GetNext()
			var name = instance.GetDeclaration().GetName()
			name = sts.TrimSuffix(name, "Like")
			v.instances_.SetValue(name, instance)
		}
	}
}

func (v *validator_) extractModules(model ast.ModelLike) {
	var imports = model.GetOptionalImports()
	if col.IsDefined(imports) {
		var iterator = imports.GetModules().GetModules().GetIterator()
		for iterator.HasNext() {
			var module = iterator.GetNext()
			// The modules are alphabetized by path, not by name.
			var path = module.GetPath()
			v.imports_.SetValue(path, module)
		}
	}
}

func (v *validator_) extractTypes(model ast.ModelLike) {
	var types = model.GetOptionalTypes()
	if col.IsDefined(types) {
		var iterator = types.GetTypes().GetIterator()
		for iterator.HasNext() {
			var type_ = iterator.GetNext()
			var name = type_.GetDeclaration().GetName()
			v.types_.SetValue(name, type_)
		}
	}
}

func (v *validator_) validateAbstraction(abstraction ast.AbstractionLike) {
	// Validate the optional prefix.
	var prefix = abstraction.GetOptionalPrefix()
	if col.IsDefined(prefix) {
		v.validatePrefix(prefix)
	}

	// Validate the optional alias.
	var alias = abstraction.GetOptionalAlias()
	if col.IsDefined(alias) {
		v.validateAlias(alias)
	}

	// Record the name of the abstraction.
	var name = abstraction.GetName()
	v.abstractions_.SetValue(name, abstraction)

	// Validate any generic arguments.
	var genericArguments = abstraction.GetOptionalGenericArguments()
	if col.IsDefined(genericArguments) {
		var arguments = genericArguments.GetArguments()
		v.validateArguments(arguments)
	}
}

func (v *validator_) validateAbstractions(abstractions ast.AbstractionsLike) {
	var iterator = abstractions.GetAbstractions().GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.validateAbstraction(abstraction)
	}
}

func (v *validator_) validateAlias(alias ast.AliasLike) {
	var name = alias.GetName()
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
		"Found an unknown module alias name: %v",
		name,
	)
	panic(message)
}

func (v *validator_) validateArgument(argument ast.ArgumentLike) {
	var abstraction = argument.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateArguments(arguments ast.ArgumentsLike) {
	// Validate the first argument.
	var argument = arguments.GetArgument()
	v.validateArgument(argument)

	// Validate any additional arguments.
	var additionalArguments = arguments.GetAdditionalArguments()
	var iterator = additionalArguments.GetIterator()
	for iterator.HasNext() {
		var additionalArgument = iterator.GetNext()
		argument = additionalArgument.GetArgument()
		v.validateArgument(argument)
	}
}

func (v *validator_) validateAspect(aspect ast.AspectLike) {
	// Validate the declaration.
	var declaration = aspect.GetDeclaration()
	v.validateDeclaration(declaration)

	// Validate the methods.
	var methods = aspect.GetMethods()
	v.validateMethods(methods)

	// Verify that this aspect is actually used in this model.
	var name = declaration.GetName()
	var abstraction = v.abstractions_.GetValue(name)
	if col.IsUndefined(abstraction) {
		var message = fmt.Sprintf(
			"The following aspect is never used in this model: %v",
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
	var parameter = attribute.GetOptionalParameter()
	var result = attribute.GetOptionalAbstraction()
	switch {
	case sts.HasPrefix(name, "Set"):
		// An attribute "setter" has only a parameter.
		v.validateParameter(parameter)
	case sts.HasPrefix(name, "Get"):
		// An attribute "getter" has only a result.
		v.validateAbstraction(result)
	default:
		// Otherwise, assume it is an attribute "to be" question.
		v.validateBoolean(result)
	}
}

func (v *validator_) validateAttributes(
	classname string,
	attributes ast.AttributesLike,
) {
	var attribute ast.AttributeLike
	var iterator = attributes.GetAttributes().GetIterator()
	for iterator.HasNext() {
		attribute = iterator.GetNext()
		v.validateAttribute(attribute)
	}

	// Validate the GetClass() method.
	iterator.ToStart()
	attribute = iterator.GetNext()
	var name = attribute.GetName()
	if name != "GetClass" {
		var message = fmt.Sprintf(
			"The %v class is missing a GetClass() method.",
			classname,
		)
		panic(message)
	}
	var abstraction = attribute.GetOptionalAbstraction()
	if classname+"ClassLike" != abstraction.GetName() {
		var message = fmt.Sprintf(
			"The GetClass() method for the %v class has the wrong result type.",
			classname,
		)
		panic(message)
	}
}

func (v *validator_) validateBoolean(abstraction ast.AbstractionLike) {
	var prefix = abstraction.GetOptionalPrefix()
	if col.IsDefined(prefix) {
		panic("A boolean type does not have a prefix.")
	}
	var alias = abstraction.GetOptionalAlias()
	if col.IsDefined(alias) {
		panic("A boolean type does not have an alias.")
	}
	var name = abstraction.GetName()
	if name != "bool" {
		panic("A question attribute must have a boolean type.")
	}
	var arguments = abstraction.GetOptionalGenericArguments()
	if col.IsDefined(arguments) {
		panic("A boolean type is not a generic type.")
	}
}

func (v *validator_) validateClass(class ast.ClassLike) {
	// Validate the declaration.
	var declaration = class.GetDeclaration()
	v.validateDeclaration(declaration)

	// Validate the constructors.
	var constructors = class.GetConstructors()
	if col.IsDefined(constructors) {
		v.validateConstructors(constructors)
	}

	// Validate the constants.
	var constants = class.GetOptionalConstants()
	if col.IsDefined(constants) {
		v.validateConstants(constants)
	}

	// Validate the functions.
	var functions = class.GetOptionalFunctions()
	if col.IsDefined(functions) {
		v.validateFunctions(functions)
	}
}

func (v *validator_) validateClasses(model ast.ModelLike) {
	var iterator = v.classes_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var name = association.GetKey()
		var instance = v.instances_.GetValue(name)
		if col.IsUndefined(instance) {
			var message = fmt.Sprintf(
				"The following instance interface is missing: %v",
				name,
			)
			panic(message)
		}
		var class = association.GetValue()
		v.validateClass(class)
	}
}

func (v *validator_) validateConstant(constant ast.ConstantLike) {
	var abstraction = constant.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstants(constants ast.ConstantsLike) {
	var iterator = constants.GetConstants().GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.validateConstant(constant)
	}
}

func (v *validator_) validateConstructor(constructor ast.ConstructorLike) {
	// Validate any parameters.
	var parameters = constructor.GetOptionalParameters()
	if col.IsDefined(parameters) {
		v.validateParameters(parameters)
	}

	// Validate the return type.
	var abstraction = constructor.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateConstructors(constructors ast.ConstructorsLike) {
	var iterator = constructors.GetConstructors().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.validateConstructor(constructor)
	}
}

func (v *validator_) validateDeclaration(declaration ast.DeclarationLike) {
	// Validate any generic parameters.
	var genericParameters = declaration.GetOptionalGenericParameters()
	if col.IsDefined(genericParameters) {
		var parameters = genericParameters.GetParameters()
		v.validateParameters(parameters)
	}
}

func (v *validator_) validateEnumeration(enumeration ast.EnumerationLike) {
	// Validate the enumerated type, not the values.
	var values = enumeration.GetValues()
	var value = values.GetValue()
	var abstraction = value.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateFunction(function ast.FunctionLike) {
	// Validate any parameters.
	var parameters = function.GetOptionalParameters()
	if col.IsDefined(parameters) {
		v.validateParameters(parameters)
	}

	// Validate the result type.
	var result = function.GetResult()
	v.validateResult(result)
}

func (v *validator_) validateFunctional(functional ast.FunctionalLike) {
	// Validate the declaration.
	var declaration = functional.GetDeclaration()
	v.validateDeclaration(declaration)

	// Validate any parameters.
	var parameters = functional.GetOptionalParameters()
	if col.IsDefined(parameters) {
		v.validateParameters(parameters)
	}

	// Validate the result type.
	var result = functional.GetResult()
	v.validateResult(result)

	// Verify that this functional is actually used in this model.
	var name = declaration.GetName()
	var abstraction = v.abstractions_.GetValue(name)
	if col.IsUndefined(abstraction) {
		var message = fmt.Sprintf(
			"The following functional is never used in this model: %v",
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
	var iterator = functions.GetFunctions().GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.validateFunction(function)
	}
}

func (v *validator_) validateInstance(instance ast.InstanceLike) {
	// Validate the instance declaration.
	var declaration = instance.GetDeclaration()
	v.validateDeclaration(declaration)

	// Validate the instance attribute methods.
	var attributes = instance.GetAttributes()
	var classname = sts.TrimSuffix(declaration.GetName(), "Like")
	v.validateAttributes(classname, attributes)

	// Validate the instance abstraction methods.
	var abstractions = instance.GetOptionalAbstractions()
	if col.IsDefined(abstractions) {
		v.validateAbstractions(abstractions)
	}

	// Validate the instance public methods.
	var methods = instance.GetOptionalMethods()
	if col.IsDefined(methods) {
		v.validateMethods(methods.GetMethods())
	}
}

func (v *validator_) validateInstances(model ast.ModelLike) {
	var iterator = v.instances_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var name = association.GetKey()
		var class = v.classes_.GetValue(name)
		if col.IsUndefined(class) {
			var message = fmt.Sprintf(
				"The following class interface is missing: %v",
				name,
			)
			panic(message)
		}
		var instance = association.GetValue()
		v.validateInstance(instance)
	}
}

func (v *validator_) validateMethod(method ast.MethodLike) {
	// Validate any method parameters.
	var parameters = method.GetOptionalParameters()
	if col.IsDefined(parameters) {
		v.validateParameters(parameters)
	}

	// Validate any method results.
	var result = method.GetOptionalResult()
	if col.IsDefined(result) {
		v.validateResult(result)
	}
}

func (v *validator_) validateMethods(methods abs.Sequential[ast.MethodLike]) {
	var iterator = methods.GetIterator()
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

func (v *validator_) validateParameter(parameter ast.ParameterLike) {
	var abstraction = parameter.GetAbstraction()
	v.validateAbstraction(abstraction)
}

func (v *validator_) validateParameters(parameters ast.ParametersLike) {
	// Validate the first parameter.
	var parameter = parameters.GetParameter()
	v.validateParameter(parameter)

	// Validate any additional parameters.
	var additionalParameters = parameters.GetAdditionalParameters()
	if col.IsDefined(additionalParameters) {
		var iterator = additionalParameters.GetIterator()
		for iterator.HasNext() {
			var parameter = iterator.GetNext()
			v.validateParameter(parameter.GetParameter())
		}
	}
}

func (v *validator_) validatePrefix(prefix ast.PrefixLike) {
	// Nothing to validate.
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
	// Validate the declaration.
	var declaration = type_.GetDeclaration()
	v.validateDeclaration(declaration)

	// Validate the abstract type.
	var abstraction = type_.GetAbstraction()
	v.validateAbstraction(abstraction)

	// Validate any enumeration values.
	var enumeration = type_.GetOptionalEnumeration()
	if col.IsDefined(enumeration) {
		v.validateEnumeration(enumeration)
	}

	// Verify that this type is actually used in this model.
	var name = declaration.GetName()
	abstraction = v.abstractions_.GetValue(name)
	if col.IsUndefined(abstraction) {
		var message = fmt.Sprintf(
			"The following type is never used in this model: %v",
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
