package API

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10/translations/id"
)

type Connection struct {
  ID int32 
  // Websocket connection 
  // Game game  
}

type ReceiveMessage struct {
  State string 
  Turn  int32
  isComplete bool 
  Winner int32
}

type SendMessage struct {
  ID int32 
  Move string 
  Auth string
}

var connections = []Connection {
  {ID:0}
}

func main(){
  router := gin.Default()

  router.Run("localhost:8080")
}
