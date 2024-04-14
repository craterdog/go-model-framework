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

package model_test

import (
	fmt "fmt"
	mod "github.com/craterdog/go-model-framework/v2"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestRoundtrips(t *tes.T) {
	var files, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var parser = mod.Parser().Make()
		var validator = mod.Validator().Make()
		var formatter = mod.Formatter().Make()
		var filename = testDirectory + file.Name()
		if sts.HasSuffix(filename, ".gomn") {
			fmt.Println(filename)
			var bytes, err = osx.ReadFile(filename)
			if err != nil {
				panic(err)
			}
			var expected = string(bytes)
			var model = parser.ParseSource(expected)
			validator.ValidateModel(model)
			var actual = formatter.FormatModel(model)
			ass.Equal(t, expected, actual)
		}
	}
}
