package types

import "time"

// Declare your type
type Duration time.Duration

// Make it implement flag.Value
func (d *Duration) Set(v string) error {
	parsed, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*d = Duration(parsed)
	return nil
}

func (d *Duration) String() string {
	duration := time.Duration(*d)
	return duration.String()
}

// Declare your type
type DateTime struct {
	time.Time
}

// Make it implement flag.Value
func (d *DateTime) Set(v string) error {
	parsed, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	//parsed.
	*d = DateTime{parsed}

	//*d.
	return nil
}

func (d *DateTime) String() string {
	//duration := time.Duration(*d)
	return d.Time.String()
}

func (_ *DateTime) IsDefault() bool {
	return true
}
