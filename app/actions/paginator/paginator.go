package paginator

import (
	"math"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/defaults"
)

// Pagination struct
type Pagination struct {
	Page       int
	TotalPages int
	PerPage    int
	Offset     int

	CurrentEntriesSize int
	TotalEntriesSize   int
}

// NewPagination instance a pagination
func NewPagination(limit, offset, currentSize int32, totalRecords int64) Pagination {
	pages := math.Ceil(float64(totalRecords) / float64(limit))
	page := math.Ceil(float64(offset)/float64(limit)) + 1

	result := Pagination{
		Page:               int(page),
		TotalPages:         int(pages),
		PerPage:            int(limit),
		Offset:             int(offset),
		CurrentEntriesSize: int(currentSize),
		TotalEntriesSize:   int(totalRecords),
	}

	return result
}

// ApplyLimitAndOffset applies the limit and offset obtained through the context to the given query
func ApplyLimitAndOffset(q *pop.Query, c buffalo.Context) error {
	page, pagErr := strconv.Atoi(defaults.String(c.Param("page"), "1"))
	if pagErr != nil {
		return pagErr
	}

	limit, limErr := strconv.Atoi(defaults.String(c.Param("Limit"), "10"))
	if limErr != nil {
		return limErr
	}

	offset := ((page - 1) * limit)

	q.Paginator.PerPage = limit
	q.Paginator.Offset = offset

	return nil
}
