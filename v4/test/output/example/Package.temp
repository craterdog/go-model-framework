/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "example" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package example

// Types

/*
<ConstrainedType> is a constrained type representing...
*/
type <ConstrainedType> <primitiveType>

const (
	<1stValue> <ConstrainedType> = iota
	<2ndValue>
	<3rdValue>
	...
)
...

// Functionals

/*
<FunctionName>Function is a functional type that defines the signature for any
function that...
*/
type <FunctionName>Function func(<Parameters>) <AbstractType>
...

// Aspects

/*
<AspectName> is an aspect interface that defines a set of method signatures
that must be supported by each instance of a <aspect name> concrete class.
*/
type <AspectName> interface {
	// Methods
	<MethodName>(<Parameters>) <AbstractType>
	...
}
...

// Classes

/*
<ClassName>ClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
<class-name>-like concrete class.
*/
type <ClassName>ClassLike interface {
	// Constants
	<ConstantName>() <AbstractType>
	...

	// Constructors
	Make<FromContext>(<Parameters>) <ClassName>Like
	...

	// Functions
	<FunctionName>(<Parameters>) <AbstractType>
	...
}
...

// Instances

/*
<ClassName>Like is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a <class-name>-like concrete class.
*/
type <ClassName>Like interface {
	// Attributes
	Get<AttributeName>() <AttributeType>
	Set<AttributeName>(<attributeName> <AttributeType>)
	...

	// Abstractions
	<AspectName>
	...

	// Methods
	<MethodName>(<Parameters>) <AbstractType>
	...
}
...
