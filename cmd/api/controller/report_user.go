/*
  Copyright (C) 2019 - 2021 MWSOFT
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
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ctrl "github.com/superhero-match/superhero-report-user/cmd/api/model"
	"go.uber.org/zap"
)

// Match publishes new match on Kafka for it to be save to DB.
func (ctl *Controller) ReportUser(c *gin.Context) {
	var req ctrl.Report

	err := c.BindJSON(&req)
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type Report",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	err = ctl.Service.StoreReport(req)
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed while executing service.StoreReport()",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func checkError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})

		return true
	}

	return false
}