package uc

import (
	"github.com/aircraft-catalog/domain"
)

// GET aircraft?name=a380
// GET /autocomplete?name=asdfasdfsf
func (cl *coreLogic) GetAircraftSheetsByName(name string) []*domain.AircraftSheet {
	return cl.AircraftSheetProvider.FindByName(name)
}

// GET /aircrafts
func (cl *coreLogic) GetAircraftSheets() []*domain.AircraftSheet {
	return cl.AircraftSheetProvider.FindAll()
}

// PUT /aircraft
func (cl *coreLogic) UpdateAircraftSheet(sheet *domain.AircraftSheet) (int, error) {
	return cl.AircraftSheetProvider.Update(sheet)
}

// GET /aircraft/:id
func (cl *coreLogic) GetAircraftDetails(ID int) *domain.AircraftSheet {
	return cl.AircraftSheetProvider.FindByID(ID)
}

// DELETE /aircraft/:id
func (cl *coreLogic) RemoveAircraftSheet(ID int) error {
	return cl.AircraftSheetProvider.Remove(ID)
}
