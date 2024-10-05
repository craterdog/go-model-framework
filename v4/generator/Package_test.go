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

package generator_test

import (
	gen "github.com/craterdog/go-model-framework/v4/generator"
	gra "github.com/craterdog/go-model-framework/v4/grammar"
	//ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

func TestGenerateClasses(t *tes.T) {
	var creator = gen.Classes().Make()

	// Read in the test model file.
	//var filename = "../../../go-test-framework/v4/ast/Package.go"
	var filename = "../ast/Package.go"
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)

	// Parse the source code for the model.
	var parser = gra.Parser().Make()
	var model = parser.ParseSource(source)

	// Validate the model.
	var validator = gra.Validator().Make()
	validator.ValidateModel(model)

	// Generate the classes.
	creator.GenerateModelClasses(model)
}

func TestCreateClassType(t *tes.T) {
	// Create a new class type model.
	var creator = gen.Generator().Make()
	var name = "example"
	var copyright = "Copyright (c) ACME Inc.  All Rights Reserved."
	creator.CreateClassType(name, copyright)
}

func TestCreateGenericType(t *tes.T) {
	// Create a new generic type model.
	var creator = gen.Generator().Make()
	var name = "example"
	var copyright string
	creator.CreateGenericType(name, copyright)
}

func TestCreateClassStructure(t *tes.T) {
	// Create a new class structure model.
	var creator = gen.Generator().Make()
	var name = "example"
	var copyright string
	creator.CreateClassStructure(name, copyright)
}

func TestCreateGenericStructure(t *tes.T) {
	// Create a new generic structure model.
	var creator = gen.Generator().Make()
	var name = "example"
	var copyright string
	creator.CreateGenericStructure(name, copyright)
}
