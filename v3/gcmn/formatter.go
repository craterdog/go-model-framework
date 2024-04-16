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

package gcmn

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
	reg "regexp"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	// This class does not initialize any private class constants.
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	// This class does not define any private class constants.
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	depth_  int
	result_ sts.Builder
}

// Public

func (v *formatter_) FormatAbstraction(abstraction AbstractionLike) string {
	v.formatAbstraction(abstraction)
	return v.getResult()
}

func (v *formatter_) FormatArguments(arguments col.ListLike[AbstractionLike]) string {
	v.formatArguments(arguments)
	return v.getResult()
}

func (v *formatter_) FormatGenerics(parameters col.ListLike[ParameterLike]) string {
	v.formatGenerics(parameters)
	return v.getResult()
}

func (v *formatter_) FormatMethod(method MethodLike) string {
	v.formatMethod(method)
	return v.getResult()
}

func (v *formatter_) FormatModel(model ModelLike) string {
	v.formatModel(model)
	return v.getResult()
}

func (v *formatter_) FormatParameter(parameter ParameterLike) string {
	v.formatParameter(parameter)
	return v.getResult()
}

func (v *formatter_) FormatParameterNames(parameters col.ListLike[ParameterLike]) string {
	v.formatParameterNames(parameters)
	return v.getResult()
}

func (v *formatter_) FormatParameters(parameters col.ListLike[ParameterLike]) string {
	v.formatParameters(parameters)
	return v.getResult()
}

func (v *formatter_) FormatResult(result ResultLike) string {
	v.formatResult(result)
	return v.getResult()
}

// Private

func (v *formatter_) appendNewline() {
	var separator = "\n"
	var indentation = "\t"
	for level := 0; level < v.depth_; level++ {
		separator += indentation
	}
	v.appendString(separator)
}

func (v *formatter_) appendString(s string) {
	v.result_.WriteString(s)
}

func (v *formatter_) fixComment(comment string) string {
	var matcher = reg.MustCompile("\n    ")
	comment = matcher.ReplaceAllString(comment, "\n\t")
	return comment
}

func (v *formatter_) formatAbstraction(abstraction AbstractionLike) {
	var prefix = abstraction.GetPrefix()
	if prefix != nil {
		v.formatPrefix(prefix)
	}
	var identifier = abstraction.GetIdentifier()
	v.appendString(identifier)
	var arguments = abstraction.GetArguments()
	if arguments != nil {
		v.appendString("[")
		v.formatArguments(arguments)
		v.appendString("]")
	}
}

func (v *formatter_) formatAbstractions(abstractions col.ListLike[AbstractionLike]) {
	v.appendNewline()
	v.appendString("// Abstractions")
	var iterator = abstractions.GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.appendNewline()
		v.formatAbstraction(abstraction)
	}
}

func (v *formatter_) formatArguments(arguments col.ListLike[AbstractionLike]) {
	var size = arguments.GetSize()
	if size > 2 {
		v.depth_++
		v.appendNewline()
	}
	var iterator = arguments.GetIterator()
	var abstraction = iterator.GetNext()
	v.formatAbstraction(abstraction)
	for iterator.HasNext() {
		abstraction = iterator.GetNext()
		v.appendString(",")
		if size > 2 {
			v.appendNewline()
		} else {
			v.appendString(" ")
		}
		v.formatAbstraction(abstraction)
	}
	if size > 2 {
		v.appendString(",")
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) formatAspect(aspect AspectLike) {
	var declaration = aspect.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	var methods = aspect.GetMethods()
	if methods != nil {
		v.formatMethods(methods)
	}
	v.depth_--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatAspects(aspects col.ListLike[AspectLike]) {
	v.appendNewline()
	v.appendString("// Aspects")
	v.appendNewline()
	var iterator = aspects.GetIterator()
	for iterator.HasNext() {
		var aspect = iterator.GetNext()
		v.formatAspect(aspect)
		v.appendNewline()
	}
}

func (v *formatter_) formatAttribute(attribute AttributeLike) {
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

func (v *formatter_) formatAttributes(attributes col.ListLike[AttributeLike]) {
	v.appendNewline()
	v.appendString("// Attributes")
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		v.appendNewline()
		v.formatAttribute(attribute)
	}
}

func (v *formatter_) formatClass(class ClassLike) {
	var declaration = class.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	var hasContent bool
	var constants = class.GetConstants()
	if constants != nil && !constants.IsEmpty() {
		v.formatConstants(constants)
		hasContent = true
	}
	var constructors = class.GetConstructors()
	if constructors != nil && !constructors.IsEmpty() {
		if hasContent {
			v.appendString("\n")
		}
		v.formatConstructors(constructors)
		hasContent = true
	}
	var functions = class.GetFunctions()
	if functions != nil && !functions.IsEmpty() {
		if hasContent {
			v.appendString("\n")
		}
		v.formatFunctions(functions)
	}
	v.depth_--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatClasses(classes col.ListLike[ClassLike]) {
	v.appendNewline()
	v.appendString("// Classes")
	v.appendNewline()
	var iterator = classes.GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		v.formatClass(class)
		v.appendNewline()
	}
}

func (v *formatter_) formatConstant(constant ConstantLike) {
	var identifier = constant.GetIdentifier()
	v.appendString(identifier)
	v.appendString("() ")
	var abstraction = constant.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatConstants(constants col.ListLike[ConstantLike]) {
	v.appendNewline()
	v.appendString("// Constants")
	var iterator = constants.GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.appendNewline()
		v.formatConstant(constant)
	}
}

func (v *formatter_) formatConstructor(constructor ConstructorLike) {
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

func (v *formatter_) formatConstructors(constructors col.ListLike[ConstructorLike]) {
	v.appendNewline()
	v.appendString("// Constructors")
	var iterator = constructors.GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.appendNewline()
		v.formatConstructor(constructor)
	}
}

func (v *formatter_) formatDeclaration(declaration DeclarationLike) {
	v.appendNewline()
	var comment = declaration.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
	v.appendString("type ")
	var identifier = declaration.GetIdentifier()
	v.appendString(identifier)
	var parameters = declaration.GetParameters()
	if parameters != nil {
		v.appendString("[")
		v.formatGenerics(parameters)
		v.appendString("]")
	}
}

func (v *formatter_) formatEnumeration(enumeration EnumerationLike) {
	v.appendNewline()
	v.appendString("const (")
	v.depth_++
	v.appendNewline()
	var parameter = enumeration.GetParameter()
	v.formatParameter(parameter)
	v.appendString(" = iota")
	var iterator = enumeration.GetIdentifiers().GetIterator()
	for iterator.HasNext() {
		var identifier = iterator.GetNext()
		v.appendNewline()
		v.appendString(identifier)
	}
	v.depth_--
	v.appendNewline()
	v.appendString(")")
}

func (v *formatter_) formatFunction(function FunctionLike) {
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

func (v *formatter_) formatFunctional(functional FunctionalLike) {
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
}

func (v *formatter_) formatFunctionals(functionals col.ListLike[FunctionalLike]) {
	v.appendNewline()
	v.appendString("// Functionals")
	v.appendNewline()
	var iterator = functionals.GetIterator()
	for iterator.HasNext() {
		var functional = iterator.GetNext()
		v.formatFunctional(functional)
		v.appendNewline()
	}
}

func (v *formatter_) formatFunctions(functions col.ListLike[FunctionLike]) {
	v.appendNewline()
	v.appendString("// Functions")
	var iterator = functions.GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.appendNewline()
		v.formatFunction(function)
	}
}

func (v *formatter_) formatGenerics(parameters col.ListLike[ParameterLike]) {
	var iterator = parameters.GetIterator()
	var parameter = iterator.GetNext()
	v.formatParameter(parameter)
	for iterator.HasNext() {
		parameter = iterator.GetNext()
		v.appendString(", ")
		v.formatParameter(parameter)
	}
}

func (v *formatter_) formatHeader(header HeaderLike) {
	var comment = header.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
	v.appendString("package ")
	var identifier = header.GetIdentifier()
	v.appendString(identifier)
}

func (v *formatter_) formatInstance(instance InstanceLike) {
	var declaration = instance.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	var hasContent bool
	var attributes = instance.GetAttributes()
	if attributes != nil && !attributes.IsEmpty() {
		v.formatAttributes(attributes)
		hasContent = true
	}
	var abstractions = instance.GetAbstractions()
	if abstractions != nil && !abstractions.IsEmpty() {
		if hasContent {
			v.appendString("\n")
		}
		v.formatAbstractions(abstractions)
		hasContent = true
	}
	var methods = instance.GetMethods()
	if methods != nil && !methods.IsEmpty() {
		if hasContent {
			v.appendString("\n")
		}
		v.formatMethods(methods)
	}
	v.depth_--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatInstances(instances col.ListLike[InstanceLike]) {
	v.appendNewline()
	v.appendString("// Instances")
	v.appendNewline()
	var iterator = instances.GetIterator()
	for iterator.HasNext() {
		var instance = iterator.GetNext()
		v.formatInstance(instance)
		v.appendNewline()
	}
}

func (v *formatter_) formatMethod(method MethodLike) {
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

func (v *formatter_) formatMethods(methods col.ListLike[MethodLike]) {
	v.appendNewline()
	v.appendString("// Methods")
	var iterator = methods.GetIterator()
	for iterator.HasNext() {
		var method = iterator.GetNext()
		v.appendNewline()
		v.formatMethod(method)
	}
}

func (v *formatter_) formatModule(module ModuleLike) {
	var identifier = module.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var text = module.GetText()
	v.appendString(text)
}

func (v *formatter_) formatModules(modules col.ListLike[ModuleLike]) {
	v.appendNewline()
	v.appendString("import (")
	v.depth_++
	var iterator = modules.GetIterator()
	for iterator.HasNext() {
		var module = iterator.GetNext()
		v.appendNewline()
		v.formatModule(module)
	}
	v.depth_--
	v.appendNewline()
	v.appendString(")")
	v.appendNewline()
}

func (v *formatter_) formatNotice(notice NoticeLike) {
	var comment = notice.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
}

func (v *formatter_) formatModel(model ModelLike) {
	var notice = model.GetNotice()
	v.formatNotice(notice)
	var header = model.GetHeader()
	v.formatHeader(header)
	v.appendNewline()
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

func (v *formatter_) formatParameter(parameter ParameterLike) {
	var identifier = parameter.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var abstraction = parameter.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatParameterName(parameter ParameterLike) {
	var identifier = parameter.GetIdentifier()
	v.appendString(identifier)
}

func (v *formatter_) formatParameterNames(parameters col.ListLike[ParameterLike]) {
	var size = parameters.GetSize()
	if size > 2 {
		v.depth_++
		v.appendNewline()
	}
	var iterator = parameters.GetIterator()
	var parameter = iterator.GetNext()
	v.formatParameterName(parameter)
	for iterator.HasNext() {
		parameter = iterator.GetNext()
		v.appendString(",")
		if size > 2 {
			v.appendNewline()
		} else {
			v.appendString(" ")
		}
		v.formatParameterName(parameter)
	}
	if size > 2 {
		v.appendString(",")
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) formatParameters(parameters col.ListLike[ParameterLike]) {
	var size = parameters.GetSize()
	if size > 1 {
		v.depth_++
		v.appendNewline()
	}
	var iterator = parameters.GetIterator()
	var parameter = iterator.GetNext()
	v.formatParameter(parameter)
	for iterator.HasNext() {
		parameter = iterator.GetNext()
		v.appendString(",")
		if size > 1 {
			v.appendNewline()
		} else {
			v.appendString(" ")
		}
		v.formatParameter(parameter)
	}
	if size > 1 {
		v.appendString(",")
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) formatPrefix(prefix PrefixLike) {
	var identifier = prefix.GetIdentifier()
	switch prefix.GetType() {
	case AliasPrefix:
		v.appendString(identifier)
		v.appendString(".")
	case ArrayPrefix:
		v.appendString("[]")
	case ChannelPrefix:
		v.appendString("chan ")
	case MapPrefix:
		v.appendString("map[")
		v.appendString(identifier)
		v.appendString("]")
	}
}

func (v *formatter_) formatResult(result ResultLike) {
	var abstraction = result.GetAbstraction()
	if abstraction != nil {
		v.formatAbstraction(abstraction)
	} else {
		v.appendString("(")
		var parameters = result.GetParameters()
		v.formatParameters(parameters)
		v.appendString(")")
	}
}

func (v *formatter_) formatType(type_ TypeLike) {
	var declaration = type_.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" ")
	var abstraction = type_.GetAbstraction()
	v.formatAbstraction(abstraction)
	var enumeration = type_.GetEnumeration()
	if enumeration != nil {
		v.appendNewline()
		v.formatEnumeration(enumeration)
	}
}

func (v *formatter_) formatTypes(types col.ListLike[TypeLike]) {
	v.appendNewline()
	v.appendString("// Types")
	v.appendNewline()
	var iterator = types.GetIterator()
	for iterator.HasNext() {
		var type_ = iterator.GetNext()
		v.formatType(type_)
		v.appendNewline()
	}
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}
