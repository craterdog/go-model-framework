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

package generator

// Templates

const complexTemplate_ = `/*
................................................................................
<Copyright>
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "<name>" provides...

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package <name>

// Types

/*
Units is a constrained type representing the possible notational forms for the
complex number.
*/
type Form uint8

const (
	Polar Form = iota
	Rectangular
)

// Functionals

/*
NormFunction[V any] is a functional type that defines the signature for any
mathematical norm function.
*/
type NormFunction[V any] func(
	value V,
) float64

// Classes

/*
ComplexClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each complex-like concrete
class.
*/
type ComplexClassLike interface {
	// Constructor
	Make(
		realPart float64,
		imaginaryPart float64,
		form Form,
	) ComplexLike
	MakeFromValue(
		value complex128,
	) ComplexLike

	// Constant
	Zero() ComplexLike
	Infinity() ComplexLike

	// Function
	Inverse(
		value ComplexLike,
	) ComplexLike
	Sum(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Difference(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Reciprocal(
		value ComplexLike,
	) ComplexLike
	Product(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Quotient(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Norm(
		function NormFunction[ComplexLike],
		value ComplexLike,
	) float64
}

// Instances

/*
ComplexLike is an instance interface that defines the complete set of attributes,
abstractions and methods that must be supported by each instance of a concrete
complex-like class.
*/
type ComplexLike interface {
	// Public
	GetClass() ComplexClassLike
	IsReal() bool
	IsImaginary() bool

	// Attribute
	GetRealPart() float64
	GetImaginaryPart() float64
	GetForm() Form
	SetForm(
		form Form,
	)

	// Aspect
	Continuous
}

// Aspects

/*
Continuous is an aspect interface that defines a set of method signatures
that must be supported by each instance of a continuous concrete class.
*/
type Continuous interface {
	IsZero() bool
	IsDiscrete() bool
	IsInfinity() bool
}`
