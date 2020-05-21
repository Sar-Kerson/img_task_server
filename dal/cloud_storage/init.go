package cloud_storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

const (
	BuckerName = "ylq_server"
	URL        = "https://storage.googleapis.com/ylq_server/%s"
)

var (
	StorageClient *storage.Client
)

func init() {
	// Creates a client.
	var err error
	ctx := context.TODO()
	StorageClient, err = storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

func Upload(ctx context.Context, buf []byte, objName string) (string, error) {
	if len(buf) == 0 {
		return "", nil
	}

	r := bytes.NewReader(buf)

	if r == nil {
		return "", errors.New("failed to new reader")
	}

	wc := StorageClient.Bucket(BuckerName).Object(objName).NewWriter(ctx)
	if _, err := io.Copy(wc, r); err != nil {
		log.Println(err)
		return "", err
	}

	if err := wc.Close(); err != nil {
		log.Println(err)
		return "", err
	}

	return formatUrl(objName), nil
}

func Download(ctx context.Context, objName string) ([]byte, error) {
	if len(objName) == 0 {
		return []byte{}, nil
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	rc, err := StorageClient.Bucket(BuckerName).Object(objName).NewReader(ctx)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	return data, nil
}

func formatUrl(objName string) string {
	return fmt.Sprintf(URL, objName)
}
