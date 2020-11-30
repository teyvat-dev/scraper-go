package wikicharacterscrapertypes

// Prisma

// CharacterProfileVoiceActorPrisma JSON struct of characterProfileVoiceActor, suitable for prisma upload
type CharacterProfileVoiceActorPrisma struct {
	EN string `json:"en"`
	CN string `json:"cn"`
	JP string `json:"jp"`
	KR string `json:"kr"`
}

// CharacterProfilePrisma JSON struct of characterProfile, suitable for prisma upload
type CharacterProfilePrisma struct {
	Affiliation string `json:"affiliation"`
	// CardImage   string `json:"cardImage"`
	// IconImage   string `json:"iconImage"`
	Birthday      string                           `json:"birthday"`
	Constellation string                           `json:"constellation"`
	Overview      string                           `json:"overview"`
	Story         interface{}                      `json:"story"`
	VoiceActor    CharacterProfileVoiceActorPrisma `json:"voiceActor"`
	VoiceLines    interface{}                      `json:"voiceLines"`
	Region        string                           `json:"region"`
	SpecialtyDish string                           `json:"specialtyDish"`
	Vision        string                           `json:"string"`
}

// CharacterTalentsPrisma JSON struct of talents, suitable for prisma upload
type CharacterTalentsPrisma struct {
	Description interface{} `json:"description"`
	Details     interface{} `json:"details"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	// TalentLevelUpMaterial string `json:"talentLevelUpMaterial"`
}

// CharacterPrisma JSON struct of character, suitable for prisma upload
type CharacterPrisma struct {
	Name           string                    `json:"name"`
	Constellations []*CharacterConstellation `json:"constellations"`
	Overview       string                    `json:"overview"`
	Rarity         int                       `json:"rarity"`
	Stats          interface{}               `json:"stats"`
	// Image   string `json:"image"`
	Weapon   string   `json:"weapon"`
	Elements []string `json:"elements"`
	// Ascensions CharacterAscensionsPrisma `json:"ascensions"`
	// Sex     string `json:"sex"`
	CharacterProfile CharacterProfilePrisma    `json:"characterProfile"`
	Talents          []*CharacterTalentsPrisma `json:"talents"`
}

// CharacterPrismaPayload JSON map of character, suitable for upload parser
type CharacterPrismaPayload map[string]CharacterPrisma

// Processing Types

type CharacterTableInfo struct {
	Rarity  int
	Image   string
	Name    string
	Element string
	Weapon  string
	Sex     string
	Nation  string
}

type CharacterProfileInfo struct {
	Name           string
	Image          string
	Introduction   string
	Personality    string
	Birthday       string
	Constellation  string
	Affiliation    string
	Dish           string
	VoiceEN        string
	VoiceCN        string
	VoiceJP        string
	VoiceKR        string
	Talents        []*CharacterTalent
	Constellations []*CharacterConstellation
}

type CharacterTalent struct {
	Type string
	Name string
	Icon string
	Info string
}

type CharacterConstellation struct {
	Level  int
	Name   string
	Effect string
}
