// Package user. The memory repository for users
package user

import (
	"e-commerce-system/internal/model/user"
	"fmt"
)

const maxSize = 100

type MemoryRepository struct {
	data map[int]*user.User
	size int
}

func CreateMemoryRepository() *MemoryRepository {
	mr := &MemoryRepository{}
	mr.data = make(map[int]*user.User, maxSize)
	mr.size = 0

	return mr
}

func (mr *MemoryRepository) Add(user *user.User) error {
	if _, ok := mr.data[user.ID]; ok {
		return fmt.Errorf("user %d is already exists", user.ID)
	}

	mr.data[user.ID] = user
	mr.size++

	return nil
}

func (mr *MemoryRepository) Delete(id int) error {
	if _, ok := mr.data[id]; !ok {
		return fmt.Errorf("user %d not found", id)
	}

	delete(mr.data, id)

	return nil
}

func (mr *MemoryRepository) GetById(id int) (error, *user.User) {
	if _, ok := mr.data[id]; !ok {
		return fmt.Errorf("user %d not found", id), nil
	}

	return nil, mr.data[id]
}

func (mr *MemoryRepository) List() []*user.User {
	list := make([]*user.User, 0, mr.size)

	for _, u := range mr.data {
		list = append(list, u)
	}

	return list
}
