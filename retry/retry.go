package retry

import "time"

const (
	DEFAULT_RETRY_INTERVAL = 3 * time.Second
	DEFAULT_RETRY_TIMES    = 3
)
