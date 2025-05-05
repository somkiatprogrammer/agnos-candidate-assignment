package middlewares

import (
	"assignment/initializers"
  "assignment/models"
  
	"database/sql"
	"fmt"
  "net/http"
  "os"
  "strings"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

func AuthMiddleWare() gin.HandlerFunc {

  return func(c *gin.Context) {

    authHeader := c.GetHeader("Authorization")

    if authHeader == "" {
    	c.JSON(http.StatusUnauthorized, gin.H{
        "status": "error",
        "message": "Authorization header is missing",
    	})
    	c.AbortWithStatus(http.StatusUnauthorized)
    	return
    }

    authToken := strings.Split(authHeader, " ")

    if len(authToken) != 2 || authToken[0] != "Bearer" {
    	c.JSON(http.StatusUnauthorized, gin.H{
        "status": "error",
        "message": "Invalid token format",
      })
    	c.AbortWithStatus(http.StatusUnauthorized)
    	return
    }

    tokenString := authToken[1]
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
    		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    	}
    	return []byte(os.Getenv("SECRET")), nil
    })
    if err != nil || !token.Valid {
    	c.JSON(http.StatusUnauthorized, gin.H{
        "status": "error",
        "message": "Invalid or expired token",
      })
    	c.AbortWithStatus(http.StatusUnauthorized)
    	return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
    	c.JSON(http.StatusUnauthorized, gin.H{
        "status": "error",
        "message": "Invalid token",
      })
    	c.Abort()
    	return
    }

    if float64(time.Now().Unix()) > claims["exp"].(float64) {
    	c.JSON(http.StatusUnauthorized, gin.H{
        "status": "error",
        "message": "token expired",
      })
    	c.AbortWithStatus(http.StatusUnauthorized)
    	return
    }

    var staff models.Staff

    sqlStatement := `SELECT * FROM staffs WHERE ID=$1;`
		row := initializers.DB.QueryRow(sqlStatement, claims["id"])
  	err = row.Scan(&staff.ID, &staff.Username, 
	    &staff.Password, &staff.Hospital)

  	switch err {
    	case sql.ErrNoRows:
	      c.JSON(http.StatusUnauthorized, gin.H{
          "status": "error",
          "message": "Cannot Login!",
	      })

	    case nil:
			  c.Set("currentStaff", staff)
				c.Next()

	    default:
	      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	        "status": "error",
	        "message": err.Error(),
	      })

	      panic(err)
	  } 

  }

}