package model

type Seller struct {
	Id              string
	Name            string
	CreatedAuctions []string
}

func NewSeller() *Seller {
	return &Seller{}
}

func (s *Seller) GetId() string {
	return s.Id
}
func (s *Seller) GetName() string {
	return s.Name
}
func (s *Seller) GetCreatedAuctions() []string {
	return s.CreatedAuctions
}

func (s *Seller) SetId(v string) *Seller {
	s.Id = v
	return s
}
func (s *Seller) SetName(v string) *Seller {
	s.Name = v
	return s
}
func (s *Seller) SetCreatedAuctions(v []string) *Seller {
	s.CreatedAuctions = v
	return s
}
