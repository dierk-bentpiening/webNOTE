package responses

type EntityCreatedSuccessFullyJSON struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Title 	 string `json:"title"`
	DateTime string `json:"DateTime"`
}