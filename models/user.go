package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string     `json:"name"`
	WalletAddr     string     `json:"wallet_addr"`
	Role           string     `json:"role"`
	Salary         string     `json:"salary"`
	WorkDays       int        `json:"work_days"`
	LeaveDays      int        `json:"leave_days"`
	LateDays       int        `json:"late_days"`
	EarlyLeaveDays int        `json:"early_leave_days"`
	CheckinTime    *time.Time `json:"checkin_time,omitempty"`
	CheckoutTime   *time.Time `json:"checkout_time,omitempty"`
}

type Attendance struct {
	WorkDays       int `json:"work_days"`
	LeaveDays      int `json:"leave_days"`
	LateDays       int `json:"late_days"`
	EarlyLeaveDays int `json:"early_leave_days"`
}

type UserJoinRequest struct {
	Name    string `json:"name"`
	Refcode string `json:"refcode"`
}

type UserJoinResponse struct {
	WalletAddr string `json:"wallet_addr"`
	WalletKey  string `json:"wallet_key"`
}

type UserCheckInRequest struct {
	WalletAddr string `json:"wallet_addr"`
}

type UserCheckInResponse struct {
	Success bool `json:"success"`
}

type UserCheckOutRequest struct {
	WalletAddr string `json:"wallet_addr"`
}

type UserCheckOutResponse struct {
	Success bool `json:"success"`
}

type UserAttendanceRequest struct {
	WalletAddr string `json:"wallet_addr"`
}

type UserAttendanceResponse struct {
	Attendance Attendance `json:"attendance"`
}

type UserSalaryRequest struct {
	WalletAddr string `json:"wallet_addr"`
}

type UserSalaryResponse struct {
	Salary string `json:"salary"`
}

type RootUpdateUserRequest struct {
	WalletAddr string `json:"wallet_addr"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Salary     string `json:"salary"`
}

type RootUpdateUserResponse struct {
	Success bool `json:"success"`
}

type RootDeleteUserRequest struct {
	WalletAddr string `json:"wallet_addr"`
}

type RootDeleteUserResponse struct {
	Success bool `json:"success"`
}

type RootCreateRefCodeResponse struct {
	Refcode string `json:"refcode"`
}

type RootCompanyOverviewResponse struct {
	TotalUsers int    `json:"total_users"`
	Users      []User `json:"users"`
}
