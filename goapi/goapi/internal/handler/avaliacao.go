package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"
)

func PostAvaliacao(c *gin.Context) {
	pizzaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	var novaAvaliacao models.Avaliacao
	if err := c.ShouldBindJSON(&novaAvaliacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, pizzas := range data.Pizzas {
		if pizzas.ID == pizzaID {
			pizzas.Avaliacao = append(pizzas.Avaliacao, novaAvaliacao)
			data.Pizzas[i] = pizzas
			data.SavePizza()
			c.JSON(http.StatusCreated, pizzas)
			return
		}
	}

	if err := service.ValidateNotaAvaliacao(novaAvaliacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"erro": errors.New("pizza not found").Error()})
}
