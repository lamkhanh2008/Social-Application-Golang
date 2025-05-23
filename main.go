package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"social_todo/common"
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

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData TodoItemCreation

		if err := c.ShouldBind(&itemData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&itemData).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		fmt.Println(itemData)
		c.JSON(http.StatusOK, gin.H{
			"data": itemData.Id,
		})
	}
}

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var itemData TodoItem

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).First(&itemData).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": itemData,
		})
	}
}

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var itemData TodoItemUpdate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := ctx.ShouldBind(&itemData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
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

func HardDeleteItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err = db.Table(TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
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

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var result []TodoItem
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()
		fmt.Println(paging)
		if err := db.Table(TodoItem{}.TableName()).Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
