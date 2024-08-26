package common

import "time"

// UintPtr converts a float64 to a pointer to uint.
func UintPtr(f float64) *uint {
	i := uint(f)
	return &i
}

// TimePtr converts a string to a pointer to time.Time.
func TimePtr(s string) *time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil
	}
	return &t
}
