package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

//Save file to a folder and send message to a RabbitMQ
func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File\n" + err.Error())
		fmt.Println(err)
		return
	}
	defer file.Close()

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

	//Send the filepath to rabbitMQ
	initRabbit(path)

}

// Create a temporary file within our temp-images directory that follows
// a particular naming pattern
func saveFile(file []byte) string {
	err := os.MkdirAll("./temp-images", os.ModePerm)
	if  err!=nil{
		fmt.Println("Could not create a folder: " + err.Error())
	}
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
