package entity

// NewEntityService to set up the entity service
func NewEntityService() IEntity {
	return &SEntity{}
}
