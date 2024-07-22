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

package module_test

import (
	fmt "fmt"
	col "github.com/craterdog/go-model-framework/v4"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

var testModels = []string{
	"ast/Package.go",
	"grammar/Package.go",
}

func TestRoundTrips(t *tes.T) {
	for _, modelfile := range testModels {
		var bytes, err = osx.ReadFile(modelfile)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var parser = col.Parser()
		var model = parser.ParseSource(source)
		var formatter = col.Formatter()
		var actual = formatter.FormatModel(model)
		ass.Equal(t, actual, source)
		var validator = col.Validator()
		validator.ValidateModel(model)
	}
}

func TestAstGeneration(t *tes.T) {
	var filename = "ast/Package.go"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	var parser = col.Parser()
	var model = parser.ParseSource(source)
	var generator = col.Generator()
	var iterator = model.GetClasses().GetClasses().GetIterator()
	for iterator.HasNext() {
		var class = iterator.GetNext()
		var name = sts.ToLower(sts.TrimSuffix(
			class.GetDeclaration().GetName(),
			"ClassLike",
		))
		source = generator.GenerateClass(model, name)
		bytes = []byte(source)
		var filename = "ast/" + name + ".go"
		var err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func TestClassType(t *tes.T) {
	var generator = col.Generator()
	var name = "example"

	// Generate a new class type model with a default copyright.
	var copyright string
	var model = generator.CreateClassType(name, copyright)

	// Validate the class type model.
	var validator = col.Validator()
	validator.ValidateModel(model)

	// Format the class type model.
	var formatter = col.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the class type model.
	var parser = col.Parser()
	model = parser.ParseSource(source)

	// Generate a concrete class for the class type model.
	name = "angle"
	source = generator.GenerateClass(model, name)
	fmt.Printf("ANGLE SOURCE:\n%v\n", source)
}

func TestGenericType(t *tes.T) {
	var generator = col.Generator()
	var name = "example"

	// Generate a new generic type model with a default copyright.
	var copyright string
	var model = generator.CreateGenericType(name, copyright)

	// Validate the generic type model.
	var validator = col.Validator()
	validator.ValidateModel(model)

	// Format the generic type model.
	var formatter = col.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the generic type model.
	var parser = col.Parser()
	model = parser.ParseSource(source)

	// Generate the concrete classes for the generic type model.
	name = "array"
	source = generator.GenerateClass(model, name)
	fmt.Printf("ARRAY SOURCE:\n%v\n", source)
}

func TestClassStructure(t *tes.T) {
	var generator = col.Generator()
	var name = "example"

	// Generate a new class structure model with a default copyright.
	var copyright string
	var model = generator.CreateClassStructure(name, copyright)

	// Validate the class structure model.
	var validator = col.Validator()
	validator.ValidateModel(model)

	// Format the class structure model.
	var formatter = col.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the class structure model.
	var parser = col.Parser()
	model = parser.ParseSource(source)

	// Generate a concrete class for the class structure model.
	name = "complex"
	source = generator.GenerateClass(model, name)
	fmt.Printf("COMPLEX SOURCE:\n%v\n", source)
}

func TestGenericStructure(t *tes.T) {
	var generator = col.Generator()
	var name = "example"

	// Generate a new generic structure model with a default copyright.
	var copyright string
	var model = generator.CreateGenericStructure(name, copyright)

	// Validate the generic structure model.
	var validator = col.Validator()
	validator.ValidateModel(model)

	// Format the generic structure model.
	var formatter = col.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the generic structure model.
	var parser = col.Parser()
	model = parser.ParseSource(source)

	// Generate the concrete classes for the generic structure model.
	name = "association"
	source = generator.GenerateClass(model, name)
	fmt.Printf("ASSOCIATION SOURCE:\n%v\n", source)
	name = "catalog"
	source = generator.GenerateClass(model, name)
	fmt.Printf("CATALOG SOURCE:\n%v\n", source)
}
