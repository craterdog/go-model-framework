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

package grammar_test

import (
	fmt "fmt"
	gra "github.com/craterdog/go-model-framework/v4/grammar"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var filenames = []string{
	"../ast/Package.go",
	"../grammar/Package.go",
}

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	for _, filename := range filenames {
		fmt.Printf("   %v\n", filename)
		// Read in the class model file.
		var bytes, err = osx.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)

		// Parse the source code for the class model.
		var parser = gra.Parser().Make()
		var model = parser.ParseSource(source)

		// Validate the class model.
		var validator = gra.Validator().Make()
		validator.ValidateModel(model)

		// Format the class model.
		var formatter = gra.Formatter().Make()
		var actual = formatter.FormatModel(model)
		ass.Equal(t, source, actual)
	}
	fmt.Println("Done.")
}
