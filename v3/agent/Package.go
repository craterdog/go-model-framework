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

/*
Package "agent" provides the ability to generate Go class files based on a
Go Package.go file that follows the format shown in the following code template:
  - https://github.com/craterdog/go-model-framework/blob/main/models/Package.go

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package agent

import ()

// Types

// Classes

/*
GeneratorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete generator-like class.
*/
type GeneratorClassLike interface {
	// Constructors
	Make() GeneratorLike
}

// Instances

/*
GeneratorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete generator-like class.
*/
type GeneratorLike interface {
	// Methods
	CreateModel(
		directory string,
		name string,
		copyright string,
	)
	GeneratePackage(directory string)
}
