package envoyer

type Pivot struct {
	ProjectID      int    `json:"project_id"`
	CollaboratorID int    `json:"collaborator_id"`
	Status         string `json:"status"`
}

type Collaborator struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pivot `json:"pivot"`
}

type CollaboratorResponse struct {
	Collaborator `json:"collaborator"`
}

type Collaborators []Collaborator

type CollaboratorsResponse struct {
	Collaborators `json:"collaborators"`
}