package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/Eri-Vadi/go_kitchen/internal/domain/entity"
	"github.com/Eri-Vadi/go_kitchen/internal/infrastracture/logger"
)

var (
	foods    []entity.Food
	onceFood sync.Once
)

func GetFoods() []entity.Food {
	onceFood.Do(func() {
		foodsHolder := struct {
			Data []entity.Food `json:"foods"`
		}{}

		jsonFile, _ := os.Open("./config/foods.json")
		byteValue, _ := ioutil.ReadAll(jsonFile)

		if err := json.Unmarshal(byteValue, &foodsHolder); err != nil {
			logger.LogPanicF("Could not unmarshal foods config: %v", err)
		}

		foods = foodsHolder.Data
	})

	return foods
}
