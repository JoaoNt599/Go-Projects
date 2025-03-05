package models

import "time"

type ServiceRrequest struct {
	ID          int
	RequestDate time.Time
	StartDate   *time.Time
	EndDate     *time.Time
	Client      Client
	Attendant   Attendant
	Technical   *Technical
	Service     Service
	Materials   []Material
}

func (s *ServiceRrequest) AssignTechnician(technical Technical) {
	s.Technical = &technical
}

func (s *ServiceRrequest) StartService() {
	now := time.Now()
	s.StartDate = &now
}

func (s *ServiceRrequest) CompleteService() {
	now := time.Now()
	s.EndDate = &now
}

func (s *ServiceRrequest) AddMaterial(material Material) {
	s.Materials = append(s.Materials, material)
}
