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

// nolint
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

var parserClass = &parserClass_{
	// Initialize the class constants.
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
	// Define the class constants.
	queueSize_ uint
	stackSize_ uint
}

// Constructors

func (c *parserClass_) Make() ParserLike {
	return &parser_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type parser_ struct {
	// Define the instance attributes.
	class_  *parserClass_
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Public

func (v *parser_) GetClass() ParserClassLike {
	return v.class_
}

func (v *parser_) ParseSource(source string) ast.ModelLike {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the model.
	var model, token, ok = v.parseModel()
	if !ok {
		var message = v.formatError(token, "Model")
		panic(message)
	}

	// Found the model.
	return model
}

// Private

func (v *parser_) parseAbstraction() (
	abstraction ast.AbstractionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse an optional prefix rule.
	var optionalPrefix ast.PrefixLike
	optionalPrefix, _, ok = v.parsePrefix()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Abstraction")
			panic(message)
		} else {
			// This is not a single abstraction rule.
			return abstraction, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional suffix rule.
	var optionalSuffix ast.SuffixLike
	optionalSuffix, _, ok = v.parseSuffix()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse an optional arguments rule.
	var optionalArguments ast.ArgumentsLike
	optionalArguments, _, ok = v.parseArguments()
	if ok {
		ruleFound_ = true
	}

	// Found a single abstraction rule.
	ruleFound_ = true
	abstraction = ast.Abstraction().Make(
		optionalPrefix,
		name,
		optionalSuffix,
		optionalArguments,
	)
	return abstraction, token, ruleFound_
}

func (v *parser_) parseAdditionalArgument() (
	additionalArgument ast.AdditionalArgumentLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AdditionalArgument")
			panic(message)
		} else {
			// This is not a single additionalArgument rule.
			return additionalArgument, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AdditionalArgument")
			panic(message)
		} else {
			// This is not a single additionalArgument rule.
			return additionalArgument, token, false
		}
	}
	ruleFound_ = true

	// Found a single additionalArgument rule.
	ruleFound_ = true
	additionalArgument = ast.AdditionalArgument().Make(argument)
	return additionalArgument, token, ruleFound_
}

func (v *parser_) parseAdditionalConstraint() (
	additionalConstraint ast.AdditionalConstraintLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AdditionalConstraint")
			panic(message)
		} else {
			// This is not a single additionalConstraint rule.
			return additionalConstraint, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single constraint rule.
	var constraint ast.ConstraintLike
	constraint, token, ok = v.parseConstraint()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AdditionalConstraint")
			panic(message)
		} else {
			// This is not a single additionalConstraint rule.
			return additionalConstraint, token, false
		}
	}
	ruleFound_ = true

	// Found a single additionalConstraint rule.
	ruleFound_ = true
	additionalConstraint = ast.AdditionalConstraint().Make(constraint)
	return additionalConstraint, token, ruleFound_
}

func (v *parser_) parseAdditionalValue() (
	additionalValue ast.AdditionalValueLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AdditionalValue")
			panic(message)
		} else {
			// This is not a single additionalValue rule.
			return additionalValue, token, false
		}
	}
	ruleFound_ = true

	// Found a single additionalValue rule.
	ruleFound_ = true
	additionalValue = ast.AdditionalValue().Make(name)
	return additionalValue, token, ruleFound_
}

func (v *parser_) parseArgument() (
	argument ast.ArgumentLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Argument")
			panic(message)
		} else {
			// This is not a single argument rule.
			return argument, token, false
		}
	}
	ruleFound_ = true

	// Found a single argument rule.
	ruleFound_ = true
	argument = ast.Argument().Make(abstraction)
	return argument, token, ruleFound_
}

func (v *parser_) parseArray() (
	array ast.ArrayLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Array")
			panic(message)
		} else {
			// This is not a single array rule.
			return array, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Array")
			panic(message)
		} else {
			// This is not a single array rule.
			return array, token, false
		}
	}
	ruleFound_ = true

	// Found a single array rule.
	ruleFound_ = true
	array = ast.Array().Make()
	return array, token, ruleFound_
}

func (v *parser_) parseAspect() (
	aspect ast.AspectLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Aspect")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspect, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Aspect")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspect, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Aspect")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspect, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited method rules.
	var methods = col.List[ast.MethodLike]()
methodsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var method ast.MethodLike
		method, token, ok = v.parseMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single aspect rule.
					return aspect, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Aspect")
				message += "The number of method rules must be at least 1."
				panic(message)
			default:
				break methodsLoop
			}
		}
		methods.AppendValue(method)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Aspect")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspect, token, false
		}
	}
	ruleFound_ = true

	// Found a single aspect rule.
	ruleFound_ = true
	aspect = ast.Aspect().Make(
		declaration,
		methods,
	)
	return aspect, token, ruleFound_
}

func (v *parser_) parseAspectDefinitions() (
	aspectDefinitions ast.AspectDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Aspects" delimiter.
	_, token, ok = v.parseDelimiter("// Aspects")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectDefinitions")
			panic(message)
		} else {
			// This is not a single aspectDefinitions rule.
			return aspectDefinitions, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited aspect rules.
	var aspects = col.List[ast.AspectLike]()
aspectsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var aspect ast.AspectLike
		aspect, token, ok = v.parseAspect()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single aspectDefinitions rule.
					return aspectDefinitions, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "AspectDefinitions")
				message += "The number of aspect rules must be at least 1."
				panic(message)
			default:
				break aspectsLoop
			}
		}
		aspects.AppendValue(aspect)
	}

	// Found a single aspectDefinitions rule.
	ruleFound_ = true
	aspectDefinitions = ast.AspectDefinitions().Make(aspects)
	return aspectDefinitions, token, ruleFound_
}

func (v *parser_) parseAspectInterfaces() (
	aspectInterfaces ast.AspectInterfacesLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Aspect" delimiter.
	_, token, ok = v.parseDelimiter("// Aspect")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectInterfaces")
			panic(message)
		} else {
			// This is not a single aspectInterfaces rule.
			return aspectInterfaces, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited interface rules.
	var interfaces = col.List[ast.InterfaceLike]()
interfacesLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var interface_ ast.InterfaceLike
		interface_, token, ok = v.parseInterface()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single aspectInterfaces rule.
					return aspectInterfaces, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "AspectInterfaces")
				message += "The number of interface rules must be at least 1."
				panic(message)
			default:
				break interfacesLoop
			}
		}
		interfaces.AppendValue(interface_)
	}

	// Found a single aspectInterfaces rule.
	ruleFound_ = true
	aspectInterfaces = ast.AspectInterfaces().Make(interfaces)
	return aspectInterfaces, token, ruleFound_
}

func (v *parser_) parseAccessor() (
	accessor ast.AccessorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single getter rule.
	var getter ast.GetterLike
	getter, token, ok = v.parseGetter()
	if ok {
		// Found a single getter accessor.
		accessor = ast.Accessor().Make(getter)
		return accessor, token, true
	}

	// Attempt to parse a single setter rule.
	var setter ast.SetterLike
	setter, token, ok = v.parseSetter()
	if ok {
		// Found a single setter accessor.
		accessor = ast.Accessor().Make(setter)
		return accessor, token, true
	}

	// This is not a single accessor rule.
	return accessor, token, false
}

func (v *parser_) parseGetter() (
	getter ast.GetterLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	var nameToken TokenLike
	name, nameToken, ok = v.parseToken(NameToken)
	if !ok {
		// This is not a single getter rule.
		return getter, nameToken, false
	}

	// Attempt to parse a single "(" delimiter.
	var leftToken TokenLike
	_, leftToken, ok = v.parseDelimiter("(")
	if !ok {
		// This is not a single getter rule.
		v.putBack(nameToken)
		return getter, leftToken, false
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		// This is not a single getter rule.
		v.putBack(leftToken)
		v.putBack(nameToken)
		return getter, token, false
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Getter")
			panic(message)
		} else {
			// This is not a single getter rule.
			return getter, token, false
		}
	}
	ruleFound_ = true

	// Found a single getter rule.
	ruleFound_ = true
	getter = ast.Getter().Make(
		name,
		abstraction,
	)
	return getter, token, ruleFound_
}

func (v *parser_) parseSetter() (
	setter ast.SetterLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Setter")
			panic(message)
		} else {
			// This is not a single setter rule.
			return setter, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Setter")
			panic(message)
		} else {
			// This is not a single setter rule.
			return setter, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single parameter rule.
	var parameter ast.ParameterLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Setter")
			panic(message)
		} else {
			// This is not a single setter rule.
			return setter, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Setter")
			panic(message)
		} else {
			// This is not a single setter rule.
			return setter, token, false
		}
	}
	ruleFound_ = true

	// Found a single setter rule.
	setter = ast.Setter().Make(
		name,
		parameter,
	)
	return setter, token, ruleFound_
}

func (v *parser_) parseAttributeMethods() (
	attributeMethods ast.AttributeMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Attribute" delimiter.
	_, token, ok = v.parseDelimiter("// Attribute")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AttributeMethods")
			panic(message)
		} else {
			// This is not a single attributeMethods rule.
			return attributeMethods, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited accessor rules.
	var accessors = col.List[ast.AccessorLike]()
accessorsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var accessor ast.AccessorLike
		accessor, token, ok = v.parseAccessor()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single attributeMethods rule.
					return attributeMethods, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "AttributeMethods")
				message += "The number of accessor rules must be at least 1."
				panic(message)
			default:
				break accessorsLoop
			}
		}
		accessors.AppendValue(accessor)
	}

	// Found a single attributeMethods rule.
	ruleFound_ = true
	attributeMethods = ast.AttributeMethods().Make(accessors)
	return attributeMethods, token, ruleFound_
}

func (v *parser_) parseChannel() (
	channel ast.ChannelLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "chan" delimiter.
	_, token, ok = v.parseDelimiter("chan")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Channel")
			panic(message)
		} else {
			// This is not a single channel rule.
			return channel, token, false
		}
	}
	ruleFound_ = true

	// Found a single channel rule.
	ruleFound_ = true
	channel = ast.Channel().Make()
	return channel, token, ruleFound_
}

func (v *parser_) parseClass() (
	class ast.ClassLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Class")
			panic(message)
		} else {
			// This is not a single class rule.
			return class, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Class")
			panic(message)
		} else {
			// This is not a single class rule.
			return class, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Class")
			panic(message)
		} else {
			// This is not a single class rule.
			return class, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single classMethods rule.
	var classMethods ast.ClassMethodsLike
	classMethods, token, ok = v.parseClassMethods()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Class")
			panic(message)
		} else {
			// This is not a single class rule.
			return class, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Class")
			panic(message)
		} else {
			// This is not a single class rule.
			return class, token, false
		}
	}
	ruleFound_ = true

	// Found a single class rule.
	ruleFound_ = true
	class = ast.Class().Make(
		declaration,
		classMethods,
	)
	return class, token, ruleFound_
}

func (v *parser_) parseClassDefinitions() (
	classDefinitions ast.ClassDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Classes" delimiter.
	_, token, ok = v.parseDelimiter("// Classes")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ClassDefinitions")
			panic(message)
		} else {
			// This is not a single classDefinitions rule.
			return classDefinitions, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited class rules.
	var classes = col.List[ast.ClassLike]()
classesLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var class ast.ClassLike
		class, token, ok = v.parseClass()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single classDefinitions rule.
					return classDefinitions, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "ClassDefinitions")
				message += "The number of class rules must be at least 1."
				panic(message)
			default:
				break classesLoop
			}
		}
		classes.AppendValue(class)
	}

	// Found a single classDefinitions rule.
	ruleFound_ = true
	classDefinitions = ast.ClassDefinitions().Make(classes)
	return classDefinitions, token, ruleFound_
}

func (v *parser_) parseClassMethods() (
	classMethods ast.ClassMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single constructorMethods rule.
	var constructorMethods ast.ConstructorMethodsLike
	constructorMethods, token, ok = v.parseConstructorMethods()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ClassMethods")
			panic(message)
		} else {
			// This is not a single classMethods rule.
			return classMethods, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional constantMethods rule.
	var optionalConstantMethods ast.ConstantMethodsLike
	optionalConstantMethods, _, ok = v.parseConstantMethods()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse an optional functionMethods rule.
	var optionalFunctionMethods ast.FunctionMethodsLike
	optionalFunctionMethods, _, ok = v.parseFunctionMethods()
	if ok {
		ruleFound_ = true
	}

	// Found a single classMethods rule.
	ruleFound_ = true
	classMethods = ast.ClassMethods().Make(
		constructorMethods,
		optionalConstantMethods,
		optionalFunctionMethods,
	)
	return classMethods, token, ruleFound_
}

func (v *parser_) parseConstant() (
	constant ast.ConstantLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constant")
			panic(message)
		} else {
			// This is not a single constant rule.
			return constant, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constant")
			panic(message)
		} else {
			// This is not a single constant rule.
			return constant, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constant")
			panic(message)
		} else {
			// This is not a single constant rule.
			return constant, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constant")
			panic(message)
		} else {
			// This is not a single constant rule.
			return constant, token, false
		}
	}
	ruleFound_ = true

	// Found a single constant rule.
	ruleFound_ = true
	constant = ast.Constant().Make(
		name,
		abstraction,
	)
	return constant, token, ruleFound_
}

func (v *parser_) parseConstantMethods() (
	constantMethods ast.ConstantMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Constant" delimiter.
	_, token, ok = v.parseDelimiter("// Constant")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ConstantMethods")
			panic(message)
		} else {
			// This is not a single constantMethods rule.
			return constantMethods, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited constant rules.
	var constants = col.List[ast.ConstantLike]()
constantsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var constant ast.ConstantLike
		constant, token, ok = v.parseConstant()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single constantMethods rule.
					return constantMethods, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "ConstantMethods")
				message += "The number of constant rules must be at least 1."
				panic(message)
			default:
				break constantsLoop
			}
		}
		constants.AppendValue(constant)
	}

	// Found a single constantMethods rule.
	ruleFound_ = true
	constantMethods = ast.ConstantMethods().Make(constants)
	return constantMethods, token, ruleFound_
}

func (v *parser_) parseConstructor() (
	constructor ast.ConstructorLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constructor")
			panic(message)
		} else {
			// This is not a single constructor rule.
			return constructor, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constructor")
			panic(message)
		} else {
			// This is not a single constructor rule.
			return constructor, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single constructor rule.
					return constructor, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Constructor")
				message += "The number of parameter rules must be at least 0."
				panic(message)
			default:
				break parametersLoop
			}
		}
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constructor")
			panic(message)
		} else {
			// This is not a single constructor rule.
			return constructor, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constructor")
			panic(message)
		} else {
			// This is not a single constructor rule.
			return constructor, token, false
		}
	}
	ruleFound_ = true

	// Found a single constructor rule.
	ruleFound_ = true
	constructor = ast.Constructor().Make(
		name,
		parameters,
		abstraction,
	)
	return constructor, token, ruleFound_
}

func (v *parser_) parseConstructorMethods() (
	constructorMethods ast.ConstructorMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Constructor" delimiter.
	_, token, ok = v.parseDelimiter("// Constructor")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ConstructorMethods")
			panic(message)
		} else {
			// This is not a single constructorMethods rule.
			return constructorMethods, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited constructor rules.
	var constructors = col.List[ast.ConstructorLike]()
constructorsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var constructor ast.ConstructorLike
		constructor, token, ok = v.parseConstructor()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single constructorMethods rule.
					return constructorMethods, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "ConstructorMethods")
				message += "The number of constructor rules must be at least 1."
				panic(message)
			default:
				break constructorsLoop
			}
		}
		constructors.AppendValue(constructor)
	}

	// Found a single constructorMethods rule.
	ruleFound_ = true
	constructorMethods = ast.ConstructorMethods().Make(constructors)
	return constructorMethods, token, ruleFound_
}

func (v *parser_) parseDeclaration() (
	declaration ast.DeclarationLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Declaration")
			panic(message)
		} else {
			// This is not a single declaration rule.
			return declaration, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "type" delimiter.
	_, token, ok = v.parseDelimiter("type")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Declaration")
			panic(message)
		} else {
			// This is not a single declaration rule.
			return declaration, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Declaration")
			panic(message)
		} else {
			// This is not a single declaration rule.
			return declaration, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional constraints rule.
	var optionalConstraints ast.ConstraintsLike
	optionalConstraints, _, ok = v.parseConstraints()
	if ok {
		ruleFound_ = true
	}

	// Found a single declaration rule.
	ruleFound_ = true
	declaration = ast.Declaration().Make(
		comment,
		name,
		optionalConstraints,
	)
	return declaration, token, ruleFound_
}

func (v *parser_) parseEnumeration() (
	enumeration ast.EnumerationLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "const" delimiter.
	_, token, ok = v.parseDelimiter("const")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Enumeration")
			panic(message)
		} else {
			// This is not a single enumeration rule.
			return enumeration, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Enumeration")
			panic(message)
		} else {
			// This is not a single enumeration rule.
			return enumeration, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single value rule.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Enumeration")
			panic(message)
		} else {
			// This is not a single enumeration rule.
			return enumeration, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited additionalValue rules.
	var additionalValues = col.List[ast.AdditionalValueLike]()
additionalValuesLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var additionalValue ast.AdditionalValueLike
		additionalValue, token, ok = v.parseAdditionalValue()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single enumeration rule.
					return enumeration, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Enumeration")
				message += "The number of additionalValue rules must be at least 0."
				panic(message)
			default:
				break additionalValuesLoop
			}
		}
		additionalValues.AppendValue(additionalValue)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Enumeration")
			panic(message)
		} else {
			// This is not a single enumeration rule.
			return enumeration, token, false
		}
	}
	ruleFound_ = true

	// Found a single enumeration rule.
	ruleFound_ = true
	enumeration = ast.Enumeration().Make(
		value,
		additionalValues,
	)
	return enumeration, token, ruleFound_
}

func (v *parser_) parseFunction() (
	function ast.FunctionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Function")
			panic(message)
		} else {
			// This is not a single function rule.
			return function, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Function")
			panic(message)
		} else {
			// This is not a single function rule.
			return function, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single function rule.
					return function, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Function")
				message += "The number of parameter rules must be at least 0."
				panic(message)
			default:
				break parametersLoop
			}
		}
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Function")
			panic(message)
		} else {
			// This is not a single function rule.
			return function, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single result rule.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Function")
			panic(message)
		} else {
			// This is not a single function rule.
			return function, token, false
		}
	}
	ruleFound_ = true

	// Found a single function rule.
	ruleFound_ = true
	function = ast.Function().Make(
		name,
		parameters,
		result,
	)
	return function, token, ruleFound_
}

func (v *parser_) parseFunctionMethods() (
	functionMethods ast.FunctionMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Function" delimiter.
	_, token, ok = v.parseDelimiter("// Function")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "FunctionMethods")
			panic(message)
		} else {
			// This is not a single functionMethods rule.
			return functionMethods, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited function rules.
	var functions = col.List[ast.FunctionLike]()
functionsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var function ast.FunctionLike
		function, token, ok = v.parseFunction()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single functionMethods rule.
					return functionMethods, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "FunctionMethods")
				message += "The number of function rules must be at least 1."
				panic(message)
			default:
				break functionsLoop
			}
		}
		functions.AppendValue(function)
	}

	// Found a single functionMethods rule.
	ruleFound_ = true
	functionMethods = ast.FunctionMethods().Make(functions)
	return functionMethods, token, ruleFound_
}

func (v *parser_) parseFunctional() (
	functional ast.FunctionalLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Functional")
			panic(message)
		} else {
			// This is not a single functional rule.
			return functional, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "func" delimiter.
	_, token, ok = v.parseDelimiter("func")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Functional")
			panic(message)
		} else {
			// This is not a single functional rule.
			return functional, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Functional")
			panic(message)
		} else {
			// This is not a single functional rule.
			return functional, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single functional rule.
					return functional, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Functional")
				message += "The number of parameter rules must be at least 0."
				panic(message)
			default:
				break parametersLoop
			}
		}
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Functional")
			panic(message)
		} else {
			// This is not a single functional rule.
			return functional, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single result rule.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Functional")
			panic(message)
		} else {
			// This is not a single functional rule.
			return functional, token, false
		}
	}
	ruleFound_ = true

	// Found a single functional rule.
	ruleFound_ = true
	functional = ast.Functional().Make(
		declaration,
		parameters,
		result,
	)
	return functional, token, ruleFound_
}

func (v *parser_) parseFunctionalDefinitions() (
	functionalDefinitions ast.FunctionalDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Functionals" delimiter.
	_, token, ok = v.parseDelimiter("// Functionals")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "FunctionalDefinitions")
			panic(message)
		} else {
			// This is not a single functionalDefinitions rule.
			return functionalDefinitions, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited functional rules.
	var functionals = col.List[ast.FunctionalLike]()
functionalsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var functional ast.FunctionalLike
		functional, token, ok = v.parseFunctional()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single functionalDefinitions rule.
					return functionalDefinitions, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "FunctionalDefinitions")
				message += "The number of functional rules must be at least 1."
				panic(message)
			default:
				break functionalsLoop
			}
		}
		functionals.AppendValue(functional)
	}

	// Found a single functionalDefinitions rule.
	ruleFound_ = true
	functionalDefinitions = ast.FunctionalDefinitions().Make(functionals)
	return functionalDefinitions, token, ruleFound_
}

func (v *parser_) parseArguments() (
	arguments ast.ArgumentsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Arguments")
			panic(message)
		} else {
			// This is not a single arguments rule.
			return arguments, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Arguments")
			panic(message)
		} else {
			// This is not a single arguments rule.
			return arguments, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited additionalArgument rules.
	var additionalArguments = col.List[ast.AdditionalArgumentLike]()
additionalArgumentsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var additionalArgument ast.AdditionalArgumentLike
		additionalArgument, token, ok = v.parseAdditionalArgument()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single arguments rule.
					return arguments, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Arguments")
				message += "The number of additionalArgument rules must be at least 0."
				panic(message)
			default:
				break additionalArgumentsLoop
			}
		}
		additionalArguments.AppendValue(additionalArgument)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Arguments")
			panic(message)
		} else {
			// This is not a single arguments rule.
			return arguments, token, false
		}
	}
	ruleFound_ = true

	// Found a single arguments rule.
	ruleFound_ = true
	arguments = ast.Arguments().Make(
		argument,
		additionalArguments,
	)
	return arguments, token, ruleFound_
}

func (v *parser_) parseConstraint() (
	constraint ast.ConstraintLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constraint")
			panic(message)
		} else {
			// This is not a single constraint rule.
			return constraint, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constraint")
			panic(message)
		} else {
			// This is not a single constraint rule.
			return constraint, token, false
		}
	}
	ruleFound_ = true

	// Found a single constraint rule.
	ruleFound_ = true
	constraint = ast.Constraint().Make(
		name,
		abstraction,
	)
	return constraint, token, ruleFound_
}

func (v *parser_) parseConstraints() (
	constraints ast.ConstraintsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constraints")
			panic(message)
		} else {
			// This is not a single constraints rule.
			return constraints, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single constraint rule.
	var constraint ast.ConstraintLike
	constraint, token, ok = v.parseConstraint()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constraints")
			panic(message)
		} else {
			// This is not a single constraints rule.
			return constraints, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited additionalConstraint rules.
	var additionalConstraints = col.List[ast.AdditionalConstraintLike]()
additionalConstraintsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var additionalConstraint ast.AdditionalConstraintLike
		additionalConstraint, token, ok = v.parseAdditionalConstraint()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single constraints rule.
					return constraints, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Constraints")
				message += "The number of additionalConstraint rules must be at least 0."
				panic(message)
			default:
				break additionalConstraintsLoop
			}
		}
		additionalConstraints.AppendValue(additionalConstraint)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Constraints")
			panic(message)
		} else {
			// This is not a single constraints rule.
			return constraints, token, false
		}
	}
	ruleFound_ = true

	// Found a single constraints rule.
	ruleFound_ = true
	constraints = ast.Constraints().Make(
		constraint,
		additionalConstraints,
	)
	return constraints, token, ruleFound_
}

func (v *parser_) parseHeader() (
	header ast.HeaderLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Header")
			panic(message)
		} else {
			// This is not a single header rule.
			return header, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "package" delimiter.
	_, token, ok = v.parseDelimiter("package")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Header")
			panic(message)
		} else {
			// This is not a single header rule.
			return header, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Header")
			panic(message)
		} else {
			// This is not a single header rule.
			return header, token, false
		}
	}
	ruleFound_ = true

	// Found a single header rule.
	ruleFound_ = true
	header = ast.Header().Make(
		comment,
		name,
	)
	return header, token, ruleFound_
}

func (v *parser_) parseImports() (
	imports ast.ImportsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "import" delimiter.
	_, token, ok = v.parseDelimiter("import")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Imports")
			panic(message)
		} else {
			// This is not a single imports rule.
			return imports, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Imports")
			panic(message)
		} else {
			// This is not a single imports rule.
			return imports, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited module rules.
	var modules = col.List[ast.ModuleLike]()
modulesLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var module ast.ModuleLike
		module, token, ok = v.parseModule()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single imports rule.
					return imports, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Imports")
				message += "The number of module rules must be at least 1."
				panic(message)
			default:
				break modulesLoop
			}
		}
		modules.AppendValue(module)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Imports")
			panic(message)
		} else {
			// This is not a single imports rule.
			return imports, token, false
		}
	}
	ruleFound_ = true

	// Found a single imports rule.
	ruleFound_ = true
	imports = ast.Imports().Make(modules)
	return imports, token, ruleFound_
}

func (v *parser_) parseInstance() (
	instance ast.InstanceLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Instance")
			panic(message)
		} else {
			// This is not a single instance rule.
			return instance, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Instance")
			panic(message)
		} else {
			// This is not a single instance rule.
			return instance, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Instance")
			panic(message)
		} else {
			// This is not a single instance rule.
			return instance, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single instanceMethods rule.
	var instanceMethods ast.InstanceMethodsLike
	instanceMethods, token, ok = v.parseInstanceMethods()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Instance")
			panic(message)
		} else {
			// This is not a single instance rule.
			return instance, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Instance")
			panic(message)
		} else {
			// This is not a single instance rule.
			return instance, token, false
		}
	}
	ruleFound_ = true

	// Found a single instance rule.
	ruleFound_ = true
	instance = ast.Instance().Make(
		declaration,
		instanceMethods,
	)
	return instance, token, ruleFound_
}

func (v *parser_) parseInstanceDefinitions() (
	instanceDefinitions ast.InstanceDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Instances" delimiter.
	_, token, ok = v.parseDelimiter("// Instances")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "InstanceDefinitions")
			panic(message)
		} else {
			// This is not a single instanceDefinitions rule.
			return instanceDefinitions, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited instance rules.
	var instances = col.List[ast.InstanceLike]()
instancesLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var instance ast.InstanceLike
		instance, token, ok = v.parseInstance()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single instanceDefinitions rule.
					return instanceDefinitions, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "InstanceDefinitions")
				message += "The number of instance rules must be at least 1."
				panic(message)
			default:
				break instancesLoop
			}
		}
		instances.AppendValue(instance)
	}

	// Found a single instanceDefinitions rule.
	ruleFound_ = true
	instanceDefinitions = ast.InstanceDefinitions().Make(instances)
	return instanceDefinitions, token, ruleFound_
}

func (v *parser_) parseInstanceMethods() (
	instanceMethods ast.InstanceMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single publicMethods rule.
	var publicMethods ast.PublicMethodsLike
	publicMethods, token, ok = v.parsePublicMethods()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "InstanceMethods")
			panic(message)
		} else {
			// This is not a single instanceMethods rule.
			return instanceMethods, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional attributeMethods rule.
	var optionalAttributeMethods ast.AttributeMethodsLike
	optionalAttributeMethods, _, ok = v.parseAttributeMethods()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse an optional aspectInterfaces rule.
	var optionalAspectInterfaces ast.AspectInterfacesLike
	optionalAspectInterfaces, _, ok = v.parseAspectInterfaces()
	if ok {
		ruleFound_ = true
	}

	// Found a single instanceMethods rule.
	ruleFound_ = true
	instanceMethods = ast.InstanceMethods().Make(
		publicMethods,
		optionalAttributeMethods,
		optionalAspectInterfaces,
	)
	return instanceMethods, token, ruleFound_
}

func (v *parser_) parseInterface() (
	interface_ ast.InterfaceLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Interface")
			panic(message)
		} else {
			// This is not a single interface rule.
			return interface_, token, false
		}
	}
	ruleFound_ = true

	// Found a single interface rule.
	ruleFound_ = true
	interface_ = ast.Interface().Make(
		abstraction,
	)
	return interface_, token, ruleFound_
}

func (v *parser_) parseInterfaceDefinitions() (
	interfaceDefinitions ast.InterfaceDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single classDefinitions rule.
	var classDefinitions ast.ClassDefinitionsLike
	classDefinitions, token, ok = v.parseClassDefinitions()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "InterfaceDefinitions")
			panic(message)
		} else {
			// This is not a single interfaceDefinitions rule.
			return interfaceDefinitions, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single instanceDefinitions rule.
	var instanceDefinitions ast.InstanceDefinitionsLike
	instanceDefinitions, token, ok = v.parseInstanceDefinitions()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "InterfaceDefinitions")
			panic(message)
		} else {
			// This is not a single interfaceDefinitions rule.
			return interfaceDefinitions, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional aspectDefinitions rule.
	var optionalAspectDefinitions ast.AspectDefinitionsLike
	optionalAspectDefinitions, _, ok = v.parseAspectDefinitions()
	if ok {
		ruleFound_ = true
	}

	// Found a single interfaceDefinitions rule.
	ruleFound_ = true
	interfaceDefinitions = ast.InterfaceDefinitions().Make(
		classDefinitions,
		instanceDefinitions,
		optionalAspectDefinitions,
	)
	return interfaceDefinitions, token, ruleFound_
}

func (v *parser_) parseMap() (
	map_ ast.MapLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "map" delimiter.
	_, token, ok = v.parseDelimiter("map")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Map")
			panic(message)
		} else {
			// This is not a single map rule.
			return map_, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Map")
			panic(message)
		} else {
			// This is not a single map rule.
			return map_, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Map")
			panic(message)
		} else {
			// This is not a single map rule.
			return map_, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Map")
			panic(message)
		} else {
			// This is not a single map rule.
			return map_, token, false
		}
	}
	ruleFound_ = true

	// Found a single map rule.
	ruleFound_ = true
	map_ = ast.Map().Make(name)
	return map_, token, ruleFound_
}

func (v *parser_) parseMethod() (
	method ast.MethodLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Method")
			panic(message)
		} else {
			// This is not a single method rule.
			return method, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Method")
			panic(message)
		} else {
			// This is not a single method rule.
			return method, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single method rule.
					return method, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Method")
				message += "The number of parameter rules must be at least 0."
				panic(message)
			default:
				break parametersLoop
			}
		}
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Method")
			panic(message)
		} else {
			// This is not a single method rule.
			return method, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional result rule.
	var optionalResult ast.ResultLike
	optionalResult, _, ok = v.parseResult()
	if ok {
		ruleFound_ = true
	}

	// Found a single method rule.
	ruleFound_ = true
	method = ast.Method().Make(
		name,
		parameters,
		optionalResult,
	)
	return method, token, ruleFound_
}

func (v *parser_) parseModel() (
	model ast.ModelLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single moduleDefinition rule.
	var moduleDefinition ast.ModuleDefinitionLike
	moduleDefinition, token, ok = v.parseModuleDefinition()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Model")
			panic(message)
		} else {
			// This is not a single model rule.
			return model, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single primitiveDefinitions rule.
	var primitiveDefinitions ast.PrimitiveDefinitionsLike
	primitiveDefinitions, token, ok = v.parsePrimitiveDefinitions()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Model")
			panic(message)
		} else {
			// This is not a single model rule.
			return model, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single interfaceDefinitions rule.
	var interfaceDefinitions ast.InterfaceDefinitionsLike
	interfaceDefinitions, token, ok = v.parseInterfaceDefinitions()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Model")
			panic(message)
		} else {
			// This is not a single model rule.
			return model, token, false
		}
	}
	ruleFound_ = true

	// Found a single model rule.
	ruleFound_ = true
	model = ast.Model().Make(
		moduleDefinition,
		primitiveDefinitions,
		interfaceDefinitions,
	)
	return model, token, ruleFound_
}

func (v *parser_) parseModule() (
	module ast.ModuleLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Module")
			panic(message)
		} else {
			// This is not a single module rule.
			return module, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single path token.
	var path string
	path, token, ok = v.parseToken(PathToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Module")
			panic(message)
		} else {
			// This is not a single module rule.
			return module, token, false
		}
	}
	ruleFound_ = true

	// Found a single module rule.
	ruleFound_ = true
	module = ast.Module().Make(
		name,
		path,
	)
	return module, token, ruleFound_
}

func (v *parser_) parseModuleDefinition() (
	moduleDefinition ast.ModuleDefinitionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single notice rule.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ModuleDefinition")
			panic(message)
		} else {
			// This is not a single moduleDefinition rule.
			return moduleDefinition, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single header rule.
	var header ast.HeaderLike
	header, token, ok = v.parseHeader()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ModuleDefinition")
			panic(message)
		} else {
			// This is not a single moduleDefinition rule.
			return moduleDefinition, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional imports rule.
	var optionalImports ast.ImportsLike
	optionalImports, _, ok = v.parseImports()
	if ok {
		ruleFound_ = true
	}

	// Found a single moduleDefinition rule.
	ruleFound_ = true
	moduleDefinition = ast.ModuleDefinition().Make(
		notice,
		header,
		optionalImports,
	)
	return moduleDefinition, token, ruleFound_
}

func (v *parser_) parseNone() (
	none ast.NoneLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "None")
			panic(message)
		} else {
			// This is not a single none rule.
			return none, token, false
		}
	}
	ruleFound_ = true

	// Found a single none rule.
	ruleFound_ = true
	none = ast.None().Make(newline)
	return none, token, ruleFound_
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Notice")
			panic(message)
		} else {
			// This is not a single notice rule.
			return notice, token, false
		}
	}
	ruleFound_ = true

	// Found a single notice rule.
	ruleFound_ = true
	notice = ast.Notice().Make(comment)
	return notice, token, ruleFound_
}

func (v *parser_) parseParameter() (
	parameter ast.ParameterLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Parameter")
			panic(message)
		} else {
			// This is not a single parameter rule.
			return parameter, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Parameter")
			panic(message)
		} else {
			// This is not a single parameter rule.
			return parameter, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Parameter")
			panic(message)
		} else {
			// This is not a single parameter rule.
			return parameter, token, false
		}
	}
	ruleFound_ = true

	// Found a single parameter rule.
	ruleFound_ = true
	parameter = ast.Parameter().Make(
		name,
		abstraction,
	)
	return parameter, token, ruleFound_
}

func (v *parser_) parseParameterized() (
	parameterized ast.ParameterizedLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Parameterized")
			panic(message)
		} else {
			// This is not a single parameterized rule.
			return parameterized, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single parameterized rule.
					return parameterized, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "Parameterized")
				message += "The number of parameter rules must be at least 1."
				panic(message)
			default:
				break parametersLoop
			}
		}
		parameters.AppendValue(parameter)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Parameterized")
			panic(message)
		} else {
			// This is not a single parameterized rule.
			return parameterized, token, false
		}
	}
	ruleFound_ = true

	// Found a single parameterized rule.
	ruleFound_ = true
	parameterized = ast.Parameterized().Make(parameters)
	return parameterized, token, ruleFound_
}

func (v *parser_) parsePrefix() (
	prefix ast.PrefixLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single array rule.
	var array ast.ArrayLike
	array, token, ok = v.parseArray()
	if ok {
		// Found a single array prefix.
		prefix = ast.Prefix().Make(array)
		return prefix, token, true
	}

	// Attempt to parse a single map rule.
	var map_ ast.MapLike
	map_, token, ok = v.parseMap()
	if ok {
		// Found a single map prefix.
		prefix = ast.Prefix().Make(map_)
		return prefix, token, true
	}

	// Attempt to parse a single channel rule.
	var channel ast.ChannelLike
	channel, token, ok = v.parseChannel()
	if ok {
		// Found a single channel prefix.
		prefix = ast.Prefix().Make(channel)
		return prefix, token, true
	}

	// This is not a single prefix rule.
	return prefix, token, false
}

func (v *parser_) parsePrimitiveDefinitions() (
	primitiveDefinitions ast.PrimitiveDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse an optional typeDefinitions rule.
	var optionalTypeDefinitions ast.TypeDefinitionsLike
	optionalTypeDefinitions, _, ok = v.parseTypeDefinitions()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse an optional functionalDefinitions rule.
	var optionalFunctionalDefinitions ast.FunctionalDefinitionsLike
	optionalFunctionalDefinitions, _, ok = v.parseFunctionalDefinitions()
	if ok {
		ruleFound_ = true
	}

	// Found a single primitiveDefinitions rule.
	ruleFound_ = true
	primitiveDefinitions = ast.PrimitiveDefinitions().Make(
		optionalTypeDefinitions,
		optionalFunctionalDefinitions,
	)
	return primitiveDefinitions, token, ruleFound_
}

func (v *parser_) parsePublicMethods() (
	publicMethods ast.PublicMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Public" delimiter.
	_, token, ok = v.parseDelimiter("// Public")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "PublicMethods")
			panic(message)
		} else {
			// This is not a single publicMethods rule.
			return publicMethods, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited method rules.
	var methods = col.List[ast.MethodLike]()
methodsLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var method ast.MethodLike
		method, token, ok = v.parseMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single publicMethods rule.
					return publicMethods, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "PublicMethods")
				message += "The number of method rules must be at least 1."
				panic(message)
			default:
				break methodsLoop
			}
		}
		methods.AppendValue(method)
	}

	// Found a single publicMethods rule.
	ruleFound_ = true
	publicMethods = ast.PublicMethods().Make(methods)
	return publicMethods, token, ruleFound_
}

func (v *parser_) parseResult() (
	result ast.ResultLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single none rule.
	var none ast.NoneLike
	none, token, ok = v.parseNone()
	if ok {
		// Found a single none result.
		result = ast.Result().Make(none)
		return result, token, true
	}

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if ok {
		// Found a single abstraction result.
		result = ast.Result().Make(abstraction)
		return result, token, true
	}

	// Attempt to parse a single parameterized rule.
	var parameterized ast.ParameterizedLike
	parameterized, token, ok = v.parseParameterized()
	if ok {
		// Found a single parameterized result.
		result = ast.Result().Make(parameterized)
		return result, token, true
	}

	// This is not a single result rule.
	return result, token, false

}

func (v *parser_) parseSuffix() (
	suffix ast.SuffixLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "." delimiter.
	_, token, ok = v.parseDelimiter(".")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Suffix")
			panic(message)
		} else {
			// This is not a single suffix rule.
			return suffix, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Suffix")
			panic(message)
		} else {
			// This is not a single suffix rule.
			return suffix, token, false
		}
	}
	ruleFound_ = true

	// Found a single suffix rule.
	ruleFound_ = true
	suffix = ast.Suffix().Make(name)
	return suffix, token, ruleFound_
}

func (v *parser_) parseType() (
	type_ ast.TypeLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single declaration rule.
	var declaration ast.DeclarationLike
	declaration, token, ok = v.parseDeclaration()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Type")
			panic(message)
		} else {
			// This is not a single type rule.
			return type_, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Type")
			panic(message)
		} else {
			// This is not a single type rule.
			return type_, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse an optional enumeration rule.
	var optionalEnumeration ast.EnumerationLike
	optionalEnumeration, _, ok = v.parseEnumeration()
	if ok {
		ruleFound_ = true
	}

	// Found a single type rule.
	ruleFound_ = true
	type_ = ast.Type().Make(
		declaration,
		abstraction,
		optionalEnumeration,
	)
	return type_, token, ruleFound_
}

func (v *parser_) parseTypeDefinitions() (
	typeDefinitions ast.TypeDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Types" delimiter.
	_, token, ok = v.parseDelimiter("// Types")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "TypeDefinitions")
			panic(message)
		} else {
			// This is not a single typeDefinitions rule.
			return typeDefinitions, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited type rules.
	var types = col.List[ast.TypeLike]()
typesLoop:
	for numberFound_ := 0; numberFound_ < unlimited; numberFound_++ {
		var type_ ast.TypeLike
		type_, token, ok = v.parseType()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single typeDefinitions rule.
					return typeDefinitions, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "TypeDefinitions")
				message += "The number of type rules must be at least 1."
				panic(message)
			default:
				break typesLoop
			}
		}
		types.AppendValue(type_)
	}

	// Found a single typeDefinitions rule.
	ruleFound_ = true
	typeDefinitions = ast.TypeDefinitions().Make(types)
	return typeDefinitions, token, ruleFound_
}

func (v *parser_) parseValue() (
	value ast.ValueLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Value")
			panic(message)
		} else {
			// This is not a single value rule.
			return value, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Value")
			panic(message)
		} else {
			// This is not a single value rule.
			return value, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "=" delimiter.
	_, token, ok = v.parseDelimiter("=")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Value")
			panic(message)
		} else {
			// This is not a single value rule.
			return value, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "iota" delimiter.
	_, token, ok = v.parseDelimiter("iota")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "Value")
			panic(message)
		} else {
			// This is not a single value rule.
			return value, token, false
		}
	}
	ruleFound_ = true

	// Found a single value rule.
	ruleFound_ = true
	value = ast.Value().Make(
		name,
		abstraction,
	)
	return value, token, ruleFound_
}

func (v *parser_) parseDelimiter(expectedValue string) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the right delimiter.
			return value, token, true
		}
		v.putBack(token)
	}

	// This is not the right delimiter.
	return value, token, false
}

func (v *parser_) parseToken(tokenType TokenType) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	token = v.getNextToken()
	for token != nil {
		// Check the token type.
		switch token.GetType() {
		case tokenType:
			// Found the right token type.
			value = token.GetValue()
			return value, token, true
		case SpaceToken, NewlineToken:
			// Ignore any unspecified whitespace.
			token = v.getNextToken()
		default:
			// This is not the right token type.
			v.putBack(token)
			return value, token, false
		}
	}

	// We are at the end-of-file marker.
	return value, token, false
}

func (v *parser_) formatError(token TokenLike, ruleName string) string {
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
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if col.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(ruleName string) string {
	return syntax_.GetValue(ruleName)
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError(token, "")
		panic(message)
	}

	return token
}

func (v *parser_) putBack(token TokenLike) {
	v.next_.AddValue(token)
}

// PRIVATE GLOBALS

// Constants

const unlimited = 4294967295 // Default to a reasonable value.

var syntax_ = col.Catalog[string, string](
	map[string]string{
		"Model":                `ModuleDefinition PrimitiveDefinitions InterfaceDefinitions`,
		"ModuleDefinition":     `Notice Header Imports?`,
		"PrimitiveDefinitions": `TypeDefinitions? FunctionalDefinitions?`,
		"InterfaceDefinitions": `ClassDefinitions InstanceDefinitions AspectDefinitions?`,
		"Notice":               `comment`,
		"Header":               `comment "package" name`,
		"Imports":              `"import" "(" Module+ ")"`,
		"Module":               `name path`,
		"TypeDefinitions":      `"// Types" Type+`,
		"Type":                 `Declaration Abstraction Enumeration?`,
		"Declaration":          `comment "type" name Constraints?`,
		"Constraints":          `"[" Parameter+ "]"`,
		"Parameter":            `name Abstraction ","`,
		"Abstraction":          `Prefix? name Suffix? Arguments?`,
		"Prefix": `
  - Array
  - Map
  - Channel`,
		"Array":                 `"[" "]"`,
		"Map":                   `"map" "[" name "]"`,
		"Channel":               `"chan"`,
		"Suffix":                `"." name`,
		"Arguments":             `"[" Argument AdditionalArgument* "]"`,
		"Argument":              `Abstraction`,
		"AdditionalArgument":    `"," Argument`,
		"Enumeration":           `"const" "(" Value AdditionalValue* ")"`,
		"Value":                 `name Abstraction "=" "iota"`,
		"AdditionalValue":       `name`,
		"FunctionalDefinitions": `"// Functionals" Functional+`,
		"Functional":            `Declaration "func" "(" Parameter* ")" Result`,
		"Result": `
  - None
  - Abstraction
  - Parameterized`,
		"None":                `newline`,
		"Parameterized":       `"(" Parameter+ ")"`,
		"ClassDefinitions":    `"// Classes" Class+`,
		"Class":               `Declaration "interface" "{" ClassMethods "}"`,
		"ClassMethods":        `ConstructorMethods ConstantMethods? FunctionMethods?`,
		"ConstructorMethods":  `"// Constructor" Constructor+`,
		"Constructor":         `name "(" Parameter* ")" Abstraction`,
		"ConstantMethods":     `"// Constant" Constant+`,
		"Constant":            `name "(" ")" Abstraction`,
		"FunctionMethods":     `"// Function" Function+`,
		"Function":            `name "(" Parameter* ")" Result`,
		"InstanceDefinitions": `"// Instances" Instance+`,
		"Instance":            `Declaration "interface" "{" InstanceMethods "}"`,
		"InstanceMethods":     `PublicMethods AttributeMethods? AspectInterfaces?`,
		"PublicMethods":       `"// Public" Method+`,
		"Method":              `name "(" Parameter* ")" Result?`,
		"AttributeMethods":    `"// Attribute" Attribute+`,
		"Attribute":           `name "(" Parameter? ")" Abstraction?`,
		"AspectInterfaces":    `"// Aspect" Interface+`,
		"Interface":           `Abstraction newline`,
		"AspectDefinitions":   `"// Aspects" Aspect+`,
		"Aspect":              `Declaration "interface" "{" Method+ "}"`,
	},
)
