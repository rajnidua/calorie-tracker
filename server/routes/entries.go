package routes

import("")
var entryCollection *mongo.Collection = openCollection(Client, "calories")

func AddEntry(c *gin.Context){

}

func GetEnteries(c *gin.Context){
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.second)
	var enteries []bson.M
	cursor,err := entryCollection.Find(ctx, bson.M{})
	if err != nil{
		c.JSON(http.StatusInternalServer, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	
	}
if err =	cursor.All(ctx,&enteries); err!=nil{
	c.JSON(http.StatusInternalServer,gin.H{"error":err.Error()})
	fmt.Println(err)
	return
}
defer cancel(
	fmt.Println(enteries)
	c.JSON(http.StatusOK,enteries)
)
}

func GetEnteriesByIngredient(c *gin.Context){

}

func GetEntryById(c *gin.Context){

}
func UpdateIngredient(c *gin.Context){

}

func UpdateEntry(c *gin.Context){}

func DeleteEntry(c *gin.Context){
entryID := c.Params.ByName("id")
docID ,_ := primitive.ObjectIDFromHex(entryID)

var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
result, err := entryCollection.DeleteOne(ctx,bson.M("_id":docID))
if err != nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	fmt.Print(err)
}
defer cancel()
c.JSON(http.StatusOK,result.DeletedCount)
}

