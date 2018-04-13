package api

import (
	"wechatHimsAPI/lib"
	"wechatHimsAPI/model"

	"github.com/gin-gonic/gin"
)

func init() {
	patientInfo := Router.Group("/patient_infos")

	patientInfo.GET("/get_queue_data", func(c *gin.Context) {
		result := gin.H{}
		patientInfos := []model.PatientInfo{}
		currentUser := &model.User{HospitalID: 2}
		currentTime := c.Query("current_time")
		sql := `select * from (select pi.id,pi.examine_finish_flag,pi.hospital_no,pi.film_no,pi.patient_name,pi.patient_sex,pi.patient_phone,pi.project_id,pi.project_name,
		    pi.body_description,pi.diagnose_info,pi.state,pi.appointment_at,pi.order_code,r.can_print,r.is_print,pid.patient_name as en_patient_name from patient_infos pi
		    left join reports r on pi.id = r.patient_info_id left join patient_info_dicoms pid on pi.id = pid.patient_info_id
		    where pi.hospital_id = ?
		    union
		    select '' as id,'' as hospital_no,'' as examine_finish_flag,'' as film_no,o.patient_name,null as patient_sex,o.patient_phone,o.project_id,o.project_name,
		    o.body_description,o.diagnose_info,o.state as state,o.appointment_at,o.order_code,false as can_print,false as is_print,'' as en_patient_name from orders o
		    left join patient_infos tpi on o.hospital_id = tpi.hospital_id and o.order_code = tpi.order_code
		    where tpi.id is null and o.hospital_id = ?) p
			where p.state in (0,1,2,3) and DATE(p.appointment_at) = DATE(?)`
		err := lib.DB.Raw(sql, currentUser.HospitalID, currentUser.HospitalID, currentTime).Scan(&patientInfos).Error

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
}