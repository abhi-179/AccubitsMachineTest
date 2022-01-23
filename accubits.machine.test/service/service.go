package service

import (
	"AccubitsMachineTest/accubits.machine.test/logger"
	models "AccubitsMachineTest/accubits.machine.test/model"
	repo "AccubitsMachineTest/accubits.machine.test/repository"

	"context"
	"net/url"
)

//	"encoding/json"

type testUsecase struct {
	repository repo.TestRepository
}

func NewtestUsecase(repo repo.TestRepository) TestUsecase {
	return &testUsecase{
		repository: repo,
	}
}

/*******************************************SaveCallLogs*******************************************/
func (r *testUsecase) SaveData(ctx context.Context, req models.Req) (*models.Response, error) {
	logger.Logger.Info("Request received from save data service part.")
	return r.repository.SaveData(ctx, req)
}

/**********************************************Fetch All Call Logs**********************************/
func (r *testUsecase) FetchData(ctx context.Context, source string, query url.Values) (*models.Response, error) {
	logger.Logger.Info("Request received from fetch data service part.")
	return r.repository.FetchData(ctx, source, query)
}
