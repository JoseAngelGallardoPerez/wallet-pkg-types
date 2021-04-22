package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type DataJSON map[string]interface{}

func (d DataJSON) Set(key string, value interface{}) {
	d[key] = value
}

func (d DataJSON) Get(key string) (value interface{}, exists bool) {
	if value, exists = d[key]; exists {
		return value, exists
	}
	return nil, exists
}

// MustGet returns the value for the given key if it exists, otherwise it panics.
func (d DataJSON) MustGet(key string) interface{} {
	if value, exists := d.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

// GetString returns the value associated with the key as a string.
func (d DataJSON) GetString(key string) (s string) {
	if val, ok := d.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (d DataJSON) GetBool(key string) (b bool) {
	if val, ok := d.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (d DataJSON) GetInt(key string) (i int) {
	if val, ok := d.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (d DataJSON) GetInt64(key string) (i64 int64) {
	if val, ok := d.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (d DataJSON) GetFloat64(key string) (f64 float64) {
	if val, ok := d.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

// GetTime returns the value associated with the key as time.
func (d DataJSON) GetTime(key string) (t time.Time) {
	if val, ok := d.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (d DataJSON) GetDuration(key string) (duration time.Duration) {
	if val, ok := d.Get(key); ok && val != nil {
		duration, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (d DataJSON) GetStringSlice(key string) (ss []string) {
	if val, ok := d.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (d DataJSON) GetStringMap(key string) (sm map[string]interface{}) {
	if val, ok := d.Get(key); ok && val != nil {
		sm, _ = val.(map[string]interface{})
	}
	return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (d DataJSON) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := d.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (d DataJSON) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val, ok := d.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}

func (d DataJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(d))
}

func (d *DataJSON) UnmarshalJSON(data []byte) error {
	ref := map[string]interface{}{}
	err := json.Unmarshal(data, &ref)
	if err != nil {
		return err
	}
	*d = ref
	return nil
}

// Scan implements the sql.Scanner interface for database deserialization.
func (d *DataJSON) Scan(value interface{}) error {
	if value == nil {
		*d = nil
		return nil
	}
	if d == nil {
		*d = DataJSON{}
	}
	switch v := value.(type) {
	case []byte:
		return d.UnmarshalJSON(v)
	case string:
		return nil
	}
	return errors.New("invalid scan source")
}

func (d DataJSON) String() string {
	val, err := d.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(val)
}

// Value implements the driver.Valuer interface for database serialization.
func (d DataJSON) Value() (driver.Value, error) {
	return d.MarshalJSON()
}

func (d DataJSON) IsNull() bool {
	return d == nil
}

func (d DataJSON) IsEmpty() bool {
	return d.IsNull() || len(d) == 0
}

func (d DataJSON) Equals(meta DataJSON) bool {
	if &d == &meta {
		return true
	}
	if len(d) != len(meta) {
		return false
	}
	for k, v := range d {
		if v2, ok := meta[k]; ok {
			if v != v2 {
				return false
			}
		}
	}
	return true
}
