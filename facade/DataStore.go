package facade

type DataStore struct {
	DB
	Cache
}

func New() *DataStore {
	return &DataStore{
		DB:    DB{},
		Cache: Cache{},
	}
}

func (d *DataStore) Create(id int) {
	d.DB.Create(id)
	d.Cache.Create(id)
}

func (d *DataStore) Update(id int) {
	d.DB.Update(id)
	d.Cache.Update(id)
}

func (d *DataStore) Delete(id int) {
	d.DB.Delete(id)
	d.Cache.Delete(id)
}
