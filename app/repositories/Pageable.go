package repositories

import (
	"gorm.io/gorm"
)

const (
	zero = iota
	one
	two
	three
)

type Pageable interface {
	getNumberOfItemsToSkip() int
	getSize() int
}

type Page[t any] struct {
	elements []*t
}

type pageRequest struct {
	pageNumber int
	size       int
	skip       int
}

func NewPageAble(pageNumber int, size int) Pageable {
	return &pageRequest{
		pageNumber: pageNumber,
		size:       size,
		skip:       computeNumberOfItemsToSkip(pageNumber, size),
	}
}

func (page *Page[t]) GetElements() []*t {
	return page.elements
}

func (pageRequest *pageRequest) getNumberOfItemsToSkip() int {
	return pageRequest.skip
}

func (pageRequest *pageRequest) getSize() int {
	return pageRequest.size
}

func computeNumberOfItemsToSkip(pageNumber int, size int) (numberOfItemsToSkip int) {
	var isRequestForPageOne = pageNumber == one
	if isRequestForPageOne {
		return pageNumber - one
	}
	numberOfItemsToSkip = (pageNumber - one) * size
	return numberOfItemsToSkip
}

func getPage[t any](db *gorm.DB, pageable Pageable) *Page[t] {
	skip := pageable.getNumberOfItemsToSkip() - one
	var size = pageable.getSize()
	var items []*t
	db.Offset(skip).Limit(size).Find(&items)
	return &Page[t]{
		elements: items,
	}
}
