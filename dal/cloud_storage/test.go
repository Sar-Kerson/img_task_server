package cloud_storage

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

// https://cloud.google.com/storage/docs/access-public-data
// https://storage.googleapis.com/[BUCKET_NAME]/[OBJECT_NAME]
func upload(filePath, obj string) {
	ctx := context.Background()
	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	bucketName := "ylq_server"

	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := client.Bucket(bucketName).Object(obj).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		log.Println(err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Println(err)
		return
	}
	log.Println("upload good")
}

func download(obj string) {
	ctx := context.Background()
	bucketName := "ylq_server"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	rc, err := client.Bucket(bucketName).Object(obj).NewReader(ctx)
	if err != nil {
		log.Println(err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Println(err)
	}
	log.Println("get!!", string(data))
}

func main() {
	upload("/Users/eric/Downloads/selfie2anime/trainB/3396.jpg", "test.jpg")
	download("test.jpg")
}
