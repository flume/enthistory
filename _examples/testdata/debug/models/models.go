package models

type Uint64 uint64

type InfoStruct struct {
	FirstAppearance string `json:"firstAppearance"`
	LastAppearance  string `json:"lastAppearance"`
}

type SpeciesType string

const (
	SpeciesTypeHuman   SpeciesType = "HUMAN"
	SpeciesTypeDog     SpeciesType = "DOG"
	SpeciesTypePenguin SpeciesType = "PENGUIN"
	SpeciesTypeVampire SpeciesType = "VAMPIRE"
)
