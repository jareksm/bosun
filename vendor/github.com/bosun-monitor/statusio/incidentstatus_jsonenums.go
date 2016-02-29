// generated by jsonenums -type=IncidentStatus; DO NOT EDIT

package statusio

import (
	"encoding/json"
	"fmt"
)

var (
	_IncidentStatusNameToValue = map[string]IncidentStatus{
		"Investigating": Investigating,
		"Identified":    Identified,
		"Monitoring":    Monitoring,
		"Resolved":      Resolved,
		"PostMortem":    PostMortem,
	}

	_IncidentStatusValueToName = map[IncidentStatus]string{
		Investigating: "Investigating",
		Identified:    "Identified",
		Monitoring:    "Monitoring",
		Resolved:      "Resolved",
		PostMortem:    "PostMortem",
	}
)

func init() {
	var v IncidentStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_IncidentStatusNameToValue = map[string]IncidentStatus{
			interface{}(Investigating).(fmt.Stringer).String(): Investigating,
			interface{}(Identified).(fmt.Stringer).String():    Identified,
			interface{}(Monitoring).(fmt.Stringer).String():    Monitoring,
			interface{}(Resolved).(fmt.Stringer).String():      Resolved,
			interface{}(PostMortem).(fmt.Stringer).String():    PostMortem,
		}
	}
}

// MarshalJSON is generated so IncidentStatus satisfies json.Marshaler.
func (r IncidentStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _IncidentStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid IncidentStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so IncidentStatus satisfies json.Unmarshaler.
func (r *IncidentStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("IncidentStatus should be a string, got %s", data)
	}
	v, ok := _IncidentStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid IncidentStatus %q", s)
	}
	*r = v
	return nil
}