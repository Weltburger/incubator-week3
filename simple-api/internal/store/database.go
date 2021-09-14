package store

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	psgs "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

type Database struct {
	GormDB *gorm.DB
	cityRepository *CityRepository
	classRepository *ClassRepository
	countryRepository *CountryRepository
	flightRepository *FlightRepository
	lunchRepository *LunchRepository
	passengerRepository *PassengerRepository
	ticketRepository *TicketRepository
	sync.Once
}

func (db *Database) CityRepository() *CityRepository {
	if db.cityRepository != nil {
		return db.cityRepository
	}

	db.cityRepository = &CityRepository{database: db}

	return db.cityRepository
}

func (db *Database) ClassRepository() *ClassRepository {
	if db.classRepository != nil {
		return db.classRepository
	}

	db.classRepository = &ClassRepository{database: db}

	return db.classRepository
}

func (db *Database) CountryRepository() *CountryRepository {
	if db.countryRepository != nil {
		return db.countryRepository
	}

	db.countryRepository = &CountryRepository{database: db}

	return db.countryRepository
}

func (db *Database) FlightRepository() *FlightRepository {
	if db.flightRepository != nil {
		return db.flightRepository
	}

	db.flightRepository = &FlightRepository{database: db}

	return db.flightRepository
}

func (db *Database) LunchRepository() *LunchRepository {
	if db.lunchRepository != nil {
		return db.lunchRepository
	}

	db.lunchRepository = &LunchRepository{database: db}

	return db.lunchRepository
}

func (db *Database) PassengerRepository() *PassengerRepository {
	if db.passengerRepository != nil {
		return db.passengerRepository
	}

	db.passengerRepository = &PassengerRepository{database: db}

	return db.passengerRepository
}

func (db *Database) TicketRepository() *TicketRepository {
	if db.ticketRepository != nil {
		return db.ticketRepository
	}

	db.ticketRepository = &TicketRepository{database: db}

	return db.ticketRepository
}

func GetDB() *Database {
	gormDB := startDB()

	return &Database{GormDB: gormDB}
}

func startDB() *gorm.DB {
	//dbUrl := os.Getenv("DB_URL")
	//sqlDB, err := sql.Open("postgres", dbUrl)
	sqlDB, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/airport?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}


	return gormDB
}

func (s *Database) Migrate() {
	s.Do(func() {
		db, _ := s.GormDB.DB()
		driver, err := psgs.WithInstance(db, &psgs.Config{})
		fsrc, err := (&file.File{}).Open("file://migrations")
		if err != nil {
			log.Fatal(err)
		}

		m, err := migrate.NewWithInstance(
			"file",
			fsrc,
			"postgres",
			driver,
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	})
}
