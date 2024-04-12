package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/datumforge/datum/pkg/utils/storage/s3"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

func main() {
	tests3()
}

func tests3() {
	storage, _ := s3.NewStorage(s3.Config{
		AccessKeyID:     "",
		SecretAccessKey: "",
		Region:          "us-east-1",
		Bucket:          "datumtestbucket",
	}, "keyNamespace")

	pdf, _ := os.ReadFile("pkg/utils/pdf/example/zacinvoice.pdf")

	buff := bytes.NewBuffer(pdf)

	id := ulids.New().String()

	if err := storage.SaveQuick(context.TODO(), buff, id); err != nil {
		panic(err)
	}

	if err := os.WriteFile("testwriting.pdf", pdf, 0644); err != nil {
		panic(err)
	}

	url, err := storage.PresignedURL(id, "pdf")
	if err != nil {
		panic(err)
	}

	fmt.Print(url)
}
