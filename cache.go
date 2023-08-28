package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

type Cache struct {
	cache map[string]float64
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]float64),
	}
}

func CheckValidCacheKey(key string) error {
	if key == "" {
		return fmt.Errorf("Key cannot be empty")
	}
	for _, b := range(key) {
		if !unicode.IsLetter(b) {
			return fmt.Errorf("Key must only contain letters")
		}
	}
	commands := getCommands()
	key = strings.ToLower(key)
	_, exists := commands[key]
	if exists {
		return fmt.Errorf("%s already exists as a command", key)
	}
	constants := getConstants()
	_, exists = constants[key]
	if exists {
		return fmt.Errorf("%s already exists as a constant", key)
	}
	return nil
}

func (c *Cache) Set(key string, value float64) error {
	key = strings.ToLower(key)
	err := CheckValidCacheKey(key)
	if err != nil {
		return err
	}
	_, exists := c.cache[key]
	if exists {
		log.Printf("Key %s already exists, will override to %f", key, value)
	} else {
		log.Printf("Key %s is new, will set it to %f", key, value)
	}
	c.cache[key] = value
	return nil
}

func (c *Cache) Get(key string) (float64, error) {
	value, ok := c.cache[key]
	if !ok {
		return 0, fmt.Errorf("Key not found")
	}
	return value, nil
}

