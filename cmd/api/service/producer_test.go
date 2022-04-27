package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/superhero-match/superhero-report-user/cmd/api/model"
	pm "github.com/superhero-match/superhero-report-user/internal/producer/model"
	"testing"
)

var shouldGenerateEncodeError = false

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

	var encoderValue interface{}
	encoderValue = report

	if shouldGenerateEncodeError {
		encoderValue = make(chan int)
	}

	err = json.NewEncoder(&sb).Encode(encoderValue)
	if err != nil {
		return fmt.Errorf("encoder error")
	}

	return nil
}

func TestService_StoreReport(t *testing.T) {
	producer := &mockProducer{storeReport: mockPublishStoreReport}
	mockService := service{
		Producer:   producer,
		TimeFormat: "2006-01-02T15:04:05",
	}

	tests := []struct {
		producer                mockProducer
		report                  model.Report
		willGenerateEncodeError bool
		shouldReturnError       bool
		expected                error
	}{
		{
			producer: mockProducer{
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "id-2",
				Reason:          "unit testing",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       false,
			expected:                nil,
		},
		{
			producer: mockProducer{
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "",
				ReportedUserID:  "id-2",
				Reason:          "unit testing",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("the reporting user id is invalid"),
		},
		{
			producer: mockProducer{
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "",
				Reason:          "unit testing",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("the reported user id is invalid"),
		},
		{
			producer: mockProducer{
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "id-2",
				Reason:          "",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("report reason is empty"),
		},
		{
			producer: mockProducer{
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "id-2",
				Reason:          "unit testing",
			},
			willGenerateEncodeError: true,
			shouldReturnError:       true,
			expected:                fmt.Errorf("encoder error"),
		},
	}

	for _, test := range tests {
		shouldGenerateEncodeError = false

		if test.willGenerateEncodeError {
			shouldGenerateEncodeError = true
		}

		err := mockService.StoreReport(test.report)
		if test.shouldReturnError && err.Error() != test.expected.Error() {
			t.Fatal(err)
		}

		if test.shouldReturnError == false && err != nil {
			t.Fatal(err)
		}
	}
}
