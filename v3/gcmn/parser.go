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
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v3"
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
		tokens_: col.Queue[TokenLike]().MakeWithCapacity(c.queueSize_),
		next_:   col.Stack[TokenLike]().MakeWithCapacity(c.stackSize_),
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Public

func (v *parser_) ParseSource(source string) ModelLike {
	// The scanner runs in a separate Go routine.
	v.source_ = source
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse a model.
	var model, token, ok = v.parseModel()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Model",
			"Source",
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
		message += v.generateGrammar("EOF",
			"Source",
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
func (v *parser_) generateGrammar(expected string, names ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, name := range names {
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			name,
			grammar[name],
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
	abstraction AbstractionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an optional prefix.
	var prefix PrefixLike
	prefix, _, ok = v.parsePrefix()
	var identifier string
	if ok {
		// Attempt to parse an identifier.
		identifier, token, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("Arguments",
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
	var arguments col.ListLike[AbstractionLike]
	if ok {
		// Attempt to parse a sequence of arguments.
		arguments, token, ok = v.parseArguments()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("Arguments",
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
			message += v.generateGrammar("]",
				"Abstraction",
				"Prefix",
				"Arguments",
			)
			panic(message)
		}
	}

	// Found an abstraction.
	abstraction = Abstraction().MakeWithAttributes(prefix, identifier, arguments)
	return abstraction, token, true
}

func (v *parser_) parseAbstractions() (
	abstractions col.ListLike[AbstractionLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Abstractions")
	if !ok {
		// This is not a sequence of abstractions.
		return abstractions, token, false
	}

	// Attempt to parse at least one abstraction.
	var abstraction AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Abstraction",
			"Abstractions",
			"Abstraction",
		)
		panic(message)
	}
	abstractions = col.List[AbstractionLike]().Make()
	for ok {
		abstractions.AppendValue(abstraction)
		abstraction, token, ok = v.parseAbstraction()
	}

	// Found a sequence of abstractions.
	return abstractions, token, true
}

func (v *parser_) parseArguments() (
	arguments col.ListLike[AbstractionLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse at least one abstraction.
	var abstraction AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		// This is not a sequence of arguments.
		return arguments, token, false
	}
	arguments = col.List[AbstractionLike]().Make()
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
	aspect AspectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an aspect.
		return aspect, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"interface"`,
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
		message += v.generateGrammar("{",
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
		message += v.generateGrammar("}",
			"Aspect",
			"Declaration",
			"Methods",
		)
		panic(message)
	}

	// Found an aspect.
	aspect = Aspect().MakeWithAttributes(declaration, methods)
	return aspect, token, true
}

func (v *parser_) parseAspects() (
	aspects col.ListLike[AspectLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Aspects")
	if !ok {
		// This is not a sequence of aspects.
		return aspects, token, false
	}

	// Attempt to parse at least one aspect.
	var aspect AspectLike
	aspect, token, ok = v.parseAspect()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Aspect",
			"Aspects",
			"Aspect",
		)
		panic(message)
	}
	aspects = col.List[AspectLike]().Make()
	for ok {
		aspects.AppendValue(aspect)
		aspect, token, ok = v.parseAspect()
	}

	// Found a sequence of aspects.
	return aspects, token, true
}

func (v *parser_) parseAttribute() (
	attribute AttributeLike,
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
		message += v.generateGrammar("(",
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
		message += v.generateGrammar(")",
			"Attribute",
			"Parameter",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an optional abstraction.
	var abstraction, _, _ = v.parseAbstraction()

	// Found a attribute.
	attribute = Attribute().MakeWithAttributes(identifier, parameter, abstraction)
	return attribute, token, true
}

func (v *parser_) parseAttributes() (
	attributes col.ListLike[AttributeLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Attributes")
	if !ok {
		// This is not a sequence of attributes.
		return attributes, token, false
	}

	// Attempt to parse at least one attribute.
	var attribute AttributeLike
	attribute, token, ok = v.parseAttribute()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Attribute",
			"Attributes",
			"Attribute",
		)
		panic(message)
	}
	attributes = col.List[AttributeLike]().Make()
	for ok {
		attributes.AppendValue(attribute)
		attribute, token, ok = v.parseAttribute()
	}

	// Found a sequence of attributes.
	return attributes, token, true
}

func (v *parser_) parseClass() (
	class ClassLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a class.
		return class, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"interface"`,
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
		message += v.generateGrammar("{",
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
		message += v.generateGrammar("}",
			"Class",
			"Declaration",
			"Constants",
			"Constructors",
			"Functions",
		)
		panic(message)
	}

	// Found a class.
	class = Class().MakeWithAttributes(declaration, constants, constructors, functions)
	return class, token, true
}

func (v *parser_) parseClasses() (
	classes col.ListLike[ClassLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Classes")
	if !ok {
		// This is not a sequence of classes.
		return classes, token, false
	}

	// Attempt to parse at least one class.
	var class ClassLike
	class, token, ok = v.parseClass()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Class",
			"Classes",
			"Class",
		)
		panic(message)
	}
	classes = col.List[ClassLike]().Make()
	for ok {
		classes.AppendValue(class)
		class, token, ok = v.parseClass()
	}

	// Found a sequence of classes.
	return classes, token, true
}

func (v *parser_) parseConstant() (
	constant ConstantLike,
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
		message += v.generateGrammar("(",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an abstraction.
	var abstraction AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Abstraction",
			"Constant",
			"Abstraction",
		)
		panic(message)
	}

	// Found a constant.
	constant = Constant().MakeWithAttributes(identifier, abstraction)
	return constant, token, true
}

func (v *parser_) parseConstants() (
	constants col.ListLike[ConstantLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Constants")
	if !ok {
		// This is not a sequence of constants.
		return constants, token, false
	}

	// Attempt to parse at least one constant.
	var constant ConstantLike
	constant, token, ok = v.parseConstant()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Constant",
			"Constants",
			"Constant",
		)
		panic(message)
	}
	constants = col.List[ConstantLike]().Make()
	for ok {
		constants.AppendValue(constant)
		constant, token, ok = v.parseConstant()
	}

	// Found a sequence of constants.
	return constants, token, true
}

func (v *parser_) parseConstructor() (
	constructor ConstructorLike,
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
		message += v.generateGrammar("(",
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
		message += v.generateGrammar(")",
			"Constructor",
			"Parameters",
			"Abstraction",
		)
		panic(message)
	}

	// Attempt to parse an abstraction.
	var abstraction AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Abstraction",
			"Constructor",
			"Parameters",
			"Abstraction",
		)
		panic(message)
	}

	// Found a constructor.
	constructor = Constructor().MakeWithAttributes(identifier, parameters, abstraction)
	return constructor, token, true
}

func (v *parser_) parseConstructors() (
	constructors col.ListLike[ConstructorLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Constructors")
	if !ok {
		// This is not a sequence of constructors.
		return constructors, token, false
	}

	// Attempt to parse at least one constructor.
	var constructor ConstructorLike
	constructor, token, ok = v.parseConstructor()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Constructor",
			"Constructors",
			"Constructor",
		)
		panic(message)
	}
	constructors = col.List[ConstructorLike]().Make()
	for ok {
		constructors.AppendValue(constructor)
		constructor, token, ok = v.parseConstructor()
	}

	// Found a sequence of constructors.
	return constructors, token, true
}

func (v *parser_) parseDeclaration() (
	declaration DeclarationLike,
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
		message += v.generateGrammar(`"type"`,
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
		message += v.generateGrammar("identifier",
			"Declaration",
			"Parameters",
		)
		panic(message)
	}

	// Attempt to parse an optional sequence of parameters.
	_, token, ok = v.parseToken(DelimiterToken, "[")
	var parameters col.ListLike[ParameterLike]
	if ok {
		parameters, token, ok = v.parseParameters()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("Parameters",
				"Declaration",
				"Parameters",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("]",
				"Declaration",
				"Parameters",
			)
			panic(message)
		}
	}

	// Found a declaration.
	declaration = Declaration().MakeWithAttributes(comment, identifier, parameters)
	return declaration, token, true
}

func (v *parser_) parseEnumeration() (
	enumeration EnumerationLike,
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
		message += v.generateGrammar("(",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse a parameter.
	var parameter ParameterLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Parameter",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "=")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("=",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	_, token, ok = v.parseToken(IdentifierToken, "iota")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"iota"`,
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Attempt to parse a sequence of identifiers.
	var identifier string
	var identifiers = col.List[string]().Make()
	identifier, _, ok = v.parseToken(IdentifierToken, "")
	for ok {
		identifiers.AppendValue(identifier)
		identifier, _, ok = v.parseToken(IdentifierToken, "")
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"Enumeration",
			"Parameter",
		)
		panic(message)
	}

	// Found an enumeration.
	enumeration = Enumeration().MakeWithAttributes(parameter, identifiers)
	return enumeration, token, true
}

func (v *parser_) parseFunction() (
	function FunctionLike,
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
		message += v.generateGrammar("(",
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
		message += v.generateGrammar(")",
			"Function",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse a result.
	var result ResultLike
	result, token, ok = v.parseResult()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Result",
			"Function",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Found a function.
	function = Function().MakeWithAttributes(identifier, parameters, result)
	return function, token, true
}

func (v *parser_) parseFunctional() (
	functional FunctionalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a functional.
		return functional, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "func")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"func"`,
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
		message += v.generateGrammar("(",
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
		message += v.generateGrammar(")",
			"Functional",
			"Declaration",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse a result.
	var result ResultLike
	result, token, ok = v.parseResult()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Result",
			"Functional",
			"Declaration",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Found a functional.
	functional = Functional().MakeWithAttributes(declaration, parameters, result)
	return functional, token, true
}

func (v *parser_) parseFunctionals() (
	functionals col.ListLike[FunctionalLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Functionals")
	if !ok {
		// This is not a sequence of functionals.
		return functionals, token, false
	}

	// Attempt to parse at least one functional.
	var functional FunctionalLike
	functional, token, ok = v.parseFunctional()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Functional",
			"Functionals",
			"Functional",
		)
		panic(message)
	}
	functionals = col.List[FunctionalLike]().Make()
	for ok {
		functionals.AppendValue(functional)
		functional, token, ok = v.parseFunctional()
	}

	// Found a sequence of functionals.
	return functionals, token, true
}

func (v *parser_) parseFunctions() (
	functions col.ListLike[FunctionLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Functions")
	if !ok {
		// This is not a sequence of functions.
		return functions, token, false
	}

	// Attempt to parse at least one function.
	var function FunctionLike
	function, token, ok = v.parseFunction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Function",
			"Functions",
			"Function",
		)
		panic(message)
	}
	functions = col.List[FunctionLike]().Make()
	for ok {
		functions.AppendValue(function)
		function, token, ok = v.parseFunction()
	}

	// Found a sequence of functions.
	return functions, token, true
}

func (v *parser_) parseHeader() (
	header HeaderLike,
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
		message += v.generateGrammar(`"package"`,
			"Header",
		)
		panic(message)
	}

	// Attempt to parse an identifier.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken, "")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("identifier",
			"Header",
		)
		panic(message)
	}

	// Found a header.
	header = Header().MakeWithAttributes(comment, identifier)
	return header, token, true
}

func (v *parser_) parseInstance() (
	instance InstanceLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not an instance.
		return instance, token, false
	}

	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "interface")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(`"interface"`,
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
		message += v.generateGrammar("{",
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
		message += v.generateGrammar("}",
			"Instance",
			"Declaration",
			"Attributes",
			"Abstractions",
			"Methods",
		)
		panic(message)
	}

	// Found an instance.
	instance = Instance().MakeWithAttributes(declaration, attributes, abstractions, methods)
	return instance, token, true
}

func (v *parser_) parseInstances() (
	instances col.ListLike[InstanceLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Instances")
	if !ok {
		// This is not a sequence of instances.
		return instances, token, false
	}

	// Attempt to parse at least one instance.
	var instance InstanceLike
	instance, token, ok = v.parseInstance()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Instance",
			"Instances",
			"Instance",
		)
		panic(message)
	}
	instances = col.List[InstanceLike]().Make()
	for ok {
		instances.AppendValue(instance)
		instance, token, ok = v.parseInstance()
	}

	// Found a sequence of instances.
	return instances, token, true
}

func (v *parser_) parseMethod() (
	method MethodLike,
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
		message += v.generateGrammar("(",
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
		message += v.generateGrammar(")",
			"Method",
			"Parameters",
			"Result",
		)
		panic(message)
	}

	// Attempt to parse an optional result.
	var result, _, _ = v.parseResult()

	// Found a method.
	method = Method().MakeWithAttributes(identifier, parameters, result)
	return method, token, true
}

func (v *parser_) parseMethods() (
	methods col.ListLike[MethodLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Methods")
	if !ok {
		// This is not a sequence of methods.
		return methods, token, false
	}

	// Attempt to parse at least one method.
	var method MethodLike
	method, token, ok = v.parseMethod()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Method",
			"Methods",
			"Method",
		)
		panic(message)
	}
	methods = col.List[MethodLike]().Make()
	for ok {
		methods.AppendValue(method)
		method, token, ok = v.parseMethod()
	}

	// Found a sequence of methods.
	return methods, token, true
}

func (v *parser_) parseModel() (
	model ModelLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a notice.
	var notice NoticeLike
	notice, token, ok = v.parseNotice()
	if !ok {
		// This is not model.
		return model, token, false
	}

	// Attempt to parse a header.
	var header HeaderLike
	header, token, ok = v.parseHeader()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Header",
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
	model = Model().MakeWithAttributes(
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
	module ModuleLike,
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
		message += v.generateGrammar("text",
			"Module",
		)
		panic(message)
	}

	// Found a module.
	module = Module().MakeWithAttributes(identifier, text)
	return module, token, true
}

func (v *parser_) parseModules() (
	modules col.ListLike[ModuleLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a literal.
	_, token, ok = v.parseToken(IdentifierToken, "import")
	if !ok {
		// This is not a sequence of modules.
		return modules, token, false
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("(",
			"Modules",
			"Module",
		)
		panic(message)
	}

	// Attempt to parse one or more modules.
	var module ModuleLike
	module, token, ok = v.parseModule()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Module",
			"Modules",
			"Module",
		)
		panic(message)
	}
	modules = col.List[ModuleLike]().Make()
	for ok {
		modules.AppendValue(module)
		module, _, ok = v.parseModule()
	}

	// Attempt to parse a delimiter.
	_, token, ok = v.parseToken(DelimiterToken, ")")
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar(")",
			"Modules",
			"Module",
		)
		panic(message)
	}

	// Found a sequence of modules.
	return modules, token, true
}

func (v *parser_) parseNotice() (
	notice NoticeLike,
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
	notice = Notice().MakeWithComment(comment)
	return notice, token, true
}

func (v *parser_) parseParameter() (
	parameter ParameterLike,
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
	var abstraction AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Abstraction",
			"Parameter",
			"Abstraction",
		)
		panic(message)
	}

	// Found a parameter.
	parameter = Parameter().MakeWithAttributes(identifier, abstraction)
	return parameter, token, true
}

func (v *parser_) parseParameters() (
	parameters col.ListLike[ParameterLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse at least one parameter.
	var parameter ParameterLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		// This is not a sequence of parameters.
		return parameters, token, false
	}
	parameters = col.List[ParameterLike]().Make()
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
	prefix PrefixLike,
	token TokenLike,
	ok bool,
) {
	var identifier string
	var prefixType PrefixType

	// Attempt to parse an array prefix.
	var delimiterToken TokenLike
	_, delimiterToken, ok = v.parseToken(DelimiterToken, "[")
	if ok {
		// Attempt to parse a delimiter.
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if ok {
			prefixType = ArrayPrefix
			prefix = Prefix().MakeWithAttributes(identifier, prefixType)
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
			message += v.generateGrammar("[",
				"Prefix",
			)
			panic(message)
		}
		identifier, token, ok = v.parseToken(IdentifierToken, "")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("identifier",
				"Prefix",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, "]")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("]",
				"Prefix",
			)
			panic(message)
		}
		prefixType = MapPrefix
		prefix = Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// Attempt to parse a channel prefix.
	_, token, ok = v.parseToken(IdentifierToken, "chan")
	if ok {
		prefixType = ChannelPrefix
		prefix = Prefix().MakeWithAttributes(identifier, prefixType)
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
		prefixType = AliasPrefix
		prefix = Prefix().MakeWithAttributes(identifier, prefixType)
		return prefix, token, true
	}

	// This is not a prefix.
	return prefix, identifierToken, false
}

func (v *parser_) parseResult() (
	result ResultLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse an abstraction.
	var abstraction AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if ok {
		// Found an abstraction result.
		result = Result().MakeWithAbstraction(abstraction)
		return result, token, true
	}

	// Attempt to parse a sequence of parameters.
	_, token, ok = v.parseToken(DelimiterToken, "(")
	var parameters col.ListLike[ParameterLike]
	if ok {
		parameters, token, ok = v.parseParameters()
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar("Parameters",
				"Result",
				"Abstraction",
				"Parameters",
			)
			panic(message)
		}
		_, token, ok = v.parseToken(DelimiterToken, ")")
		if !ok {
			var message = v.formatError(token)
			message += v.generateGrammar(")",
				"Result",
				"Abstraction",
				"Parameters",
			)
			panic(message)
		}

		// Found a named parameters result.
		result = Result().MakeWithParameters(parameters)
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
	type_ TypeLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a declaration.
	var declaration DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		// This is not a type.
		return type_, token, false
	}

	// Attempt to parse an abstraction.
	var abstraction AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Abstraction",
			"Type",
			"Declaration",
			"Abstraction",
			"Enumeration",
		)
		panic(message)
	}

	// Attempt to parse an optional enumeration.
	var enumeration EnumerationLike
	enumeration, token, _ = v.parseEnumeration()

	// Found a type.
	type_ = Type().MakeWithAttributes(declaration, abstraction, enumeration)
	return type_, token, true
}

func (v *parser_) parseTypes() (
	types col.ListLike[TypeLike],
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a note.
	_, token, ok = v.parseToken(NoteToken, "// Types")
	if !ok {
		// This is not a sequence of types.
		return types, token, false
	}

	// Attempt to parse at least one type.
	var type_ TypeLike
	type_, token, ok = v.parseType()
	if !ok {
		var message = v.formatError(token)
		message += v.generateGrammar("Type",
			"Types",
			"Type",
		)
		panic(message)
	}
	types = col.List[TypeLike]().Make()
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

var grammar = map[string]string{
	"Source":      `Model EOL* EOF  ! Terminated with an end-of-file marker.`,
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
