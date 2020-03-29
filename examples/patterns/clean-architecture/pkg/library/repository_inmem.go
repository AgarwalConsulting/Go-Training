package library

import (
	"algogrit.com/library/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[entity.ID]*entity.Book
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.Book{}
	return &IRepo{
		m: m,
	}
}

//Store a Book
func (r *IRepo) Store(a *entity.Book) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a Book
func (r *IRepo) Find(id entity.ID) (*entity.Book, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//Search Books
// func (r *IRepo) Search(query string) ([]*entity.Book, error) {
// 	var d []*entity.Book
// 	for _, j := range r.m {
// 		if strings.Contains(strings.ToLower(j.Name), query) {
// 			d = append(d, j)
// 		}
// 	}
// 	if len(d) == 0 {
// 		return nil, entity.ErrNotFound
// 	}

// 	return d, nil
// }

//FindAll Books
func (r *IRepo) FindAll() ([]*entity.Book, error) {
	var d []*entity.Book
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Book
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}
