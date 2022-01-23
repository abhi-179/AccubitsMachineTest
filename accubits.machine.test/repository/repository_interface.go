package repository

import (
	models "AccubitsMachineTest/accubits.machine.test/model"
	"context"
	"net/url"
)

type TestRepository interface {
	SaveData(ctx context.Context, req models.Req) (*models.Response, error)
	FetchData(ctx context.Context, source string, query url.Values) (*models.Response, error)
}
