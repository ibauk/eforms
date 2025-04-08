package main

import (
	"fmt"
	"math/rand"

	"gorm.io/gorm"
)

type OtpModel struct {
	gorm.Model //This auto generates CreatedAt and ID fields
	UserEmail  string
	Token      string
}

func OTPGenerate(db *gorm.DB, userEmail string, otpLength int) (token string, err error) {
	db.AutoMigrate(&OtpModel{})

	x := OTPgenerateOTP(otpLength)
	token = x

	err = db.Create(&OtpModel{
		UserEmail: userEmail,
		Token:     token,
	}).Error

	return

}

func OTPgenerateOTP(otp_length int) string {
	otp := ""

	min := 0
	max := 9

	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < otp_length; i++ {
		generatedOtp := fmt.Sprintf("%d", rand.Intn(max-min+1)+min)
		otp += generatedOtp
	}
	return otp
}

func OTPValid(db *gorm.DB, userEmail string, token string) bool {
	var foundOTP OtpModel

	fmt.Printf("Checking %v with %v\n", userEmail, token)
	err := db.Model(OtpModel{}).Where("token = ?", token).Where("user_email = ?", userEmail).Where("deleted_at is null").Find(&foundOTP).Error
	if err != nil {
		return false
	}
	if foundOTP.UserEmail == userEmail && foundOTP.Token == token {
		fmt.Printf("%s One Time Password Validated \n", userEmail)
		fmt.Printf("%v\n", foundOTP.ID)
		db.Delete(&foundOTP, foundOTP.ID)
		return true
	}
	fmt.Println("Nope")
	return false
}
