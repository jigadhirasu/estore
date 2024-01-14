package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
tags: this.tags,
uuid: this.uuid,
ip: this.ip,
instance_name: this.instanceName,
version: this.version,
created_time: this.createdTime,
*/

type G struct {
	A string
	B string
	C string
}

func (g G) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("A", g.A)
	enc.AddString("B", g.B)
	enc.AddString("C", g.C)
	return nil
}

func TestZap(t *testing.T) {

	a()

}

func a() {
	al := Access().With(zap.String("AAA", "OK"))

	al.Info("/api/v2/role", zap.String("method", "POST"))

	al.Error("/api/v2/role", zap.String("method", "POST"))

	ol := Out()

	ol.Debug("/api/v2/role", zap.String("method", "POST"))

}
