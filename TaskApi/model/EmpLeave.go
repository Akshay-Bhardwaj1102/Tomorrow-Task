package model

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ImageName    string `json:"imageName"`
	ImageContent []byte `json:"-"`
	EmpLeaveID   uint   ` json:"-"`
}

type EmpLeave struct {
	//EmpID     uint      `gorm:"primary_key;unique_index;type:integer"`
	gorm.Model
	Name      string    `json:"name"`
	LeaveType string    `json:"leaveType"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	TeamName  string    `json:"teamName"`
	Reporter  string    `json:"reporter"`
	Image     File      `gorm:"foreignkey:EmpLeaveID" json:"image"`
	//gorm:"foreignkey:EmpLeaveID"

	//EmpLeaveID uint `gorm:"primary key;unique_index;index;type:integer ON CONFLICT (EmpLeaveID) DO UPDATE SET EmpLeaveID = EXCLUDED.EmpLeaveID"`
}

// type File struct {
// 	ImageName    string `json:"imageName"`
// 	ImageContent []byte `json:"-"`
// 	EmpLeaveID   uint   `gorm:"foreignkey:EmpLeaveID" json:"-"`
// }

// type EmpLeave struct {
// 	EmpLeaveID uint      `gorm:"primary_key;unique_index;type:integer"`
// 	Name       string    `json:"name"`
// 	LeaveType  string    `json:"leaveType"`
// 	StartDate  time.Time `json:"startDate"`
// 	EndDate    time.Time `json:"endDate"`
// 	TeamName   string    `json:"teamName"`
// 	Reporter   string    `json:"reporter"`
// 	Image      File      ` json:"image"`
// 	//gorm:"foreignkey:EmpLeaveID"

// 	//EmpLeaveID uint `gorm:"primary key;unique_index;index;type:integer ON CONFLICT (EmpLeaveID) DO UPDATE SET EmpLeaveID = EXCLUDED.EmpLeaveID"`
// }
