package database

import (
	"database/sql"
	"testing"

	"github.com/alextavella/gointensivo/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec(`create table orders (
		id varchar(255) not null, 
		price float not null, 
		tax float not null, 
		final_price float not null
	)`)
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestSavingOrder() {
	order, err := entity.NewOrder("123", 10.0, 1.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.SaveOrder(order)
	suite.NoError(err)
}
