package main

import (
	"fmt"
	"technical-assistence/internal/models"
	"technical-assistence/internal/services"
)

func main() {
	client := models.Client{"John Neto", "123.456.789-00", "Street A, 123", "11000000000"}
	attendant := models.Attendant{"Maria Oliver", "987.654.321-00"}
	technical := models.Technical{"Kevin Lima", "111.222.333-44", "11000000000"}
	service := models.Service{1, "Screen change", 300.0}
	material := models.Material{101, "LCD Screen", 150.0}

	request := models.ServiceRrequest{ID: 1, Client: client, Attendant: attendant, Service: service}
	request.AssignTechnician(technical)
	request.StartService()
	request.AddMaterial(material)
	request.CompleteService()

	services.ViewRequest(request)
	services.ListMaterials(request.Materials)

	fmt.Println("Process completed!")
}
