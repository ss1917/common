package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"time"
)

func LoginAuthentication(sso_uri string) (name string, id, regionid, status int) {

	type Userinfo struct {
		Username string `json:"username"`
		User_id  int    `json:"user_id"`
		RegionId int    `json:"regionid"`
		status   int    `json:"status"`
	}
	var user Userinfo
	req := httplib.Post(sso_uri)
	req.SetTimeout(2*time.Second, 3*time.Second)
	req.Param("auth_key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InNzcyIsInVzZXJpZCI6MiwicmVnaW9uaWQiOjAsImV4cCI6MTQ5OTc1MzE2OSwianRpIjoiMTU2MTg3MTgwNjAiLCJpc3MiOiJzaGVuc2h1byIsIm5iZiI6MTQ5OTE0ODI2OX0.OZko2PZYZ46wmgsZD0lWViYX4dvFEXoSadOVPIYr1IA")
	str, _ := req.String()
	json.Unmarshal([]byte(str), &user)
	name = user.Username
	id = user.User_id
	status = user.status
	regionid = user.RegionId
	return
}
