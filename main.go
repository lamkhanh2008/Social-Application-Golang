package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	appCtx "social_todo/component/appContext"
	"social_todo/component/tokenprovider/jwt"
	"social_todo/middleware"
	ginitem "social_todo/module/item/transport/gin"
	"social_todo/module/upload"
	"social_todo/module/user/storage"
	ginuser "social_todo/module/user/transport/gin"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"status;"`
	CreatedAt   *time.Time `json:"created_at, omitempty" gorm:"column:created_at;"`
	UpdatedAt   *time.Time `json:"updated_at, omitempty" gorm:"column:updated_at;"`
}

type TodoItemCreation struct {
	Id          int     `json:"id" gorm:"column:id;"`
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"status;"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

func (TodoItem) TableName() string {
	return "todo_items"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_Conn")
	jwtSecret := os.Getenv("JWT_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	appContext := appCtx.NewAppContext(db)
	db = appContext.GetMaiDBConnection()
	authStore := storage.NewSQLStore(db)
	tokenProvider := jwt.NewTokenJWTProvider("jwt", jwtSecret)
	db = db.Debug()
	r := gin.Default()
	r.Static("/static", "./static")
	v1 := r.Group("/v1")

	v1.Use(middleware.Recover())

	{
		v1.POST("/login", ginuser.Logic(db, tokenProvider))
		v1.PUT("/upload", upload.Upload(db))
		v1.POST("/register", ginuser.Register(db))
		v1.GET("/profile", middleware.RequiredAuth(authStore, tokenProvider), ginuser.Profile())
		items := v1.Group("/items")
		{
			items.POST("/", middleware.RequiredAuth(authStore, tokenProvider), ginitem.CreateItem(db))
			// items.GET("/", ListItem(db))
			items.GET("/", middleware.RequiredAuth(authStore, tokenProvider), ginitem.GetItemByID(db))
			items.PATCH("/:id", middleware.RequiredAuth(authStore, tokenProvider), ginitem.UpdateItemById(db))
			items.DELETE("/:id", SoftDeleteItem(db))
			items.GET("/list", middleware.RequiredAuth(authStore, tokenProvider), ginitem.GetListItems(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}

func SoftDeleteItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		deletedStatus := "Deleted"
		itemData := TodoItemUpdate{
			Status: &deletedStatus,
		}
		if err = db.Where("id = ?", id).Updates(&itemData).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
