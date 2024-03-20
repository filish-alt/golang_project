package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createaccount struct{
	owner string `json:"owner" binding:"required"`
	credit string `json:"credit" binding:"required"`
}

func (server *Server)CreateAccount(c *gin.Context) {
	var account createaccount
   if err :=c.ShouldBindJSON(&account); err !=nil{
       c.JSON(http.StatusBadRequest, errorResponse(err)) 
	   return
   }
}

type getaccount struct{
	id int `uri:"id" binding:"required, min=1"`
}

func(store *Server) getAccount(c *gin.Context){
  var getrequest getaccount
   if err := c.BindUri(&getrequest); err!=nil{

   }
}