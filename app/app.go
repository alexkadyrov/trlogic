package app

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	DB     *gorm.DB
	Logger *logrus.Logger
)
