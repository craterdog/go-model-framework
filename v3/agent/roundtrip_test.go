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
	gcm "github.com/craterdog/go-model-framework/v3/gcmn"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const inputDirectory = "./input/"

func TestRoundtrips(t *tes.T) {
	var files, err = osx.ReadDir(inputDirectory)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var parser = gcm.Parser().Make()
		var validator = gcm.Validator().Make()
		var formatter = gcm.Formatter().Make()
		var filename = inputDirectory + file.Name()
		if sts.HasSuffix(filename, ".gcmn") {
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
