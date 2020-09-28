/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package mapper

import (
	"github.com/superhero-match/superhero-report-user/cmd/api/model"
	pm "github.com/superhero-match/superhero-report-user/internal/producer/model"
)

// MapAPIReportToProducer maps API Report model to Producer Report model.
func MapAPIReportToProducer(report model.Report, createdAt string) pm.Report {
	return pm.Report{
		ReportingUserID: report.ReportingUserID,
		ReportedUserID:  report.ReportedUserID,
		Reason:          report.Reason,
		CreatedAt:       createdAt,
	}
}
