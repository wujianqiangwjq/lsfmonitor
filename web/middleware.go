package web

import (
	"fmt"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v2"
)

const (
	ldapserver = "127.0.0.1"
	port       = "389"
	secert     = "sdfsdfsdfeswfesfdss33e"
	dn         = "uid=admin,dc=wujq,dc=com"
	gbn        = "ou=Group,dc=wujq,dc=com"
	pbn        = "ou=People,dc=wujq,dc=com"
	lpass      = "wujqopen"
)

func init() {

}
func Auth(username, password string) bool {
	ldc, _ := ldap.Dial("tcp", fmt.Sprintf("%s:%s", ldapserver, port))
	user := fmt.Sprintf("cn=%s,ou=People,dc=wujq,dc=com", username)
	err := ldc.Bind(user, password)
	if err != nil {
		return false
	} else {
		return true
	}
}
func Authmiddle(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	if token == "" {
		token = c.Request.Header.Get("Authorization")
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}
	}
	if token == "" {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{"id": 203, "message": "donot get token value"})
	} else {
		flag := Auth(ParseToken(token))
		if !flag {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"id": 203, "message": "username or passowrd error"})
		}
	}

}
func GetUsersByGroup(group string) []string {
	var data []string
	ldc, _ := ldap.Dial("tcp", fmt.Sprintf("%s:%s", ldapserver, port))

	sq := ldap.NewSearchRequest(gbn, ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false, fmt.Sprintf("(&(objectClass=posixGroup)(cn=%s))", group), []string{"gidNumber"}, nil)

	if rs, ok := ldc.Search(sq); ok == nil && len(rs.Entries) > 0 {
		gid := rs.Entries[0].GetAttributeValue("gidNumber")
		usq := ldap.NewSearchRequest(pbn, ldap.ScopeWholeSubtree, ldap.DerefInSearching, 0, 0, false, fmt.Sprintf("(&(objectClass=posixAccount)(gidNumber=%s))", gid), []string{"cn"}, nil)
		if urs, uok := ldc.Search(usq); uok == nil {
			for _, item := range urs.Entries {
				data = append(data, item.GetAttributeValue("cn"))
			}

		}
	}
	return data

}
