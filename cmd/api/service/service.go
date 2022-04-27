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

package service

import (
	"github.com/superhero-match/superhero-report-user/cmd/api/model"
	"github.com/superhero-match/superhero-report-user/internal/config"
	"github.com/superhero-match/superhero-report-user/internal/producer"
)

// Service interface defines service methods.
type Service interface {
	StoreReport(report model.Report) error
}

// service holds all the different services that are used when handling request.
type service struct {
	Producer   producer.Producer
	TimeFormat string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (Service, error) {
	return &service{
		Producer:   producer.NewProducer(cfg),
		TimeFormat: cfg.App.TimeFormat,
	}, nil
}
