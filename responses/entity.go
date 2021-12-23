package responses

type EntityCreatedSuccessFullyJSON struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	DateTime string `json:"DateTime"`
}

type EntityAllreadyExistJSON struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	DateTime string `json:"DateTime"`
}
