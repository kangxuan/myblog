package main

import (
	"myblog/models"
	"myblog/settings"
)

func main() {
	settings.SetUp()
	models.SetUp()
}
