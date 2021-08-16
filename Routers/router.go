package Routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// User 结构体定义
type User struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

func Setup() *gin.Engine {

	router := gin.Default()
	//	router.LoadHTMLGlob("Views/**/*")

	// 雖說是spa, 但每一頁都需要設定, 不然重整會出現404
	router.Use(static.Serve("/", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/home", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/sport", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/lottery", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/slot", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/esport", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/chess", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/register", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/login", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/logout", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/game/lottery", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/game/sport", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/game/slot", static.LocalFile("./Views/paper-kit-react/build", true)))
	router.Use(static.Serve("/game/chess", static.LocalFile("./Views/paper-kit-react/build", true)))

	// session
	store := cookie.NewStore([]byte("gssecret"))
	router.Use(sessions.Sessions("mysession", store))

	//authController := Controller.NewAuthController()
	//router.GET("/", authController.LoginPage)

	// todo , 頁面使用的API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.GET("/session", func(c *gin.Context) {

			// 初始化session对象
			session := sessions.Default(c)

			log.Println(session.Get("auth"))

			if session.Get("auth") == "1" {
				c.JSON(http.StatusOK, gin.H{
					"status":  "1",
					"message": "已登入",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "0",
					"message": "未登入",
				})
			}
		})

		api.GET("/logout", func(c *gin.Context) {

			// 初始化session对象
			session := sessions.Default(c)
			session.Set("auth", "0")
			session.Save()

			c.JSON(http.StatusOK, gin.H{
				"status":  "1",
				"message": "已登出",
			})
		})

		api.POST("/login", func(c *gin.Context) {

			// 初始化user struct
			u := User{}
			if c.ShouldBind(&u) == nil {
				// 绑定成功， 打印请求参数
				log.Println(u.Account)
				log.Println(u.Password)
			}

			if u.Account == "admin" && u.Password == "12345" {

				// 初始化session
				session := sessions.Default(c)

				// 設置session
				session.Set("account", u.Account)
				session.Set("auth", "1")

				// 保存session
				session.Save()

				c.JSON(http.StatusOK, gin.H{
					"status":  "1",
					"message": "success",
					"account": "admin",
					"balance": "999999",
				})
			} else {

				c.JSON(http.StatusOK, gin.H{
					"status":  "0",
					"message": "帳號密碼錯誤！",
				})
			}

		})

		api.POST("/register", func(c *gin.Context) {

			// 初始化user struct
			u := User{}
			if c.ShouldBind(&u) == nil {
				// 绑定成功， 打印请求参数
				log.Println(u.Account)
				log.Println(u.Password)
			}

			// 初始化session
			session := sessions.Default(c)

			// 設置session
			session.Set("account", u.Account)
			session.Set("auth", "1")

			// 保存session
			session.Save()

			c.JSON(http.StatusOK, gin.H{
				"status":  "1",
				"message": "success",
				"account": "admin",
				"balance": "999999",
			})

		})
	}

	return router
}
