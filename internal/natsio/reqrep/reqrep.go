package reqrep

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/model/requestmodel"
	"github.com/vothanhdo2602/hicon/external/util/natstil"
	"github.com/vothanhdo2602/hicon/external/util/pjson"
	"github.com/vothanhdo2602/hicon/internal/natsio"
	"time"
)

func UpdateConfig(ctx context.Context, data *requestmodel.UpdateConfig) error {
	var (
		resp natstil.BaseResponse[any]
	)

	req, _ := pjson.Marshal(ctx, data)
	msg, err := natsio.GetNC().Request(natsio.GetUpdateConfigSubject(), req, 10*time.Second)
	if err != nil {
		return err
	}

	err = pjson.Unmarshal(ctx, msg.Data, &resp)
	if err != nil {
		return err
	}

	return nil
}

func FindByPrimaryKeys(ctx context.Context, data *requestmodel.FindByPrimaryKeys) (*natstil.BaseResponse[any], error) {
	var (
		resp natstil.BaseResponse[any]
	)

	req, _ := pjson.Marshal(ctx, data)
	msg, err := natsio.GetNC().Request(natsio.GetFindByPrimaryKeysSubject(), req, 10*time.Second)
	if err != nil {
		return nil, err
	}

	err = pjson.Unmarshal(ctx, msg.Data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
