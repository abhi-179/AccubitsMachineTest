package service

import (
	"context"
	"net/url"
	models "AccubitsMachineTest/accubits.machine.test/model"
)

type TestUsecase interface {
	SaveData(ctx context.Context, req models.Req) (*models.Response, error)
	FetchData(ctx context.Context, source string, query url.Values) (*models.Response, error)
}
