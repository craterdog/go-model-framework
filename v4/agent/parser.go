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

var parserClass = &parserClass_{
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
	queueSize_ int
	stackSize_ int
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{
		tokens_: gcf.Queue[TokenLike](c.queueSize_),
		next_:   gcf.Stack[TokenLike](c.stackSize_),
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	class_  ParserClassLike
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Attributes

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

// Public

func (v *parser_) ParseSource(source string) gcm.ModelLike {
	// The scanner runs in a separate Go routine.
	v.source_ = source
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse a model.
	var model, token, ok = v.parseModel()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Model",
			"Gcmn",
			"Model",
		)
		panic(message)
	}

	// Attempt to parse optional end-of-line characters.
	for ok {
		_, _, ok = v.parseToken(EOLToken, "")
	}

	// Attempt to parse the end-of-file marker.
	_, token, ok = v.parseToken(EOFToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("EOF",
			"Gcmn",
			"Model",
		)
		panic(message)
	}

	// Found a model.
	return model
}

// Private

/*
This private instance method returns an error message containing the context for
a parsing error.
*/
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

/*
This private instance method is useful when creating scanner and parser error
messages that include the required grammatical rules.
*/
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

/*
This private instance method attempts to read the next token from the token
stream and return it.
*/
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
	abstraction gcm.AbstractionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional prefix.
	var prefix gcm.PrefixLike
	prefix, _, ok = v.parsePrefix()
	var identifier string
	if ok {
		// Attempt to parse an identifier.
		identifier, token, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("Arguments",
				"Abstraction",
				"Prefix",
				"Arguments",
			)
			panic(message)
		}
	} else {
		// Attempt to parse an identifier.
		var identifierToken TokenLike
		identifier, identifierToken, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			// This is not an abstraction.
			return abstraction, identifierToken, false
		}
		var delimiterToken TokenLike
		_, delimiterToken, ok = v.parseToken(DelimiterToken, "(")
		if ok {
			// The identifier is the next method name not an abstraction.
			v.putBack(delimiterToken)
			v.putBack(identifierToken)
			return abstraction, identifierToken, false
		}

	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	var arguments = gcf.List[gcm.AbstractionLike]()
	if ok {
		// Attempt to parse a sequence of arguments.
		arguments, token, ok = v.parseArguments()
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("Arguments",
				"Abstraction",
				"Prefix",
				"Arguments",
			)
			panic(message)
		}

		// Attempt to parse a delimiter.
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("]",
				"Abstraction",
				"Prefix",
				"Arguments",
			)
			panic(message)
		}
	}

	// Found an abstraction.
	abstraction = gcm.Abstraction().MakeWithAttributes(prefix, identifier, arguments)
	return abstraction, token, true
}

func (v *parser_) parseAbstractions() (
	abstractions col.ListLike[gcm.AbstractionLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	abstractions = gcf.List[gcm.AbstractionLike]()
	_, token, ok = v.parseToken(NoteToken, "// Abstractions")
	if !ok {
		// This is not a sequence of abstractions.
		return abstractions, token, false
	}

	// Attempt to parse at least one abstraction.
	var abstraction gcm.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Abstraction",
			"Abstractions",
			"Abstraction",
		)
		panic(message)
	}
	for ok {
		abstractions.AppendValue(abstraction)
		abstraction, token, ok = v.parseAbstraction()
	}

	// Found a sequence of abstractions.
	return abstractions, token, true
}

func (v *parser_) parseArguments() (
	arguments col.ListLike[gcm.AbstractionLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse at least one abstraction.
	arguments = gcf.List[gcm.AbstractionLike]()
	var abstraction gcm.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		// This is not a sequence of arguments.
		return arguments, token, false
	}
	for ok {
		arguments.AppendValue(abstraction)
		_, token, ok = v.parseToken(DelimiterToken, ",")
		if ok {
			abstraction, token, ok = v.parseAbstraction()
		}
	}

	// Found a sequence of arguments.
	return arguments, token, true
}

func (v *parser_) parseAspect() (
	aspect gcm.AspectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration gcm.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an aspect.
		return aspect, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"interface"`,
			"Aspect",
			"Declaration",
			"Methods",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("{",
			"Aspect",
			"Declaration",
			"Methods",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of methods.
	var methods, _, _ = v.parseMethods()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Aspect",
			"Declaration",
			"Methods",
		)
		panic(message)
	}

	// Found an aspect.
	aspect = gcm.Aspect().MakeWithAttributes(declaration, methods)
	return aspect, token, true
}

func (v *parser_) parseAspects() (
	aspects col.ListLike[gcm.AspectLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	aspects = gcf.List[gcm.AspectLike]()
	_, token, ok = v.parseToken(NoteToken, "// Aspects")
	if !ok {
		// This is not a sequence of aspects.
		return aspects, token, false
	}

	// Attempt to parse at least one aspect.
	var aspect gcm.AspectLike
	aspect, token, ok = v.parseAspect()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Aspect",
			"Aspects",
			"Aspect",
		)
		panic(message)
	}
	for ok {
		aspects.AppendValue(aspect)
		aspect, token, ok = v.parseAspect()
	}

	// Found a sequence of aspects.
	return aspects, token, true
}

func (v *parser_) parseAttribute() (
	attribute gcm.AttributeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a attribute.
		return attribute, token, false
	}

	// Attempt to parse a delimiter.
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

	// Attempt to parse a delimiter.
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

	// Found a attribute.
	attribute = gcm.Attribute().MakeWithAttributes(identifier, parameter, abstraction)
	return attribute, token, true
}

func (v *parser_) parseAttributes() (
	attributes col.ListLike[gcm.AttributeLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	attributes = gcf.List[gcm.AttributeLike]()
	_, token, ok = v.parseToken(NoteToken, "// Attributes")
	if !ok {
		// This is not a sequence of attributes.
		return attributes, token, false
	}

	// Attempt to parse at least one attribute.
	var attribute gcm.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Attribute",
			"Attributes",
			"Attribute",
		)
		panic(message)
	}
	for ok {
		attributes.AppendValue(attribute)
		attribute, token, ok = v.parseAttribute()
	}

	// Found a sequence of attributes.
	return attributes, token, true
}

func (v *parser_) parseClass() (
	class gcm.ClassLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration gcm.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a class.
		return class, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"interface"`,
			"Class",
			"Declaration",
			"Constants",
			"Constructors",
			"Functions",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "{")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("{",
			"Class",
			"Declaration",
			"Constants",
			"Constructors",
			"Functions",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of constants.
	var constants, _, _ = v.parseConstants()

	// Attempt to parse an optional sequence of constructors.
	var constructors, _, _ = v.parseConstructors()

	// Attempt to parse an optional sequence of functions.
	var functions, _, _ = v.parseFunctions()

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "}")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("}",
			"Class",
			"Declaration",
			"Constants",
			"Constructors",
			"Functions",
		)
		panic(message)
	}

	// Found a class.
	class = gcm.Class().MakeWithAttributes(declaration, constants, constructors, functions)
	return class, token, true
}

func (v *parser_) parseClasses() (
	classes col.ListLike[gcm.ClassLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	classes = gcf.List[gcm.ClassLike]()
	_, token, ok = v.parseToken(NoteToken, "// Classes")
	if !ok {
		// This is not a sequence of classes.
		return classes, token, false
	}

	// Attempt to parse at least one class.
	var class gcm.ClassLike
	class, token, ok = v.parseClass()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Class",
			"Classes",
			"Class",
		)
		panic(message)
	}
	for ok {
		classes.AppendValue(class)
		class, token, ok = v.parseClass()
	}

	// Found a sequence of classes.
	return classes, token, true
}

func (v *parser_) parseConstant() (
	constant gcm.ConstantLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a constant.
		return constant, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an abstraction.
	var abstraction gcm.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Abstraction",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Found a constant.
	constant = gcm.Constant().MakeWithAttributes(identifier, abstraction)
	return constant, token, true
}

func (v *parser_) parseConstants() (
	constants col.ListLike[gcm.ConstantLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	constants = gcf.List[gcm.ConstantLike]()
	_, token, ok = v.parseToken(NoteToken, "// Constants")
	if !ok {
		// This is not a sequence of constants.
		return constants, token, false
	}

	// Attempt to parse at least one constant.
	var constant gcm.ConstantLike
	constant, token, ok = v.parseConstant()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Constant",
			"Constants",
			"Constant",
		)
		panic(message)
	}
	for ok {
		constants.AppendValue(constant)
		constant, token, ok = v.parseConstant()
	}

	// Found a sequence of constants.
	return constants, token, true
}

func (v *parser_) parseConstructor() (
	constructor gcm.ConstructorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a constructor.
		return constructor, token, false
	}

	// Attempt to parse a delimiter.
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

	// Attempt to parse a delimiter.
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

	// Attempt to parse an abstraction.
	var abstraction gcm.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Abstraction",
			"Constructor",
			"Parameters",
			"Abstraction",
		)
		panic(message)
	}

	// Found a constructor.
	constructor = gcm.Constructor().MakeWithAttributes(identifier, parameters, abstraction)
	return constructor, token, true
}

func (v *parser_) parseConstructors() (
	constructors col.ListLike[gcm.ConstructorLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	constructors = gcf.List[gcm.ConstructorLike]()
	_, token, ok = v.parseToken(NoteToken, "// Constructors")
	if !ok {
		// This is not a sequence of constructors.
		return constructors, token, false
	}

	// Attempt to parse at least one constructor.
	var constructor gcm.ConstructorLike
	constructor, token, ok = v.parseConstructor()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Constructor",
			"Constructors",
			"Constructor",
		)
		panic(message)
	}
	for ok {
		constructors.AppendValue(constructor)
		constructor, token, ok = v.parseConstructor()
	}

	// Found a sequence of constructors.
	return constructors, token, true
}

func (v *parser_) parseDeclaration() (
	declaration gcm.DeclarationLike,
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

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "type")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"type"`,
			"Declaration",
			"Parameters",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("identifier",
			"Declaration",
			"Parameters",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	var parameters = gcf.List[gcm.ParameterLike]()
	if ok {
		parameters, token, ok = v.parseParameters()
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("Parameters",
				"Declaration",
				"Parameters",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("]",
				"Declaration",
				"Parameters",
			)
			panic(message)
		}
	}

	// Found a declaration.
	declaration = gcm.Declaration().MakeWithAttributes(comment, identifier, parameters)
	return declaration, token, true
}

func (v *parser_) parseEnumeration() (
	enumeration gcm.EnumerationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "const")
	if !ok {
		// This is not an enumeration.
		return enumeration, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse a parameter.
	var parameter gcm.ParameterLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Parameter",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "=")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("=",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	_, token, ok = v.parseToken(IdentifierToken, "iota")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"iota"`,
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse a sequence of identifiers.
	var identifier string
	var identifiers = gcf.List[string]()
	identifier, _, ok = v.parseToken(IdentifierToken, "")
	for ok {
		identifiers.AppendValue(identifier)
		identifier, _, ok = v.parseToken(IdentifierToken, "")
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Found an enumeration.
	enumeration = gcm.Enumeration().MakeWithAttributes(parameter, identifiers)
	return enumeration, token, true
}

func (v *parser_) parseFunction() (
	function gcm.FunctionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a function.
		return function, token, false
	}

	// Attempt to parse a delimiter.
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

	// Attempt to parse a delimiter.
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
	var result gcm.ResultLike
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
	function = gcm.Function().MakeWithAttributes(identifier, parameters, result)
	return function, token, true
}

func (v *parser_) parseFunctional() (
	functional gcm.FunctionalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration gcm.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a functional.
		return functional, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "func")
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

	// Attempt to parse a delimiter.
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

	// Attempt to parse a delimiter.
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
	var result gcm.ResultLike
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
	functional = gcm.Functional().MakeWithAttributes(declaration, parameters, result)
	return functional, token, true
}

func (v *parser_) parseFunctionals() (
	functionals col.ListLike[gcm.FunctionalLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	functionals = gcf.List[gcm.FunctionalLike]()
	_, token, ok = v.parseToken(NoteToken, "// Functionals")
	if !ok {
		// This is not a sequence of functionals.
		return functionals, token, false
	}

	// Attempt to parse at least one functional.
	var functional gcm.FunctionalLike
	functional, token, ok = v.parseFunctional()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Functional",
			"Functionals",
			"Functional",
		)
		panic(message)
	}
	for ok {
		functionals.AppendValue(functional)
		functional, token, ok = v.parseFunctional()
	}

	// Found a sequence of functionals.
	return functionals, token, true
}

func (v *parser_) parseFunctions() (
	functions col.ListLike[gcm.FunctionLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	functions = gcf.List[gcm.FunctionLike]()
	_, token, ok = v.parseToken(NoteToken, "// Functions")
	if !ok {
		// This is not a sequence of functions.
		return functions, token, false
	}

	// Attempt to parse at least one function.
	var function gcm.FunctionLike
	function, token, ok = v.parseFunction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Function",
			"Functions",
			"Function",
		)
		panic(message)
	}
	for ok {
		functions.AppendValue(function)
		function, token, ok = v.parseFunction()
	}

	// Found a sequence of functions.
	return functions, token, true
}

func (v *parser_) parseHeader() (
	header gcm.HeaderLike,
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

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "package")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(`"package"`,
			"Header",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("identifier",
			"Header",
		)
		panic(message)
	}

	// Found a header.
	header = gcm.Header().MakeWithAttributes(comment, identifier)
	return header, token, true
}

func (v *parser_) parseInstance() (
	instance gcm.InstanceLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration gcm.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an instance.
		return instance, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
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

	// Attempt to parse a delimiter.
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

	// Attempt to parse an optional sequence of attributes.
	var attributes, _, _ = v.parseAttributes()

	// Attempt to parse an optional sequence of abstractions.
	var abstractions, _, _ = v.parseAbstractions()

	// Attempt to parse an optional sequence of methods.
	var methods, _, _ = v.parseMethods()

	// Attempt to parse a delimiter.
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
	instance = gcm.Instance().MakeWithAttributes(declaration, attributes, abstractions, methods)
	return instance, token, true
}

func (v *parser_) parseInstances() (
	instances col.ListLike[gcm.InstanceLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	instances = gcf.List[gcm.InstanceLike]()
	_, token, ok = v.parseToken(NoteToken, "// Instances")
	if !ok {
		// This is not a sequence of instances.
		return instances, token, false
	}

	// Attempt to parse at least one instance.
	var instance gcm.InstanceLike
	instance, token, ok = v.parseInstance()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Instance",
			"Instances",
			"Instance",
		)
		panic(message)
	}
	for ok {
		instances.AppendValue(instance)
		instance, token, ok = v.parseInstance()
	}

	// Found a sequence of instances.
	return instances, token, true
}

func (v *parser_) parseMethod() (
	method gcm.MethodLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a method.
		return method, token, false
	}

	// Attempt to parse a delimiter.
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

	// Attempt to parse a delimiter.
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
	var result, _, _ = v.parseResult()

	// Found a method.
	method = gcm.Method().MakeWithAttributes(identifier, parameters, result)
	return method, token, true
}

func (v *parser_) parseMethods() (
	methods col.ListLike[gcm.MethodLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	methods = gcf.List[gcm.MethodLike]()
	_, token, ok = v.parseToken(NoteToken, "// Methods")
	if !ok {
		// This is not a sequence of methods.
		return methods, token, false
	}

	// Attempt to parse at least one method.
	var method gcm.MethodLike
	method, token, ok = v.parseMethod()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Method",
			"Methods",
			"Method",
		)
		panic(message)
	}
	for ok {
		methods.AppendValue(method)
		method, token, ok = v.parseMethod()
	}

	// Found a sequence of methods.
	return methods, token, true
}

func (v *parser_) parseModel() (
	model gcm.ModelLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a notice.
	var notice gcm.NoticeLike
	notice, token, ok = v.parseNotice()
	if !ok {
		// This is not model.
		return model, token, false
	}

	// Attempt to parse a header.
	var header gcm.HeaderLike
	header, token, ok = v.parseHeader()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Header",
			"Model",
			"Notice",
			"Header",
			"Modules",
			"Types",
			"Functionals",
			"Aspects",
			"Classes",
			"Instances",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of modules.
	var modules, _, _ = v.parseModules()

	// Attempt to parse an optional sequence of types.
	var types, _, _ = v.parseTypes()

	// Attempt to parse an optional sequence of functionals.
	var functionals, _, _ = v.parseFunctionals()

	// Attempt to parse an optional sequence of aspects.
	var aspects, _, _ = v.parseAspects()

	// Attempt to parse an optional sequence of classes.
	var classes, _, _ = v.parseClasses()

	// Attempt to parse an optional sequence of instances.
	var instances, _, _ = v.parseInstances()

	// Found a model.
	model = gcm.Model().MakeWithAttributes(
		notice,
		header,
		modules,
		types,
		functionals,
		aspects,
		classes,
		instances,
	)
	return model, token, true
}

func (v *parser_) parseModule() (
	module gcm.ModuleLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a module.
		return module, token, false
	}

	// Attempt to parse text.
	var text string
	text, token, ok = v.parseToken(TextToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("text",
			"Module",
		)
		panic(message)
	}

	// Found a module.
	module = gcm.Module().MakeWithAttributes(identifier, text)
	return module, token, true
}

func (v *parser_) parseModules() (
	modules col.ListLike[gcm.ModuleLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a literal.
	modules = gcf.List[gcm.ModuleLike]()
	_, token, ok = v.parseToken(IdentifierToken, "import")
	if !ok {
		// This is not a sequence of modules.
		return modules, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("(",
			"Modules",
			"Module",
		)
		panic(message)
	}

	// Attempt to parse one or more modules.
	var module gcm.ModuleLike
	module, token, ok = v.parseModule()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Module",
			"Modules",
			"Module",
		)
		panic(message)
	}
	for ok {
		modules.AppendValue(module)
		module, _, ok = v.parseModule()
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax(")",
			"Modules",
			"Module",
		)
		panic(message)
	}

	// Found a sequence of modules.
	return modules, token, true
}

func (v *parser_) parseNotice() (
	notice gcm.NoticeLike,
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
	notice = gcm.Notice().MakeWithComment(comment)
	return notice, token, true
}

func (v *parser_) parseParameter() (
	parameter gcm.ParameterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		// This is not a parameter.
		return parameter, token, false
	}

	// Attempt to parse an abstraction.
	var abstraction gcm.AbstractionLike
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
	parameter = gcm.Parameter().MakeWithAttributes(identifier, abstraction)
	return parameter, token, true
}

func (v *parser_) parseParameters() (
	parameters col.ListLike[gcm.ParameterLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse at least one parameter.
	parameters = gcf.List[gcm.ParameterLike]()
	var parameter gcm.ParameterLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		// This is not a sequence of parameters.
		return parameters, token, false
	}
	for ok {
		parameters.AppendValue(parameter)
		_, token, ok = v.parseToken(DelimiterToken, ",")
		if ok {
			parameter, token, ok = v.parseParameter()
		}
	}

	// Found a sequence of parameters.
	return parameters, token, true
}

func (v *parser_) parsePrefix() (
	prefix gcm.PrefixLike,
	token TokenLike,
	ok bool,
) {
	var identifier string
	var prefixType gcm.PrefixType

	// Attempt to parse an array prefix.
	var delimiterToken TokenLike
	_, delimiterToken, ok = v.parseToken(DelimiterToken, "[")
	if ok {
		// Attempt to parse a delimiter.
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if ok {
			prefixType = gcm.ArrayPrefix
			prefix = gcm.Prefix().MakeWithAttributes(identifier, prefixType)
			return prefix, token, true
		}
		v.putBack(delimiterToken)
		return prefix, token, false
	}

	// Attempt to parse a map prefix.
	_, _, ok = v.parseToken(IdentifierToken, "map")
	if ok {
		_, token, ok = v.parseToken(DelimiterToken, "[")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("[",
				"Prefix",
			)
			panic(message)
		}
		identifier, token, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("identifier",
				"Prefix",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("]",
				"Prefix",
			)
			panic(message)
		}
		prefixType = gcm.MapPrefix
		prefix = gcm.Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// Attempt to parse a channel prefix.
	_, token, ok = v.parseToken(IdentifierToken, "chan")
	if ok {
		prefixType = gcm.ChannelPrefix
		prefix = gcm.Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// Attempt to parse an alias prefix.
	var identifierToken TokenLike
	identifier, identifierToken, ok = v.parseToken(IdentifierToken, "")
	if ok {
		_, token, ok = v.parseToken(DelimiterToken, ".")
		if !ok {
			v.putBack(identifierToken)
			return prefix, token, false
		}
		prefixType = gcm.AliasPrefix
		prefix = gcm.Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// This is not a prefix.
	return prefix, identifierToken, false
}

func (v *parser_) parseResult() (
	result gcm.ResultLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an abstraction.
	var abstraction gcm.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if ok {
		// Found an abstraction result.
		result = gcm.Result().MakeWithAbstraction(abstraction)
		return result, token, true
	}

	// Attempt to parse a sequence of parameters.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	var parameters col.ListLike[gcm.ParameterLike]
	if ok {
		parameters, token, ok = v.parseParameters()
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax("Parameters",
				"Result",
				"Abstraction",
				"Parameters",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, ")")
		if !ok {
			var message = v.formatError(token)
			message += v.generateSyntax(")",
				"Result",
				"Abstraction",
				"Parameters",
			)
			panic(message)
		}

		// Found a named parameters result.
		result = gcm.Result().MakeWithParameters(parameters)
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
	// Attempt to parse a specific token.
	token = v.getNextToken()
	value = token.GetValue()
	if token.GetType() == expectedType {
		var constrained = len(expectedValue) > 0
		if !constrained || value == expectedValue {
			// Found the expected token.
			return value, token, true
		}
	}

	// This is not the right token.
	v.putBack(token)
	return "", token, false
}

func (v *parser_) parseType() (
	type_ gcm.TypeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration gcm.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a type.
		return type_, token, false
	}

	// Attempt to parse an abstraction.
	var abstraction gcm.AbstractionLike
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
	var enumeration gcm.EnumerationLike
	enumeration, token, _ = v.parseEnumeration()

	// Found a type.
	type_ = gcm.Type().MakeWithAttributes(declaration, abstraction, enumeration)
	return type_, token, true
}

func (v *parser_) parseTypes() (
	types col.ListLike[gcm.TypeLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	types = gcf.List[gcm.TypeLike]()
	_, token, ok = v.parseToken(NoteToken, "// Types")
	if !ok {
		// This is not a sequence of types.
		return types, token, false
	}

	// Attempt to parse at least one type.
	var type_ gcm.TypeLike
	type_, token, ok = v.parseType()
	if !ok {
		var message = v.formatError(token)
		message += v.generateSyntax("Type",
			"Types",
			"Type",
		)
		panic(message)
	}
	for ok {
		types.AppendValue(type_)
		type_, token, ok = v.parseType()
	}

	// Found a sequence of types.
	return types, token, true
}

func (v *parser_) putBack(token TokenLike) {
	//fmt.Printf("Put Back %v\n", token)
	v.next_.AddValue(token)
}

var syntax = map[string]string{
	"Gcmn":        `Model EOL* EOF  ! Terminated with an end-of-file marker.`,
	"Model":       `Notice Header Modules? Types? Functionals? Aspects? Classes? Instances?`,
	"Notice":      `comment`,
	"Header":      `comment "package" identifier`,
	"Modules":     `"import" "(" Module+ ")"`,
	"Module":      `identifier text`,
	"Types":       `"// Types" Type+`,
	"Type":        `Declaration Abstraction Enumeration?`,
	"Declaration": `comment "type" identifier ("[" Parameters "]")?`,
	"Parameters":  `Parameter ("," Parameter)* ","?`,
	"Parameter":   `identifier Abstraction`,
	"Abstraction": `Prefix? identifier ("[" Arguments "]")?`,
	"Prefix": `
    "[" "]"
    "map" "[" identifier "]"
    "chan"
    identifier "."`,
	"Arguments":    `Abstraction ("," Abstraction)* ","?`,
	"Enumeration":  `"const" "(" Parameter "=" "iota" identifier* ")"`,
	"Functionals":  `"// Functionals" Functional+`,
	"Functional":   `Declaration "func" "(" Parameters? ")" Result`,
	"Result":       `Abstraction | "(" Parameters ")"`,
	"Aspects":      `"// Aspects" Aspect+`,
	"Aspect":       `Declaration "interface" "{" Methods? "}"`,
	"Classes":      `"// Classes" Class+`,
	"Class":        `Declaration "interface" "{" Constants? Constructors? Functions? "}"`,
	"Constants":    `"// Constants" Constant+`,
	"Constant":     `identifier "(" ")" Abstraction`,
	"Constructors": `"// Constructors" Constructor+`,
	"Constructor":  `identifier "(" Parameters? ")" Abstraction`,
	"Functions":    `"// Functions" Function+`,
	"Function":     `identifier "(" Parameters? ")" Result`,
	"Instances":    `"// Instances" Instance+`,
	"Instance":     `Declaration "interface" "{" Attributes? Abstractions? Methods? "}"`,
	"Attributes":   `"// Attributes" Attribute+`,
	"Attribute":    `identifier "(" Parameter? ")" Abstraction?`,
	"Abstractions": `"// Abstractions" Abstraction+`,
	"Methods":      `"// Methods" Method+`,
	"Method":       `identifier "(" Parameters? ")" Result?`,
}
