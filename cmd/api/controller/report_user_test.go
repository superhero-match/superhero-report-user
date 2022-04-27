/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-report-user/cmd/api/model"
	"github.com/superhero-match/superhero-report-user/cmd/api/service/mapper"
	pm "github.com/superhero-match/superhero-report-user/internal/producer/model"
)

type mockProducer struct {
	storeReport func(producer *kafka.Writer, report pm.Report) error
}

func (m *mockProducer) Close() error {
	return nil
}

func (m *mockProducer) StoreReport(report pm.Report) error {
	return m.storeReport(nil, report)
}

func mockPublishStoreReport(producer *kafka.Writer, report pm.Report) error {
	err := report.Validate()
	if err != nil {
		return err
	}

	var sb bytes.Buffer

	err = json.NewEncoder(&sb).Encode(report)
	if err != nil {
		return fmt.Errorf("encoder error")
	}

	return nil
}

type mockService struct {
	mProducer mockProducer
}

func (srv *mockService) Close() error {
	return srv.mProducer.Close()
}

func (srv *mockService) StoreReport(report model.Report) error {
	return srv.mProducer.StoreReport(mapper.MapAPIReportToProducer(
		report,
		time.Now().UTC().Format("2006-01-02T15:04:05"),
	))
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func TestController_ReportUser(t *testing.T) {
	mockProd := mockProducer{
		storeReport: mockPublishStoreReport,
	}

	mService := &mockService{
		mProducer: mockProd,
	}

	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatal(err)
	}

	defer logger.Sync()

	mockController := &Controller{
		Service:    mService,
		Logger:     logger,
		TimeFormat: "2006-01-02T15:04:05",
	}

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(
		ctx,
		map[string]interface{}{
			"reportingUserID": "test-id-1",
			"reportedUserID":  "test-id-2",
			"reason":          "unit testing",
		},
	)

	mockController.ReportUser(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}
