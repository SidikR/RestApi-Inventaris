package seeder

import (
	"main/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Seed(db *gorm.DB) {
	// Buat pengguna dengan peran "admin"
	adminUser := model.User{
		Email:    "admin@gmail.com",
		Username: "admin123",
		Role:     "admin",
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123@"), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password")
	}
	adminUser.Password = string(hashedPassword)

	// Tambahkan pengguna admin ke database
	if err := db.Create(&adminUser).Error; err != nil {
		panic("Failed to seed admin user")
	}

	// Buat pengguna dengan peran "user"
	user := model.User{
		Email:    "user@gmail.com",
		Username: "user123",
		Role:     "user",
	}

	// Hash password untuk pengguna "user"
	hashedUserPassword, err := bcrypt.GenerateFromPassword([]byte("user123@"), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash user password")
	}
	user.Password = string(hashedUserPassword)

	// Tambahkan pengguna "user" ke database
	if err := db.Create(&user).Error; err != nil {
		panic("Failed to seed user")
	}

	medicine := model.Medicine{
		MedicineName:        "Parasetamol",
		MedicineDescription: "Obat Sakit Panas",
		Stock:               1,
		Unit:                "tablet",
		Expired:             "20/042025",
		Price:               2000,
	}
	db.Create(&medicine)

	division := model.Division{
		DivisionName:        "Poli",
		DivisionDescription: "Terdapat beberapa POLI",
	}
	db.Create(&division)

	room := model.Room{
		RoomName:        "A1",
		RoomDescription: "Ruangannya bersih",
		Location:        "Lantai 1",
		DivisionId:      division.ID,
		// Division:        division,
	}
	db.Create(&room)

	inventory := model.Inventory{
		InventoryName:        "Stetoskop",
		InventoryDescription: "Untuk Mengukur detak Jantung",
		Qty:                  10,
		Condition:            "Layak",
		RoomId:               room.ID,
	}
	db.Create(&inventory)

	doctor := model.Doctor{
		DoctorName:   "Rahmad Sidik",
		Email:        "sidikellampungi@gmail.com",
		Role:         "Dokter Bedah",
		JoinAt:       "2022-12-12",
		Image:        "sidik.jpg",
		Status:       "Senior",
		MobileNumber: "089528596517",
		Address:      "Sidomulyo",
	}
	db.Create(&doctor)

	poly := model.Poly{
		PolyName:        "Poli Umum",
		PolyDescription: "Ruangannya bersih",
		RoomId:          room.ID,
		// Room:            room,
		DoctorIds: []string{"sasdas", "sdasda"},
	}
	db.Create(&poly)

}
