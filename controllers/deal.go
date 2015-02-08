package controllers

import (
  "time"
  "fmt"
  "strconv"
  "encoding/json"
  "jostler/models"
)


type TickController struct {
  baseController
}

func (this *TickController) AddDeal() {
  var tick models.Tick
  err := json.Unmarshal(this.Ctx.Input.RequestBody, &tick)
  if err != nil {
    this.Response(ERROR, fmt.Sprintf("Incorrect JSON: %v", err))
    return
  }
  if invalid := this.Validate(&tick); invalid != nil {
    this.Response(FAIL, invalid)
    return
  }
  tick.User = this.user
  tick.Created = time.Now()
  tick.Modified = tick.Created
  _, err = this.db.Insert(&tick)
  if err != nil {
    this.Response(ERROR, err)
  } else {
    this.Response(SUCCESS, map[string]interface{}{"tick": tick})
  }
}

func (this *TickController) AllDeals() {
  var ticks []*models.Tick
  this.db.QueryTable("tick").Filter("user_id", this.user.Id).All(&ticks);
  this.Response(SUCCESS, map[string]interface{}{"ticks": ticks})
}

func (this *TickController) DealInfo() {
  id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
  if err != nil {
    this.Response(FAIL, map[string]interface{}{"id": "Must be an integer"})
    return
  }
  tick := models.Tick{Id: id}
  err = this.db.Read(&tick)
  if err != nil {
    this.Response(FAIL, map[string]interface{}{"id": fmt.Sprintf("Tick with id %v not found", id)})
  } else {
    this.Response(SUCCESS, map[string]interface{}{"tick": tick})
  }
}

func (this *TickController) DealUpdate() {
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  tick := models.Tick{Id: id}
  err := this.db.Read(&tick)
  if err != nil {
    this.Response(FAIL, map[string]interface{}{"id": fmt.Sprintf("Tick with id %v not found", id)})
    return
  }
  err = json.Unmarshal(this.Ctx.Input.RequestBody, &tick)
  if err != nil {
    this.Response(ERROR, fmt.Sprintf("Incorrect JSON: %v", err))
    return
  }
  if invalid := this.Validate(&tick); invalid != nil {
    this.Response(FAIL, invalid)
    return
  }
  tick.Modified = time.Now()
  if _, err = this.db.Update(&tick); err == nil {
    this.Response(SUCCESS, map[string]interface{}{"tick": tick})
  } else {
    this.Response(ERROR, err)
  }
}

func (this *TickController) DealDelete() {
  id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
  if _, err := this.db.Delete(&models.Deal{Id: id}); err == nil {
    this.Response(SUCCESS, nil)
  } else {
    this.Response(ERROR, err)
  }
}