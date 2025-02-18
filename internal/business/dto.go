package business

type CreateBusinessRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}

type GetBusinessRequest struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Rating      float64   `json:"rating"`
    Description string    `json:"description"`
}