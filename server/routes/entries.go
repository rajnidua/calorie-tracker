package routes

import("context"
"fmt"
"time"
"github.com/gin-gonic/gin"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
"gopkg.in/mgo.v2/bson")
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

func AddEntry(c *gin.Context){
var ctx,cancel = context.WithTimeout(context.Background(),100*time.second)
var entry models.Entry


if err : = c.BindJSON(&entry); err!=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	fmt.println(err)
	return
}
validationErr:=validate.Struct(entry)
if validationErr !=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":validationErr.Error()})
	fmt.Println(validationErr)
	return
}
entry.ID =primitive.NewObjectID()
result,insertErr := entryCollection.InsertOne(ctx,entry)
if insertErr != nil{
	msg := fmt.Sprintf("order item was not created")
	c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
	fmt.Println(insertErr)
	return
}
defer cancel()
c.JSON(http.StatusOK,result)
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
ingredient:=c.Params.ByName("id")
var ctx,cancel= context.withTimeout(context.Background((),100*time.Second))
var enteries []bson.M
cursor,err := entryCollection.Find(ctx,bson.M{"ingredients":ingredient})
err!:=nil{
	c.JSON(http.statusInternalServerError,gin.H{"error":err.Error()})
	fmt.Println(err)
	return
}
if err =	cursor.All(ctx,&enteries); err!=nil{
	c.JSON(http.StatusInternalServer,gin.H{"error":err.Error()})
	fmt.Println(err)
	return
}
defer cancel()
fmt.Println(enteries)
c.JSON(http.statusOK,enteries)
}

func GetEntryById(c *gin.Context){
EntryID :=c.Params.ByName("id")
docID, _ :=primitive.ObjectIDFromHex(EntryID)
var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
var entry bson.M
if err := entryCollection.FindOne(ctx,bson.M{"_id":docID}).Decode(&entry); err !=nil{c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
fmt.Println(err)
return}
defer canceel()
fmt.Println(entry)
c.JSON(http.StatusOK,entry)
}
func UpdateIngredient(c *gin.Context){
entryID:= c.Params.ByName("id")
	docID, _ := primitve.ObjectIDFromHex(entryID)
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
	 
	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
var ingredient Ingredient
	if err : = c.BindJSON(&entry); err!=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":validationErr.Error()})
	fmt.println(validationErr)
	return
}
result,err:=entryCollection.UpdateOne(ctx,bson.M{"_id":docID}),
bson.D{"$set",bson.D{{"ingredients",ingredient.Ingredients}}}},
if err !=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	fmt.Println(err)
	return
}
defer cancel()
c.JSON(http.StatusOK,result.ModifiedCount)
}

func UpdateEntry(c *gin.Context){
	entryID:= c.Params.ByName("id")
	docID, _ := primitve.ObjectIDFromHex(entryID)
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
	var entry model.Entry

	if err : = c.BindJSON(&entry); err!=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":validationErr.Error()})
	fmt.println(validationErr)
	return
}
result,err := entryCollection.ReplaceOne(ctx,bson.M{"_id":docID},bson.M{"dish":entry.Dish,
"fat":entry.Fat,
"ingredients":entry.Ingredients,
"calories":entry.Calories},)
if err != nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	fmt.Println(err)
	return
}
defer cancel()
c.JSON(http.StatusOK,result.ModifiedCount)

}

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

