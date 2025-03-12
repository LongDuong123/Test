package repository

import (
	"time"

	"example.com/models"
	"example.com/mysql"
)

type UserRepository interface {
	AddUser(name string, walletAddr string) error
	CheckIn(walletAddr string) error
	GetCheckInTime(walletAddr string) (*time.Time, error)
	UpdateWorkDays(walletAddr string) error
	UpdateEarlyLeaveDays(walletAddr string) error
	UpdateLateDays(walletAddr string) error
	GetAttendance(walletAddr string) (*models.Attendance, error)
	GetSalary(walletAddr string) (string, error)
	UpdateUser(name string, walletAddr string, role string, salary string) error
	DeleteUser(walletAddr string) error
	GetCompanyOverview() ([]models.User, error)
}

type userRepository struct {
	Database *mysql.MySQL
}

func NewUserRepository(database *mysql.MySQL) UserRepository {
	return &userRepository{
		Database: database,
	}
}

func (r *userRepository) AddUser(name string, walletAddr string) error {
	result := r.Database.DB.Exec("INSERT INTO users (name, wallet_addr) VALUES (?, ?)", name, walletAddr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) CheckIn(walletAddr string) error {
	result := r.Database.DB.Exec("UPDATE users SET checkin_time = ? WHERE wallet_addr = ?", time.Now(), walletAddr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) GetCheckInTime(walletAddr string) (*time.Time, error) {
	var checkinTime time.Time
	result := r.Database.DB.Raw("SELECT checkin_time FROM users WHERE wallet_addr = ?", walletAddr).Scan(&checkinTime)
	if result.Error != nil {
		return nil, result.Error
	}
	return &checkinTime, nil
}

func (r *userRepository) UpdateWorkDays(walletAddr string) error {
	result := r.Database.DB.Exec("UPDATE users SET work_days = COALESCE(work_days, 0) + 1 WHERE wallet_addr = ?", walletAddr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) UpdateEarlyLeaveDays(walletAddr string) error {
	result := r.Database.DB.Exec("UPDATE users SET early_leave_days = COALESCE(early_leave_days, 0) + 1 WHERE wallet_addr = ?", walletAddr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) UpdateLateDays(walletAddr string) error {
	result := r.Database.DB.Exec("UPDATE users SET late_days = COALESCE(late_days, 0) + 1 WHERE wallet_addr = ?", walletAddr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) UpdateUser(name string, walletAddr string, role string, salary string) error {
	result := r.Database.DB.Exec("UPDATE users SET name = ?, role = ?, salary = ? WHERE wallet_addr = ?", name, role, salary, walletAddr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) DeleteUser(walletAddr string) error {
	result := r.Database.DB.Exec("DELETE FROM users WHERE wallet_addr = ?", walletAddr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) GetAttendance(walletAddr string) (*models.Attendance, error) {
	var attendance models.Attendance
	result := r.Database.DB.Raw("SELECT work_days, leave_days, late_days, early_leave_days FROM users WHERE wallet_addr = ?", walletAddr).Scan(&attendance)
	if result.Error != nil {
		return nil, result.Error
	}
	return &attendance, nil
}

func (r *userRepository) GetSalary(walletAddr string) (string, error) {
	var salary string
	result := r.Database.DB.Raw("SELECT COALESCE(salary, '') FROM users WHERE wallet_addr = ?", walletAddr).Scan(&salary)
	if result.Error != nil {
		return "", result.Error
	}
	return salary, nil
}

func (r *userRepository) GetCompanyOverview() ([]models.User, error) {
	var users []models.User
	result := r.Database.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
