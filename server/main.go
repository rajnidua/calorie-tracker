package main


import("os"
"github.com/rajnidua/calorie-tracker/routes"
"github.com/gin-contrib/cors"
"github.com/gin-gonic/gin")

func main(){
	port := os.Getenv("PORT")
	if port ==""{
		port = "8000"
}

router := gin.New()
router.User(gin.Logger())
router.Use(cors.Default())

router.POST("/entry/create",routes.AddEntry)
router.GET("/enteries",routes.GetEntries)
router.GET("/entry/:id/",routes.GetEntryByID)
router.GET("ingredient/:ingredient",routes.GetEnteriesByIngredient)

router.PUT("/entry/update/:id",routes.UpdateEntry)
router.PUT("/ingredient/update/:id",routes.UpdateIngredients)
router.DELETE("/entry/delete/:id",routes.DeleteEntry)
router.RUN(":"+port)
}