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

func TestCreateClassType(t *tes.T) {
	// Create a new class type model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright = "Copyright (c) ACME Inc.  All Rights Reserved."
	var model = generator.CreateClassType(name, copyright)

	// Generate a class from the class type model.
	var source = generator.GenerateClass(model, "angle")
	ass.Equal(t, angle, source)
}

func TestCreateGenericType(t *tes.T) {
	// Create a new generic type model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright string
	var model = generator.CreateGenericType(name, copyright)

	// Generate a class from the generic type model.
	var source = generator.GenerateClass(model, "array")
	ass.Equal(t, array, source)
}

func TestCreateClassStructure(t *tes.T) {
	// Create a new class structure model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright string
	var model = generator.CreateClassStructure(name, copyright)

	// Generate a class from the class structure model.
	var source = generator.GenerateClass(model, "complex")
	ass.Equal(t, complex_, source)
}

func TestCreateGenericStructure(t *tes.T) {
	// Create a new generic structure model.
	var generator = age.Generator().Make()
	var name = "example"
	var copyright string
	var model = generator.CreateGenericStructure(name, copyright)

	// Generate the classes from the generic structure model.
	var source = generator.GenerateClass(model, "association")
	ass.Equal(t, association, source)
	source = generator.GenerateClass(model, "catalog")
	ass.Equal(t, catalog, source)
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
	return Angle()
}

// Angular

func (v angle_) AsNormalized() AngleLike {
	var result_ AngleLike
	// TBA - Implement the method.
	return result_
}

func (v angle_) InUnits(units Units) float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

// Public

func (v angle_) AsFloat() float64 {
	var result_ float64
	// TBA - Implement the method.
	return result_
}

func (v angle_) IsZero() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Private
`

const array = `/*
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

var arrayClass = map[string]any{}
var arrayMutex syn.Mutex

// Function

func Array[V any]() ArrayClassLike[V] {
	// Generate the name of the bound class type.
	var result_ ArrayClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	arrayMutex.Lock()
	var value = arrayClass[name]
	switch actual := value.(type) {
	case *arrayClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &arrayClass_[V]{
			// Initialize class constants.
		}
		arrayClass[name] = result_
	}
	arrayMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type arrayClass_[V any] struct {
	// Define class constants.
	defaultRanker_ RankingFunction[V]
}

// Constants

func (c *arrayClass_[V]) DefaultRanker() RankingFunction[V] {
	return c.defaultRanker_
}

// Constructors

func (c *arrayClass_[V]) MakeFromValue(value []V) ArrayLike[V] {
	// TBA - Validate the value.
	return array_[V](value)
}

func (c *arrayClass_[V]) MakeFromSequence(values Sequential[V]) ArrayLike[V] {
	var result_ ArrayLike[V]
	// TBA - Implement the method.
	return result_
}

func (c *arrayClass_[V]) MakeFromSize(size int) ArrayLike[V] {
	var result_ ArrayLike[V]
	// TBA - Implement the method.
	return result_
}

// INSTANCE METHODS

// Target

type array_[V any] []V

// Attributes

func (v array_[V]) GetClass() ArrayClassLike[V] {
	return Array[V]()
}

// Accessible[V]

func (v array_[V]) GetValue(index uint) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v array_[V]) GetValues(
	first uint,
	last uint,
) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

// Sequential[V]

func (v array_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v array_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v array_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

// Updatable[V]

func (v array_[V]) SetValue(
	index uint,
	value V,
) {
	// TBA - Implement the method.
}

func (v array_[V]) SetValues(
	index uint,
	values Sequential[V],
) {
	// TBA - Implement the method.
}

// Public

func (v array_[V]) SortValues() {
	// TBA - Implement the method.
}

func (v array_[V]) SortValuesWithRanker(ranker RankingFunction[V]) {
	// TBA - Implement the method.
}

// Private
`

const complex_ = `/*
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

const association = `/*
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

var associationClass = map[string]any{}
var associationMutex syn.Mutex

// Function

func Association[K comparable, V any]() AssociationClassLike[K, V] {
	// Generate the name of the bound class type.
	var result_ AssociationClassLike[K, V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	associationMutex.Lock()
	var value = associationClass[name]
	switch actual := value.(type) {
	case *associationClass_[K, V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &associationClass_[K, V]{
			// Initialize class constants.
		}
		associationClass[name] = result_
	}
	associationMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type associationClass_[K comparable, V any] struct {
	// Define class constants.
}

// Constructors

func (c *associationClass_[K, V]) MakeWithAttributes(
	key K,
	value V,
) AssociationLike[K, V] {
	return &association_[K, V]{
		// Initialize instance attributes.
		class_: c,
		key_: key,
		value_: value,
	}
}

// INSTANCE METHODS

// Target

type association_[K comparable, V any] struct {
	// Define instance attributes.
	class_ AssociationClassLike[K, V]
	key_ K
	value_ V
}

// Attributes

func (v *association_[K, V]) GetClass() AssociationClassLike[K, V] {
	return v.class_
}

func (v *association_[K, V]) GetKey() K {
	return v.key_
}

func (v *association_[K, V]) GetValue() V {
	return v.value_
}

func (v *association_[K, V]) SetValue(value V) {
	v.value_ = value
}

// Private
`

const catalog = `/*
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

var catalogClass = map[string]any{}
var catalogMutex syn.Mutex

// Function

func Catalog[K comparable, V any]() CatalogClassLike[K, V] {
	// Generate the name of the bound class type.
	var result_ CatalogClassLike[K, V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	catalogMutex.Lock()
	var value = catalogClass[name]
	switch actual := value.(type) {
	case *catalogClass_[K, V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &catalogClass_[K, V]{
			// Initialize class constants.
		}
		catalogClass[name] = result_
	}
	catalogMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type catalogClass_[K comparable, V any] struct {
	// Define class constants.
	defaultRanker_ RankingFunction[AssociationLike[K, V]]
}

// Constants

func (c *catalogClass_[K, V]) DefaultRanker() RankingFunction[AssociationLike[K, V]] {
	return c.defaultRanker_
}

// Constructors

func (c *catalogClass_[K, V]) Make() CatalogLike[K, V] {
	return &catalog_[K, V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *catalogClass_[K, V]) MakeFromArray(associations []AssociationLike[K, V]) CatalogLike[K, V] {
	return &catalog_[K, V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *catalogClass_[K, V]) MakeFromMap(associations map[K]V) CatalogLike[K, V] {
	return &catalog_[K, V]{
		// Initialize instance attributes.
		class_: c,
	}
}

func (c *catalogClass_[K, V]) MakeFromSequence(associations Sequential[AssociationLike[K, V]]) CatalogLike[K, V] {
	return &catalog_[K, V]{
		// Initialize instance attributes.
		class_: c,
	}
}

// Functions

func (c *catalogClass_[K, V]) Extract(
	catalog CatalogLike[K, V],
	keys Sequential[K],
) CatalogLike[K, V] {
	var result_ CatalogLike[K, V]
	// TBA - Implement the function.
	return result_
}

func (c *catalogClass_[K, V]) Merge(
	first CatalogLike[K, V],
	second CatalogLike[K, V],
) CatalogLike[K, V] {
	var result_ CatalogLike[K, V]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type catalog_[K comparable, V any] struct {
	// Define instance attributes.
	class_ CatalogClassLike[K, V]
}

// Attributes

func (v *catalog_[K, V]) GetClass() CatalogClassLike[K, V] {
	return v.class_
}

// Associative[K, V]

func (v *catalog_[K, V]) GetKeys() Sequential[K] {
	var result_ Sequential[K]
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) GetValue(key K) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) RemoveValue(key K) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) SetValue(
	key K,
	value V,
) {
	// TBA - Implement the method.
}

// Sequential[AssociationLike[K, V]]

func (v *catalog_[K, V]) AsArray() []AssociationLike[K, V] {
	var result_ []AssociationLike[K, V]
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *catalog_[K, V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *catalog_[K, V]) SortValues() {
	// TBA - Implement the method.
}

func (v *catalog_[K, V]) SortValuesWithRanker(ranker RankingFunction[AssociationLike[K, V]]) {
	// TBA - Implement the method.
}

// Private
`
