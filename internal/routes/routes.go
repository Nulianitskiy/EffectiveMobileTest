package routes

import (
	"GoTimeTracker/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/allPeople", controller.GetAllPeople)
	r.POST("/people", controller.AddPeople)
	r.PUT("/people", controller.UpdatePeople)
	r.DELETE("/people", controller.DeletePeople)

	r.POST("/task", controller.AddTask)
	r.PUT("/taskAssign", controller.AssignPeopleOnTask)
	r.PUT("/taskStart", controller.StartTask)
	r.PUT("/taskEnd", controller.EndTask)
	r.GET("/task", controller.GetTasks)

}
