package Models
import (
	"freshers_bootcamp/Day4/Config"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]Customer) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}