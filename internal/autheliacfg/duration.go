package autheliacfg

import (
	"github.com/brunnels/authelia-oidc-operator/pkg/duration"
)

// Duration is an alias for pkg/duration.Duration for backwards compatibility.
type Duration = duration.Duration

// StandardizeDurationString is an alias for pkg/duration.StandardizeDurationString for backwards compatibility.
var StandardizeDurationString = duration.StandardizeDurationString

// IsStringInSlice is an alias for pkg/duration.IsStringInSlice for backwards compatibility.
var IsStringInSlice = duration.IsStringInSlice

// Duration unit types.
const (
	DurationUnitDays   = duration.DurationUnitDays
	DurationUnitWeeks  = duration.DurationUnitWeeks
	DurationUnitMonths = duration.DurationUnitMonths
	DurationUnitYears  = duration.DurationUnitYears
)

// Number of hours in particular measurements of time.
const (
	HoursInDay   = duration.HoursInDay
	HoursInWeek  = duration.HoursInWeek
	HoursInMonth = duration.HoursInMonth
	HoursInYear  = duration.HoursInYear
)

