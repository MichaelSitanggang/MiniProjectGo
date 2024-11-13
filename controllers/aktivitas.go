package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/gin-gonic/gin"
)

type ControllerAktip struct {
	sc services.UseCaseAktivitas
}

func NewControlAktipitas(sc services.UseCaseAktivitas) *ControllerAktip {
	return &ControllerAktip{sc: sc}
}

func (ca ControllerAktip) GetAllAktip(c *gin.Context) {
	aktip, err := ca.sc.AllAktivitas()
	if err != nil {
		c.JSON(500, gin.H{"error": "Data tidak ada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Aktivitas": aktip})
}
