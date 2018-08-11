package eager

import (
	"github.com/spf13/cast"
)

//
func (c *config) Get(key string) interface{} {
	storge, ok := c.storge.(*map[string]interface{})
	if !ok {
		return nil
	}
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	return (*storge)[key]
}

//
func (c *config) GetE(key string) (interface{}, error) {
	storge, ok := c.storge.(*map[string]interface{})
	if !ok {
		return nil, ErrAssert
	}
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	return (*storge)[key], nil
}

//
func (c *config) GetString(key string) string {
	return cast.ToString(c.Get(key))
}

//
func (c *config) GetStringE(key string) (string, error) {
	value, err := c.GetE(key)
	if err != nil {
		return "", ErrAssert
	}

	return cast.ToStringE(value)
}

//
func (c *config) GetInt8(key string) int8 {
	return cast.ToInt8(c.Get(key))
}

//
func (c *config) GetInt8E(key string) (int8, error) {
	value, err := c.GetE(key)
	if err != nil {
		return 0, ErrAssert
	}

	return cast.ToInt8E(value)
}

//
func (c *config) GetInt16(key string) int16 {
	return cast.ToInt16(c.Get(key))
}

//
func (c *config) GetInt16E(key string) (int16, error) {
	value, err := c.GetE(key)
	if err != nil {
		return 0, ErrAssert
	}

	return cast.ToInt16E(value)
}

//
func (c *config) GetInt32(key string) int32 {
	return cast.ToInt32(c.Get(key))
}

//
func (c *config) GetInt32E(key string) (int32, error) {
	value, err := c.GetE(key)
	if err != nil {
		return 0, ErrAssert
	}

	return cast.ToInt32E(value)
}

//
func (c *config) GetInt64(key string) int64 {
	return cast.ToInt64(c.Get(key))
}

//
func (c *config) GetInt64E(key string) (int64, error) {
	value, err := c.GetE(key)
	if err != nil {
		return 0, ErrAssert
	}

	return cast.ToInt64E(value)
}

//
func (c *config) GetFloat32(key string) float32 {
	return cast.ToFloat32(c.Get(key))
}

//
func (c *config) GetFloat32E(key string) (float32, error) {
	value, err := c.GetE(key)
	if err != nil {
		return 0, ErrAssert
	}

	return cast.ToFloat32E(value)
}

//
func (c *config) GetFloat64(key string) float64 {
	return cast.ToFloat64(c.Get(key))
}

//
func (c *config) GetFloat64E(key string) (float64, error) {
	value, err := c.GetE(key)
	if err != nil {
		return 0, ErrAssert
	}

	return cast.ToFloat64E(value)
}

//
func (c *config) GetBool(key string) bool {
	return cast.ToBool(c.Get(key))
}

//
func (c *config) GetBoolE(key string) (bool, error) {
	value, err := c.GetE(key)
	if err != nil {
		return false, ErrAssert
	}

	return cast.ToBoolE(value)
}
