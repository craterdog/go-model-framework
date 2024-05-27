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
	col "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	reg "regexp"
	sts "strings"
)

// CLASS ACCESS

// Reference

var formatterClass = &formatterClass_{
	defaultMaximum_: 8,
}

// Function

func Formatter() FormatterClassLike {
	return formatterClass
}

// CLASS METHODS

// Target

type formatterClass_ struct {
	defaultMaximum_ int
}

// Constants

func (c *formatterClass_) DefaultMaximum() int {
	return c.defaultMaximum_
}

// Constructors

func (c *formatterClass_) Make() FormatterLike {
	return &formatter_{
		maximum_: c.defaultMaximum_,
	}
}

func (c *formatterClass_) MakeWithMaximum(maximum int) FormatterLike {
	if maximum < 0 {
		maximum = c.defaultMaximum_
	}
	return &formatter_{
		class_:   c,
		maximum_: maximum,
	}
}

// INSTANCE METHODS

// Target

type formatter_ struct {
	class_   FormatterClassLike
	depth_   int
	maximum_ int
	result_  sts.Builder
}

// Attributes

func (v *formatter_) GetClass() FormatterClassLike {
	return v.class_
}

func (v *formatter_) GetDepth() int {
	return v.depth_
}

func (v *formatter_) GetMaximum() int {
	return v.maximum_
}

// Public

func (v *formatter_) FormatAbstraction(abstraction ast.AbstractionLike) string {
	v.formatAbstraction(abstraction)
	return v.getResult()
}

func (v *formatter_) FormatArguments(arguments col.ListLike[ast.AbstractionLike]) string {
	v.formatArguments(arguments)
	return v.getResult()
}

func (v *formatter_) FormatGenerics(parameters col.ListLike[ast.ParameterLike]) string {
	v.formatGenerics(parameters)
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

func (v *formatter_) FormatParameterNames(parameters col.ListLike[ast.ParameterLike]) string {
	v.formatParameterNames(parameters)
	return v.getResult()
}

func (v *formatter_) FormatParameters(parameters col.ListLike[ast.ParameterLike]) string {
	v.formatParameters(parameters)
	return v.getResult()
}

func (v *formatter_) FormatResult(result ast.ResultLike) string {
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

func (v *formatter_) formatAbstraction(abstraction ast.AbstractionLike) {
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

func (v *formatter_) formatAbstractions(abstractions col.ListLike[ast.AbstractionLike]) {
	v.appendNewline()
	v.appendString("// Abstractions")
	var iterator = abstractions.GetIterator()
	for iterator.HasNext() {
		var abstraction = iterator.GetNext()
		v.appendNewline()
		v.formatAbstraction(abstraction)
	}
}

func (v *formatter_) formatArguments(arguments col.ListLike[ast.AbstractionLike]) {
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

func (v *formatter_) formatAspect(aspect ast.AspectLike) {
	var declaration = aspect.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	var methods = aspect.GetMethods()
	v.formatMethods(methods)
	v.depth_--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatAspects(aspects col.ListLike[ast.AspectLike]) {
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

func (v *formatter_) formatAttributes(attributes col.ListLike[ast.AttributeLike]) {
	v.appendNewline()
	v.appendString("// Attributes")
	var iterator = attributes.GetIterator()
	for iterator.HasNext() {
		var attribute = iterator.GetNext()
		v.appendNewline()
		v.formatAttribute(attribute)
	}
}

func (v *formatter_) formatClass(class ast.ClassLike) {
	var declaration = class.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	var hasContent bool
	var constants = class.GetConstants()
	if constants != nil {
		v.formatConstants(constants)
		hasContent = true
	}
	var constructors = class.GetConstructors()
	if constructors != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatConstructors(constructors)
		hasContent = true
	}
	var functions = class.GetFunctions()
	if functions != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatFunctions(functions)
	}
	v.depth_--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatClasses(classes col.ListLike[ast.ClassLike]) {
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

func (v *formatter_) formatConstant(constant ast.ConstantLike) {
	var identifier = constant.GetIdentifier()
	v.appendString(identifier)
	v.appendString("() ")
	var abstraction = constant.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatConstants(constants col.ListLike[ast.ConstantLike]) {
	v.appendNewline()
	v.appendString("// Constants")
	var iterator = constants.GetIterator()
	for iterator.HasNext() {
		var constant = iterator.GetNext()
		v.appendNewline()
		v.formatConstant(constant)
	}
}

func (v *formatter_) formatConstructor(constructor ast.ConstructorLike) {
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

func (v *formatter_) formatConstructors(constructors col.ListLike[ast.ConstructorLike]) {
	v.appendNewline()
	v.appendString("// Constructors")
	var iterator = constructors.GetIterator()
	for iterator.HasNext() {
		var constructor = iterator.GetNext()
		v.appendNewline()
		v.formatConstructor(constructor)
	}
}

func (v *formatter_) formatDeclaration(declaration ast.DeclarationLike) {
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

func (v *formatter_) formatEnumeration(enumeration ast.EnumerationLike) {
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

func (v *formatter_) formatFunction(function ast.FunctionLike) {
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

func (v *formatter_) formatFunctionals(functionals col.ListLike[ast.FunctionalLike]) {
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

func (v *formatter_) formatFunctions(functions col.ListLike[ast.FunctionLike]) {
	v.appendNewline()
	v.appendString("// Functions")
	var iterator = functions.GetIterator()
	for iterator.HasNext() {
		var function = iterator.GetNext()
		v.appendNewline()
		v.formatFunction(function)
	}
}

func (v *formatter_) formatGenerics(parameters col.ListLike[ast.ParameterLike]) {
	var iterator = parameters.GetIterator()
	var parameter = iterator.GetNext()
	v.formatParameter(parameter)
	for iterator.HasNext() {
		parameter = iterator.GetNext()
		v.appendString(", ")
		v.formatParameter(parameter)
	}
}

func (v *formatter_) formatHeader(header ast.HeaderLike) {
	var comment = header.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
	v.appendString("package ")
	var identifier = header.GetIdentifier()
	v.appendString(identifier)
}

func (v *formatter_) formatInstance(instance ast.InstanceLike) {
	var declaration = instance.GetDeclaration()
	v.formatDeclaration(declaration)
	v.appendString(" interface {")
	v.depth_++
	var hasContent bool
	var attributes = instance.GetAttributes()
	if attributes != nil {
		v.formatAttributes(attributes)
		hasContent = true
	}
	var abstractions = instance.GetAbstractions()
	if abstractions != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatAbstractions(abstractions)
		hasContent = true
	}
	var methods = instance.GetMethods()
	if methods != nil {
		if hasContent {
			v.appendString("\n")
		}
		v.formatMethods(methods)
	}
	v.depth_--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) formatInstances(instances col.ListLike[ast.InstanceLike]) {
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

func (v *formatter_) formatMethods(methods col.ListLike[ast.MethodLike]) {
	v.appendNewline()
	v.appendString("// Methods")
	var iterator = methods.GetIterator()
	for iterator.HasNext() {
		var method = iterator.GetNext()
		v.appendNewline()
		v.formatMethod(method)
	}
}

func (v *formatter_) formatModule(module ast.ModuleLike) {
	var identifier = module.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var text = module.GetText()
	v.appendString(text)
}

func (v *formatter_) formatModules(modules col.ListLike[ast.ModuleLike]) {
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

func (v *formatter_) formatNotice(notice ast.NoticeLike) {
	var comment = notice.GetComment()
	comment = v.fixComment(comment)
	v.appendString(comment)
}

func (v *formatter_) formatModel(model ast.ModelLike) {
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

func (v *formatter_) formatParameter(parameter ast.ParameterLike) {
	var identifier = parameter.GetIdentifier()
	v.appendString(identifier)
	v.appendString(" ")
	var abstraction = parameter.GetAbstraction()
	v.formatAbstraction(abstraction)
}

func (v *formatter_) formatParameterName(parameter ast.ParameterLike) {
	var identifier = parameter.GetIdentifier()
	v.appendString(identifier)
}

func (v *formatter_) formatParameterNames(parameters col.ListLike[ast.ParameterLike]) {
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

func (v *formatter_) formatParameters(parameters col.ListLike[ast.ParameterLike]) {
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

func (v *formatter_) formatPrefix(prefix ast.PrefixLike) {
	var identifier = prefix.GetIdentifier()
	switch prefix.GetType() {
	case ast.AliasPrefix:
		v.appendString(identifier)
		v.appendString(".")
	case ast.ArrayPrefix:
		v.appendString("[]")
	case ast.ChannelPrefix:
		v.appendString("chan ")
	case ast.MapPrefix:
		v.appendString("map[")
		v.appendString(identifier)
		v.appendString("]")
	}
}

func (v *formatter_) formatResult(result ast.ResultLike) {
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

func (v *formatter_) formatType(type_ ast.TypeLike) {
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

func (v *formatter_) formatTypes(types col.ListLike[ast.TypeLike]) {
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
