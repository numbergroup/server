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

func getUint64(val, key string, maxVal, defaultVal uint64) (uint64, error) {
	if len(val) == 0 {
		return defaultVal, nil
	}
	out, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to parse int64 param \"%s\" value \"%s\"", key, val)
	}
	if maxVal != 0 && out > maxVal {
		return 0, errors.Errorf("param \"%s\" too large, must be less than %d", key, maxVal)
	}
	return out, nil
}

// GetUint64Param gets the uint64 parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of 0 means no max.
func GetUint64Param(c *gin.Context, key string, maxVal, defaultVal uint64) (uint64, error) {
	return getUint64(c.Param(key), key, maxVal, defaultVal)
}

// GetUint64Query gets the uint64 query parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of 0 means no max.
func GetUint64Query(c *gin.Context, key string, maxVal, defaultVal uint64) (uint64, error) {
	return getUint64(c.Query(key), key, maxVal, defaultVal)
}

func getBool(val, key string, defaultVal bool) (bool, error) {
	if len(val) == 0 {
		return defaultVal, nil
	}
	out, err := strconv.ParseBool(val)
	if err != nil {
		return false, errors.Wrapf(err, "failed to parse bool param \"%s\" value \"%s\"", key, val)
	}
	return out, nil
}

// GetBoolParam gets the bool parameter from the gin context, and returns the value or the default value if the parameter is not set.
func GetBoolParam(c *gin.Context, key string, defaultVal bool) (bool, error) {
	return getBool(c.Param(key), key, defaultVal)
}

// GetBoolQuery gets the bool query parameter from the gin context, and returns the value or the default value if the parameter is not set.
func GetBoolQuery(c *gin.Context, key string, defaultVal bool) (bool, error) {
	return getBool(c.Query(key), key, defaultVal)
}

// GetInt32Param gets the int parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of -1 means no max.
func GetInt32Param(c *gin.Context, key string, maxVal, defaultVal int32) (int32, error) {
	val, err := getInt64(c.Param(key), key, int64(maxVal), int64(defaultVal))
	if err != nil {
		return 0, err
	}

	if val > 0x7FFFFFFF {
		return 0, errors.New("param is too large")
	}
	if val < -0x80000000 {
		return 0, errors.New("param is too small")
	}
	return int32(val), nil
}

// GetInt32Query gets the int query parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of -1 means no max.
func GetInt32Query(c *gin.Context, key string, maxVal, defaultVal int32) (int32, error) {
	val, err := getInt64(c.Query(key), key, int64(maxVal), int64(defaultVal))
	if err != nil {
		return 0, err
	}

	if val > 0x7FFFFFFF {
		return 0, errors.New("param is too large")
	}
	if val < -0x80000000 {
		return 0, errors.New("param is too small")
	}
	return int32(val), nil
}

// GetIntParam gets the int parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of -1 means no max.
func GetIntParam(c *gin.Context, key string, maxVal, defaultVal int) (int, error) {
	if strconv.IntSize == 32 { //  32-bit system
		val, err := GetInt32Param(c, key, int32(maxVal), int32(defaultVal))
		return int(val), err
	}

	val, err := GetInt64Param(c, key, int64(maxVal), int64(defaultVal))
	return int(val), err
}

// GetIntQuery gets the int query parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of -1 means no max.
func GetIntQuery(c *gin.Context, key string, maxVal, defaultVal int) (int, error) {
	if strconv.IntSize == 32 { //  32-bit system
		val, err := GetInt32Query(c, key, int32(maxVal), int32(defaultVal))
		return int(val), err
	}
	val, err := GetInt64Query(c, key, int64(maxVal), int64(defaultVal))
	return int(val), err
}

// GetUint32Param gets the int parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of 0 means no max.
func GetUint32Param(c *gin.Context, key string, maxVal, defaultVal uint32) (uint32, error) {
	val, err := GetUint64Param(c, key, uint64(maxVal), uint64(defaultVal))
	if err != nil {
		return 0, err
	}

	if val > 0xFFFFFFFF {
		return 0, errors.New("param is too large")
	}
	return uint32(val), nil
}

// GetInt32Query gets the int query parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of 0 means no max.
func GetUint32Query(c *gin.Context, key string, maxVal, defaultVal uint32) (uint32, error) {
	val, err := GetUint64Query(c, key, uint64(maxVal), uint64(defaultVal))
	if err != nil {
		return 0, err
	}

	if val > 0xFFFFFFFF {
		return 0, errors.New("param is too large")
	}
	return uint32(val), nil
}

// GetIntParam gets the int parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of 0 means no max.
func GetUintParam(c *gin.Context, key string, maxVal, defaultVal uint) (uint, error) {
	if strconv.IntSize == 32 { //  32-bit system
		val, err := GetUint32Param(c, key, uint32(maxVal), uint32(defaultVal))
		return uint(val), err
	}

	val, err := GetUint64Param(c, key, uint64(maxVal), uint64(defaultVal))
	return uint(val), err
}

// GetIntQuery gets the int query parameter from the gin context, and returns the value or the default value if the parameter is not set.
// maxVal of 0 means no max.
func GetUintQuery(c *gin.Context, key string, maxVal, defaultVal uint) (uint, error) {
	if strconv.IntSize == 32 { //  32-bit system
		val, err := GetUint32Query(c, key, uint32(maxVal), uint32(defaultVal))
		return uint(val), err
	}
	val, err := GetUint64Query(c, key, uint64(maxVal), uint64(defaultVal))
	return uint(val), err
}
