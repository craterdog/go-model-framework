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
	age "github.com/craterdog/go-collection-framework/v4/agent"
	fmt "fmt"
	syn "sync"
)

// CLASS ACCESS

// Reference

var queueClass = map[string]any{}
var queueMutex syn.Mutex

// Function

func Queue[V Value]() QueueClassLike[V] {
	// Generate the name of the bound class type.
	var result_ QueueClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	queueMutex.Lock()
	var value = queueClass[name]
	switch actual := value.(type) {
	case *queueClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &queueClass_[V]{
			// Any private class constants should be initialized here.
		}
		queueClass[name] = result_
	}
	queueMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type queueClass_[V Value] struct {
	notation_ NotationLike
	defaultCapacity_ int
}

// Constants

func (c *queueClass_[V]) Notation() NotationLike {
	return c.notation_
}

func (c *queueClass_[V]) DefaultCapacity() int {
	return c.defaultCapacity_
}

// Constructors

func (c *queueClass_[V]) Make() QueueLike[V] {
	return &queue_[V]{}
}

func (c *queueClass_[V]) MakeFromArray(values []V) QueueLike[V] {
	return &queue_[V]{}
}

func (c *queueClass_[V]) MakeFromSequence(values Sequential[V]) QueueLike[V] {
	return &queue_[V]{}
}

func (c *queueClass_[V]) MakeFromSource(source string) QueueLike[V] {
	return &queue_[V]{}
}

func (c *queueClass_[V]) MakeWithCapacity(capacity int) QueueLike[V] {
	return &queue_[V]{
		capacity_: capacity,
	}
}

// Functions

func (c *queueClass_[V]) Fork(
	group Synchronized,
	input QueueLike[V],
	size int,
) Sequential[QueueLike[V]] {
	var result_ Sequential[QueueLike[V]]
	// TBA - Implement the function.
	return result_
}

func (c *queueClass_[V]) Join(
	group Synchronized,
	inputs Sequential[QueueLike[V]],
) QueueLike[V] {
	var result_ QueueLike[V]
	// TBA - Implement the function.
	return result_
}

func (c *queueClass_[V]) Split(
	group Synchronized,
	input QueueLike[V],
	size int,
) Sequential[QueueLike[V]] {
	var result_ Sequential[QueueLike[V]]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type queue_[V Value] struct {
	class_ QueueClassLike[V]
	capacity_ int
}

// Attributes

func (v *queue_[V]) GetClass() QueueClassLike[V] {
	return v.class_
}

func (v *queue_[V]) GetCapacity() int {
	return v.capacity_
}

// Limited[V]

func (v *queue_[V]) AddValue(value V) {
	// TBA - Implement the method.
}

func (v *queue_[V]) RemoveAll() {
	// TBA - Implement the method.
}

// Sequential[V]

func (v *queue_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *queue_[V]) GetIterator() age.IteratorLike[V] {
	var result_ age.IteratorLike[V]
	// TBA - Implement the method.
	return result_
}

func (v *queue_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *queue_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Public

func (v *queue_[V]) CloseQueue() {
	// TBA - Implement the method.
}

func (v *queue_[V]) RemoveHead() (
	head V,
	ok bool,
) {
	// TBA - Implement the method.
	return
}

// Private
