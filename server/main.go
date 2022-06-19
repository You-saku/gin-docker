package main

import (
	// Gin用
	"net/http"
	"github.com/gin-gonic/gin"

	// gorm用
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func main() {
	/*
	* DB接続確認
	* Todo: DBへの接続情報は環境変数で管理
	* DB接続が確認できたのでコメントアウト
	*/
	// user := "user"
    // pass := "secret"
    // protocol := "tcp(db:3306)" // 文法 => tco(コンテナ名:ポート番号) ハマるポイント
    // dbName := "develop"


	// type Sample struct {
	// 	gorm.Model
	// 	Code  string
	// 	Price uint
	// }
	// connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// db.AutoMigrate(&Sample{})


	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello gin!!",
		})
	})

	sampleGroup := r.Group("/api") // APIのグループ化
	sampleGroup.Use(sampleMiddleware()) // ミドルウェアを設定
	sampleGroup.GET("/ok", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "API!!",
			})
		})

	r.Run(":80")
}

// 自作のミドルウェア
func sampleMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

		flag := c.Request.Header["Variable"][0] // headerのkeyがなぜか配列でくる
		if(flag == "yes") {
			c.JSON(400, gin.H{"message": "NotFound."})
			c.Abort()
		}
		
		c.Next()
	}
}
