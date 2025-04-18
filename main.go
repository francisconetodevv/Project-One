package main

import (
	"exame/model"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Desafio: Sistema de Gerenciamento de Biblioteca com Tipos Diversos")

	frotaEletrics := model.Frota[model.EletricCar]{}
	frotaEletrics.AddVehicle(
		model.EletricCar{
			CarModel:       "Tesla Model S",
			BatteryCapacit: 100,
			Autonomy:       600,
			Review:         time.Now().AddDate(0, -6, 0),
		},

		time.Now(),
		15000,
	)

	frotaCombustao := model.Frota[model.CombustionCar]{}
	frotaCombustao.AddVehicle(
		model.CombustionCar{
			CarModel:   "Fiat Uno",
			PowerHorse: 1000,
			Fuel:       "Gasolina",
			Review:     time.Now().AddDate(0, -8, 0),
		},

		time.Now(),
		45000,
	)

	fmt.Println(model.Report(frotaEletrics))
	fmt.Println("\n" + model.Report(frotaEletrics))
	fmt.Println(frotaCombustao.Vehicle)
}
