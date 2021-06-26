package db

import (
	"log"
)

func (database *DB) CreatorTable(table interface{}, tableName string) {
	db := database.DB
	hasTable := db.Migrator().HasTable(tableName)
	if !hasTable {
		err := db.Migrator().CreateTable(table)
		if err != nil {
			log.Fatalf("Could not migrate: %v", err.Error())
		}
	}
}

func (database *DB) CreatorConstraint(constraint interface{}, constraintName string) {
	db := database.DB
	hasConstraint := db.Migrator().HasConstraint(constraint, constraintName)
	if !hasConstraint {
		err := db.Migrator().CreateConstraint(constraint, constraintName)
		if err != nil {
			log.Fatalf("Could not create foreign key: %v", err.Error())
		}
	}
}

func Migrate(database *DB)  {
	database.CreatorTable(&User{}, "user")
	database.CreatorTable(&DoctorAppointment{}, "doctor_appointment")
	database.CreatorTable(&Registrant{}, "registrant")
	//database.CreatorConstraint(&User{}, "UserRegistrants")
	//database.CreatorConstraint(&User{}, "fk_users_registrant")
	//database.CreatorConstraint(&DoctorAppointment{}, "DoctorRegistrants")
	//database.CreatorConstraint(&DoctorAppointment{}, "fk_doctor_appointments_registrant")
	log.Printf("Migration did run successfully")
}
