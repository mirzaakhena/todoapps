package payload

import (
	"todoapps/shared/driver"
)

type Payload struct {
	Data      any                    `json:"data"`
	Publisher driver.ApplicationData `json:"publisher"`
	TraceID   string                 `json:"traceId"`
}
