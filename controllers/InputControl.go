package controllers

import (
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
		c.JSON(http.StatusUnauthorized, gin.H{"Status": "Gagal", "Kondisi": false, "error": "User ID not found in context"})
		return
	}
	userID, ok := user_id.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": "Gagal", "Kondisi": false, "error": "Invalid user ID"})
		return
	}
	aktivitas, err := ic.inputUse.Findall(int(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Gagal", "Kondisi": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "Aktivitas": aktivitas})
}

func (ic *InputControl) CreatedAktivitas(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	userID, _ := user_id.(int)
	var input entities.Input_aktivitas
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Gagal", "Kondisi": false, "error": err.Error()})
		return
	}
	input.User_id = userID
	if err := ic.inputUse.CreateAktip(userID, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Server eror"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "Berhasil", "Kondisi": true, "message": "Data Berhasil Ditambah"})
}

func (ic *InputControl) UpdatedAktivitas(c *gin.Context) {
	user_id, err := c.Get("user_id")
	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": "Gagal", "Kondisi": false, "message": "User id tidak ditemuka"})
		return
	}
	userID, _ := user_id.(int)
	aktivitasID, _ := strconv.Atoi(c.Param("id"))
	var aktivitas entities.Input_aktivitas

	if err := c.ShouldBindBodyWithJSON(&aktivitas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Eror pada data"})
		return
	}

	if err := ic.inputUse.UpdateAktip(aktivitasID, userID, &aktivitas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Server eror"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "Message": "Data berhasil diubah"})
}

func (ic *InputControl) DeletedAktivitas(c *gin.Context) {
	user, _ := c.Get("user_id")
	user_id, ok := user.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Eror user_id"})
		return
	}
	aktivitasID, _ := strconv.Atoi(c.Param("id"))
	if err := ic.inputUse.DeleteAktip(aktivitasID, user_id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Data gagal dihapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "Message": "Berhasil Dihapus"})
}

func (ic *InputControl) GetHistoryAll(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": "Gagal", "Kondisi": false, "error": "User ID not found in context"})
		return
	}
	userID, ok := user_id.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": "Gagal", "Kondisi": false, "error": "Invalid user ID"})
		return
	}
	history, err := ic.inputUse.FindAllHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Gagal", "Kondisi": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "History": history})
}
