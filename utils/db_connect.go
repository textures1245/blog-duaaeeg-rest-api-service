package utils

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
)

func DbConnect() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
}
