package schemas

import "time"

// Schema represents the results of an API request
type Schema interface {
	GetCachedFor() (d time.Duration)
}
