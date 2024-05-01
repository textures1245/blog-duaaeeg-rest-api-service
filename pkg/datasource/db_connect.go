package datasource

import (
	"github.com/textures1245/BlogDuaaeeg-backend/db"
)

func DbConnect() *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	return client
}
