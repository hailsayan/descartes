package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	name  string
	price float64
	count int
}

type Store struct {
	products []Product
}

func NewStore() *Store {
	return &Store{
		products: []Product{},
	}
}

func (s *Store) AddProduct(name string, price float64, count int) error {
	if price <= 0.0 {
		return errors.New("price should be positive")
	}
	if count <= 0 {
		return errors.New("count should be positive")
	}
	for _, product := range s.products {
		if product.name == strings.ToLower(name) {
			return errors.New(fmt.Sprintf("%s already exists", name))
		}
	}

	s.products = append(s.products, Product{strings.ToLower(name), price, count})
	return nil
}

func (s *Store) GetProductCount(name string) (int, error) {
	for _, v := range s.products {
		if v.name == strings.ToLower(name) {
			return v.count, nil
		}
	}
	return 0, errors.New("invalid product name")
}

func (s *Store) GetProductPrice(name string) (float64, error) {
	for _, v := range s.products {
		if v.name == strings.ToLower(name) {
			return v.price, nil
		}
	}

	return 0.0, errors.New("invalid product name")
}

func (s *Store) Order(name string, count int) error {
	if count <= 0 {
		return errors.New("count should be positive")
	}
	for pi, v := range s.products {
		if v.name == strings.ToLower(name) {
			if v.count == 0 {
				return errors.New(fmt.Sprintf("there is no %s in the store", name))
			}
			if v.count < count {
				return errors.New(fmt.Sprintf("not enough %s in the store. there are %d left", name, v.count))
			}
			s.products[pi].count -= count
			return nil
		}
	}
	return errors.New("invalid product name")
}

func (s *Store) ProductsList() ([]string, error) {
	var existingProducts = make([]string, 0)

	for _, v := range s.products {
		if v.count > 0 {
			existingProducts = append(existingProducts, v.name)
		}
	}

	if len(existingProducts) == 0 {
		return nil, errors.New("store is empty")
	}

	sort.Strings(existingProducts)
	return existingProducts, nil
}
