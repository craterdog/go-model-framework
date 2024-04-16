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

var collatorClass = map[string]any{}
var collatorMutex syn.Mutex

// Function

func Collator[V Value]() CollatorClassLike[V] {
	// Generate the name of the bound class type.
	var result_ CollatorClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	collatorMutex.Lock()
	var value = collatorClass[name]
	switch actual := value.(type) {
	case *collatorClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &collatorClass_[V]{
			// This class has no private constants to initialize.
		}
		collatorClass[name] = result_
	}
	collatorMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type collatorClass_[V Value] struct {
	defaultMaximum_ int
}

// Constants

func (c *collatorClass_[V]) DefaultMaximum() int {
	return c.defaultMaximum_
}

// Constructors

func (c *collatorClass_[V]) Make() CollatorLike[V] {
	return &collator_[V]{}
}

func (c *collatorClass_[V]) MakeWithMaximum(maximum int) CollatorLike[V] {
	return &collator_[V]{
		maximum_: maximum,
	}
}

// Functions

// INSTANCE METHODS

// Target

type collator_[V Value] struct {
	depth_ int
	maximum_ int
}

// Attributes

func (v *collator_[V]) GetDepth() int {
	return v.depth_
}

func (v *collator_[V]) GetMaximum() int {
	return v.maximum_
}

// Public

func (v *collator_[V]) CompareValues(
	first V,
	second V,
) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *collator_[V]) RankValues(
	first V,
	second V,
) int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

// Private
