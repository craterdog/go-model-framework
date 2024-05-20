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
	osx "os"
	sts "strings"
	tes "testing"
)

const packageName = "example"

func TestInitialization(t *tes.T) {
	fmt.Println("Initialization Test")
	var err = osx.RemoveAll(outputDirectory)
	if err != nil {
		panic(err)
	}
	err = osx.MkdirAll(outputDirectory, 0755)
	if err != nil {
		panic(err)
	}

	var copyright string
	var generator = age.Generator().Make()
	generator.CreateModel(outputDirectory, packageName, copyright)
	fmt.Println()
}

func TestGeneration(t *tes.T) {
	fmt.Println("Generation Test")
	var files, err = osx.ReadDir(inputDirectory)
	if err != nil {
		panic(err)
	}

	var generator = age.Generator().Make()
	for _, file := range files {
		var fileSuffix = ".gcmn"
		var fileName = sts.TrimSuffix(file.Name(), fileSuffix)
		fmt.Println("\t" + fileName + "/Package.go")
		var bytes, err = osx.ReadFile(inputDirectory + file.Name())
		if err != nil {
			panic(err)
		}
		var directoryName = outputDirectory + fileName + "/"
		err = osx.RemoveAll(directoryName)
		if err != nil {
			panic(err)
		}
		err = osx.MkdirAll(directoryName, 0755)
		if err != nil {
			panic(err)
		}
		err = osx.WriteFile(directoryName+"Package.go", bytes, 0644)
		if err != nil {
			panic(err)
		}
		generator.GeneratePackage(directoryName)
	}
	fmt.Println()
}
