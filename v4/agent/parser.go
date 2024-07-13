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
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	sts "strings"
)

// CLASS ACCESS

// Reference

var parserClass = &parserClass_{
	// Initialize class constants.
	queueSize_: 16,
	stackSize_: 4,
}

// Function

func Parser() ParserClassLike {
	return parserClass
}

// CLASS METHODS

// Target

type parserClass_ struct {
	// Define class constants.
	queueSize_ uint
	stackSize_ uint
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{
		// Initialize instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	// Define instance attributes.
	class_  ParserClassLike
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) ParseSource(source string) ast.ModelLike {
	source = sts.ReplaceAll(source, "\t", "    ")
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse a model.
	var model, token, ok = v.parseModel()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Model",
			"Model",
			"Notice",
			"Header",
			"Imports",
			"Types",
			"Functionals",
			"Classes",
			"Instances",
			"Aspects",
		)
		panic(message)
	}

	// Attempt to parse the end-of-file marker.
	_, token, ok = v.parseToken(EOFToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("EOF",
			"Model",
			"Notice",
			"Header",
			"Imports",
			"Types",
			"Functionals",
			"Classes",
			"Instances",
			"Aspects",
		)
		panic(message)
	}

	// Found a model.
	return model
}

// Private

func (v *parser_) formatError(token TokenLike) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		Scanner().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source line with the error in it.
	message += "\033[36m"
	if line > 1 {
		message += fmt.Sprintf("%04d: ", line-1) + string(lines[line-2]) + "\n"
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count = 0
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < len(lines) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"

	return message
}

func (v *parser_) generateSyntax(expected string, names ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, name := range names {
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			name,
			syntax[name],
		)
	}
	return message
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
	if !ok {
		panic("The token channel terminated without an EOF token.")
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError(token)
		panic(message)
	}

	return token
}

func (v *parser_) parseAbstraction() (
	abstraction ast.AbstractionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional prefix.
	var prefix, _, _ = v.parsePrefix()

	// Attempt to parse an optional alias.
	var alias, _, _ = v.parseAlias()

	// Attempt to parse the name of the abstraction.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		if col.IsDefined(prefix) || col.IsDefined(alias) {
			var message = v.formatError(token)
			message += v.generateSyntax("name",
				"Abstraction",
				"Prefix",
				"GenericArguments",
			)
			panic(message)
		}
		// This is not an abstraction.
		return abstraction, token, false
	}

	// Check if the name is actually a method name for the next declaration.
	var delimiterToken TokenLike
	_, delimiterToken, ok = v.parseToken(DelimiterToken, "(")
	if ok {
		// This is not an abstraction, put back the delimiter and name tokens.
		v.putBack(delimiterToken)
		v.putBack(token)
		return abstraction, token, false
	}

	// Attempt to parse optional generic arguments.
	var genericArguments, _, _ = v.parseGenericArguments()

	// Found an abstraction.
	abstraction = ast.Abstraction().Make(prefix, alias, name, genericArguments)
	return abstraction, token, true
}

func (v *parser_) parseAbstractions() (
	abstractions ast.AbstractionsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Abstractions")
	if !ok {
		// This is not a sequence of abstractions.
		return abstractions, token, false
	}

	// Attempt to parse one or more abstractions.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Abstraction",
			"Abstractions",
			"Abstraction",
		)
		panic(message)
	}
	var list = col.List[ast.AbstractionLike]()
	for ok {
		list.AppendValue(abstraction)
		abstraction, token, ok = v.parseAbstraction()
	}

	// Found a sequence of abstractions.
	abstractions = ast.Abstractions().Make(note, list)
	return abstractions, token, true
}

func (v *parser_) parseAdditionalArgument() (
	additionalArgument ast.AdditionalArgumentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the trailing "," for the previous argument.
	_, token, ok = v.parseToken(DelimiterToken, ",")
	if !ok {
		// This is not an additional argument.
		return additionalArgument, token, false
	}

	// Attempt to parse an additional argument.
	var argument ast.ArgumentLike
	argument, _, ok = v.parseArgument()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Argument",
			"AdditionalArgument",
			"Argument",
		)
		panic(message)
	}

	// Found an additional argument.
	additionalArgument = ast.AdditionalArgument().Make(argument)
	return additionalArgument, token, true
}

func (v *parser_) parseAdditionalParameter() (
	additionalParameter ast.AdditionalParameterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the trailing "," for the previous parameter.
	_, token, ok = v.parseToken(DelimiterToken, ",")
	if !ok {
		// This is not an additional parameter.
		return additionalParameter, token, false
	}

	// Attempt to parse an additional parameter.
	var parameter ast.ParameterLike
	parameter, _, ok = v.parseParameter()
	if !ok {
		// This is not an additional parameter, put back the comma token.
		v.putBack(token)
		return additionalParameter, token, false
	}

	// Found an additional parameter.
	additionalParameter = ast.AdditionalParameter().Make(parameter)
	return additionalParameter, token, true
}

func (v *parser_) parseAdditionalValue() (
	additionalValue ast.AdditionalValueLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the value.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not an additional value.
		return additionalValue, token, false
	}

	// Found an additional value.
	additionalValue = ast.AdditionalValue().Make(name)
	return additionalValue, token, true
}

func (v *parser_) parseAlias() (
	alias ast.AliasLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a module name abbreviation.
	var name string
	var nameToken TokenLike
	name, nameToken, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not an alias.
		return alias, nameToken, false
	}

	// Attempt to parse the trailing ".".
	_, token, ok = v.parseToken(DelimiterToken, ".")
	if !ok {
		// This is not an alias, put back the name token.
		v.putBack(nameToken)
		return alias, token, false
	}

	// Found an alias.
	alias = ast.Alias().Make(name)
	return alias, token, true
}

func (v *parser_) parseArgument() (
	argument ast.ArgumentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an abstraction.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		// This is not an argument.
		return argument, token, false
	}

	// Found an argument.
	argument = ast.Argument().Make(abstraction)
	return argument, token, true
}

func (v *parser_) parseArguments() (
	arguments ast.ArgumentsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an argument.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	if !ok {
		// This is not a sequence of arguments.
		return arguments, token, false
	}

	// Attempt to parse zero or more additional arguments.
	var additionalArguments = col.List[ast.AdditionalArgumentLike]()
	var additionalArgument ast.AdditionalArgumentLike
	additionalArgument, token, ok = v.parseAdditionalArgument()
	for ok {
		additionalArguments.AppendValue(additionalArgument)
		additionalArgument, token, ok = v.parseAdditionalArgument()
	}

	// Found a sequence of arguments.
	arguments = ast.Arguments().Make(argument, additionalArguments)
	return arguments, token, true
}

func (v *parser_) parseArray() (
	array ast.ArrayLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening "[" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if !ok {
		// This is not an array.
		return array, token, false
	}

	// Attempt to parse the closing "]" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"Array",
		)
		panic(message)
	}

	// Found an array.
	array = ast.Array().Make()
	return array, token, true
}

func (v *parser_) parseAspect() (
	aspect ast.AspectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an aspect.
		return aspect, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(NameToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"interface"`,
			"Aspect",
			"Declaration",
			"Method",
		)
		panic(message)
	}

	// Attempt to parse the opening "{" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("{",
			"Aspect",
			"Declaration",
			"Method",
		)
		panic(message)
	}

	// Attempt to parse one or more methods.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Methods",
			"Aspect",
			"Declaration",
			"Method",
		)
		panic(message)
	}
	var methods = col.List[ast.MethodLike]()
	for ok {
		methods.AppendValue(method)
		method, _, ok = v.parseMethod()
	}

	// Attempt to parse the closing "}" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Aspect",
			"Declaration",
			"Method",
		)
		panic(message)
	}

	// Found an aspect.
	aspect = ast.Aspect().Make(declaration, methods)
	return aspect, token, true
}

func (v *parser_) parseAspects() (
	aspects ast.AspectsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Aspects")
	if !ok {
		// This is not a sequence of aspects.
		return aspects, token, false
	}

	// Attempt to parse one or more aspects.
	var aspect ast.AspectLike
	aspect, token, ok = v.parseAspect()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Aspect",
			"Aspects",
			"Aspect",
		)
		panic(message)
	}
	var list = col.List[ast.AspectLike]()
	for ok {
		list.AppendValue(aspect)
		aspect, token, ok = v.parseAspect()
	}

	// Found a sequence of aspects.
	list.SortValuesWithRanker(func(first, second ast.AspectLike) col.Rank {
		var firstName = first.GetDeclaration().GetName()
		var secondName = second.GetDeclaration().GetName()
		switch {
		case firstName < secondName:
			return col.LesserRank
		case firstName > secondName:
			return col.GreaterRank
		default:
			return col.EqualRank
		}
	})
	aspects = ast.Aspects().Make(note, list)
	return aspects, token, true
}

func (v *parser_) parseAttribute() (
	attribute ast.AttributeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the attribute.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not an attribute.
		return attribute, token, false
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Attribute",
			"Parameter",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional parameter.
	var parameter, _, _ = v.parseParameter()

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Attribute",
			"Parameter",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional abstraction.
	var abstraction, _, _ = v.parseAbstraction()

	// Found an attribute.
	attribute = ast.Attribute().Make(name, parameter, abstraction)
	return attribute, token, true
}

func (v *parser_) parseAttributes() (
	attributes ast.AttributesLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Attributes")
	if !ok {
		// This is not a sequence of attributes.
		return attributes, token, false
	}

	// Attempt to parse one or more attributes.
	var attribute ast.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Attribute",
			"Attributes",
			"Attribute",
		)
		panic(message)
	}
	var list = col.List[ast.AttributeLike]()
	for ok {
		list.AppendValue(attribute)
		attribute, token, ok = v.parseAttribute()
	}

	// Found a sequence of attributes.
	attributes = ast.Attributes().Make(note, list)
	return attributes, token, true
}

func (v *parser_) parseChannel() (
	channel ast.ChannelLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the "chan" keyword.
	_, token, ok = v.parseToken(NameToken, "chan")
	if !ok {
		// This is not a channel.
		return channel, token, false
	}

	// Found a channel.
	channel = ast.Channel().Make()
	return channel, token, true
}

func (v *parser_) parseClass() (
	class ast.ClassLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a class.
		return class, token, false
	}

	// Attempt to parse the "interface" keyword.
	_, token, ok = v.parseToken(NameToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"interface"`,
			"Class",
			"Declaration",
			"Constructors",
			"Constants",
			"Functions",
		)
		panic(message)
	}

	// Attempt to parse the opening "{" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("{",
			"Class",
			"Declaration",
			"Constructors",
			"Constants",
			"Functions",
		)
		panic(message)
	}

	// Attempt to parse a sequence of constructors.
	var constructors ast.ConstructorsLike
	constructors, token, ok = v.parseConstructors()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("// Constructors",
			"Class",
			"Declaration",
			"Constructors",
			"Constants",
			"Functions",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of constants.
	var constants, _, _ = v.parseConstants()

	// Attempt to parse an optional sequence of functions.
	var functions, _, _ = v.parseFunctions()

	// Attempt to parse the closing "}" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Class",
			"Declaration",
			"Constructors",
			"Constants",
			"Functions",
		)
		panic(message)
	}

	// Found a class.
	class = ast.Class().Make(declaration, constructors, constants, functions)
	return class, token, true
}

func (v *parser_) parseClasses() (
	classes ast.ClassesLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Classes")
	if !ok {
		// This is not a sequence of classes.
		return classes, token, false
	}

	// Attempt to parse one or more classes.
	var class ast.ClassLike
	class, token, ok = v.parseClass()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Class",
			"Classes",
			"Class",
		)
		panic(message)
	}
	var list = col.List[ast.ClassLike]()
	for ok {
		list.AppendValue(class)
		class, token, ok = v.parseClass()
	}

	// Found a sequence of classes.
	list.SortValuesWithRanker(func(first, second ast.ClassLike) col.Rank {
		var firstName = first.GetDeclaration().GetName()
		var secondName = second.GetDeclaration().GetName()
		switch {
		case firstName < secondName:
			return col.LesserRank
		case firstName > secondName:
			return col.GreaterRank
		default:
			return col.EqualRank
		}
	})
	classes = ast.Classes().Make(note, list)
	return classes, token, true
}

func (v *parser_) parseConstant() (
	constant ast.ConstantLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the constant.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not a constant.
		return constant, token, false
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional abstraction.
	var abstraction, _, _ = v.parseAbstraction()

	// Found a constant.
	constant = ast.Constant().Make(name, abstraction)
	return constant, token, true
}

func (v *parser_) parseConstants() (
	constants ast.ConstantsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Constants")
	if !ok {
		// This is not a sequence of constants.
		return constants, token, false
	}

	// Attempt to parse one or more constants.
	var constant ast.ConstantLike
	constant, token, ok = v.parseConstant()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Constant",
			"Constants",
			"Constant",
		)
		panic(message)
	}
	var list = col.List[ast.ConstantLike]()
	for ok {
		list.AppendValue(constant)
		constant, token, ok = v.parseConstant()
	}

	// Found a sequence of constants.
	constants = ast.Constants().Make(note, list)
	return constants, token, true
}

func (v *parser_) parseConstructor() (
	constructor ast.ConstructorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the constructor.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not a constructor.
		return constructor, token, false
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Constructor",
			"Parameters",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	var parameters, _, _ = v.parseParameters()

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Constructor",
			"Parameters",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional abstraction.
	var abstraction, _, _ = v.parseAbstraction()

	// Found a constructor.
	constructor = ast.Constructor().Make(name, parameters, abstraction)
	return constructor, token, true
}

func (v *parser_) parseConstructors() (
	constructors ast.ConstructorsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Constructors")
	if !ok {
		// This is not a sequence of constructors.
		return constructors, token, false
	}

	// Attempt to parse one or more constructors.
	var constructor ast.ConstructorLike
	constructor, token, ok = v.parseConstructor()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Constructor",
			"Constructors",
			"Constructor",
		)
		panic(message)
	}
	var list = col.List[ast.ConstructorLike]()
	for ok {
		list.AppendValue(constructor)
		constructor, token, ok = v.parseConstructor()
	}

	// Found a sequence of constructors.
	constructors = ast.Constructors().Make(note, list)
	return constructors, token, true
}

func (v *parser_) parseDeclaration() (
	declaration ast.DeclarationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a comment.
	var comment string
	comment, token, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not a declaration.
		return declaration, token, false
	}

	// Attempt to parse the "type" keyword.
	_, token, ok = v.parseToken(NameToken, "type")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"type"`,
			"Declaration",
			"GenericParameters",
		)
		panic(message)
	}

	// Attempt to parse the name of the type declaration.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("name",
			"Declaration",
			"GenericParameters",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of generic parameters.
	var genericParameters, _, _ = v.parseGenericParameters()

	// Found a declaration.
	declaration = ast.Declaration().Make(comment, name, genericParameters)
	return declaration, token, true
}

func (v *parser_) parseEnumeration() (
	enumeration ast.EnumerationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the "const" keyword.
	_, token, ok = v.parseToken(NameToken, "const")
	if !ok {
		// This is not an enumeration.
		return enumeration, token, false
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Enumeration",
			"Values",
		)
		panic(message)
	}

	// Attempt to parse the enumerated values.
	var values ast.ValuesLike
	values, token, ok = v.parseValues()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Values",
			"Enumeration",
			"Values",
		)
		panic(message)
	}

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Enumeration",
			"Values",
		)
		panic(message)
	}

	// Found an enumeration.
	enumeration = ast.Enumeration().Make(values)
	return enumeration, token, true
}

func (v *parser_) parseFunction() (
	function ast.FunctionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the function.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not a function.
		return function, token, false
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Function",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	var parameters, _, _ = v.parseParameters()

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Function",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse a result.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Result",
			"Function",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Found a function.
	function = ast.Function().Make(name, parameters, result)
	return function, token, true
}

func (v *parser_) parseFunctional() (
	functional ast.FunctionalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a functional.
		return functional, token, false
	}

	// Attempt to parse the "func" keyword.
	_, token, ok = v.parseToken(NameToken, "func")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"func"`,
			"Functional",
			"Declaration",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Functional",
			"Declaration",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	var parameters, _, _ = v.parseParameters()

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Functional",
			"Declaration",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse a result.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Result",
			"Functional",
			"Declaration",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Found a functional.
	functional = ast.Functional().Make(declaration, parameters, result)
	return functional, token, true
}

func (v *parser_) parseFunctionals() (
	functionals ast.FunctionalsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Functionals")
	if !ok {
		// This is not a sequence of functionals.
		return functionals, token, false
	}

	// Attempt to parse one or more functionals.
	var functional ast.FunctionalLike
	functional, token, ok = v.parseFunctional()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Functional",
			"Functionals",
			"Functional",
		)
		panic(message)
	}
	var list = col.List[ast.FunctionalLike]()
	for ok {
		list.AppendValue(functional)
		functional, token, ok = v.parseFunctional()
	}

	// Found a sequence of functionals.
	list.SortValuesWithRanker(func(first, second ast.FunctionalLike) col.Rank {
		var firstName = first.GetDeclaration().GetName()
		var secondName = second.GetDeclaration().GetName()
		switch {
		case firstName < secondName:
			return col.LesserRank
		case firstName > secondName:
			return col.GreaterRank
		default:
			return col.EqualRank
		}
	})
	functionals = ast.Functionals().Make(note, list)
	return functionals, token, true
}

func (v *parser_) parseFunctions() (
	functions ast.FunctionsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Functions")
	if !ok {
		// This is not a sequence of functions.
		return functions, token, false
	}

	// Attempt to parse one or more functions.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Function",
			"Functions",
			"Function",
		)
		panic(message)
	}
	var list = col.List[ast.FunctionLike]()
	for ok {
		list.AppendValue(function)
		function, token, ok = v.parseFunction()
	}

	// Found a sequence of functions.
	functions = ast.Functions().Make(note, list)
	return functions, token, true
}

func (v *parser_) parseGenericArguments() (
	genericArguments ast.GenericArgumentsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening "[" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if !ok {
		// This is not a sequence of generic arguments.
		return genericArguments, token, false
	}

	// Attempt to parse the arguments.
	var arguments ast.ArgumentsLike
	arguments, token, ok = v.parseArguments()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Arguments",
			"GenericArguments",
			"Arguments",
		)
		panic(message)
	}

	// Attempt to parse the closing "]" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"GenericArguments",
			"Arguments",
		)
		panic(message)
	}

	// Found a sequence of generic arguments.
	genericArguments = ast.GenericArguments().Make(arguments)
	return genericArguments, token, true
}

func (v *parser_) parseGenericParameters() (
	genericParameters ast.GenericParametersLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening "[" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if !ok {
		// This is not a sequence of generic parameters.
		return genericParameters, token, false
	}

	// Attempt to parse the parameters.
	var parameters ast.ParametersLike
	parameters, token, ok = v.parseParameters()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Parameters",
			"GenericParameters",
			"Parameters",
		)
		panic(message)
	}

	// Attempt to parse the closing "]" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"GenericParameters",
			"Parameters",
		)
		panic(message)
	}

	// Found a sequence of generic parameters.
	genericParameters = ast.GenericParameters().Make(parameters)
	return genericParameters, token, true
}

func (v *parser_) parseHeader() (
	header ast.HeaderLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a comment.
	var comment string
	comment, token, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not a header.
		return header, token, false
	}

	// Attempt to parse the "package" keyword.
	_, token, ok = v.parseToken(NameToken, "package")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"package"`,
			"Header",
		)
		panic(message)
	}

	// Attempt to parse the name of the package.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("name",
			"Header",
		)
		panic(message)
	}

	// Found a header.
	header = ast.Header().Make(comment, name)
	return header, token, true
}

func (v *parser_) parseImports() (
	imports ast.ImportsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the "import" keyword.
	_, token, ok = v.parseToken(NameToken, "import")
	if !ok {
		// This is not a sequence of imported modules.
		return imports, token, false
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Imports",
			"Modules",
		)
		panic(message)
	}

	// Attempt to parse a sequence of imported modules.
	var modules ast.ModulesLike
	modules, token, ok = v.parseModules()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Modules",
			"Imports",
			"Modules",
		)
		panic(message)
	}

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Imports",
			"Modules",
		)
		panic(message)
	}

	// Found a sequence of imported modules.
	imports = ast.Imports().Make(modules)
	return imports, token, true
}

func (v *parser_) parseInstance() (
	instance ast.InstanceLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an instance.
		return instance, token, false
	}

	// Attempt to parse the "interface" keyword.
	_, token, ok = v.parseToken(NameToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"interface"`,
			"Instance",
			"Declaration",
			"Attributes",
			"Abstractions",
			"Methods",
		)
		panic(message)
	}

	// Attempt to parse the opening "{" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("{",
			"Instance",
			"Declaration",
			"Attributes",
			"Abstractions",
			"Methods",
		)
		panic(message)
	}

	// Attempt to parse a sequence of attributes.
	var attributes ast.AttributesLike
	attributes, token, ok = v.parseAttributes()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("{",
			"Instance",
			"Declaration",
			"Attributes",
			"Abstractions",
			"Methods",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of abstractions.
	var abstractions, _, _ = v.parseAbstractions()

	// Attempt to parse an optional sequence of methods.
	var methods, _, _ = v.parseMethods()

	// Attempt to parse the closing "}" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Instance",
			"Declaration",
			"Attributes",
			"Abstractions",
			"Methods",
		)
		panic(message)
	}

	// Found an instance.
	instance = ast.Instance().Make(declaration, attributes, abstractions, methods)
	return instance, token, true
}

func (v *parser_) parseInstances() (
	instances ast.InstancesLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Instances")
	if !ok {
		// This is not a sequence of instances.
		return instances, token, false
	}

	// Attempt to parse one or more instances.
	var instance ast.InstanceLike
	instance, token, ok = v.parseInstance()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Instance",
			"Instances",
			"Instance",
		)
		panic(message)
	}
	var list = col.List[ast.InstanceLike]()
	for ok {
		list.AppendValue(instance)
		instance, token, ok = v.parseInstance()
	}

	// Found a sequence of instances.
	list.SortValuesWithRanker(func(first, second ast.InstanceLike) col.Rank {
		var firstName = first.GetDeclaration().GetName()
		var secondName = second.GetDeclaration().GetName()
		switch {
		case firstName < secondName:
			return col.LesserRank
		case firstName > secondName:
			return col.GreaterRank
		default:
			return col.EqualRank
		}
	})
	instances = ast.Instances().Make(note, list)
	return instances, token, true
}

func (v *parser_) parseMap() (
	map_ ast.MapLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the "map" keyword.
	_, token, ok = v.parseToken(NameToken, "map")
	if !ok {
		// This is not a map.
		return map_, token, false
	}

	// Attempt to parse the opening "[" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("[",
			"Map",
		)
		panic(message)
	}

	// Attempt to parse the name of the map key type.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("name",
			"Map",
		)
		panic(message)
	}

	// Attempt to parse the closing "]" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "]")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("]",
			"Map",
		)
		panic(message)
	}

	// Found a map.
	map_ = ast.Map().Make(name)
	return map_, token, true
}

func (v *parser_) parseMethod() (
	method ast.MethodLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the method.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not a method.
		return method, token, false
	}

	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Method",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	var parameters, _, _ = v.parseParameters()

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Method",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse an optional result.
	var result ast.ResultLike
	result, _, _ = v.parseResult()

	// Found a method.
	method = ast.Method().Make(name, parameters, result)
	return method, token, true
}

func (v *parser_) parseMethods() (
	methods ast.MethodsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Methods")
	if !ok {
		// This is not a sequence of methods.
		return methods, token, false
	}

	// Attempt to parse one or more methods.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Method",
			"Methods",
			"Method",
		)
		panic(message)
	}
	var list = col.List[ast.MethodLike]()
	for ok {
		list.AppendValue(method)
		method, token, ok = v.parseMethod()
	}

	// Found a sequence of methods.
	methods = ast.Methods().Make(note, list)
	return methods, token, true
}

func (v *parser_) parseModel() (
	model ast.ModelLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a notice.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	if !ok {
		// This is not a model.
		return model, token, false
	}

	// Attempt to parse a header.
	var header ast.HeaderLike
	header, token, ok = v.parseHeader()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Header",
			"Model",
			"Notice",
			"Header",
			"Imports",
			"Types",
			"Functionals",
			"Classes",
			"Instances",
			"Aspects",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of imports.
	var imports, _, _ = v.parseImports()

	// Attempt to parse an optional sequence of types.
	var types, _, _ = v.parseTypes()

	// Attempt to parse an optional sequence of functionals.
	var functionals, _, _ = v.parseFunctionals()

	// Attempt to parse an optional sequence of classes.
	var classes, _, _ = v.parseClasses()

	// Attempt to parse an optional sequence of instances.
	var instances, _, _ = v.parseInstances()

	// Attempt to parse an optional sequence of aspects.
	var aspects, _, _ = v.parseAspects()

	// Found a model.
	model = ast.Model().Make(
		notice,
		header,
		imports,
		types,
		functionals,
		classes,
		instances,
		aspects,
	)
	return model, token, true
}

func (v *parser_) parseModule() (
	module ast.ModuleLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the module.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not a module.
		return module, token, false
	}

	// Attempt to parse the path of the module.
	var path string
	path, token, ok = v.parseToken(PathToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("path",
			"Module",
		)
		panic(message)
	}

	// Found a module.
	module = ast.Module().Make(name, path)
	return module, token, true
}

func (v *parser_) parseModules() (
	modules ast.ModulesLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a sequence of modules.
	var list = col.List[ast.ModuleLike]()
	var module ast.ModuleLike
	module, token, ok = v.parseModule()
	for ok {
		list.AppendValue(module)
		module, token, ok = v.parseModule()
	}

	// Found a sequence of modules.
	list.SortValuesWithRanker(func(first, second ast.ModuleLike) col.Rank {
		var firstPath = first.GetPath()
		var secondPath = second.GetPath()
		switch {
		case firstPath < secondPath:
			return col.LesserRank
		case firstPath > secondPath:
			return col.GreaterRank
		default:
			return col.EqualRank
		}
	})
	modules = ast.Modules().Make(list)
	return modules, token, true
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a comment.
	var comment string
	comment, token, ok = v.parseToken(CommentToken, "")
	if !ok {
		// This is not a notice.
		return notice, token, false
	}

	// Found a notice.
	notice = ast.Notice().Make(comment)
	return notice, token, true
}

func (v *parser_) parseParameter() (
	parameter ast.ParameterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the parameter.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not a parameter.
		return parameter, token, false
	}

	// Attempt to parse the abstract type of the parameter.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Abstraction",
			"Parameter",
			"Abstraction",
		)
		panic(message)
	}

	// Found a parameter.
	parameter = ast.Parameter().Make(name, abstraction)
	return parameter, token, true
}

func (v *parser_) parseParameterized() (
	parameterized ast.ParameterizedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the opening "(" bracket.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		// This is not a parameterized result.
		return parameterized, token, false
	}

	// Attempt to parse a sequence of parameters.
	var parameters ast.ParametersLike
	parameters, token, ok = v.parseParameters()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Parameters",
			"Parameterized",
			"Parameters",
		)
		panic(message)
	}

	// Attempt to parse the closing ")" bracket.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Parameterized",
			"Parameters",
		)
		panic(message)
	}

	// Found a parameterized result.
	parameterized = ast.Parameterized().Make(parameters)
	return parameterized, token, true
}

func (v *parser_) parseParameters() (
	parameters ast.ParametersLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a parameter.
	var parameter ast.ParameterLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		// This is not a sequence of parameters.
		return parameters, token, false
	}

	// Attempt to parse zero or more additional parameters.
	var additionalParameters = col.List[ast.AdditionalParameterLike]()
	var additionalParameter ast.AdditionalParameterLike
	additionalParameter, token, ok = v.parseAdditionalParameter()
	for ok {
		additionalParameters.AppendValue(additionalParameter)
		additionalParameter, token, ok = v.parseAdditionalParameter()
	}

	// Attempt to parse an optional trailing "," for the last parameter.
	v.parseToken(DelimiterToken, ",")

	// Found a sequence of parameters.
	parameters = ast.Parameters().Make(parameter, additionalParameters)
	return parameters, token, true
}

func (v *parser_) parsePrefix() (
	prefix ast.PrefixLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an array prefix.
	var array ast.ArrayLike
	array, token, ok = v.parseArray()
	if ok {
		prefix = ast.Prefix().Make(array)
		return prefix, token, true
	}

	// Attempt to parse an map prefix.
	var map_ ast.MapLike
	map_, token, ok = v.parseMap()
	if ok {
		prefix = ast.Prefix().Make(map_)
		return prefix, token, true
	}

	// Attempt to parse an channel prefix.
	var channel ast.ChannelLike
	channel, token, ok = v.parseChannel()
	if ok {
		prefix = ast.Prefix().Make(channel)
		return prefix, token, true
	}

	// This is not a prefix.
	return prefix, token, false
}

func (v *parser_) parseResult() (
	result ast.ResultLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an abstract result type.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if ok {
		result = ast.Result().Make(abstraction)
		return result, token, true
	}

	// Attempt to parse a parameterized result type.
	var parameterized ast.ParameterizedLike
	parameterized, token, ok = v.parseParameterized()
	if ok {
		result = ast.Result().Make(parameterized)
		return result, token, true
	}

	// This is not a result.
	return result, token, false
}

func (v *parser_) parseToken(expectedType TokenType, expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	token = v.getNextToken()
	if token.GetType() == expectedType {
		// Found the right token type.
		value = token.GetValue()
		if col.IsUndefined(expectedValue) || value == expectedValue {
			// Found the right token value.
			return value, token, true
		}
	}

	// This is not the right token.
	v.putBack(token)
	return value, token, false
}

func (v *parser_) parseType() (
	type_ ast.TypeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a type.
		return type_, token, false
	}

	// Attempt to parse an abstraction.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Abstraction",
			"Type",
			"Declaration",
			"Abstraction",
			"Enumeration",
		)
		panic(message)
	}

	// Attempt to parse an optional enumeration.
	var enumeration ast.EnumerationLike
	enumeration, _, _ = v.parseEnumeration()

	// Found a type.
	type_ = ast.Type().Make(declaration, abstraction, enumeration)
	return type_, token, true
}

func (v *parser_) parseTypes() (
	types ast.TypesLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	var note string
	note, token, ok = v.parseToken(NoteToken, "// Types")
	if !ok {
		// This is not a sequence of types.
		return types, token, false
	}

	// Attempt to parse one or more types.
	var type_ ast.TypeLike
	type_, token, ok = v.parseType()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Type",
			"Types",
			"Type",
		)
		panic(message)
	}
	var list = col.List[ast.TypeLike]()
	for ok {
		list.AppendValue(type_)
		type_, token, ok = v.parseType()
	}

	// Found a sequence of types.
	list.SortValuesWithRanker(func(first, second ast.TypeLike) col.Rank {
		var firstName = first.GetDeclaration().GetName()
		var secondName = second.GetDeclaration().GetName()
		switch {
		case firstName < secondName:
			return col.LesserRank
		case firstName > secondName:
			return col.GreaterRank
		default:
			return col.EqualRank
		}
	})
	types = ast.Types().Make(note, list)
	return types, token, true
}

func (v *parser_) parseValue() (
	value ast.ValueLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse the name of the enumerated value.
	var name string
	name, token, ok = v.parseToken(NameToken, "")
	if !ok {
		// This is not an enumerated value.
		return value, token, false
	}

	// Attempt to parse the abstract type of the enumerated value.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Abstraction",
			"Value",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse the "=" operator.
	_, token, ok = v.parseToken(DelimiterToken, "=")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("=",
			"Value",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse the "iota" keyword.
	_, token, ok = v.parseToken(NameToken, "iota")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"iota"`,
			"Value",
			"Abstraction",
		)
		panic(message)
	}

	// Found an enumerated value.
	value = ast.Value().Make(name, abstraction)
	return value, token, true
}

func (v *parser_) parseValues() (
	values ast.ValuesLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a value.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if !ok {
		// This is not a sequence of values.
		return values, token, false
	}

	// Attempt to parse zero or more additional values.
	var additionalValues = col.List[ast.AdditionalValueLike]()
	var additionalValue ast.AdditionalValueLike
	additionalValue, token, ok = v.parseAdditionalValue()
	for ok {
		additionalValues.AppendValue(additionalValue)
		additionalValue, token, ok = v.parseAdditionalValue()
	}

	// Found a sequence of values.
	values = ast.Values().Make(value, additionalValues)
	return values, token, true
}

func (v *parser_) putBack(token TokenLike) {
	//fmt.Printf("Put Back %v\n", token)
	v.next_.AddValue(token)
}

var syntax = map[string]string{
	"Model":               `Notice Header Imports? Types? Functionals? Classes Instances Aspects? EOF`,
	"Notice":              `comment`,
	"Header":              `comment "package" name`,
	"Imports":             `"import" "(" Modules ")"`,
	"Modules":             `Module+`,
	"Module":              `name path`,
	"Types":               `note Type+`,
	"Type":                `Declaration Abstraction Enumeration?`,
	"Declaration":         `comment "type" name GenericParameters?`,
	"GenericParameters":   `"[" Parameters "]"`,
	"Parameters":          `Parameter AdditionalParameter* ","?`,
	"Parameter":           `name Abstraction`,
	"AdditionalParameter": `"," Parameter`,
	"Abstraction":         `Prefix? Alias? name GenericArguments?`,
	"Prefix": `
    Array
    Map
    Channel`,
	"Array":              `"[" "]"`,
	"Map":                `"map" "[" name "]"`,
	"Channel":            `"chan"`,
	"Alias":              `name "."`,
	"GenericArguments":   `"[" Arguments "]"`,
	"Arguments":          `Argument AdditionalArgument*`,
	"Argument":           `Abstraction`,
	"AdditionalArgument": `"," Argument`,
	"Enumeration":        `"const" "(" Values ")"`,
	"Values":             `Value AdditionalValue*`,
	"Value":              `name Abstraction "=" "iota"`,
	"AdditionalValue":    `name`,
	"Functionals":        `note Functional+`,
	"Functional":         `Declaration "func" "(" Parameters? ")" Result`,
	"Result": `
    Abstraction
    Parameterized`,
	"Parameterized": `"(" Parameters ")"`,
	"Classes":       `note Class+`,
	"Class":         `Declaration "interface" "{" Constructors Constants? Functions? "}"`,
	"Constructors":  `note Constructor+`,
	"Constructor":   `name "(" Parameters? ")" Abstraction`,
	"Constants":     `note Constant+`,
	"Constant":      `name "(" ")" Abstraction`,
	"Functions":     `note Function+`,
	"Function":      `name "(" Parameters? ")" Result`,
	"Instances":     `note Instance+`,
	"Instance":      `Declaration "interface" "{" Attributes Abstractions? Methods? "}"`,
	"Attributes":    `note Attribute+`,
	"Attribute":     `name "(" Parameter? ")" Abstraction?`,
	"Abstractions":  `note Abstraction+`,
	"Methods":       `note Method+`,
	"Method":        `name "(" Parameters? ")" Result?`,
	"Aspects":       `note Aspect+`,
	"Aspect":        `Declaration "interface" "{" Method+ "}"`,
}
