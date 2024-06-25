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
	ast "github.com/craterdog/go-model-framework/v4/ast"
	reg "regexp"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// Initialize the class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	// Define the instance attributes.
	class_  FormatterClassLike
	depth_  int
	result_ sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() int {
	return v.depth_
}

// Public

func (v *formatter_) FormatAbstraction(abstraction ast.AbstractionLike) string {
	v.formatAbstraction(abstraction)
	return v.getResult()
}

func (v *formatter_) FormatArguments(arguments ast.ArgumentsLike) string {
	v.formatArguments(arguments)
	return v.getResult()
}

func (v *formatter_) FormatMethod(method ast.MethodLike) string {
	v.formatMethod(method)
	return v.getResult()
}

func (v *formatter_) FormatModel(model ast.ModelLike) string {
	v.formatModel(model)
	return v.getResult()
}

func (v *formatter_) FormatParameter(parameter ast.ParameterLike) string {
	v.formatParameter(parameter)
	return v.getResult()
}

func (v *formatter_) FormatParameters(parameters ast.ParametersLike) string {
	v.formatParameters(parameters)
	return v.getResult()
}

func (v *formatter_) FormatResult(result ast.ResultLike) string {
	v.formatResult(result)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
	for level := 0; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) fixComment(comment string) string {
	var matcher = reg.MustCompile("\n    ")
	comment = matcher.ReplaceAllString(comment, "\n\t")
	return comment
}

func (v *formatter_) formatAbstraction(abstraction ast.AbstractionLike) {
	var prefix = abstraction.GetPrefix()
	if prefix != nil {
		v.formatPrefix(prefix)
	}
	var identifier = abstraction.GetIdentifier()
	v.appendString(identifier)
	var genericArguments = abstraction.GetGenericArguments()
	if genericArguments != nil {
		v.formatGenericArguments(genericArguments)
	}
}

func (v *formatter_) formatAbstractions(abstractions ast.AbstractionsLike) {
	v.appendNewline()
	v.appendString("// Abstractions")
	var iterator = abstractions.GetAbstractions().GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.formatAbstraction(abstraction)
	}
}

func (v *formatter_) formatAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
) {
	v.appendNewline()
	var argument = additionalArgument.GetArgument()
	v.formatArgument(argument)
	v.appendString(",")
}

func (v *formatter_) formatAdditionalArguments(
	additionalArguments ast.AdditionalArgumentsLike,
) {
	v.appendString(",")
	var iterator = additionalArguments.GetAdditionalArguments().GetIterator()
	for iterator.HasNext() {
		var additionalArgument = iterator.GetNext()
		v.formatAdditionalArgument(additionalArgument)
	}
	v.appendNewline()
}

func (v *formatter_) formatAdditionalParameter(
	additionalParameter ast.AdditionalParameterLike,
) {
	v.appendNewline()
	var parameter = additionalParameter.GetParameter()
	v.formatParameter(parameter)
	v.appendString(",")
}

func (v *formatter_) formatAdditionalParameters(
	additionalParameters ast.AdditionalParametersLike,
) {
	v.appendString(",")
	var iterator = additionalParameters.GetAdditionalParameters().GetIterator()
	for iterator.HasNext() {
		var additionalParameter = iterator.GetNext()
		v.formatAdditionalParameter(additionalParameter)
	}
	v.appendNewline()
}

func (v *formatter_) formatAdditionalValue(
	additionalValue ast.AdditionalValueLike,
) {
	v.appendNewline()
	var identifier = additionalValue.GetIdentifier()
	v.appendString(identifier)
}

func (v *formatter_) formatAlias(alias ast.AliasLike) {
	var identifier = alias.GetIdentifier()
	v.appendString(identifier)
	v.appendString(".")
}

func (v *formatter_) formatArgument(argument ast.ArgumentLike) {
	var abstraction = argument.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatArguments(arguments ast.ArgumentsLike) {
	var additionalArguments = arguments.GetAdditionalArguments()
	if additionalArguments != nil {
		v.depth_++
		v.appendNewline()
	}
	var argument = arguments.GetArgument()
	v.formatArgument(argument)
	if additionalArguments != nil {
		v.formatAdditionalArguments(additionalArguments)
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) formatArray(array ast.ArrayLike) {
	v.appendString("[]")
}

func (v *formatter_) formatAspect(aspect ast.AspectLike) {
	v.appendNewline()
	var declaration = aspect.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	v.appendNewline()
	var methods = aspect.GetMethods()
	v.formatMethods(methods)
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) formatAspects(aspects ast.AspectsLike) {
	v.appendNewline()
	v.appendString("// Aspects")
	v.appendNewline()
	var iterator = aspects.GetAspects().GetIterator()
	for iterator.HasNext() {
		var aspect = iterator.GetNext()
		v.formatAspect(aspect)
	}
	v.appendNewline()
}

func (v *formatter_) formatAttribute(attribute ast.AttributeLike) {
	var identifier = attribute.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameter = attribute.GetParameter()
	if parameter != nil {
		v.formatParameter(parameter)
	}
	v.appendString(")")
	var abstraction = attribute.GetAbstraction()
	if abstraction != nil {
		v.appendString(" ")
		v.formatAbstraction(abstraction)
	}
}

func (v *formatter_) formatAttributes(attributes ast.AttributesLike) {
	v.appendNewline()
	v.appendString("// Attributes")
	var iterator = attributes.GetAttributes().GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		v.formatAttribute(attribute)
	}
}

func (v *formatter_) formatChannel(channel ast.ChannelLike) {
	v.appendString("chan ")
}

func (v *formatter_) formatClass(class ast.ClassLike) {
	v.appendNewline()
	var declaration = class.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	v.appendNewline()
	var constants = class.GetConstants()
	if constants != nil {
		v.formatConstants(constants)
	}
	var constructors = class.GetConstructors()
	if constructors != nil {
		v.formatConstructors(constructors)
	}
	var functions = class.GetFunctions()
	if functions != nil {
		v.formatFunctions(functions)
	}
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) formatClasses(classes ast.ClassesLike) {
	v.appendNewline()
	v.appendString("// Classes")
	v.appendNewline()
	var iterator = classes.GetClasses().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		v.formatClass(class)
	}
	v.appendNewline()
}

func (v *formatter_) formatConstant(constant ast.ConstantLike) {
	v.appendNewline()
	var identifier = constant.GetIdentifier()
	v.appendString(identifier)
	v.appendString("() ")
	var abstraction = constant.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatConstants(constants ast.ConstantsLike) {
	v.appendNewline()
	v.appendString("// Constants")
	var iterator = constants.GetConstants().GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.formatConstant(constant)
	}
}

func (v *formatter_) formatConstructor(constructor ast.ConstructorLike) {
	v.appendNewline()
	var identifier = constructor.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameters = constructor.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(") ")
	var abstraction = constructor.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatConstructors(constructors ast.ConstructorsLike) {
	v.appendNewline()
	v.appendString("// Constructors")
	var iterator = constructors.GetConstructors().GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.formatConstructor(constructor)
	}
}

func (v *formatter_) formatDeclaration(declaration ast.DeclarationLike) {
	var comment = declaration.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
	v.appendString("type ")
	var identifier = declaration.GetIdentifier()
	v.appendString(identifier)
	var genericParameters = declaration.GetGenericParameters()
	if genericParameters != nil {
		v.formatGenericParameters(genericParameters)
	}
}

func (v *formatter_) formatEnumeration(enumeration ast.EnumerationLike) {
	v.appendNewline()
	v.appendString("const (")
	v.depth_++
	v.appendNewline()
	var value = enumeration.GetValue()
	v.formatValue(value)
	var iterator = enumeration.GetAdditionalValues().GetIterator()
	for iterator.HasNext() {
		var additionalValue = iterator.GetNext()
		v.formatAdditionalValue(additionalValue)
	}
	v.depth_--
	v.appendNewline()
	v.appendString(")")
}

func (v *formatter_) formatFunction(function ast.FunctionLike) {
	v.appendNewline()
	var identifier = function.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameters = function.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(") ")
	var result = function.GetResult()
	v.formatResult(result)
}

func (v *formatter_) formatFunctional(functional ast.FunctionalLike) {
	v.appendNewline()
	var declaration = functional.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" func(")
	var parameters = functional.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(") ")
	var result = functional.GetResult()
	v.formatResult(result)
	v.appendNewline()
}

func (v *formatter_) formatFunctionals(functionals ast.FunctionalsLike) {
	v.appendNewline()
	v.appendString("// Functionals")
	v.appendNewline()
	var iterator = functionals.GetFunctionals().GetIterator()
	for iterator.HasNext() {
		var functional = iterator.GetNext()
		v.formatFunctional(functional)
	}
	v.appendNewline()
}

func (v *formatter_) formatFunctions(functions ast.FunctionsLike) {
	v.appendNewline()
	v.appendString("// Functions")
	var iterator = functions.GetFunctions().GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.formatFunction(function)
	}
}

func (v *formatter_) formatGenericArguments(
	genericArguments ast.GenericArgumentsLike,
) {
	v.appendString("[")
	var arguments = genericArguments.GetArguments()
	v.formatArguments(arguments)
	v.appendString("]")
}

func (v *formatter_) formatGenericParameters(
	genericParameters ast.GenericParametersLike,
) {
	v.appendString("[")
	var parameters = genericParameters.GetParameters()
	v.formatParameters(parameters)
	v.appendString("]")
}

func (v *formatter_) formatHeader(header ast.HeaderLike) {
	v.appendNewline()
	var comment = header.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
	v.appendString("package ")
	var identifier = header.GetIdentifier()
	v.appendString(identifier)
	v.appendNewline()
}

func (v *formatter_) formatInstance(instance ast.InstanceLike) {
	v.appendNewline()
	var declaration = instance.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	v.appendNewline()
	var attributes = instance.GetAttributes()
	if attributes != nil {
		v.formatAttributes(attributes)
	}
	var abstractions = instance.GetAbstractions()
	if abstractions != nil {
		v.formatAbstractions(abstractions)
	}
	var methods = instance.GetMethods()
	if methods != nil {
		v.formatMethods(methods)
	}
	v.depth_--
	v.appendNewline()
	v.appendString("}")
	v.appendNewline()
}

func (v *formatter_) formatInstances(instances ast.InstancesLike) {
	v.appendNewline()
	v.appendString("// Instances")
	v.appendNewline()
	var iterator = instances.GetInstances().GetIterator()
	for iterator.HasNext() {
		var instance = iterator.GetNext()
		v.formatInstance(instance)
	}
	v.appendNewline()
}

func (v *formatter_) formatMap(map_ ast.MapLike) {
	var identifier = map_.GetIdentifier()
	v.appendString("map[")
	v.appendString(identifier)
	v.appendString("]")
}

func (v *formatter_) formatMethod(method ast.MethodLike) {
	var identifier = method.GetIdentifier()
	v.appendString(identifier)
	v.appendString("(")
	var parameters = method.GetParameters()
	if parameters != nil {
		v.formatParameters(parameters)
	}
	v.appendString(")")
	var result = method.GetResult()
	if result != nil {
		v.appendString(" ")
		v.formatResult(result)
	}
}

func (v *formatter_) formatMethods(methods ast.MethodsLike) {
	v.appendNewline()
	v.appendString("// Methods")
	var iterator = methods.GetMethods().GetIterator()
	for iterator.HasNext() {
		var method = iterator.GetNext()
		v.formatMethod(method)
	}
}

func (v *formatter_) formatModel(model ast.ModelLike) {
	var notice = model.GetNotice()
	v.formatNotice(notice)
	var header = model.GetHeader()
	v.formatHeader(header)
	var modules = model.GetModules()
	if modules != nil {
		v.formatModules(modules)
	}
	var types = model.GetTypes()
	if types != nil {
		v.formatTypes(types)
	}
	var functionals = model.GetFunctionals()
	if functionals != nil {
		v.formatFunctionals(functionals)
	}
	var aspects = model.GetAspects()
	if aspects != nil {
		v.formatAspects(aspects)
	}
	var classes = model.GetClasses()
	if classes != nil {
		v.formatClasses(classes)
	}
	var instances = model.GetInstances()
	if instances != nil {
		v.formatInstances(instances)
	}
}

func (v *formatter_) formatModule(module ast.ModuleLike) {
	v.appendNewline()
	var identifier = module.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var text = module.GetText()
	v.appendString(text)
}

func (v *formatter_) formatModules(modules ast.ModulesLike) {
	v.appendNewline()
	v.appendString("import (")
	v.depth_++
	var iterator = modules.GetModules().GetIterator()
	for iterator.HasNext() {
		var module = iterator.GetNext()
		v.formatModule(module)
	}
	v.depth_--
	v.appendNewline()
	v.appendString(")")
	v.appendNewline()
}

func (v *formatter_) formatNotice(notice ast.NoticeLike) {
	var comment = notice.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
}

func (v *formatter_) formatParameter(parameter ast.ParameterLike) {
	var identifier = parameter.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var abstraction = parameter.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatParameterized(result ast.ParameterizedLike) {
	v.appendString("(")
	var parameters = result.GetParameters()
	v.formatParameters(parameters)
	v.appendString(")")
}

func (v *formatter_) formatParameters(parameters ast.ParametersLike) {
	var additionalParameters = parameters.GetAdditionalParameters()
	if additionalParameters != nil {
		v.depth_++
		v.appendNewline()
	}
	var parameter = parameters.GetParameter()
	v.formatParameter(parameter)
	if additionalParameters != nil {
		v.formatAdditionalParameters(additionalParameters)
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) formatPrefix(prefix ast.PrefixLike) {
	switch actual := prefix.GetAny().(type) {
	case ast.ArrayLike:
		v.formatArray(actual)
	case ast.MapLike:
		v.formatMap(actual)
	case ast.ChannelLike:
		v.formatChannel(actual)
	case ast.AliasLike:
		v.formatAlias(actual)
	default:
		var message = fmt.Sprintf(
			"Attempted to format an unknown prefix type: %T",
			actual,
		)
		panic(message)
	}
}

func (v *formatter_) formatResult(result ast.ResultLike) {
	switch actual := result.GetAny().(type) {
	case ast.AbstractionLike:
		v.formatAbstraction(actual)
	case ast.ParameterizedLike:
		v.formatParameterized(actual)
	default:
		var message = fmt.Sprintf(
			"Attempted to format an unknown result type: %T",
			actual,
		)
		panic(message)
	}
}

func (v *formatter_) formatType(type_ ast.TypeLike) {
	v.appendNewline()
	var declaration = type_.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" ")
	var abstraction = type_.GetAbstraction()
	v.formatAbstraction(abstraction)
	var enumeration = type_.GetEnumeration()
	if enumeration != nil {
		v.formatEnumeration(enumeration)
	}
	v.appendNewline()
}

func (v *formatter_) formatTypes(types ast.TypesLike) {
	v.appendNewline()
	v.appendString("// Types")
	v.appendNewline()
	var iterator = types.GetTypes().GetIterator()
	for iterator.HasNext() {
		var type_ = iterator.GetNext()
		v.formatType(type_)
	}
	v.appendNewline()
}

func (v *formatter_) formatValue(value ast.ValueLike) {
	v.appendNewline()
	var identifier = value.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var abstraction = value.GetAbstraction()
	v.formatAbstraction(abstraction)
	v.appendString(" = iota")
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
