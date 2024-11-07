package client

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm/logger"
	"project/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsertAmenity_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	AmenityClient = &amenityClient{}

	amenity := model.Amenity{
		Id:   1,
		Name: "Pool",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SET IDENTITY_INSERT "amenities" ON;INSERT INTO "amenities" ("name","id") OUTPUT INSERTED."id" VALUES (@p1,@p2);SET IDENTITY_INSERT "amenities" OFF;`).
		WithArgs(amenity.Name, amenity.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result := AmenityClient.InsertAmenity(amenity)

	a.Equal(amenity, result)
	a.Equal(1, amenity.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetAmenityById_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	AmenityClient = &amenityClient{}

	amenity := model.Amenity{
		Id:   1,
		Name: "Pool",
	}

	mock.ExpectQuery(`SELECT * FROM "amenities" WHERE id = @p1 ORDER BY "amenities"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`).
		WithArgs(amenity.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(amenity.Id, amenity.Name))

	result := AmenityClient.GetAmenityById(amenity.Id)

	a.Equal(amenity, result)
	a.Equal(1, amenity.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetAmenityByName_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	AmenityClient = &amenityClient{}

	amenity := model.Amenity{
		Id:   1,
		Name: "Pool",
	}

	mock.ExpectQuery(`SELECT * FROM "amenities" WHERE name = @p1 ORDER BY "amenities"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`).
		WithArgs(amenity.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(amenity.Id, amenity.Name))

	result := AmenityClient.GetAmenityByName(amenity.Name)

	a.Equal(amenity, result)
	a.Equal(amenity.Name, result.Name)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetAmenities_Client(t *testing.T) {
	a := assert.New(t)

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Failed to create mock database")
	}
	defer db.Close()

	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		DriverName: "sqlserver",
		Conn:       db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Connection failed to open")
	}

	Db = gormDB
	AmenityClient = &amenityClient{}

	amenities := model.Amenities{
		model.Amenity{
			Id:   1,
			Name: "Pool",
		},

		model.Amenity{
			Id:   2,
			Name: "Breakfast",
		},
	}

	mock.ExpectQuery(`SELECT * FROM "amenities"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(amenities[0].Id, amenities[0].Name).
			AddRow(amenities[1].Id, amenities[1].Name))

	result := AmenityClient.GetAmenities()

	a.Equal(amenities, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
