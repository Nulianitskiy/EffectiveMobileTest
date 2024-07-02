package controller

import (
	dbase "GoTimeTracker/database"
	"GoTimeTracker/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type FilterParams struct {
	Page     int                    `json:"page" binding:"required"`
	PageSize int                    `json:"page_size" binding:"required"`
	Filters  map[string]interface{} `json:"filters"`
}

func GetAllPeople(ctx *gin.Context) {
	var filterParams FilterParams
	if err := ctx.ShouldBindJSON(&filterParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := dbase.GetInstance()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	people, err := db.GetAllPeople(filterParams.Page, filterParams.PageSize, filterParams.Filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, people)
}

func AddPeople(ctx *gin.Context) {
	passport := strings.Split(ctx.Query("passportNumber"), " ")
	serie, _ := strconv.Atoi(passport[0])
	number, _ := strconv.Atoi(passport[1])

	people := model.People{
		PassportSerie:  serie,
		PassportNumber: number,
	}

	db, err := dbase.GetInstance()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	err = db.AddPeople(people)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func UpdatePeople(ctx *gin.Context) {
	var p model.People
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := dbase.GetInstance()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	err = db.UpdatePeople(p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func DeletePeople(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := dbase.GetInstance()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
		return
	}

	err = db.DeletePeople(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
