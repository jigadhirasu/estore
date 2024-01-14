package logger

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

func Tags(tags string) zap.Field {
	return zap.String("tags", tags)
}

func JSON(key string, data any) zap.Field {
	b, _ := json.Marshal(data)
	return zap.Any(key, json.RawMessage(b))
}

func Runtime(d time.Duration) zap.Field {
	return zap.Int64("runtime", int64(d/time.Millisecond))
}
