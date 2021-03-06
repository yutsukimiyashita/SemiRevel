package daos

import (
	"SemiRevel/app/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

//InitDB 関数
func InitDB() {
	// gormでMySQL接続
	// 失敗したらerrに格納される
	// Openの第二引数は {username}:{password}@/{dbname}?charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", "revel:revel@tcp(revel_db:3306)/semirevel?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	db.DB()
	// マイグレード
	db.AutoMigrate(&models.Material{})
	db.AutoMigrate(&models.User{})

	// dbをDB(*gorm.DB)として外からも使えるようにしてあげます
	// (関数の外でvarで宣言してるのでこうすれば他のコントローラーからもDB.hoge()って感じで使える)
	DB = db

}
