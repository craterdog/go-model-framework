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
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func Parser() ParserClassLike {
	return parserClass
}

// Constructor Methods

func (c *parserClass_) Make() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
		class_: c,
	}
	return instance
}

// INSTANCE INTERFACE

// Public Methods

func (v *parser_) GetClass() ParserClassLike {
	return v.getClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.ModelLike {
	// Create a scanner running in a separate Go routine.
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserClass.queueSize_)
	Scanner().Make(v.source_, v.tokens_)
	v.next_ = col.Stack[TokenLike](parserClass.stackSize_)

	// Attempt to parse the model from the token stream.
	var result_, token, ok = v.parseModel()
	if !ok {
		var message = v.formatError(token, "Model")
		panic(message)
	}
	return result_
}

// Private Methods

func (v *parser_) getClass() *parserClass_ {
	return v.class_
}

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
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
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

func (v *parser_) parseAspectDefinition() (
	aspectDefinition ast.AspectDefinitionLike,
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
			var message = v.formatError(token, "AspectDefinition")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspectDefinition, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "interface" delimiter.
	_, token, ok = v.parseDelimiter("interface")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectDefinition")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspectDefinition, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectDefinition")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspectDefinition, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited aspectMethod rules.
	var aspectMethods = col.List[ast.AspectMethodLike]()
aspectMethodsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var aspectMethod ast.AspectMethodLike
		aspectMethod, token, ok = v.parseAspectMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single aspect rule.
					return aspectDefinition, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "AspectDefinition")
				message += "The number of aspectMethod rules must be at least 1."
				panic(message)
			default:
				break aspectMethodsLoop
			}
		}
		aspectMethods.AppendValue(aspectMethod)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectDefinition")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspectDefinition, token, false
		}
	}
	ruleFound_ = true

	// Found a single aspect rule.
	ruleFound_ = true
	aspectDefinition = ast.AspectDefinition().Make(
		declaration,
		aspectMethods,
	)
	return aspectDefinition, token, ruleFound_
}

func (v *parser_) parseAspectInterface() (
	aspectInterface ast.AspectInterfaceLike,
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
			var message = v.formatError(token, "AspectInterface")
			panic(message)
		} else {
			// This is not a single interface rule.
			return aspectInterface, token, false
		}
	}
	ruleFound_ = true

	// Found a single interface rule.
	ruleFound_ = true
	aspectInterface = ast.AspectInterface().Make(
		abstraction,
	)
	return aspectInterface, token, ruleFound_
}

func (v *parser_) parseAspectMethod() (
	aspectMethod ast.AspectMethodLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single method rule.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectMethod")
			panic(message)
		} else {
			// This is not a single aspect rule.
			return aspectMethod, token, false
		}
	}
	ruleFound_ = true

	// Found a single aspectMethod rule.
	ruleFound_ = true
	aspectMethod = ast.AspectMethod().Make(
		method,
	)
	return aspectMethod, token, ruleFound_
}

func (v *parser_) parseAspectSection() (
	aspectSection ast.AspectSectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Aspect Definitions" delimiter.
	_, token, ok = v.parseDelimiter("// Aspect Definitions")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectSection")
			panic(message)
		} else {
			// This is not a single aspectSection rule.
			return aspectSection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited aspectDefinition rules.
	var aspectDefinitions = col.List[ast.AspectDefinitionLike]()
aspectDefinitionsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var aspectDefinition ast.AspectDefinitionLike
		aspectDefinition, token, ok = v.parseAspectDefinition()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single aspectSection rule.
					return aspectSection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "AspectSection")
				message += "The number of aspectDefinition rules must be at least 1."
				panic(message)
			default:
				break aspectDefinitionsLoop
			}
		}
		aspectDefinitions.AppendValue(aspectDefinition)
	}
	aspectDefinitions.SortValuesWithRanker(
		func(first, second ast.AspectDefinitionLike) col.Rank {
			var declaration = first.GetDeclaration()
			var firstName = declaration.GetName()
			declaration = second.GetDeclaration()
			var secondName = declaration.GetName()
			switch {
			case firstName < secondName:
				return col.LesserRank
			case firstName > secondName:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Found a single aspectSection rule.
	ruleFound_ = true
	aspectSection = ast.AspectSection().Make(aspectDefinitions)
	return aspectSection, token, ruleFound_
}

func (v *parser_) parseAspectSubsection() (
	aspectSubsection ast.AspectSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Aspect Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Aspect Methods")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AspectSubsection")
			panic(message)
		} else {
			// This is not a single aspectSubsection rule.
			return aspectSubsection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited interface rules.
	var interfaces = col.List[ast.AspectInterfaceLike]()
interfacesLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var aspectInterface ast.AspectInterfaceLike
		aspectInterface, token, ok = v.parseAspectInterface()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single aspectSubsection rule.
					return aspectSubsection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "AspectSubsection")
				message += "The number of interface rules must be at least 1."
				panic(message)
			default:
				break interfacesLoop
			}
		}
		interfaces.AppendValue(aspectInterface)
	}

	// Found a single aspectSubsection rule.
	ruleFound_ = true
	aspectSubsection = ast.AspectSubsection().Make(interfaces)
	return aspectSubsection, token, ruleFound_
}

func (v *parser_) parseAttributeMethod() (
	attributeMethod ast.AttributeMethodLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single getterMethod rule.
	var getterMethod ast.GetterMethodLike
	getterMethod, token, ok = v.parseGetterMethod()
	if ok {
		// Found a single getterMethod attributeMethod.
		attributeMethod = ast.AttributeMethod().Make(getterMethod)
		return attributeMethod, token, true
	}

	// Attempt to parse a single setterMethod rule.
	var setterMethod ast.SetterMethodLike
	setterMethod, token, ok = v.parseSetterMethod()
	if ok {
		// Found a single setterMethod attributeMethod.
		attributeMethod = ast.AttributeMethod().Make(setterMethod)
		return attributeMethod, token, true
	}

	// This is not a single attributeMethod rule.
	return attributeMethod, token, false
}

func (v *parser_) parseAttributeSubsection() (
	attributeSubsection ast.AttributeSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Attribute Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Attribute Methods")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "AttributeSubsection")
			panic(message)
		} else {
			// This is not a single attributeSubsection rule.
			return attributeSubsection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited attributeMethod rules.
	var attributeMethods = col.List[ast.AttributeMethodLike]()
attributeMethodsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var attributeMethod ast.AttributeMethodLike
		attributeMethod, token, ok = v.parseAttributeMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single attributeSubsection rule.
					return attributeSubsection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "AttributeSubsection")
				message += "The number of attributeMethod rules must be at least 1."
				panic(message)
			default:
				break attributeMethodsLoop
			}
		}
		attributeMethods.AppendValue(attributeMethod)
	}

	// Found a single attributeSubsection rule.
	ruleFound_ = true
	attributeSubsection = ast.AttributeSubsection().Make(attributeMethods)
	return attributeSubsection, token, ruleFound_
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
	classDefinition ast.ClassDefinitionLike,
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
			// This is not a single classDefinition rule.
			return classDefinition, token, false
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
			// This is not a single classDefinition rule.
			return classDefinition, token, false
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
			// This is not a single classDefinition rule.
			return classDefinition, token, false
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
			// This is not a single classDefinition rule.
			return classDefinition, token, false
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
			// This is not a single classDefinition rule.
			return classDefinition, token, false
		}
	}
	ruleFound_ = true

	// Found a single classDefinition rule.
	ruleFound_ = true
	classDefinition = ast.ClassDefinition().Make(
		declaration,
		classMethods,
	)
	return classDefinition, token, ruleFound_
}

func (v *parser_) parseClassMethods() (
	classMethods ast.ClassMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single constructorSubsection rule.
	var constructorSubsection ast.ConstructorSubsectionLike
	constructorSubsection, token, ok = v.parseConstructorSubsection()
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

	// Attempt to parse an optional constantSubsection rule.
	var optionalConstantSubsection ast.ConstantSubsectionLike
	optionalConstantSubsection, _, ok = v.parseConstantSubsection()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse an optional functionSubsection rule.
	var optionalFunctionSubsection ast.FunctionSubsectionLike
	optionalFunctionSubsection, _, ok = v.parseFunctionSubsection()
	if ok {
		ruleFound_ = true
	}

	// Found a single classMethods rule.
	ruleFound_ = true
	classMethods = ast.ClassMethods().Make(
		constructorSubsection,
		optionalConstantSubsection,
		optionalFunctionSubsection,
	)
	return classMethods, token, ruleFound_
}

func (v *parser_) parseClassSection() (
	classSection ast.ClassSectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Class Definitions" delimiter.
	_, token, ok = v.parseDelimiter("// Class Definitions")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ClassSection")
			panic(message)
		} else {
			// This is not a single classSection rule.
			return classSection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited classDefinition rules.
	var classDefinitions = col.List[ast.ClassDefinitionLike]()
classDefinitionsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var classDefinition ast.ClassDefinitionLike
		classDefinition, token, ok = v.parseClass()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single classSection rule.
					return classSection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "ClassSection")
				message += "The number of classDefinition rules must be at least 1."
				panic(message)
			default:
				break classDefinitionsLoop
			}
		}
		classDefinitions.AppendValue(classDefinition)
	}
	classDefinitions.SortValuesWithRanker(
		func(first, second ast.ClassDefinitionLike) col.Rank {
			var declaration = first.GetDeclaration()
			var firstName = sts.TrimSuffix(declaration.GetName(), "ClassLike")
			declaration = second.GetDeclaration()
			var secondName = sts.TrimSuffix(declaration.GetName(), "ClassLike")
			switch {
			case firstName < secondName:
				return col.LesserRank
			case firstName > secondName:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Found a single classSection rule.
	ruleFound_ = true
	classSection = ast.ClassSection().Make(classDefinitions)
	return classSection, token, ruleFound_
}

func (v *parser_) parseConstantMethod() (
	constantMethod ast.ConstantMethodLike,
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
			// This is not a single constantMethod rule.
			return constantMethod, token, false
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
			// This is not a single constantMethod rule.
			return constantMethod, token, false
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
			// This is not a single constantMethod rule.
			return constantMethod, token, false
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
			// This is not a single constantMethod rule.
			return constantMethod, token, false
		}
	}
	ruleFound_ = true

	// Found a single constantMethod rule.
	ruleFound_ = true
	constantMethod = ast.ConstantMethod().Make(
		name,
		abstraction,
	)
	return constantMethod, token, ruleFound_
}

func (v *parser_) parseConstantSubsection() (
	constantSubsection ast.ConstantSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Constant Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Constant Methods")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ConstantSubsection")
			panic(message)
		} else {
			// This is not a single constantSubsection rule.
			return constantSubsection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited constantMethod rules.
	var constantMethods = col.List[ast.ConstantMethodLike]()
constantMethodsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var constantMethod ast.ConstantMethodLike
		constantMethod, token, ok = v.parseConstantMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single constantSubsection rule.
					return constantSubsection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "ConstantSubsection")
				message += "The number of constantMethod rules must be at least 1."
				panic(message)
			default:
				break constantMethodsLoop
			}
		}
		constantMethods.AppendValue(constantMethod)
	}

	// Found a single constantSubsection rule.
	ruleFound_ = true
	constantSubsection = ast.ConstantSubsection().Make(constantMethods)
	return constantSubsection, token, ruleFound_
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
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
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

func (v *parser_) parseConstructorMethod() (
	constructorMethod ast.ConstructorMethodLike,
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
			// This is not a single constructorMethod rule.
			return constructorMethod, token, false
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
			// This is not a single constructorMethod rule.
			return constructorMethod, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single constructorMethod rule.
					return constructorMethod, token, false
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
			// This is not a single constructorMethod rule.
			return constructorMethod, token, false
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
			// This is not a single constructorMethod rule.
			return constructorMethod, token, false
		}
	}
	ruleFound_ = true

	// Found a single constructorMethod rule.
	ruleFound_ = true
	constructorMethod = ast.ConstructorMethod().Make(
		name,
		parameters,
		abstraction,
	)
	return constructorMethod, token, ruleFound_
}

func (v *parser_) parseConstructorSubsection() (
	constructorSubsection ast.ConstructorSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Constructor Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Constructor Methods")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "ConstructorSubsection")
			panic(message)
		} else {
			// This is not a single constructorSubsection rule.
			return constructorSubsection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited constructorMethod rules.
	var constructorMethods = col.List[ast.ConstructorMethodLike]()
constructorMethodsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var constructorMethod ast.ConstructorMethodLike
		constructorMethod, token, ok = v.parseConstructorMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single constructorSubsection rule.
					return constructorSubsection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "ConstructorSubsection")
				message += "The number of constructorMethod rules must be at least 1."
				panic(message)
			default:
				break constructorMethodsLoop
			}
		}
		constructorMethods.AppendValue(constructorMethod)
	}

	// Found a single constructorSubsection rule.
	ruleFound_ = true
	constructorSubsection = ast.ConstructorSubsection().Make(constructorMethods)
	return constructorSubsection, token, ruleFound_
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
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
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

func (v *parser_) parseFunctionMethod() (
	functionMethod ast.FunctionMethodLike,
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
			// This is not a single functionMethod rule.
			return functionMethod, token, false
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
			// This is not a single functionMethod rule.
			return functionMethod, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single functionMethod rule.
					return functionMethod, token, false
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
			// This is not a single functionMethod rule.
			return functionMethod, token, false
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
			// This is not a single functionMethod rule.
			return functionMethod, token, false
		}
	}
	ruleFound_ = true

	// Found a single functionMethod rule.
	ruleFound_ = true
	functionMethod = ast.FunctionMethod().Make(
		name,
		parameters,
		result,
	)
	return functionMethod, token, ruleFound_
}

func (v *parser_) parseFunctionSubsection() (
	functionSubsection ast.FunctionSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Function Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Function Methods")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "FunctionSubsection")
			panic(message)
		} else {
			// This is not a single functionSubsection rule.
			return functionSubsection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited functionMethod rules.
	var functionMethods = col.List[ast.FunctionMethodLike]()
functionMethodsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var functionMethod ast.FunctionMethodLike
		functionMethod, token, ok = v.parseFunctionMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single functionSubsection rule.
					return functionSubsection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "FunctionSubsection")
				message += "The number of functionMethod rules must be at least 1."
				panic(message)
			default:
				break functionMethodsLoop
			}
		}
		functionMethods.AppendValue(functionMethod)
	}

	// Found a single functionSubsection rule.
	ruleFound_ = true
	functionSubsection = ast.FunctionSubsection().Make(functionMethods)
	return functionSubsection, token, ruleFound_
}

func (v *parser_) parseFunctionalDefinition() (
	functionalDefinition ast.FunctionalDefinitionLike,
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
			return functionalDefinition, token, false
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
			return functionalDefinition, token, false
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
			return functionalDefinition, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 0 to unlimited parameter rules.
	var parameters = col.List[ast.ParameterLike]()
parametersLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var parameter ast.ParameterLike
		parameter, token, ok = v.parseParameter()
		if !ok {
			switch {
			case numberFound_ < 0:
				if !ruleFound_ {
					// This is not a single functional rule.
					return functionalDefinition, token, false
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
			return functionalDefinition, token, false
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
			return functionalDefinition, token, false
		}
	}
	ruleFound_ = true

	// Found a single functional rule.
	ruleFound_ = true
	functionalDefinition = ast.FunctionalDefinition().Make(
		declaration,
		parameters,
		result,
	)
	return functionalDefinition, token, ruleFound_
}

func (v *parser_) parseFunctionalSection() (
	functionalSection ast.FunctionalSectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Functional Definitions" delimiter.
	_, token, ok = v.parseDelimiter("// Functional Definitions")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "FunctionalSection")
			panic(message)
		} else {
			// This is not a single functionalSection rule.
			return functionalSection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited functional rules.
	var functionalDefinitions = col.List[ast.FunctionalDefinitionLike]()
functionalDefinitionsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var functionalDefinition ast.FunctionalDefinitionLike
		functionalDefinition, token, ok = v.parseFunctionalDefinition()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single functionalSection rule.
					return functionalSection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "FunctionalSection")
				message += "The number of functional rules must be at least 1."
				panic(message)
			default:
				break functionalDefinitionsLoop
			}
		}
		functionalDefinitions.AppendValue(functionalDefinition)
	}
	functionalDefinitions.SortValuesWithRanker(
		func(first, second ast.FunctionalDefinitionLike) col.Rank {
			var declaration = first.GetDeclaration()
			var firstName = sts.TrimSuffix(declaration.GetName(), "Function")
			declaration = second.GetDeclaration()
			var secondName = sts.TrimSuffix(declaration.GetName(), "Function")
			switch {
			case firstName < secondName:
				return col.LesserRank
			case firstName > secondName:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Found a single functionalSection rule.
	ruleFound_ = true
	functionalSection = ast.FunctionalSection().Make(functionalDefinitions)
	return functionalSection, token, ruleFound_
}

func (v *parser_) parseGetterMethod() (
	getterMethod ast.GetterMethodLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single name token.
	var name string
	var nameToken TokenLike
	name, nameToken, ok = v.parseToken(NameToken)
	if !ok {
		// This is not a single getterMethod rule.
		return getterMethod, nameToken, false
	}

	// Attempt to parse a single "(" delimiter.
	var leftToken TokenLike
	_, leftToken, ok = v.parseDelimiter("(")
	if !ok {
		// This is not a single getterMethod rule.
		v.putBack(nameToken)
		return getterMethod, leftToken, false
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		// This is not a single getterMethod rule.
		v.putBack(leftToken)
		v.putBack(nameToken)
		return getterMethod, token, false
	}
	ruleFound_ = true

	// Attempt to parse a single abstraction rule.
	var abstraction ast.AbstractionLike
	abstraction, token, ok = v.parseAbstraction()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "GetterMethod")
			panic(message)
		} else {
			// This is not a single getterMethod rule.
			return getterMethod, token, false
		}
	}
	ruleFound_ = true

	// Found a single getterMethod rule.
	ruleFound_ = true
	getterMethod = ast.GetterMethod().Make(
		name,
		abstraction,
	)
	return getterMethod, token, ruleFound_
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
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
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

func (v *parser_) parseInstanceDefinition() (
	instanceDefinition ast.InstanceDefinitionLike,
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
			// This is not a single instanceDefinition rule.
			return instanceDefinition, token, false
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
			// This is not a single instanceDefinition rule.
			return instanceDefinition, token, false
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
			// This is not a single instanceDefinition rule.
			return instanceDefinition, token, false
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
			// This is not a single instanceDefinition rule.
			return instanceDefinition, token, false
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
			// This is not a single instanceDefinition rule.
			return instanceDefinition, token, false
		}
	}
	ruleFound_ = true

	// Found a single instanceDefinition rule.
	ruleFound_ = true
	instanceDefinition = ast.InstanceDefinition().Make(
		declaration,
		instanceMethods,
	)
	return instanceDefinition, token, ruleFound_
}

func (v *parser_) parseInstanceSection() (
	instanceSection ast.InstanceSectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Instance Definitions" delimiter.
	_, token, ok = v.parseDelimiter("// Instance Definitions")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "InstanceSection")
			panic(message)
		} else {
			// This is not a single instanceSection rule.
			return instanceSection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited instanceDefinition rules.
	var instanceDefinitions = col.List[ast.InstanceDefinitionLike]()
instanceDefinitionsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var instanceDefinition ast.InstanceDefinitionLike
		instanceDefinition, token, ok = v.parseInstanceDefinition()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single instanceSection rule.
					return instanceSection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "InstanceSection")
				message += "The number of instanceDefinition rules must be at least 1."
				panic(message)
			default:
				break instanceDefinitionsLoop
			}
		}
		instanceDefinitions.AppendValue(instanceDefinition)
	}
	instanceDefinitions.SortValuesWithRanker(
		func(first, second ast.InstanceDefinitionLike) col.Rank {
			var declaration = first.GetDeclaration()
			var firstName = sts.TrimSuffix(declaration.GetName(), "Like")
			declaration = second.GetDeclaration()
			var secondName = sts.TrimSuffix(declaration.GetName(), "Like")
			switch {
			case firstName < secondName:
				return col.LesserRank
			case firstName > secondName:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Found a single instanceSection rule.
	ruleFound_ = true
	instanceSection = ast.InstanceSection().Make(instanceDefinitions)
	return instanceSection, token, ruleFound_
}

func (v *parser_) parseInstanceMethods() (
	instanceMethods ast.InstanceMethodsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single publicSubsection rule.
	var publicSubsection ast.PublicSubsectionLike
	publicSubsection, token, ok = v.parsePublicSubsection()
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

	// Attempt to parse an optional attributeSubsection rule.
	var optionalAttributeSubsection ast.AttributeSubsectionLike
	optionalAttributeSubsection, _, ok = v.parseAttributeSubsection()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse an optional aspectSubsection rule.
	var optionalAspectSubsection ast.AspectSubsectionLike
	optionalAspectSubsection, _, ok = v.parseAspectSubsection()
	if ok {
		ruleFound_ = true
	}

	// Found a single instanceMethods rule.
	ruleFound_ = true
	instanceMethods = ast.InstanceMethods().Make(
		publicSubsection,
		optionalAttributeSubsection,
		optionalAspectSubsection,
	)
	return instanceMethods, token, ruleFound_
}

func (v *parser_) parseInterfaceDefinitions() (
	interfaceDefinitions ast.InterfaceDefinitionsLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single classSection rule.
	var classSection ast.ClassSectionLike
	classSection, token, ok = v.parseClassSection()
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

	// Attempt to parse a single instanceSection rule.
	var instanceSection ast.InstanceSectionLike
	instanceSection, token, ok = v.parseInstanceSection()
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

	// Attempt to parse an optional aspectSection rule.
	var optionalAspectSection ast.AspectSectionLike
	optionalAspectSection, _, ok = v.parseAspectSection()
	if ok {
		ruleFound_ = true
	}

	// Found a single interfaceDefinitions rule.
	ruleFound_ = true
	interfaceDefinitions = ast.InterfaceDefinitions().Make(
		classSection,
		instanceSection,
		optionalAspectSection,
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
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
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
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
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

	// Attempt to parse an optional type section rule.
	var optionalTypeSection ast.TypeSectionLike
	optionalTypeSection, _, ok = v.parseTypeSection()
	if ok {
		ruleFound_ = true
	}

	// Attempt to parse an optional functionalSection rule.
	var optionalFunctionalSection ast.FunctionalSectionLike
	optionalFunctionalSection, _, ok = v.parseFunctionalSection()
	if ok {
		ruleFound_ = true
	}

	// Found a single primitiveDefinitions rule.
	ruleFound_ = true
	primitiveDefinitions = ast.PrimitiveDefinitions().Make(
		optionalTypeSection,
		optionalFunctionalSection,
	)
	return primitiveDefinitions, token, ruleFound_
}

func (v *parser_) parsePublicMethod() (
	publicMethod ast.PublicMethodLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single method rule.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "PublicMethod")
			panic(message)
		} else {
			// This is not a single public rule.
			return publicMethod, token, false
		}
	}
	ruleFound_ = true

	// Found a single publicMethod rule.
	ruleFound_ = true
	publicMethod = ast.PublicMethod().Make(
		method,
	)
	return publicMethod, token, ruleFound_
}

func (v *parser_) parsePublicSubsection() (
	publicSubsection ast.PublicSubsectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Public Methods" delimiter.
	_, token, ok = v.parseDelimiter("// Public Methods")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "PublicSubsection")
			panic(message)
		} else {
			// This is not a single publicSubsection rule.
			return publicSubsection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited publicMethod rules.
	var publicMethods = col.List[ast.PublicMethodLike]()
publicMethodsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var publicMethod ast.PublicMethodLike
		publicMethod, token, ok = v.parsePublicMethod()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single publicSubsection rule.
					return publicSubsection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "PublicSubsection")
				message += "The number of publicMethod rules must be at least 1."
				panic(message)
			default:
				break publicMethodsLoop
			}
		}
		publicMethods.AppendValue(publicMethod)
	}

	// Found a single publicSubsection rule.
	ruleFound_ = true
	publicSubsection = ast.PublicSubsection().Make(publicMethods)
	return publicSubsection, token, ruleFound_
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

func (v *parser_) parseSetterMethod() (
	setterMethod ast.SetterMethodLike,
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
			var message = v.formatError(token, "SetterMethod")
			panic(message)
		} else {
			// This is not a single setterMethod rule.
			return setterMethod, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "SetterMethod")
			panic(message)
		} else {
			// This is not a single setterMethod rule.
			return setterMethod, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single parameter rule.
	var parameter ast.ParameterLike
	parameter, token, ok = v.parseParameter()
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "SetterMethod")
			panic(message)
		} else {
			// This is not a single setterMethod rule.
			return setterMethod, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "SetterMethod")
			panic(message)
		} else {
			// This is not a single setterMethod rule.
			return setterMethod, token, false
		}
	}
	ruleFound_ = true

	// Found a single setterMethod rule.
	setterMethod = ast.SetterMethod().Make(
		name,
		parameter,
	)
	return setterMethod, token, ruleFound_
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

func (v *parser_) parseTypeDefinition() (
	typeDefinition ast.TypeDefinitionLike,
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
			return typeDefinition, token, false
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
			return typeDefinition, token, false
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
	typeDefinition = ast.TypeDefinition().Make(
		declaration,
		abstraction,
		optionalEnumeration,
	)
	return typeDefinition, token, ruleFound_
}

func (v *parser_) parseTypeSection() (
	typeSection ast.TypeSectionLike,
	token TokenLike,
	ok bool,
) {
	var ruleFound_ bool

	// Attempt to parse a single "// Type Definitions" delimiter.
	_, token, ok = v.parseDelimiter("// Type Definitions")
	if !ok {
		if ruleFound_ {
			// Found a syntax error.
			var message = v.formatError(token, "TypeSection")
			panic(message)
		} else {
			// This is not a single typeSection rule.
			return typeSection, token, false
		}
	}
	ruleFound_ = true

	// Attempt to parse 1 to unlimited typeDefinition rules.
	var typeDefinitions = col.List[ast.TypeDefinitionLike]()
typeDefinitionsLoop:
	for numberFound_ := 0; numberFound_ < v.getClass().unlimited_; numberFound_++ {
		var typeDefinition ast.TypeDefinitionLike
		typeDefinition, token, ok = v.parseTypeDefinition()
		if !ok {
			switch {
			case numberFound_ < 1:
				if !ruleFound_ {
					// This is not a single typeSection rule.
					return typeSection, token, false
				}
				// Found a syntax error.
				var message = v.formatError(token, "TypeSection")
				message += "The number of typeDefinition rules must be at least 1."
				panic(message)
			default:
				break typeDefinitionsLoop
			}
		}
		typeDefinitions.AppendValue(typeDefinition)
	}
	typeDefinitions.SortValuesWithRanker(
		func(first, second ast.TypeDefinitionLike) col.Rank {
			var declaration = first.GetDeclaration()
			var firstName = declaration.GetName()
			declaration = second.GetDeclaration()
			var secondName = declaration.GetName()
			switch {
			case firstName < secondName:
				return col.LesserRank
			case firstName > secondName:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Found a single typeSection rule.
	ruleFound_ = true
	typeSection = ast.TypeSection().Make(typeDefinitions)
	return typeSection, token, ruleFound_
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
	if uti.IsDefined(ruleName) {
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
	return v.getClass().syntax_.GetValue(ruleName)
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

// PRIVATE INTERFACE

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	class_  *parserClass_
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	queueSize_ uint
	stackSize_ uint
	unlimited_ int
	syntax_    abs.CatalogLike[string, string]
}

// Class Reference

var parserClass = &parserClass_{
	// Initialize the class constants.
	queueSize_: 16,
	stackSize_: 4,
	unlimited_: 4294967295, // Default to a reasonable value.
	syntax_: col.Catalog[string, string](
		map[string]string{
			"Model":                `ModuleDefinition PrimitiveDefinitions InterfaceDefinitions`,
			"ModuleDefinition":     `Notice Header Imports?`,
			"PrimitiveDefinitions": `TypeSection? FunctionalSection?`,
			"InterfaceDefinitions": `ClassSection InstanceSection AspectSection?`,
			"Notice":               `comment`,
			"Header":               `comment "package" name`,
			"Imports":              `"import" "(" Module+ ")"`,
			"Module":               `name path`,
			"TypeSection":          `"// Type Definitions" TypeDefinition+`,
			"TypeDefinition":       `Declaration Abstraction Enumeration?`,
			"Declaration":          `comment "type" name Constraints?`,
			"Constraints":          `"[" Constraint AdditionalConstraint* "]"`,
			"Constraint":           `name Abstraction`,
			"AdditionalConstraint": `"," Constraint`,
			"Abstraction":          `Prefix? name Suffix? Arguments?`,
			"Prefix": `
  - Array
  - Map
  - Channel
`,
			"Array":                `"[" "]"`,
			"Map":                  `"map" "[" name "]"`,
			"Channel":              `"chan"`,
			"Suffix":               `"." name`,
			"Arguments":            `"[" Argument AdditionalArgument* "]"`,
			"Argument":             `Abstraction`,
			"AdditionalArgument":   `"," Argument`,
			"Enumeration":          `"const" "(" Value AdditionalValue* ")"`,
			"Value":                `name Abstraction "=" "iota"`,
			"AdditionalValue":      `name`,
			"FunctionalSection":    `"// Functional Definitions" FunctionalDefinition+`,
			"FunctionalDefinition": `Declaration "func" "(" Parameter* ")" Result`,
			"Parameter":            `name Abstraction ","`,
			"Result": `
  - None
  - Abstraction
  - Parameterized
`,
			"None":                  `newline`,
			"Parameterized":         `"(" Parameter+ ")"`,
			"ClassSection":          `"// Class Definitions" ClassDefinition+`,
			"ClassDefinition":       `Declaration "interface" "{" ClassMethods "}"`,
			"ClassMethods":          `ConstructorSubsection ConstantSubsection? FunctionSubsection?`,
			"ConstructorSubsection": `"// Constructor Methods" ConstructorMethod+`,
			"ConstructorMethod":     `name "(" Parameter* ")" Abstraction`,
			"ConstantSubsection":    `"// Constant Methods" ConstantMethod+`,
			"ConstantMethod":        `name "(" ")" Abstraction`,
			"FunctionSubsection":    `"// Function Methods" FunctionMethod+`,
			"FunctionMethod":        `name "(" Parameter* ")" Result`,
			"InstanceSection":       `"// Instance Definitions" InstanceDefinition+`,
			"InstanceDefinition":    `Declaration "interface" "{" InstanceMethods "}"`,
			"InstanceMethods":       `PublicSubsection AttributeSubsection? AspectSubsection?`,
			"PublicSubsection":      `"// Public Methods" PublicMethod+`,
			"PublicMethod":          `Method`,
			"Method":                `name "(" Parameter* ")" Result?`,
			"AttributeSubsection":   `"// Attribute Methods" AttributeMethod+`,
			"AttributeMethod": `
  - GetterMethod
  - SetterMethod
`,
			"GetterMethod":     `name "(" ")" Abstraction`,
			"SetterMethod":     `name "(" Parameter ")"`,
			"AspectSubsection": `"// Aspect Interfaces" AspectInterface+`,
			"AspectInterface":  `Abstraction`,
			"AspectSection":    `"// Aspect Definitions" AspectDefinition+`,
			"AspectDefinition": `Declaration "interface" "{" AspectMethod+ "}"`,
			"AspectMethod":     `Method`,
		},
	),
}
