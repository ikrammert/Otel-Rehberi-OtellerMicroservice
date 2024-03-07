package controllers

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	RefId = logrus.Fields{
		"ref-id": uuid.NewString(), //id for controllers
	}
)
