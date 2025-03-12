package service

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/internal/wallet"
	"example.com/models"
	"example.com/repository"
)

type UserService interface {
	UserJoin(name string, refcode string) (string, string, error)
	UserCheckIn(walletAddr string) error
	UserCheckOut(walletAddr string) error
	UserAttendance(walletAddr string) (*models.Attendance, error)
	UserSalary(walletAddr string) (string, error)
	RootUpdateUser(name string, walletAddr string, role string, salary string) error
	RootDeleteUser(walletAddr string) error
	RootCreateRefCode() (string, error)
	RootCompanyOverview() ([]models.User, error)
}

type userService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	s := &userService{}
	s.Repo = repo
	return s
}

func (s *userService) UserJoin(name string, refcode string) (string, string, error) {
	walletAddr, walletKey, err := wallet.GenerateEthereumWallet() // Generate wallet address and key
	if err != nil {
		return "", "", err
	}

	err = s.Repo.AddUser(name, walletAddr) // Add user to the database
	if err != nil {
		return "", "", err
	}

	return walletAddr, walletKey, nil
}

func (s *userService) UserCheckIn(walletAddr string) error {
	return s.Repo.CheckIn(walletAddr)
}

func (s *userService) UserCheckOut(walletAddr string) error {
	checkIn, err := s.Repo.GetCheckInTime(walletAddr) // Get the check-in time
	if err != nil {
		return err
	}

	now := time.Now()
	if checkIn.Hour() > 8 || (checkIn.Hour() == 8 && checkIn.Minute() > 0) { // Check if the user is late
		return s.Repo.UpdateLateDays(walletAddr)
	}

	if now.Hour() < 17 || (now.Hour() == 17 && now.Minute() < 30) { // Check if the user is leaving early
		return s.Repo.UpdateEarlyLeaveDays(walletAddr)
	}

	return s.Repo.UpdateWorkDays(walletAddr)
}

func (s *userService) UserAttendance(walletAddr string) (*models.Attendance, error) {
	return s.Repo.GetAttendance(walletAddr)
}

func (s *userService) UserSalary(walletAddr string) (string, error) {
	return s.Repo.GetSalary(walletAddr)
}

func (s *userService) RootUpdateUser(name string, walletAddr string, role string, salary string) error {
	return s.Repo.UpdateUser(name, walletAddr, role, salary)
}

func (s *userService) RootDeleteUser(walletAddr string) error {
	return s.Repo.DeleteUser(walletAddr)
}

func (s *userService) RootCreateRefCode() (string, error) {
	refcode := "REF" + strconv.Itoa(int(time.Now().Unix()))       // Generate refcode
	if err := updateEnv("REFCODE", refcode, ".env"); err != nil { // Update the environment variable
		return "", err
	}
	return refcode, nil
}

func (s *userService) RootCompanyOverview() ([]models.User, error) {
	return s.Repo.GetCompanyOverview()
}

func updateEnv(key, value, filename string) error {
	data, _ := os.ReadFile(filename)
	lines := []string{}
	if len(data) > 0 {
		lines = strings.Split(string(data), "\n")
	}
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = fmt.Sprintf("%s=%s", key, value)
			found = true
		}
	}
	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}
	return os.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
}
