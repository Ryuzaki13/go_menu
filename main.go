package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Загрузить HTML шаблоны
	router.LoadHTMLGlob("./html/*.html")
	// Установить путь к статическим файлам (стили, скрипты и тд)
	router.Static("/r/", "./resource")
	// Установить обработчик запроса главной страницы
	router.GET("/", handlerIndex)
	// Запустить сервер локально на порту 8080
	router.Run("localhost:8080")
}

func handlerIndex(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{
		"Title": "Hello",
		"Menu":  []string{"О нас", "Продукты", "Контакты", "Форум"},
	})
}
