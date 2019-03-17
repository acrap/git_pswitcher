package profile

import (
	"os"
	"testing"
)

func TestWriteOne(t *testing.T) {
	db := JsonFileDb{path: "/tmp/test_jsondb"}
	var err error
	profile := Profile{Name: "1", Email: "jsahd@as.com"}
	os.Create(db.path)

	if err = db.WriteProfiles([]Profile{profile}); err != nil {
		t.Errorf("Can't write one profile: %s", err)
	}

	// check file
	var profiles []Profile
	if profiles, err = db.GetProfiles(); err != nil {
		t.Errorf("Can't get profiles: %s", err)
	}

	if len(profiles) != 1 {
		t.Errorf("There must be the one element")
	}

}

func TestDelete(t *testing.T) {
	db := JsonFileDb{path: "/tmp/test_jsondb_td"}
	var err error
	var profiles []Profile
	os.Create(db.path)

	if err = db.WriteProfiles([]Profile{Profile{Name: "1", Email: "jsahd@as.com"},
		Profile{Name: "2", Email: "jsahd2@as.com"},
		Profile{Name: "3", Email: "jsahd3@as.com"}}); err != nil {
		t.Errorf("Can't write one profile: %s", err)
	}
	db.DeleteProfile("2")
	profiles, err = db.GetProfiles()
	if err != nil {
		t.Errorf("Can't get the profiles: %s", err)
	}

	for num, elem := range profiles {
		if num == 0 {
			if elem.Name != "1" {
				t.Errorf("Wrong content after deleting an element: %s", err)
			}
		}
		if num == 1 {
			if elem.Name != "3" {
				t.Errorf("Wrong content after deleting an element: %s", err)
			}
		}
	}

}

func TestValidation(t *testing.T) {
	var err error
	profile := Profile{Name: "1"}
	if err = profile.SetEmail("a.user13@mail.ru"); err != nil {
		t.Errorf("email validation failed")
	}

}
