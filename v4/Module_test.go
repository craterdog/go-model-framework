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

func TestLifecycle(t *tes.T) {
	var generator = mod.Generator()
	var name = "example"

	// Generate a new class model with a default copyright.
	var copyright string
	var model = generator.CreateModel(name, copyright)

	// Validate the class model.
	var validator = mod.Validator()
	validator.ValidateModel(model)

	// Format the class model.
	var formatter = mod.Formatter()
	var source = formatter.FormatModel(model)

	// Parse the source code for the class model.
	var parser = mod.Parser()
	model = parser.ParseSource(source)

	// Generate a concrete class for the class model.
	name = "set"
	source = generator.GenerateClass(model, name)
	fmt.Printf("CLASS SOURCE:\n %v\n", source)
}
