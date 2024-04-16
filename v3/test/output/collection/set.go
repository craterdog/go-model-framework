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

package collection

import (
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var setClass = map[string]any{}
var setMutex syn.Mutex

// Function

func Set[V Value]() SetClassLike[V] {
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
			// This class has no private constants to initialize.
		}
		setClass[name] = result_
	}
	setMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type setClass_[V Value] struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *setClass_[V]) Make() SetLike[V] {
	return &set_[V]{}
}

func (c *setClass_[V]) MakeFromArray(values []V) SetLike[V] {
	return &set_[V]{}
}

func (c *setClass_[V]) MakeFromSequence(values Sequential[V]) SetLike[V] {
	return &set_[V]{}
}

func (c *setClass_[V]) MakeFromSource(
	source string,
	notation NotationLike,
) SetLike[V] {
	return &set_[V]{}
}

func (c *setClass_[V]) MakeWithCollator(collator CollatorLike[V]) SetLike[V] {
	return &set_[V]{
		collator_: collator,
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

type set_[V Value] struct {
	collator_ CollatorLike[V]
}

// Attributes

func (v *set_[V]) GetCollator() CollatorLike[V] {
	return v.collator_
}

// Accessible[V]

func (v *set_[V]) GetValue(index int) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) GetValues(
	first int,
	last int,
) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

// Flexible[V]

func (v *set_[V]) AddValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) AddValues(values Sequential[V]) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveAll() {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveValue(value V) {
	// TBA - Implement the method.
}

func (v *set_[V]) RemoveValues(values Sequential[V]) {
	// TBA - Implement the method.
}

// Searchable[V]

func (v *set_[V]) ContainsAll(values Sequential[V]) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) ContainsAny(values Sequential[V]) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) ContainsValue(value V) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) GetIndex(value V) int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

// Sequential[V]

func (v *set_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *set_[V]) GetIterator() IteratorLike[V] {
	var result_ IteratorLike[V]
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

// Private
