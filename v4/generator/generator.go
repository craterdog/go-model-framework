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

package generator

import (
	fmt "fmt"
	ast "github.com/craterdog/go-model-framework/v4/ast"
	gra "github.com/craterdog/go-model-framework/v4/grammar"
	sts "strings"
	tim "time"
)

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// Initialize the class constants.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
	// Define the class constants.
}

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{
		// Initialize the instance attributes.
		class_: c,
	}
}

// INSTANCE METHODS

// Target

type generator_ struct {
	// Define the instance attributes.
	class_ GeneratorClassLike
}

// Attributes

func (v *generator_) GetClass() GeneratorClassLike {
	return v.class_
}

// Public

func (v *generator_) CreateClassStructure(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(complexTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = gra.Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) CreateClassType(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(angleTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = gra.Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) CreateGenericStructure(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(catalogTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = gra.Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

func (v *generator_) CreateGenericType(
	name string,
	copyright string,
) ast.ModelLike {
	copyright = v.expandCopyright(copyright)
	var source = sts.ReplaceAll(arrayTemplate_, "<Copyright>", copyright)
	source = sts.ReplaceAll(source, "<name>", name)
	var parser = gra.Parser().Make()
	var model = parser.ParseSource(source)
	return model
}

// Private

func (v *generator_) expandCopyright(copyright string) string {
	var maximum = 78
	var length = len(copyright)
	if length > maximum {
		var message = fmt.Sprintf(
			"The copyright notice cannot be longer than 78 characters: %v",
			copyright,
		)
		panic(message)
	}
	if length == 0 {
		copyright = fmt.Sprintf(
			"Copyright (c) %v.  All Rights Reserved.",
			tim.Now().Year(),
		)
		length = len(copyright)
	}
	var padding = (maximum - length) / 2
	for range padding {
		copyright = " " + copyright + " "
	}
	if len(copyright) < maximum {
		copyright = " " + copyright
	}
	copyright = "." + copyright + "."
	return copyright
}
