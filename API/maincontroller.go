package API

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrade{};

type Connection struct {
  ID int64 
  Socket websocket  
  // Game game  
}

type ReceiveMessage struct {
  State string 
  Turn  int64
  isComplete bool 
  Winner int64
}

type SendMessage struct {
  ID int32 
  Move string 
  Auth string
}

var connections = []Connection {
  {ID:0},
}

func addConnection(c *gin.Context){
    id := c.Param("id")
    var newCon Connection; 
    newCon.ID, _ = strconv.ParseInt(id,10,32)
    check := true;
    for _, con := range connections {
          if(con.ID == newCon.ID){
              check = false;
              break;
          }
    }
    if(check){
      connections = append(connections, newCon)
    }

}

func createConnection(w http.ResponseWriter, r *http.Request){
  conn, err := upgrader.Upgrad(w,r, nil) 
  if err != nil {
    fmt.Println("Failed to create connection")
    return
  }

  for {
  t, msg, err := conn.ReadMessage()
  if err != nil {
    break
  }
  conn.WriteMessage(t, msg) 
  }
}


func main(){
  router := gin.Default()
  router.GET("newCon:id", addConnection)
  router.GET("/ws", func(c *gin.Context) {
    createConnection(c.Writer, c.Request)
  })
  router.Run("localhost:8080")
}
