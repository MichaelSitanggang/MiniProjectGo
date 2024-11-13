package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/gin-gonic/gin"
)

type InputControl struct {
	inputUse services.InputUseCase
}

func NewInputController(uc services.InputUseCase) *InputControl {
	return &InputControl{uc}
}

func (ic *InputControl) GetInputAktivitasAll(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	userID, ok := user_id.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		fmt.Println("ini adalah user id : ", userID)
		return
	}
	aktivitas, err := ic.inputUse.Findall(int(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aktivitas)
}

func (ic *InputControl) CreatedAktivitas(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	userID, ok := user_id.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		fmt.Println(userID)
		return
	}

	var input entities.Input_aktivitas
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.User_id = int(userID)

	if err := ic.inputUse.CreateAktip(int(userID), &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Aktivitas created successfully"})
}

func (ic *InputControl) UpdatedAktivitas(c *gin.Context) {
	user_id, err := c.Get("user_id")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User id tidak ditemuka"})
		return
	}
	userID, _ := user_id.(int)
	aktivitasID, _ := strconv.Atoi(c.Param("id"))
	var aktivitas entities.Input_aktivitas

	if err := c.ShouldBindBodyWithJSON(&aktivitas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "EROR PADA DATA"})
		return
	}

	if err := ic.inputUse.UpdateAktip(aktivitasID, userID, &aktivitas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Gagal di ubah"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Data berhasil diubah"})
}

func (ic *InputControl) DeletedAktivitas(c *gin.Context) {
	user, _ := c.Get("user_id")
	user_id, ok := user.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Eror user_id"})
		return
	}
	aktivitasID, _ := strconv.Atoi(c.Param("id"))
	if err := ic.inputUse.DeleteAktip(aktivitasID, user_id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Data gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Berhasil Dihapus"})
}
