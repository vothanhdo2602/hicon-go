package pjson

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/vothanhdo2602/hicon/external/util/log"
)

func Marshal(ctx context.Context, data interface{}) ([]byte, error) {
	r, err := json.Marshal(data)
	if err != nil {
		log.WithCtx(ctx).Error(err.Error())
	}
	return r, err
}

func Unmarshal(ctx context.Context, data []byte, output interface{}) error {
	err := json.Unmarshal(data, output)
	if err != nil {
		log.WithCtx(ctx).Error(err.Error())
		return err
	}
	return nil
}
