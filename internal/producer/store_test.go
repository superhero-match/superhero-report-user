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

package producer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/segmentio/kafka-go"

	"github.com/superhero-match/superhero-report-user/internal/producer/model"
)

var shouldGenerateEncodeError = false

func mockPublishStoreReport(producer *kafka.Writer, report model.Report) error {
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

func TestProducer_StoreReport(t *testing.T) {
	tests := []struct {
		mockProducer            producer
		report                  model.Report
		willGenerateEncodeError bool
		shouldReturnError       bool
		expected                error
	}{
		{
			mockProducer: producer{
				Producer:    nil,
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "id-2",
				Reason:          "unit testing",
				CreatedAt:       "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       false,
			expected:                nil,
		},
		{
			mockProducer: producer{
				Producer:    nil,
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "",
				ReportedUserID:  "id-2",
				Reason:          "unit testing",
				CreatedAt:       "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("the reporting user id is invalid"),
		},
		{
			mockProducer: producer{
				Producer:    nil,
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "",
				Reason:          "unit testing",
				CreatedAt:       "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("the reported user id is invalid"),
		},
		{
			mockProducer: producer{
				Producer:    nil,
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "id-2",
				Reason:          "",
				CreatedAt:       "2022-04-25T12:00:00",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("report reason is empty"),
		},
		{
			mockProducer: producer{
				Producer:    nil,
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "id-2",
				Reason:          "unit testing",
				CreatedAt:       "",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("report createdAt id is empty"),
		},
		{
			mockProducer: producer{
				Producer:    nil,
				storeReport: mockPublishStoreReport,
			},
			report: model.Report{
				ReportingUserID: "id-1",
				ReportedUserID:  "id-2",
				Reason:          "unit testing",
				CreatedAt:       "2022-04-25T12:00:00",
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

		err := test.mockProducer.StoreReport(test.report)
		if test.shouldReturnError && err.Error() != test.expected.Error() {
			t.Fatal(err)
		}

		if test.shouldReturnError == false && err != nil {
			t.Fatal(err)
		}
	}
}
