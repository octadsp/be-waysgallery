package artsdto

type ArtResponse struct {
	ID    int    `json:"id"`
	Image string `json:"image" form:"image"`
}