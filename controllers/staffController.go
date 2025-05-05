package controllers

import (
  "assignment/initializers"
  "assignment/models"

  "database/sql"
  "net/http"
  "os"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v4"
)

func StaffCreate(c *gin.Context) {

  var staff models.Staff

  if err := c.ShouldBind(&staff); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "message": err.Error(),
    })
    return
  }

  sqlStatement := `
    INSERT INTO staffs (username, password, hospital)
    VALUES ($1, $2, $3)`
  _, err := initializers.DB.Exec(sqlStatement, staff.Username, staff.Password, staff.Hospital)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "status": "error",
      "message": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "status": "success",
  })

}

func StaffLogin(c *gin.Context) {

  var authInput models.AuthInput

  if err := c.ShouldBind(&authInput); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "message": err.Error(),
    })
    return
  }

  var staffFound models.Staff

  sqlStatement := `SELECT * FROM staffs WHERE username=$1 AND password=$2 ` + 
    `AND hospital=$3;`
  row := initializers.DB.QueryRow(sqlStatement, authInput.Username, 
    authInput.Password, authInput.Hospital)
  err := row.Scan(&staffFound.ID, &staffFound.Username, 
    &staffFound.Password, &staffFound.Hospital)

  switch err {
    case sql.ErrNoRows:
      c.JSON(http.StatusBadRequest, gin.H{
        "status": "error",
        "message": "Cannot Login!",
      })

    case nil:
      generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": staffFound.ID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
      })

      token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

      if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
          "status": "error",
          "message": "failed to generate token",
        })
      }

      c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "token": token,
      })

    default:
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
        "status": "error",
        "message": err.Error(),
      })

      panic(err)
  }

}
