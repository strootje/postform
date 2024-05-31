package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	bolt "go.etcd.io/bbolt"
)

func main() {
	app := fiber.New()

	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		forms, err := tx.CreateBucketIfNotExists([]byte("forms"))
		forms.Put([]byte("contact"), []byte("123123"))
		return err
	})

	app.Post("/submit/:formId", func(c *fiber.Ctx) error {
		var value []byte
		err := db.View(func(tx *bolt.Tx) error {
			forms := tx.Bucket([]byte("forms"))
			value = forms.Get([]byte(c.Params("formId")))
			if value == nil {
				return errors.New("form not found")
			}

			return nil
		})

		if err != nil {
			return err
		}

		return c.SendString(string(value))
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
