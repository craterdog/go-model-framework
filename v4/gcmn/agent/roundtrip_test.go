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

package agent_test

import (
	fmt "fmt"
	age "github.com/craterdog/go-model-framework/v4/gcmn/agent"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const inputDirectory = "../../test/input/"
const outputDirectory = "../../test/output/"

func TestRoundtrips(t *tes.T) {
	fmt.Println("Roundtrip Test")
	var files, err = osx.ReadDir(inputDirectory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var parser = age.Parser().Make()
		var validator = age.Validator().Make()
		var formatter = age.Formatter().Make()
		var filename = file.Name()
		if sts.HasSuffix(filename, ".gcmn") {
			fmt.Println("\t" + filename)
			var bytes, err = osx.ReadFile(inputDirectory + filename)
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
	fmt.Println()
}
