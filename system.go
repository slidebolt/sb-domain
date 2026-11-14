package domain

type Time struct {
	Timestamp int64 `json:"timestamp" doc:"Unix timestamp for this observation."`
	Hour      int   `json:"hour" doc:"Local hour from 0 through 23."`
	Minute    int   `json:"minute" doc:"Local minute from 0 through 59."`
	Day       int   `json:"day" doc:"Day of the month."`
	Week      int   `json:"week" doc:"ISO week number."`
	Month     int   `json:"month" doc:"Month number."`
	Year      int   `json:"year" doc:"Four-digit year."`
}

type Location struct {
	Latitude  float64 `json:"latitude" doc:"Latitude in decimal degrees."`
	Longitude float64 `json:"longitude" doc:"Longitude in decimal degrees."`
	Altitude  float64 `json:"altitude,omitempty" doc:"Altitude above sea level."`
	City      string  `json:"city,omitempty" doc:"Locality name."`
	Country   string  `json:"country,omitempty" doc:"Country name."`
	Timezone  string  `json:"timezone,omitempty" doc:"IANA time-zone identifier."`
}

type Group struct {
	MemberCount int      `json:"member_count" doc:"Number of entities currently in the group."`
	Status      string   `json:"status,omitempty" doc:"Aggregate group status."`
	Control     []string `json:"control,omitempty" doc:"Commands supported across group members."`
}

// Diagnostic is provider-owned operational state that has no controllable
// behavior. Its open shape is intentional: transports expose different health
// counters while sharing the same provider-neutral entity type.
type Diagnostic map[string]any
