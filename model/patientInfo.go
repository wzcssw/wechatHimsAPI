package model

import (
	"errors"
	"strconv"
	"time"
	"wechatHimsAPI/lib"

	"github.com/gin-gonic/gin"
)

type PatientInfo struct {
	ID                uint64 `gorm:"primary_key"`
	ExamineFinishFlag int
	HospitalNo        string
	FilmNo            string
	PatientName       string
	PatientSex        int
	PatientPhone      string
	ProjectID         int
	ProjectName       string
	BodyDescription   string
	DiagnoseInfo      string
	State             int
	AppointmentAt     time.Time
	FinishedAt        time.Time
	OrderCode         string
	CanPrint          int
	IsPrint           int
	EnPatientName     string
}

var patientInfoStates = [6]string{"unpay", "wait_examine", "examining", "finish", "late", "cancel"}

func (patientInfo *PatientInfo) StateZH() string {
	return patientInfoStates[patientInfo.State]
}

//  这里 finished_at 字段可能没有查询
func GetQueue(hospitalID uint, currentTime string) ([]PatientInfo, error) {
	patientInfos := []PatientInfo{}
	sql := `select * from (select pi.id,pi.examine_finish_flag,pi.hospital_no,pi.film_no,pi.patient_name,pi.patient_sex,pi.patient_phone,pi.project_id,pi.project_name,
		pi.body_description,pi.diagnose_info,pi.state,pi.appointment_at,pi.finished_at,pi.order_code,r.can_print,r.is_print,pid.patient_name as en_patient_name from patient_infos pi
		left join reports r on pi.id = r.patient_info_id left join patient_info_dicoms pid on pi.id = pid.patient_info_id
		where pi.hospital_id = ?
		union
		select '' as id,'' as hospital_no,'' as examine_finish_flag,'' as film_no,o.patient_name,null as patient_sex,o.patient_phone,o.project_id,o.project_name,
		o.body_description,o.diagnose_info,o.state as state,o.appointment_at,tpi.finished_at,o.order_code,false as can_print,false as is_print,'' as en_patient_name from orders o
		left join patient_infos tpi on o.hospital_id = tpi.hospital_id and o.order_code = tpi.order_code
		where tpi.id is null and o.hospital_id = ?) p
		where p.state in (0,1,2,3) and DATE(p.appointment_at) = DATE(?)`
	err := lib.DB.Raw(sql, hospitalID, hospitalID, currentTime).Scan(&patientInfos).Error
	return patientInfos, err
}

func FinishExamine(ID int) error {
	patientInfo := PatientInfo{}
	orm := lib.DB.Where("id = ?", ID).First(&patientInfo)
	if patientInfo.ID == 0 {
		return errors.New("未找到id为" + strconv.Itoa(ID) + "的patient_info条目")
	} else {
		rowsAffected := orm.Update(gin.H{"state": 3, "finished_at": time.Now()}).RowsAffected
		if rowsAffected < 1 {
			return errors.New("更新失败: RowsAffected为" + strconv.Itoa(int(rowsAffected)))
		} else {
			return nil
		}
	}
}
