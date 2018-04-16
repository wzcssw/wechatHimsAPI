package api

import (
	"strconv"
	"wechatHimsAPI/model"

	"github.com/gin-gonic/gin"
)

func init() {
	patientInfo := Router.Group("/patient_infos")

	patientInfo.GET("/get_queue_data", func(c *gin.Context) {
		result := gin.H{}
		currentUser := model.CurrentUser(c)
		currentTime := c.Query("current_time")
		patientInfos, err := model.GetQueue(currentUser.HospitalID, currentTime)
		if err == nil {
			result["data"] = patientInfos
			result["msg"] = "OK"
			result["success"] = true
		} else {
			result["msg"] = err.Error()
			result["success"] = false
		}
		c.JSON(200, result)
	})

	patientInfo.POST("/finish_examine", func(c *gin.Context) {
		result := gin.H{}
		patientInfoIDStr := c.PostForm("id")
		patientInfoIDInt, _ := strconv.Atoi(patientInfoIDStr)
		err := model.FinishExamine(patientInfoIDInt)
		if err == nil {
			result["msg"] = "OK"
			result["success"] = true
		} else {
			result["msg"] = err.Error()
			result["success"] = false
		}
		c.JSON(200, result)
	})
}
