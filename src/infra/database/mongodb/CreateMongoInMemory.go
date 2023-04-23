package mongodb

import (
	"fmt"
	"os"
	"runtime"

	"github.com/tryvium-travels/memongo"
)

func CreateMongoTemp() *memongo.Server {
	opts := &memongo.Options{
		MongoVersion: "5.0.5",
	}
	if runtime.GOARCH == "arm64" {
		if runtime.GOOS == "darwin" {
			// Only set the custom url as workaround for arm64 macs
			opts.DownloadURL = "https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.5.tgz"
		}
	}
	mongoServer, err := memongo.StartWithOptions(opts)

	if err != nil {
		panic(err)
	}

	os.Setenv("MONGO_URL", mongoServer.URI())
	os.Setenv("DATABASE", memongo.RandomDatabase())

	fmt.Println("MONGO_URL:", os.Getenv("MONGO_URL"))
	fmt.Println("DATABASE: ", os.Getenv("DATABASE"))

	return mongoServer
}
