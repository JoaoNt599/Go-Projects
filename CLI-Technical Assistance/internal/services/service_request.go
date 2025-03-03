package services

import (
	"fmt"
	"technical-assistence/internal/models"
)

func ViewRequest(s models.ServiceRrequest) {
	fmt.Println("Request: ", s.ID)
	fmt.Println("Client: ", s.Client.Displayinformation())
	if s.Technical != nil {
		fmt.Println("Technical: ", s.Technical.Displayinformation())
	}
	fmt.Println("Service: ", s.Service.Displayinformation())
}
