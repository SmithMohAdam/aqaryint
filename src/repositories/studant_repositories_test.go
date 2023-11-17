package repositories_test

import (
	"errors"
	"testing"

	"aqaryint/src/models"
	"aqaryint/src/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateStudant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()
	dsn := "host=localhost user=postgres password=smith dbname=postgres port=5432 sslmode=disable "
	//gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatalf("Failed to initialize gorm database: %v", err)
	}

	repo := repositories.NewStudantRepository(gormDB)

	studant := &models.Student{
		Id:    1,
		Name:  "mohammed",
		Age:   28,
		City:  "khartoum",
		Class: 'A',
		//Fit:   true,
	}

	//mock.ExpectExec("INSERT INTO `students` ('name', 'age','city','class','fit') VALUES ('mohammed',28,'khartoum,'A',true')").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO students (id ,name, age, city, class) VALUES (1,'mohammed', 28, 'khartoum', 'A')").WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.CreateStudant(studant)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestGetStudant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()
	dsn := "host=localhost user=postgres password=smith dbname=postgres port=5432 sslmode=disable "

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize gorm database: %v", err)
	}

	repo := repositories.NewStudantRepository(gormDB)

	studant := &models.Student{
		Name:  "mohammed",
		Age:   28,
		City:  "khartoum",
		Class: rune('A'),
		//Fit:   true,
	}

	rows := sqlmock.NewRows([]string{"name", "age", "city", "class", "fit"}).
		AddRow(studant.Name, studant.Age, studant.City, studant.Class) //, studant.Fit

	mock.ExpectQuery("SELECT .* FROM `students`").WillReturnRows(rows)

	result, err := repo.GetStudant(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil || result.Name != studant.Name || result.Age != studant.Age || result.City != studant.City || result.Class != studant.Class { //|| result.Fit != studant.Fit
		t.Errorf("Expected student with name %v, got %v", studant.Name, result)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestUpdateStudant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize gorm database: %v", err)
	}

	repo := repositories.NewStudantRepository(gormDB)

	studant := &models.Student{
		Name:  "mohammed",
		Age:   28,
		City:  "khartoum",
		Class: rune('A'),
		//Fit:   true,
	}

	mock.ExpectExec("UPDATE `students`").WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateStudant(studant)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestDeleteStudant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize gorm database: %v", err)
	}

	repo := repositories.NewStudantRepository(gormDB)

	mock.ExpectExec("DELETE FROM `students`").WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteStudant(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestGetStudantError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to initialize gorm database: %v", err)
	}

	repo := repositories.NewStudantRepository(gormDB)

	mock.ExpectQuery("SELECT .* FROM `students`").WillReturnError(errors.New("database error"))

	_, err = repo.GetStudant(1)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
