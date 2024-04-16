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

var sorterClass = map[string]any{}
var sorterMutex syn.Mutex

// Function

func Sorter[V Value]() SorterClassLike[V] {
	// Generate the name of the bound class type.
	var result_ SorterClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	sorterMutex.Lock()
	var value = sorterClass[name]
	switch actual := value.(type) {
	case *sorterClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &sorterClass_[V]{
			// This class has no private constants to initialize.
		}
		sorterClass[name] = result_
	}
	sorterMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type sorterClass_[V Value] struct {
	defaultRanker_ RankingFunction[V]
}

// Constants

func (c *sorterClass_[V]) DefaultRanker() RankingFunction[V] {
	return c.defaultRanker_
}

// Constructors

func (c *sorterClass_[V]) Make() SorterLike[V] {
	return &sorter_[V]{}
}

func (c *sorterClass_[V]) MakeWithRanker(ranker RankingFunction[V]) SorterLike[V] {
	return &sorter_[V]{
		ranker_: ranker,
	}
}

// Functions

// INSTANCE METHODS

// Target

type sorter_[V Value] struct {
	ranker_ RankingFunction[V]
}

// Attributes

func (v *sorter_[V]) GetRanker() RankingFunction[V] {
	return v.ranker_
}

// Systematic[V]

func (v *sorter_[V]) ReverseValues(values []V) {
	// TBA - Implement the method.
}

func (v *sorter_[V]) ShuffleValues(values []V) {
	// TBA - Implement the method.
}

func (v *sorter_[V]) SortValues(values []V) {
	// TBA - Implement the method.
}

// Public

// Private
