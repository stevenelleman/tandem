package reqs

import "sg/services/public-api/internal/models"

// TODO: Consider making generic request interface, BindJSON -> Obj directly + error handling

// Note: The attribute field must be capitalized, otherwise it's private to BindJSON
type CreateNodeReq struct {
	Id string `json:"id"`
	Document string `json:"document"`

	// TODO: Should metadata be received as a byte array?
	Author string `json:"author"`
}

func (req *CreateNodeReq) Obj() *models.Node {
	return &models.Node{
		Id: req.Id,
		Author: req.Author,
		Document: req.Document,
	}
}

type CreateEdgeReq struct {
	Id string `json:"id"`
	Author string `json:"author"`
}

func (req *CreateEdgeReq) Obj() *models.Edge {
	return &models.Edge{
		Id: req.Id,
		Author: req.Author,
	}
}

type CreateScopeReq struct {
	Id string `json:"id"`
	Author string `json:"author"`
}

func (req *CreateScopeReq) Obj() *models.Scope {
	return &models.Scope{
		Id: req.Id,
		Author: req.Author,
	}
}

