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
	mod "github.com/craterdog/go-model-framework/v4"
	tes "testing"
)

func TestClassType(t *tes.T) {
	var generator = mod.Generator()
	var name = "example"

	// Generate a new class type model with a default copyright.
	var copyright string
	var model = generator.CreateClassType(name, copyright)

	// Validate the class type model.
	var validator = mod.Validator()
	validator.ValidateModel(model)

	// Format the class type model.
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the class type model.
	var parser = mod.Parser()
	model = parser.ParseSource(source)

	// Generate a concrete class for the class type model.
	name = "angle"
	source = generator.GenerateClass(model, name)
	fmt.Printf("ANGLE SOURCE:\n%v\n", source)
}

func TestGenericType(t *tes.T) {
	var generator = mod.Generator()
	var name = "example"

	// Generate a new generic type model with a default copyright.
	var copyright string
	var model = generator.CreateGenericType(name, copyright)

	// Validate the generic type model.
	var validator = mod.Validator()
	validator.ValidateModel(model)

	// Format the generic type model.
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the generic type model.
	var parser = mod.Parser()
	model = parser.ParseSource(source)

	// Generate the concrete classes for the generic type model.
	name = "array"
	source = generator.GenerateClass(model, name)
	fmt.Printf("ARRAY SOURCE:\n%v\n", source)
}

func TestClassStructure(t *tes.T) {
	var generator = mod.Generator()
	var name = "example"

	// Generate a new class structure model with a default copyright.
	var copyright string
	var model = generator.CreateClassStructure(name, copyright)

	// Validate the class structure model.
	var validator = mod.Validator()
	validator.ValidateModel(model)

	// Format the class structure model.
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the class structure model.
	var parser = mod.Parser()
	model = parser.ParseSource(source)

	// Generate a concrete class for the class structure model.
	name = "complex"
	source = generator.GenerateClass(model, name)
	fmt.Printf("COMPLEX SOURCE:\n%v\n", source)
}

func TestGenericStructure(t *tes.T) {
	var generator = mod.Generator()
	var name = "example"

	// Generate a new generic structure model with a default copyright.
	var copyright string
	var model = generator.CreateGenericStructure(name, copyright)

	// Validate the generic structure model.
	var validator = mod.Validator()
	validator.ValidateModel(model)

	// Format the generic structure model.
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the generic structure model.
	var parser = mod.Parser()
	model = parser.ParseSource(source)

	// Generate the concrete classes for the generic structure model.
	name = "association"
	source = generator.GenerateClass(model, name)
	fmt.Printf("ASSOCIATION SOURCE:\n%v\n", source)
	name = "catalog"
	source = generator.GenerateClass(model, name)
	fmt.Printf("CATALOG SOURCE:\n%v\n", source)
}
