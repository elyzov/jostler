package controllers

import (
  // "strconv"
  // "encoding/json"
  "jostler/models"
)


type UserController struct {
  restController
}

func (this *UserController) NestPrepare() {
  this.model = &models.UserModel{}
  this.model.Prepare()
}