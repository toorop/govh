package cloud

// Image repesents Image
type Image struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Visibility   string `json:"visibilty"`
	Status       string `json:"status"`
	Region       string `json:"region"`
	MinDisk      int    `json:"minDisk"`
	CreationDate string `json:"creationDate"`
	MinRAM       int    `json:"minRam"`
	User         string `json:"user"`
	Type         string `json:"type"`
}
