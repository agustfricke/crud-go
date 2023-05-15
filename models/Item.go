package models

import "gorm.io/gorm"

type Item struct {
  gorm.Model
  Body string `gorm:"type:varchar(100);not null"`
  Done bool `default:"false"`
}
