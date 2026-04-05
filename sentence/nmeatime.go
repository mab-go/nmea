package sentence

import (
	"fmt"
	"strconv"
	"strings"
)

// NMEATime represents the time of fix as reported in an NMEA sentence (typically UTC for GPS).
// The raw wire format is (h)hmmss.sss; this type unpacks that into discrete fields with no
// floating-point precision loss. Parsing rules match [SegmentParser.AsNMEATime].
//
// Second = 60 is accepted only for UTC 23:59 (a leap second) and round-trips through String().
type NMEATime struct {
	Hour        int
	Minute      int
	Second      int
	Millisecond int
}

// String returns the canonical wire encoding of t: two digits each for hour, minute, and second,
// a period, and exactly three fractional digits (e.g. "174800.864"). Inputs that omit the
// fractional part parse as millisecond zero; String always emits ".sss", so serialization
// normalizes that case to ".000".
//
// The 10-character shape (before any wider hour field) applies to values produced by successful
// parsing via [SegmentParser.AsNMEATime]; manually populated structs are not re-validated here.
func (t NMEATime) String() string {
	return fmt.Sprintf("%02d%02d%02d.%03d", t.Hour, t.Minute, t.Second, t.Millisecond)
}

// parseNMEATime parses a raw NMEA time segment string (format: (h)hmmss[.s[s[s]]]) into an
// NMEATime without any floating-point conversion. The integer part must be 4–6 digits. If more
// than three fractional digits are present, the remainder is truncated (not rounded).
func parseNMEATime(s string) (NMEATime, error) {
	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return NMEATime{}, fmt.Errorf("too many decimal points")
	}

	hour, minute, second, err := parseNMEATimeIntPart(parts[0])
	if err != nil {
		return NMEATime{}, err
	}

	ms, err := parseNMEATimeFracPart(parts)
	if err != nil {
		return NMEATime{}, err
	}

	if err := validateNMEATimeFields(hour, minute, second, ms); err != nil {
		return NMEATime{}, err
	}

	return NMEATime{Hour: hour, Minute: minute, Second: second, Millisecond: ms}, nil
}

// parseNMEATimeIntPart extracts hour, minute, and second from the integer portion of an NMEA time
// string (format: (h)hmmss — exactly 4, 5, or 6 ASCII digits).
func parseNMEATimeIntPart(intPart string) (hour, minute, second int, err error) {
	if len(intPart) < 4 || len(intPart) > 6 {
		return 0, 0, 0, fmt.Errorf("integer part must be 4 to 6 digits")
	}

	secondStr := intPart[len(intPart)-2:]
	minuteStr := intPart[len(intPart)-4 : len(intPart)-2]
	hourStr := intPart[:len(intPart)-4]
	if hourStr == "" {
		hourStr = "0"
	}

	hour, err = strconv.Atoi(hourStr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid hour")
	}

	minute, err = strconv.Atoi(minuteStr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid minute")
	}

	second, err = strconv.Atoi(secondStr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid second")
	}

	return hour, minute, second, nil
}

// parseNMEATimeFracPart extracts milliseconds from the split parts of an NMEA time string. If no
// fractional part is present (len(parts) == 1), it returns 0. If the fractional run is longer than
// three digits, only the first three are used; the rest are ignored.
func parseNMEATimeFracPart(parts []string) (int, error) {
	if len(parts) < 2 {
		return 0, nil
	}

	fracStr := parts[1]
	for len(fracStr) < 3 {
		fracStr += "0"
	}

	ms, err := strconv.Atoi(fracStr[:3])
	if err != nil {
		return 0, fmt.Errorf("invalid fractional seconds")
	}

	return ms, nil
}

// validateNMEATimeFields checks that each field is within its valid NMEA range. Second may be 60
// only for UTC 23:59 (leap second).
func validateNMEATimeFields(hour, minute, second, ms int) error {
	if hour < 0 || hour > 23 {
		return fmt.Errorf("hour %d out of range [0, 23]", hour)
	}
	if minute < 0 || minute > 59 {
		return fmt.Errorf("minute %d out of range [0, 59]", minute)
	}
	if err := validateNMEATimeSecond(hour, minute, second); err != nil {
		return err
	}
	if ms < 0 || ms > 999 {
		return fmt.Errorf("millisecond %d out of range [0, 999]", ms)
	}

	return nil
}

func validateNMEATimeSecond(hour, minute, second int) error {
	if second < 0 || second > 60 {
		return fmt.Errorf("second %d out of range [0, 60]", second)
	}
	if second == 60 && (hour != 23 || minute != 59) {
		return fmt.Errorf("second 60 is only valid at UTC 23:59 (leap second)")
	}

	return nil
}
