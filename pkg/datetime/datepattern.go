package datetime

// DatePattern provides common date formatting patterns
type DatePattern struct{}

const (
	ISO8601        = "2006-01-02T15:04:05Z07:00"           // ISO 8601 combined date and time with timezone
	ISO8601Date    = "2006-01-02"                          // ISO 8601 date only
	ISO8601Time    = "15:04:05"                            // ISO 8601 time only
	RFC1123        = "Mon, 02 Jan 2006 15:04:05 MST"       // RFC 1123 date and time
	RFC1123Z       = "Mon, 02 Jan 2006 15:04:05 -0700"     // RFC 1123 with numeric timezone
	RFC3339        = "2006-01-02T15:04:05Z07:00"           // RFC 3339 combined date and time with timezone
	RFC3339Nano    = "2006-01-02T15:04:05.999999999Z07:00" // RFC 3339 with nanosecond precision
	ANSIC          = "Mon Jan _2 15:04:05 2006"            // ANSI C's asctime() format
	UnixDate       = "Mon Jan _2 15:04:05 MST 2006"        // Unix date format
	RubyDate       = "Mon Jan 02 15:04:05 -0700 2006"      // Ruby's date format
	Kitchen        = "3:04PM"                              // The time in 12-hour AM/PM format
	Stamp          = "Jan _2 15:04:05"                     // Date and time without year
	StampMilli     = "Jan _2 15:04:05.000"                 // Date and time with millisecond precision
	StampMicro     = "Jan _2 15:04:05.000000"              // Date and time with microsecond precision
	StampNano      = "Jan _2 15:04:05.000000000"           // Date and time with nanosecond precision
	CustomDateTime = "02-01-2006 15:04:05"                 // Custom date and time format (DD-MM-YYYY HH:MM:SS)
	CustomDate     = "02-01-2006"                          // Custom date format (DD-MM-YYYY)
	CustomTime     = "15:04:05"                            // Custom time format (HH:MM:SS)
)
