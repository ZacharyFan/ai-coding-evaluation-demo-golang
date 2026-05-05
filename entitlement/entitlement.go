package entitlement

type Store struct {
	Active map[string]string
}

func (s *Store) Enable(accountID string, plan string) {
	if s.Active == nil {
		s.Active = map[string]string{}
	}
	s.Active[accountID] = plan
}

func (s *Store) IsActive(accountID string) bool {
	return s.Active != nil && s.Active[accountID] != ""
}
