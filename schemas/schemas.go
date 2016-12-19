package schemas

import "time"

type Schema interface {
	GetCachedFor() (d time.Duration)
}
