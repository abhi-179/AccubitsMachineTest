package repository

import (
	"AccubitsMachineTest/accubits.machine.test/logger"
	models "AccubitsMachineTest/accubits.machine.test/model"
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"

	"gorm.io/gorm"
)

type testRepository struct {
	DBConn *gorm.DB
}

func NewtestRepository(conn *gorm.DB) TestRepository {
	return &testRepository{
		DBConn: conn,
	}
}

/*********************************************Save Call logs****************************************/
func (r *testRepository) SaveData(ctx context.Context, req models.Req) (*models.Response, error) {
	logger.Logger.Info("Enter in to save call logs repository part.")
	Courses := make([]models.Course, 0)
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		colly.AllowedDomains("coursera.org", "www.coursera.org"),
	)
	c.OnHTML(".ais-InfiniteHits-list li", func(h *colly.HTMLElement) {
		courses := h.DOM
		course := models.Course{
			Title:    h.ChildText("h2"),
			Provider: courses.Find("div.cds-37.card-content.css-0.cds-38").Find("div.cds-37.card-info.css-0.cds-39.cds-grid-item").Find("div.cds-37.css-0.cds-38").Find("div.cds-37.css-0.cds-39.cds-grid-item").Find("span.cds-7.partner-name.css-1cxz0bb.cds-9").Text(),
			Level:    courses.Find("div.cds-37.difficulty-label-wrap.css-0.cds-38").Find("div.cds-37.css-0.cds-39.cds-grid-item").Find("span.cds-7.difficulty.css-1vjdgz.cds-9").Text(),
			Rating:   courses.Find("div.rating-enroll-wrapper").Find("div.info-item").Find("div.rc-Ratings.horizontal-box").Find("span.ratings-text").Text(),
		}
		Courses = append(Courses, course)

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting:", r.URL)
	})
	c.Visit("https://www.coursera.org/courses?query=" + req.Query + "&page=" + strconv.Itoa(req.Page))

	db := r.DBConn.CreateInBatches(&Courses, 30)
	if db.Error != nil {
		logger.Logger.WithError(db.Error).Error("Data not saved.")
		return &models.Response{IsSuccess: false, StatusCode: 400, Message: "Data not saved"}, nil
	}
	return &models.Response{IsSuccess: true, StatusCode: 200, Message: "Data saved successfully."}, nil
}

/*********************************************FetchAllCallLogs*****************************************/
func (r *testRepository) FetchData(ctx context.Context, source string, query url.Values) (*models.Response, error) {
	result := []models.Course{}
	db := r.DBConn.Where("title ILIKE ? or level ILIKE ? or rating ILIKE ?", source+"%", source+"%", source+"%").Find(&result)
	if db.Error != nil {
		logger.Logger.WithError(db.Error).Error("Data not found.")
		return &models.Response{IsSuccess: false, StatusCode: 400, Message: "Data not found"}, nil
	}
	logger.Logger.Info("Data found. ")
	return &models.Response{IsSuccess: true, StatusCode: 200, Message: "Data found.", Data: &result}, nil
}
