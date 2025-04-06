package storage

import (
	"errors"

	"github.com/zefrenchwan/m3.git/properties"
)

func InitDao(config properties.PropertiesMap) (Dao, error) {
	db := config["DBTYPE"]
	if db == "sqlite" {
		path := config["SQLITE"]
		return NewEmbeddedDao(path)
	}

	return nil, errors.New("No dao matching: " + db)
}
