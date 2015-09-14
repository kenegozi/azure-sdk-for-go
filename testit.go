package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/Azure/azure-sdk-for-go/storage"
)

func main() {
	acc := os.Getenv("AZURE_QUEUE_ACCOUNT_NAME")
	key := os.Getenv("AZURE_QUEUE_ACCOUNT_KEY")

	az, err := storage.NewBasicClient(acc, key)
	if err != nil {
		log.Println(err)
		return
	}

	blobs := az.GetBlobService()
	containerName := "gopload"
	created, err := blobs.CreateContainerIfNotExists(containerName, storage.ContainerAccessTypeContainer)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("container was created?", created)

	imagePath := `/Users/kenegozi/Dropbox/Photos/Ken/ken-2014-icon-head.jpg`
	imgData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Println(err)
		return
	}
	parameters := storage.PutBlockBlobParameters{
		CustomHeaders: map[string]string{}, //{"Content-Type": "image/jpeg", "x-ms-blob-content-type": "image/jpeg"},
	}

	parameters.ContentType = "slfkdm"

	typeofParams := reflect.TypeOf(parameters)
	valueofParams := reflect.ValueOf(parameters)
	for fix := 0; fix < typeofParams.NumField(); fix++ {
		headerValue := valueofParams.Field(fix).String()
		if headerValue == "" {
			continue
		}
		f := typeofParams.Field(fix)
		ftag := f.Tag
		headerTag := ftag.Get("header")
		aTag := ftag.Get("a")
		xTag := ftag.Get("x")
		log.Println("header z x", headerTag, aTag, xTag)

		if headerTag == "" {
			continue
		}
		headerNames := strings.Split(headerTag, ",")
		for _, hn := range headerNames {
			parameters.CustomHeaders[hn] = headerValue
		}
	}
	log.Printf("%#v\n", parameters)
	return
	//parameters.CustomHeaders["Content-Type"] = "image/jpeg"
	err = blobs.PutBlockBlob(containerName, "momo.jpg", uint64(len(imgData)), bytes.NewReader(imgData), parameters)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(blobs.GetBlobURL(containerName, "momo.jpg"))

}
