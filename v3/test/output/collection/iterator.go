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

var iteratorClass = map[string]any{}
var iteratorMutex syn.Mutex

// Function

func Iterator[V Value]() IteratorClassLike[V] {
	// Generate the name of the bound class type.
	var result_ IteratorClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	iteratorMutex.Lock()
	var value = iteratorClass[name]
	switch actual := value.(type) {
	case *iteratorClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &iteratorClass_[V]{
			// This class has no private constants to initialize.
		}
		iteratorClass[name] = result_
	}
	iteratorMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type iteratorClass_[V Value] struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *iteratorClass_[V]) MakeFromSequence(values Sequential[V]) IteratorLike[V] {
	return &iterator_[V]{}
}

// Functions

// INSTANCE METHODS

// Target

type iterator_[V Value] struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Public

func (v *iterator_[V]) GetNext() V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *iterator_[V]) GetPrevious() V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *iterator_[V]) GetSlot() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *iterator_[V]) HasNext() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *iterator_[V]) HasPrevious() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *iterator_[V]) ToEnd() {
	// TBA - Implement the method.
}

func (v *iterator_[V]) ToSlot(slot int) {
	// TBA - Implement the method.
}

func (v *iterator_[V]) ToStart() {
	// TBA - Implement the method.
}

// Private
