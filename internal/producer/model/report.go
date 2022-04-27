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

package model

import "fmt"

var (
	ErrInvalidReportingUserID = fmt.Errorf("the reporting user id is invalid")
	ErrInvalidReportedUserID  = fmt.Errorf("the reported user id is invalid")
	ErrReasonIsEmpty          = fmt.Errorf("report reason is empty")
	ErrCreatedAtIsEmpty       = fmt.Errorf("report createdAt id is empty")
)

// Report holds data about reported user.
type Report struct {
	ReportingUserID string `json:"reportingUserID"`
	ReportedUserID  string `json:"reportedUserID"`
	Reason          string `json:"reason"`
	CreatedAt       string `json:"createdAt"`
}

// Validate validates ProfilePicture data.
func (r Report) Validate() error {
	if len(r.ReportingUserID) == 0 {
		return ErrInvalidReportingUserID
	}

	if len(r.ReportedUserID) == 0 {
		return ErrInvalidReportedUserID
	}

	if len(r.Reason) == 0 {
		return ErrReasonIsEmpty
	}

	if len(r.CreatedAt) == 0 {
		return ErrCreatedAtIsEmpty
	}

	return nil
}
