package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Cache struct {
	cache map[string]float64
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]float64),
	}
}

func isAlpha(s string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z]+$`)
	return pattern.MatchString(s)
}

func CheckValidCacheKey(key string) error {
	if key == "" {
		return fmt.Errorf("Key cannot be empty")
	}
	if !isAlpha(key) {
		return fmt.Errorf("Key must only contain letters")
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

func (c *Cache) Delete(key string) error {
	delete(c.cache, key)
	return nil
}

func (c *Cache) Get(key string) (float64, bool) {
	value, exists := c.cache[key]
	return value, exists
}

func representsNumber(token string, cache *Cache) bool {
	token = strings.ToLower(token)
	_, exists := cache.cache[token]
	if exists {
		return true
	}
	constants := getConstants()
	_, exists = constants[token]
	if exists {
		return true
	}
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}