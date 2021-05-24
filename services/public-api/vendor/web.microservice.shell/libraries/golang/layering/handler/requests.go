package handler

// Note: The attribute field must be capitalized, otherwise it's private to BindJSON
type createSiloReq struct {
	State string `json:"state"`
}

type updateSiloReq struct {
	State string `json:"state"`
}
