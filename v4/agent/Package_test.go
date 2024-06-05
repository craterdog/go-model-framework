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
	age "github.com/craterdog/go-model-framework/v4/agent"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestCreateSimpleModel(t *tes.T) {
	// Create a new simple model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright = "Copyright (c) ACME Inc.  All Rights Reserved."
	var model = generator.CreateSimpleModel(name, copyright)

	// Generate a class from the simple model.
	var source = generator.GenerateClass(model, "angle")
	ass.Equal(t, simple, source)
}

func TestCreateCompoundModel(t *tes.T) {
	// Create a new compound model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright = "Copyright (c) ACME Inc.  All Rights Reserved."
	var model = generator.CreateCompoundModel(name, copyright)

	// Generate a class from the compound model.
	var source = generator.GenerateClass(model, "complex")
	ass.Equal(t, compound, source)
}

func TestCreateGenericModel(t *tes.T) {
	// Create a new generic model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright string
	var model = generator.CreateGenericModel(name, copyright)

	// Generate a class from the generic model.
	var source = generator.GenerateClass(model, "set")
	ass.Equal(t, generic, source)
}

const simple = `/*
................................................................................
.                 Copyright (c) ACME Inc.  All Rights Reserved.                .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package example

import ()

// CLASS ACCESS

// Reference

var angleClass = &angleClass_{
	// Initialize class constants.
}

// Function

func Angle() AngleClassLike {
	return angleClass
}

// CLASS METHODS

// Target

type angleClass_ struct {
	// Define class constants.
	pi_ AngleLike
	tau_ AngleLike
}

// Constants

func (c *angleClass_) Pi() AngleLike {
	return c.pi_
}

func (c *angleClass_) Tau() AngleLike {
	return c.tau_
}

// Constructors

func (c *angleClass_) MakeFromValue(value float64) AngleLike {
	// TBA - Validate the value.
	return angle_(value)
}

func (c *angleClass_) MakeFromString(value string) AngleLike {
	var result_ AngleLike
	// TBA - Implement the method.
	return result_
}

// Functions

func (c *angleClass_) Apply(
	function TrigonometricFunction,
	angle AngleLike,
) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

func (c *angleClass_) Sine(angle AngleLike) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

func (c *angleClass_) Cosine(angle AngleLike) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

func (c *angleClass_) Tangent(angle AngleLike) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type angle_ float64

// Attributes

func (v angle_) GetClass() AngleClassLike {
	return v.class_
}

// Angular

func (v angle_) AsNormalized() AngleLike {
	var result_ AngleLike
	// TBA - Implement the method.
	return result_
}

func (v angle_) AsUnits(units Units) float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

// Public

func (v angle_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v angle_) AsFloat() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v angle_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Private
`

const compound = `/*
................................................................................
.                 Copyright (c) ACME Inc.  All Rights Reserved.                .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package example

import ()

// CLASS ACCESS

// Reference

var complexClass = &complexClass_{
	// Initialize class constants.
}

// Function

func Complex() ComplexClassLike {
	return complexClass
}

// CLASS METHODS

// Target

type complexClass_ struct {
	// Define class constants.
	zero_ ComplexLike
	infinity_ ComplexLike
}

// Constants

func (c *complexClass_) Zero() ComplexLike {
	return c.zero_
}

func (c *complexClass_) Infinity() ComplexLike {
	return c.infinity_
}

// Constructors

func (c *complexClass_) MakeWithAttributes(
	realPart float64,
	imaginaryPart float64,
	form Form,
) ComplexLike {
	return &complex_{
		// Initialize instance attributes.
		class_: c,
		realPart_: realPart,
		imaginaryPart_: imaginaryPart,
		form_: form,
	}
}

func (c *complexClass_) MakeFromValue(value complex128) ComplexLike {
	return &complex_{
		// Initialize instance attributes.
		class_: c,
	}
}

// Functions

func (c *complexClass_) Inverse(value ComplexLike) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Sum(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Difference(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Reciprocal(value ComplexLike) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Product(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Quotient(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBA - Implement the function.
	return result_
}

func (c *complexClass_) Norm(
	function NormFunction[ComplexLike],
	value ComplexLike,
) float64 {
	var result_ float64
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type complex_ struct {
	// Define instance attributes.
	class_ ComplexClassLike
	realPart_ float64
	imaginaryPart_ float64
	form_ Form
}

// Attributes

func (v *complex_) GetClass() ComplexClassLike {
	return v.class_
}

func (v *complex_) GetRealPart() float64 {
	return v.realPart_
}

func (v *complex_) GetImaginaryPart() float64 {
	return v.imaginaryPart_
}

func (v *complex_) GetForm() Form {
	return v.form_
}

func (v *complex_) SetForm(form Form) {
	v.form_ = form
}

// Continuous

func (v *complex_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *complex_) IsDiscrete() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *complex_) IsInfinity() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *complex_) IsReal() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *complex_) IsImaginary() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Private
`

const generic = `/*
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

package example

import (
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var setClass = map[string]any{}
var setMutex syn.Mutex

// Function

func Set[V any]() SetClassLike[V] {
	// Generate the name of the bound class type.
	var result_ SetClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	setMutex.Lock()
	var value = setClass[name]
	switch actual := value.(type) {
	case *setClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &setClass_[V]{
			// Initialize class constants.
		}
		setClass[name] = result_
	}
	setMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type setClass_[V any] struct {
	// Define class constants.
	defaultRanker_ RankingFunction[V]
}

// Constants

func (c *setClass_[V]) DefaultRanker() RankingFunction[V] {
	return c.defaultRanker_
}

// Constructors

func (c *setClass_[V]) Make() SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *setClass_[V]) MakeFromArray(values []V) SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *setClass_[V]) MakeFromSequence(values Sequential[V]) SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *setClass_[V]) MakeWithRanker(ranker RankingFunction[V]) SetLike[V] {
	return &set_[V]{
		// Initialize instance attributes.
		class_: c,
		ranker_: ranker,
	}
}

// Functions

func (c *setClass_[V]) And(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Or(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Sans(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *setClass_[V]) Xor(
	first SetLike[V],
	second SetLike[V],
) SetLike[V] {
	var result_ SetLike[V]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type set_[V any] struct {
	// Define instance attributes.
	class_ SetClassLike[V]
	ranker_ RankingFunction[V]
}

// Attributes

func (v *set_[V]) GetClass() SetClassLike[V] {
	return v.class_
}

func (v *set_[V]) SetRanker(ranker RankingFunction[V]) {
	v.ranker_ = ranker
}

// Sequential[V]

func (v *set_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *set_[V]) AddValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveAll() {
	// TBA - Implement the method.
}

// Private
`
