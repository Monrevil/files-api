package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	path := saveFile(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	initRabbit(path)

}

// Create a temporary file within our temp-images directory that follows
// a particular naming pattern
func saveFile(file []byte) string {
	tempFile, err := ioutil.TempFile("temp-images", "file-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.Write(file)

	abs, err := filepath.Abs(tempFile.Name())
	if err != nil {
		fmt.Println("Could not get absolute filepath:", err)
	}
	return abs
}
