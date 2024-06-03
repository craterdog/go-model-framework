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

func TestCreateModel(t *tes.T) {
	// Create a new model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright = "Copyright (c) ACME Inc.  All Rights Reserved."
	var model = generator.CreateModel(name, copyright)

	// Generate a class from the model.
	var source = generator.GenerateClass(model, "angle")
	ass.Equal(t, angle, source)
}

func TestCreateGeneric(t *tes.T) {
	// Create a new generic model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright string
	var generic = generator.CreateGeneric(name, copyright)

	// Generate a generic class from the model.
	var source = generator.GenerateClass(generic, "set")
	ass.Equal(t, set, source)
}

const angle = `/*
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

func (c *angleClass_) MakeFromFloat(value float64) AngleLike {
	return &angle_{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *angleClass_) MakeFromString(value string) AngleLike {
	return &angle_{
		// Initialize instance attributes.
		class_: c,
	}
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

type angle_ struct {
	// Define instance attributes.
	class_ AngleClassLike
}

// Attributes

func (v *angle_) GetClass() AngleClassLike {
	return v.class_
}

// Angular

func (v *angle_) AsNormalized() AngleLike {
	var result_ AngleLike
	// TBA - Implement the method.
	return result_
}

func (v *angle_) InUnits(units Units) float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

// Public

func (v *angle_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *angle_) AsFloat() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

// Private
`

const set = `/*
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
