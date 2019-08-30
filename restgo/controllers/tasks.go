package controllers

import (
	"encoding/json"
	"net/http"

	"../common"
	"../models"
)

var Tasks = new(tasksController)

type tasksController struct{}

func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Create(t.Name, t.Desc)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusCreated)
}
func (tc *tasksController) Get(w http.ResponseWriter, r *http.Request) {

}
func (tc *tasksController) Show(w http.ResponseWriter, r *http.Request)   {}
func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request) {}
func (tc *tasksController) Delete(w http.ResponseWriter, r *http.Request) {}
