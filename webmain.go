package main

import (
	"encoding/json"
	//	"io/ioutil"
	"log"
	"lsfproject/jobsearch"
	"net/http"
	"web"

	"github.com/gin-gonic/gin"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	search_grpc_port = "127.0.0.1:5355"
)

var jobclient jobsearch.JobClient

func init() {
	conn, err := grpc.Dial(search_grpc_port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	jobclient = jobsearch.NewJobClient(conn)
}

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RequestData struct {
	Name   []string `form:"name" json:"name" `
	Status []string `form:"status" json:"status"`
}

func AlljobGroup(c *gin.Context) {
	var users []string
	data := c.DefaultQuery("args", "{}")
	var param RequestData
	err := json.Unmarshal([]byte(data), &param)
	if err != nil {
		c.JSON(200, err)
	} else {
		if len(param.Name) > 0 {
			for _, item := range param.Name {
				tmp := web.GetUsersByGroup(item)
				users = append(users, tmp...)
			}
			log.Println(users)
			if len(users) > 0 {
				var req jobsearch.Requestjob
				req.User = users
				req.Status = param.Status
				log.Println(req)
				res, ok := jobclient.FindJob(context.Background(), &req)
				if ok == nil {
					data, err := json.Marshal(res)
					if err == nil {
						c.JSON(200, string(data))
					} else {
						c.JSON(200, gin.H{"status": "find error"})
					}
				} else {
					c.JSON(200, gin.H{"status": "find error"})
				}
			} else {
				c.JSON(200, "{}")
			}
		} else {
			c.JSON(200, "{}")
		}
	}
}

func AlljobUser(c *gin.Context) {
	status := c.DefaultQuery("args", "{}")
	var param RequestData
	json.Unmarshal([]byte(status), &param)
	var req jobsearch.Requestjob
	req.User = param.Name
	req.Status = param.Status
	res, ok := jobclient.FindJob(context.Background(), &req)
	if ok == nil {
		data, err := json.Marshal(res)
		if err == nil {
			c.JSON(200, string(data))
		} else {
			c.JSON(200, gin.H{"status": "find error"})
		}
	} else {
		c.JSON(200, gin.H{"status": "find error"})
	}

}

//func AlljobUser(c *gin.Context) {
//	users := c.DefaultQuery("users", "all")
//	status := c.DefaultQuery("status", "all")
//}
func main() {
	r := gin.Default()

	r.POST("/api/login", func(c *gin.Context) {
		var user User
		if c.BindJSON(&user) == nil {
			username := user.Username
			passwd := user.Password

			if username != "" && passwd != "" {
				if web.Auth(username, passwd) {
					log.Println("ok********************8")
					c.JSON(http.StatusOK, gin.H{
						"token": web.GetToken(username, passwd),
					})
				} else {
					log.Println("failed********************8")
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "user validate account",
					})
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "username or passwd are empty",
				})
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "username or passwd are empty",
			})
		}
	})
	group := r.Group("/api", web.Authmiddle)
	group.GET("/alljob/group/", AlljobGroup)
	group.GET("/alljob/user/", AlljobUser)
	r.Run(":80")

}
