package api

import (
	"strconv"
)

func parseOffsetLimit(offset string, limit string) (int, int, error) {
	var pagesize int
	var page int
	if limit == "" {
		pagesize = 10
	} else {
		var err error
		pagesize, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, err
		}
	}
	if offset == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(offset)
		if err != nil {
			return 0, 0, err
		}
	}

	return page, pagesize, nil
}
