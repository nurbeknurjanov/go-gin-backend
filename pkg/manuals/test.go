package manuals

type Test struct {
	ID   *int    `json:"id"`
	Name *string `json:"name,omitempty"`
	Age  *string `json:"age"`
}
