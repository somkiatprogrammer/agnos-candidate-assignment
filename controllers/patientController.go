package controllers

import (
  "assignment/initializers"
  "assignment/models"

  "database/sql"
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)


func PatientSearchV1(c *gin.Context) {

  var patient models.Patient

  id := c.Param("id")

  sqlStatement := `SELECT * FROM patients WHERE national_id=$1 OR passport_id=$2;`
  row := initializers.DB.QueryRow(sqlStatement, id, id)
  err := row.Scan(&patient.ID, &patient.FirstNameTH, &patient.MiddleNameTH, 
    &patient.LastNameTH, &patient.FirstNameEN, &patient.MiddleNameEN, 
    &patient.LastNameEN, &patient.DateOfBirth, &patient.PatientHN, 
    &patient.NationalID, &patient.PassportID, &patient.PhoneNumber, 
    &patient.Email, &patient.Gender)

  switch err {
    case sql.ErrNoRows:
      c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "message": "No rows were returned!",
      })

    case nil:
      c.JSON(http.StatusOK, gin.H{
        "first_name_th": patient.FirstNameTH,
        "middle_name_th": patient.MiddleNameTH,
        "last_name_th": patient.LastNameTH,
        "first_name_en": patient.FirstNameEN, 
        "middle_name_en": patient.MiddleNameEN, 
        "last_name_en": patient.LastNameEN, 
        "date_of_birth": patient.DateOfBirth, 
        "patient_hn": patient.PatientHN, 
        "national_id": patient.NationalID, 
        "passport_id": patient.PassportID, 
        "phone_number": patient.PhoneNumber, 
        "email": patient.Email, 
        "gender": patient.Gender, 
      })

    default:
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
        "status": "error",
        "message": err.Error(),
      })
  }

}

func PatientSearchV2(c *gin.Context) {

  var patient models.Patient
  var patients []models.Patient
  var patientSearchInput models.PatientSearchInput

  staff, _ := c.Get("currentStaff")
  currentStaff := staff.(models.Staff)

  if err := c.ShouldBind(&patientSearchInput); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "status": "error",
      "message": err.Error(),
    })
    return
  }

  sqlStatement := `SELECT * FROM patients WHERE patient_hn=$1`

  if patientSearchInput.NationalID != "" {
    sqlStatement += fmt.Sprintf(` AND national_id='%s'`, patientSearchInput.NationalID)
  }

  if patientSearchInput.PassportID != "" {
    sqlStatement += fmt.Sprintf(` AND passport_id='%s'`, patientSearchInput.PassportID)
  }

  if patientSearchInput.FirstName != "" {
    sqlStatement += fmt.Sprintf(` AND (first_name_th='%s' OR first_name_en='%s')`, 
      patientSearchInput.FirstName, patientSearchInput.FirstName)
  }

  if patientSearchInput.MiddleName != "" {
    sqlStatement += fmt.Sprintf(` AND (middle_name_th='%s' OR middle_name_en='%s')`, 
      patientSearchInput.MiddleName, patientSearchInput.MiddleName)
  }

  if patientSearchInput.LastName != "" {
    sqlStatement += fmt.Sprintf(` AND (last_name_th='%s' OR last_name_en='%s')`, 
      patientSearchInput.LastName, patientSearchInput.LastName)
  }

  if patientSearchInput.DateOfBirth != "" {
    sqlStatement += fmt.Sprintf(` AND date_of_birth='%s'`, patientSearchInput.DateOfBirth)
  }

  if patientSearchInput.PhoneNumber != "" {
    sqlStatement += fmt.Sprintf(` AND phone_number='%s'`, patientSearchInput.PhoneNumber)
  }

  if patientSearchInput.Email != "" {
    sqlStatement += fmt.Sprintf(` AND email='%s'`, patientSearchInput.Email)
  }

  rows, err := initializers.DB.Query(sqlStatement, currentStaff.Hospital)

  if err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "status": "error",
      "message": err.Error(),
    })

    panic(err)
  }

  defer rows.Close()

  for rows.Next() {

    err = rows.Scan(&patient.ID, &patient.FirstNameTH, &patient.MiddleNameTH, 
      &patient.LastNameTH, &patient.FirstNameEN, &patient.MiddleNameEN, 
      &patient.LastNameEN, &patient.DateOfBirth, &patient.PatientHN, 
      &patient.NationalID, &patient.PassportID, &patient.PhoneNumber, 
      &patient.Email, &patient.Gender)

    if err != nil {
      c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
        "status": "error",
        "message": err.Error(),
      })
      panic(err)
    }

    patients = append(patients, patient)

  }
  
  err = rows.Err()

  if err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "status": "error",
      "message": err,
    })

    panic(err)
  }

  c.JSON(http.StatusOK, patients)

}
