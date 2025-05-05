package models

type AuthInput struct {
  Username string `form:"username" binding:"required"`
  Password string `form:"password" binding:"required"`
  Hospital string `form:"hospital" binding:"required"`
}