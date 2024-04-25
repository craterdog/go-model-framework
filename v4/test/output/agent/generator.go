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

package agent

import ()

// CLASS ACCESS

// Reference

var generatorClass = &generatorClass_{
	// Any private class constants should be initialized here.
}

// Function

func Generator() GeneratorClassLike {
	return generatorClass
}

// CLASS METHODS

// Target

type generatorClass_ struct {
}

// Constants

// Constructors

func (c *generatorClass_) Make() GeneratorLike {
	return &generator_{}
}

// Functions

// INSTANCE METHODS

// Target

type generator_ struct {
	class_ GeneratorClassLike
}

// Attributes

func (v *generator_) GetClass() GeneratorClassLike {
	return v.class_
}

// Public

func (v *generator_) CreateModel(
	directory string,
	name string,
	copyright string,
) {
	// TBA - Implement the method.
}

func (v *generator_) GeneratePackage(directory string) {
	// TBA - Implement the method.
}

// Private
