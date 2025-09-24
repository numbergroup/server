package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/numbergroup/errors"
)

func getInt64(val, key string, maxVal, defaultVal int64) (int64, error) {
	if len(val) == 0 {
		return defaultVal, nil
	}
	out, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to parse int64 param \"%s\" value \"%s\"", key, val)
	}
	if maxVal != -1 && out > maxVal {
		return 0, errors.Errorf("param \"%s\" too large, must be less than %d", key, maxVal)
	}
	return out, nil
}

// GetInt64Param gets the int64 parameter from the gin context, and returns the value or the default value if the parameter is not set.
func GetInt64Param(c *gin.Context, key string, maxVal, defaultVal int64) (int64, error) {
	return getInt64(c.Param(key), key, maxVal, defaultVal)
}

// GetInt64Query gets the int64 query parameter from the gin context, and returns the value or the default value if the parameter is not set.
func GetInt64Query(c *gin.Context, key string, maxVal, defaultVal int64) (int64, error) {
	return getInt64(c.Query(key), key, maxVal, defaultVal)
}

// GetUUIDParam gets the uuid parameter from the gin context Param call, and parses then returns the value
// or an error if the parameter is not set or invalid.
func GetUUIDParam(c *gin.Context, key string) (uuid.UUID, error) {
	val := c.Param(key)
	if len(val) == 0 {
		return uuid.Nil, errors.Errorf("param \"%s\" is required", key)
	}
	out, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, errors.Wrapf(err, "failed to parse uuid param \"%s\" value \"%s\"", key, val)
	}
	return out, nil
}

// GetPagenation takes the query parameters "page" and "pageSize" from the gin context,
// and returns the values or an error if the parameters are invalid. Defaults page to 0, and pageSize to
// defaultPageSize if not set.
// Returns page, pageSize, error
func GetPagenation(c *gin.Context, maxPageSize, defaultPageSize int64) (int64, int64, error) {
	page, err := GetInt64Query(c, "page", -1, 0)
	if err != nil {
		return 0, 0, err
	}

	if page < 0 {
		return 0, 0, errors.New("page must be greater than or equal to 0")
	}

	pageSize, err := GetInt64Query(c, "pageSize", maxPageSize, defaultPageSize)
	if err != nil {
		return 0, 0, err
	}

	if pageSize <= 0 {
		return 0, 0, errors.New("pageSize must be greater than 0")
	}

	return page, pageSize, nil
}
