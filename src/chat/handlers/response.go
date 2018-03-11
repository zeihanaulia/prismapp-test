package handlers

// StandardResponse standart error status code API pinjam.co.id
type StandardResponse struct {
	HTTPStatusCode int         `json:"http_status_code,omitempty"`
	Success        bool        `json:"success"`
	StateCode      string      `json:"state_code,omitempty"`
	Message        string      `json:"message,omitempty"`
	Data           interface{} `json:"data,omitempty"`
	Pagination     *Pagination `json:"pagination,omitempty"`
}

// Pagination optional for pagination
type Pagination struct {
	TotalCount  int `json:"total_count,omitempty"`
	TotalPages  int `json:"total_pages,omitempty"`
	CurrentPage int `json:"current_page,omitempty"`
	Limit       int `json:"limit,omitempty"`
}
