// main.go
package main

import (
	"fmt"
	"main/database"
	"main/migration"
	"main/router"
	"net/url"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := database.InitDB()
	defer db.Close()

	db.LogMode(true)

	migration.Migrate(db)
	// seeder.Seed(db)

	r := gin.Default()

	// Aktifkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Check if the origin is in the range 3000 to 3999
			originURL, err := url.Parse(origin)
			if err != nil {
				return false
			}

			port, err := strconv.Atoi(originURL.Port())
			if err != nil {
				return false
			}

			return port >= 3000 && port <= 3999
		},
	}))
	// Group router for authentication
	authRouter := r.Group("/")
	router.SetAuthRoutes(authRouter, db)

	// Group router untuk barang
	medicineRouter := r.Group("/")
	router.SetMedicineRoutes(medicineRouter, db)

	// Group router untuk barang
	patientRouter := r.Group("/")
	router.SetPatientRoutes(patientRouter, db)

	// Group router untuk barang
	receipRouter := r.Group("/")
	router.SetReceipRoutes(receipRouter, db)

	// Group router untuk barang
	registrationRouter := r.Group("/")
	router.SetRegistrationRoutes(registrationRouter, db)

	// Group router untuk barang
	roomRouter := r.Group("/")
	router.SetRoomRoutes(roomRouter, db)

	// Group router untuk barang
	bilReportRouter := r.Group("/")
	router.SetBilReportRoutes(bilReportRouter, db)

	// Group router untuk barang
	divisionRouter := r.Group("/")
	router.SetDivisionRoutes(divisionRouter, db)

	// Group router untuk barang
	doctorRouter := r.Group("/")
	router.SetDoctorRoutes(doctorRouter, db)

	// Group router untuk barang
	inPatientCareRouter := r.Group("/")
	router.SetInPatientCareRoutes(inPatientCareRouter, db)

	// Group router untuk barang
	inventoryRouter := r.Group("/")
	router.SetInventoryRoutes(inventoryRouter, db)

	// Group router untuk barang
	polyRouter := r.Group("/")
	router.SetPolyRoutes(polyRouter, db)

	// Group router untuk barang
	medicalReportRouter := r.Group("/")
	router.SetMedicalReportRoutes(medicalReportRouter, db)

	r.Run(":8080")
}
