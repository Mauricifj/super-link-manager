package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/super-link-manager/models"
	"github.com/super-link-manager/utils"
	"os"
)

func ConnectDB() *sql.DB {
	var host = os.Getenv("POSTGRES_HOST")
	var port = os.Getenv("POSTGRES_PORT")
	var user = os.Getenv("POSTGRES_USERNAME")
	var password = os.Getenv("POSTGRES_PASSWORD")
	var database = os.Getenv("POSTGRES_DB")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := sql.Open("postgres", connectionString)
	utils.CheckErr(err)
	return db
}

func Links() []models.Link {
	db := ConnectDB()

	linksQuery, err := db.Query("select * from links order by id asc")
	utils.CheckErr(err)

	var links []models.Link

	for linksQuery.Next() {
		var id string
		var linkType string
		var name string
		var price int

		err = linksQuery.Scan(
			&id,
			&linkType,
			&name,
			&price,
		)
		utils.CheckErr(err)

		link := models.Link{
			Id:       id,
			LinkType: linkType,
			Name:     name,
			Price:    price,
		}

		links = append(links, link)
	}
	defer db.Close()

	return links
}

func CreateLink(id, linkType, name string, price int) bool {
	db := ConnectDB()

	linkInsert, err := db.Prepare("insert into links (id, type, name, price) values ($1, $2, $3, $4)")
	utils.CheckErr(err)

	result, err := linkInsert.Exec(id, linkType, name, price)
	utils.CheckErr(err)

	rowsAffected, err := result.RowsAffected()
	utils.CheckErr(err)

	defer db.Close()

	return rowsAffected > 0
}

func DeleteLink(id string) bool {
	db := ConnectDB()

	productDelete, err := db.Prepare("delete from links where id=$1")
	utils.CheckErr(err)

	result, err := productDelete.Exec(id)
	utils.CheckErr(err)

	rowsAffected, err := result.RowsAffected()
	utils.CheckErr(err)

	defer db.Close()

	return rowsAffected > 0
}

func UpdateLink(id, linkType, name string, price int) bool {
	db := ConnectDB()

	updateProduct, err := db.Prepare("update links set type=$2, name=$3, price=$4 where id=$1")
	utils.CheckErr(err)

	result, err := updateProduct.Exec(id, linkType, name, price)
	utils.CheckErr(err)

	rowsAffected, err := result.RowsAffected()
	utils.CheckErr(err)

	defer db.Close()

	return rowsAffected > 0
}