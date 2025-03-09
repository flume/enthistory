package models

type Uint64 uint64

func DefaultUint64() Uint64 {
	return 0
}

type InfoStruct struct {
	FirstAppearance string `json:"firstAppearance"`
	LastAppearance  string `json:"lastAppearance"`
}

func DefaultInfoStruct() InfoStruct {
	return InfoStruct{
		FirstAppearance: "UNKNOWN",
		LastAppearance:  "UNKNOWN",
	}
}

type SpeciesType string

const (
	SpeciesTypeHuman   SpeciesType = "HUMAN"
	SpeciesTypeDog     SpeciesType = "DOG"
	SpeciesTypePenguin SpeciesType = "PENGUIN"
	SpeciesTypeVampire SpeciesType = "VAMPIRE"
	SpeciesTypeUnknown SpeciesType = "UNKNOWN"
)

func (e SpeciesType) String() string {
	return string(e)
}

func DefaultSpeciesType() SpeciesType {
	return SpeciesTypeUnknown
}
