// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

func main() {

	ctx := context.Background()
	// Sets your Google Cloud Platform project ID.

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucket := "alvardevlp07.appspot.com"

	// // Creates a Bucket instance.
	// it := client.Bucket(bucket).Objects(ctx, nil)

	// for {
	// 	attrs, err := it.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("No text found.")
	// 	}
	// 	fmt.Println(attrs.Name)
	// }

	f, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("No text found.")
	}
	defer f.Close()

	wc := client.Bucket(bucket).Object("notes.txt").NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Println("Error copying")
	}
	if err := wc.Close(); err != nil {
		fmt.Println("Conexion closed")
		fmt.Println(err)
	}

	// ctx := context.Background()
	// client, err := vision.NewImageAnnotatorClient(ctx)

	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	// defer client.Close()

	// image := vision.NewImageFromURI("gs://alvardevlp07.appspot.com/nota.jpg")
	// annotations, err := client.DetectTexts(ctx, image, nil, 10)

	// if err != nil {
	// 	log.Fatalf("Failed to detect labels: %v", err)
	// }

	// if len(annotations) == 0 {
	// 	fmt.Println("No text found.")
	// } else {
	// 	fmt.Println("Text:")
	// 	fmt.Println(annotations[0].Description)
	// }

}
