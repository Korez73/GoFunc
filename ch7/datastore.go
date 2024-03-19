package ch7

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// factory
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}
