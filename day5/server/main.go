package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Config
var mongoUri string = "mongodb://localhost:27017"
var mongoDbName string = "ticket_app_db"
var mongoCollectionTicket string = "tickets"

// Database variables
var mongoclient *mongo.Client
var ticketCollection *mongo.Collection

// Model Ticket for Collection "tickets"
type Ticket struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name       string             `json:"name" bson:"name"`
    Age        int                `json:"age" bson:"age"`
    Destination string            `json:"destination" bson:"destination"`
}

// Connect to MongoDB
func connectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var errrorConnection error
    mongoclient, errrorConnection = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
    if errrorConnection != nil {
        log.Fatal("MongoDB Connection Error:", errrorConnection)
    }

    ticketCollection = mongoclient.Database(mongoDbName).Collection(mongoCollectionTicket)
    fmt.Println("Connected to MongoDB!")
}

// POST /tickets
func bookTicket(c *gin.Context) {
    var jbodyTicket Ticket

    // Bind JSON body to jbodyTicket
    if err := c.BindJSON(&jbodyTicket); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Insert ticket into MongoDB
    result, err := ticketCollection.InsertOne(ctx, jbodyTicket)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book ticket"})
        return
    }

    // Extract the inserted ID
    ticketId, ok := result.InsertedID.(primitive.ObjectID)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse inserted ID"})
        return
    }
    jbodyTicket.ID = ticketId

    // Read the booked ticket from MongoDB
    var bookedTicket Ticket
    err = ticketCollection.FindOne(ctx, bson.M{"_id": jbodyTicket.ID}).Decode(&bookedTicket)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booked ticket"})
        return
    }

    // Return booked ticket
    c.JSON(http.StatusCreated, gin.H{
        "message": "Ticket booked successfully",
        "ticket":  bookedTicket,
    })
}

// GET /tickets (Not required in the original request, just kept for completeness)
// func readAllTickets(c *gin.Context) {
//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     cursor, err := ticketCollection.Find(ctx, bson.M{})
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
//         return
//     }
//     defer cursor.Close(ctx)

//     tickets := []Ticket{}
//     if err := cursor.All(ctx, &tickets); err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tickets"})
//         return
//     }

//     c.JSON(http.StatusOK, tickets)
// }
func readTicketbyId(c *gin.Context) {
    id := c.Param("id")

    // Convert string ID to primitive.ObjectID
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var ticket Ticket
    err = ticketCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&ticket)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message":"Ticket downloaded successfully","ticket":ticket})
}



// PUT /tickets/:id
func updateDestination(c *gin.Context) {
    id := c.Param("id")
    // Convert string ID to primitive.ObjectID
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    var jbodyTicket Ticket

    if err := c.BindJSON(&jbodyTicket); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var oldTicket Ticket

    err = ticketCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&oldTicket)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
        return
    }
    oldTicket.Destination = jbodyTicket.Destination

    result, err := ticketCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": oldTicket})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ticket"})
        return
    }

    if result.MatchedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
        return
    }
    // Return updated ticket
    c.JSON(http.StatusOK, gin.H{
        "message": "Ticket updated successfully",
        "ticket":  oldTicket,
    })
}

// DELETE /tickets/:id
func deleteTicket(c *gin.Context) {
    id := c.Param("id")
    // Convert string ID to primitive.ObjectID
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    result, errDelete := ticketCollection.DeleteOne(ctx, bson.M{"_id": objectID})
    if errDelete != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ticket"})
        return
    }

    if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}

func main() {
    // Connect to MongoDB
    connectDB()

    // Set up Gin router
    r := gin.Default()
    // CORS Configuration
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // React frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
    // Routes
    r.POST("/tickets", bookTicket)              // Book ticket
    r.PUT("/tickets/:id", updateDestination)    // Update destination
    r.DELETE("/tickets/:id", deleteTicket)      // Delete ticket
    r.GET("/tickets/:id", readTicketbyId)       //download ticket by id 

    // Start server
    r.Run(":8080")
}
