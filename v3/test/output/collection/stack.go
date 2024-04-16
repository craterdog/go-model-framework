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

var stackClass = map[string]any{}
var stackMutex syn.Mutex

// Function

func Stack[V Value]() StackClassLike[V] {
	// Generate the name of the bound class type.
	var result_ StackClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	stackMutex.Lock()
	var value = stackClass[name]
	switch actual := value.(type) {
	case *stackClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &stackClass_[V]{
			// This class has no private constants to initialize.
		}
		stackClass[name] = result_
	}
	stackMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type stackClass_[V Value] struct {
	defaultCapacity_ int
}

// Constants

func (c *stackClass_[V]) DefaultCapacity() int {
	return c.defaultCapacity_
}

// Constructors

func (c *stackClass_[V]) Make() StackLike[V] {
	return &stack_[V]{}
}

func (c *stackClass_[V]) MakeFromArray(values []V) StackLike[V] {
	return &stack_[V]{}
}

func (c *stackClass_[V]) MakeFromSequence(values Sequential[V]) StackLike[V] {
	return &stack_[V]{}
}

func (c *stackClass_[V]) MakeFromSource(
	source string,
	notation NotationLike,
) StackLike[V] {
	return &stack_[V]{}
}

func (c *stackClass_[V]) MakeWithCapacity(capacity int) StackLike[V] {
	return &stack_[V]{
		capacity_: capacity,
	}
}

// Functions

// INSTANCE METHODS

// Target

type stack_[V Value] struct {
	capacity_ int
}

// Attributes

func (v *stack_[V]) GetCapacity() int {
	return v.capacity_
}

// Limited[V]

func (v *stack_[V]) AddValue(value V) {
	// TBA - Implement the method.
}

func (v *stack_[V]) RemoveAll() {
	// TBA - Implement the method.
}

// Sequential[V]

func (v *stack_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *stack_[V]) GetIterator() IteratorLike[V] {
	var result_ IteratorLike[V]
	// TBA - Implement the method.
	return result_
}

func (v *stack_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *stack_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *stack_[V]) RemoveTop() V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

// Private
