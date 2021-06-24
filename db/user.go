package db

func (database *DB) CreateUser(user *User) (err error) {
	if err := database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
