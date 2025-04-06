package entities

type Information map[string]string

// Client or client to be.
// Anyway, person of interest
type User struct {
	Id     string
	Fields Information
}

// AddInformation upserts key value mapping for that user.
// No action if user is null
func (u *User) AddInformation(key, value string) {
	if u != nil {
		if u.Fields == nil {
			u.Fields = make(Information)
		}

		u.Fields[key] = value
	}
}

// RemoveInformation removes an information by key for that user
// If there is no more field, just set to nil
func (u *User) RemoveInformation(key string) {
	if u != nil && u.Fields != nil {
		delete(u.Fields, key)

		if len(u.Fields) == 0 {
			u.Fields = nil
		}
	}
}
