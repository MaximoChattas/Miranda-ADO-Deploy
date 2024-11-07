package client

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"project/model"
	"testing"
)

func TestInsertImage_Client(t *testing.T) {
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
	ImageClient = &imageClient{}

	image := model.Image{
		Id:      1,
		Path:    "Images/1-1.JPG",
		HotelId: 1,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SET IDENTITY_INSERT "images" ON;INSERT INTO "images" ("path","hotel_id","id") OUTPUT INSERTED."id" VALUES (@p1,@p2,@p3);SET IDENTITY_INSERT "images" OFF;`).
		WithArgs(image.Path, image.HotelId, image.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result := ImageClient.InsertImage(image)

	a.Equal(image, result)
	a.Equal(1, result.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestInsertImages_Client(t *testing.T) {
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
	ImageClient = &imageClient{}

	images := model.Images{
		model.Image{
			Id:      1,
			Path:    "Images/1-1.JPG",
			HotelId: 1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SET IDENTITY_INSERT "images" ON;INSERT INTO "images" ("path","hotel_id","id") OUTPUT INSERTED."id" VALUES (@p1,@p2,@p3);SET IDENTITY_INSERT "images" OFF;`).
		WithArgs(images[0].Path, images[0].HotelId, images[0].Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result := ImageClient.InsertImages(images)

	a.Equal(images, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetImageById_Client(t *testing.T) {
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
	ImageClient = &imageClient{}

	image := model.Image{
		Id:      1,
		Path:    "Images/1-1.JPG",
		HotelId: 1,
	}

	mock.ExpectQuery(`SELECT * FROM "images" WHERE id = @p1 ORDER BY "images"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`).
		WithArgs(image.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "path", "hotel_id"}).
			AddRow(image.Id, image.Path, image.HotelId))

	result := ImageClient.GetImageById(image.Id)

	a.Equal(image, result)
	a.Equal(1, image.Id)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetImages_Client(t *testing.T) {
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
	ImageClient = &imageClient{}

	images := model.Images{
		model.Image{
			Id:      1,
			Path:    "Images/1-1.JPG",
			HotelId: 1,
		},

		model.Image{
			Id:      2,
			Path:    "Images/1-2.JPG",
			HotelId: 1,
		},
	}

	mock.ExpectQuery(`SELECT * FROM "images"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "path", "hotel_id"}).
			AddRow(images[0].Id, images[0].Path, images[0].HotelId).
			AddRow(images[1].Id, images[1].Path, images[1].HotelId))

	result := ImageClient.GetImages()

	a.Equal(images, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}

func TestGetImagesByHotelId_Client(t *testing.T) {
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
	ImageClient = &imageClient{}

	images := model.Images{
		model.Image{
			Id:      1,
			Path:    "Images/1-1.JPG",
			HotelId: 1,
		},

		model.Image{
			Id:      2,
			Path:    "Images/1-2.JPG",
			HotelId: 1,
		},
	}

	hotelId := 1

	mock.ExpectQuery(`SELECT * FROM "images" WHERE hotel_id = @p1`).
		WithArgs(hotelId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "path", "hotel_id"}).
			AddRow(images[0].Id, images[0].Path, images[0].HotelId).
			AddRow(images[1].Id, images[1].Path, images[1].HotelId))

	result := ImageClient.GetImagesByHotelId(hotelId)

	a.Equal(images, result)

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}
}
