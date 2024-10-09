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
	fmt "fmt"
	gen "github.com/craterdog/go-model-framework/v4/generator"
	gra "github.com/craterdog/go-model-framework/v4/grammar"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var testDirectories = []string{
	"../../../go-test-framework-model/v4/ast/",
	"../../../go-test-framework-model/v4/grammar/",
	"../../../go-test-framework-model/v4/example/",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, directory := range testDirectories {
		fmt.Printf("   %v\n", directory)
		var bytes, err = osx.ReadFile(directory + "Package.go")
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var parser = gra.Parser().Make()
		var model = parser.ParseSource(source)
		var formatter = gra.Formatter().Make()
		var actual = formatter.FormatModel(model)
		ass.Equal(t, source, actual)
		var validator = gra.Validator().Make()
		validator.ValidateModel(model)
		var generator = gen.Classes().Make()
		var classes = generator.GenerateModelClasses(model).GetIterator()
		for classes.HasNext() {
			var association = classes.GetNext()
			var className = association.GetKey()
			var classSource = association.GetValue()
			bytes = []byte(classSource)
			var filename = directory + className + ".go"
			err = osx.WriteFile(filename, bytes, 0644)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("Done.")
}
