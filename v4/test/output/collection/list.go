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

var listClass = map[string]any{}
var listMutex syn.Mutex

// Function

func List[V Value]() ListClassLike[V] {
	// Generate the name of the bound class type.
	var result_ ListClassLike[V]
	var name = fmt.Sprintf("%T", result_)

	// Check for existing bound class type.
	listMutex.Lock()
	var value = listClass[name]
	switch actual := value.(type) {
	case *listClass_[V]:
		// This bound class type already exists.
		result_ = actual
	default:
		// Add a new bound class type.
		result_ = &listClass_[V]{
			// Any private class constants should be initialized here.
		}
		listClass[name] = result_
	}
	listMutex.Unlock()

	// Return a reference to the bound class type.
	return result_
}

// CLASS METHODS

// Target

type listClass_[V Value] struct {
	notation_ NotationLike
}

// Constants

func (c *listClass_[V]) Notation() NotationLike {
	return c.notation_
}

// Constructors

func (c *listClass_[V]) Make() ListLike[V] {
	return &list_[V]{}
}

func (c *listClass_[V]) MakeFromArray(values []V) ListLike[V] {
	return &list_[V]{}
}

func (c *listClass_[V]) MakeFromSequence(values Sequential[V]) ListLike[V] {
	return &list_[V]{}
}

func (c *listClass_[V]) MakeFromSource(source string) ListLike[V] {
	return &list_[V]{}
}

// Functions

func (c *listClass_[V]) Concatenate(
	first ListLike[V],
	second ListLike[V],
) ListLike[V] {
	var result_ ListLike[V]
	// TBA - Implement the function.
	return result_
}

// INSTANCE METHODS

// Target

type list_[V Value] struct {
	class_ ListClassLike[V]
}

// Attributes

func (v *list_[V]) GetClass() ListClassLike[V] {
	return v.class_
}

// Accessible[V]

func (v *list_[V]) GetValue(index int) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) GetValues(
	first int,
	last int,
) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

// Expandable[V]

func (v *list_[V]) AppendValue(value V) {
	// TBA - Implement the method.
}

func (v *list_[V]) AppendValues(values Sequential[V]) {
	// TBA - Implement the method.
}

func (v *list_[V]) InsertValue(
	slot int,
	value V,
) {
	// TBA - Implement the method.
}

func (v *list_[V]) InsertValues(
	slot int,
	values Sequential[V],
) {
	// TBA - Implement the method.
}

func (v *list_[V]) RemoveAll() {
	// TBA - Implement the method.
}

func (v *list_[V]) RemoveValue(index int) V {
	var result_ V
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) RemoveValues(
	first int,
	last int,
) Sequential[V] {
	var result_ Sequential[V]
	// TBA - Implement the method.
	return result_
}

// Searchable[V]

func (v *list_[V]) ContainsAll(values Sequential[V]) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) ContainsAny(values Sequential[V]) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) ContainsValue(value V) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) GetIndex(value V) int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

// Sequential[V]

func (v *list_[V]) AsArray() []V {
	var result_ []V
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) GetIterator() age.IteratorLike[V] {
	var result_ age.IteratorLike[V]
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) GetSize() int {
	var result_ int
	// TBA - Implement the method.
	return result_
}

func (v *list_[V]) IsEmpty() bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

// Sortable[V]

func (v *list_[V]) ReverseValues() {
	// TBA - Implement the method.
}

func (v *list_[V]) ShuffleValues() {
	// TBA - Implement the method.
}

func (v *list_[V]) SortValues() {
	// TBA - Implement the method.
}

func (v *list_[V]) SortValuesWithRanker(ranker age.RankingFunction[V]) {
	// TBA - Implement the method.
}

// Updatable[V]

func (v *list_[V]) SetValue(
	index int,
	value V,
) {
	// TBA - Implement the method.
}

func (v *list_[V]) SetValues(
	index int,
	values Sequential[V],
) {
	// TBA - Implement the method.
}

// Public

// Private
