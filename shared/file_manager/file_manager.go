package file_manager

import (
	"archive/zip"
	"bytes"
	"log"
	"os/exec"
)

func GetZip() error {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFolder(folderName string) error {
	// firstly, create the folder with the name of the playlist with all the videos to be downloaded
	err := exec.Command("bash", "-c", "mkdir", folderName).Run()
	if err != nil {
		return err
	}
	// next, add all mp3s inside the folderName folder
	err = exec.Command("bash", "-c", "mv *.mp3").Run()
	return err
}

//func CleanUp() error {
//	// TODO: add code to remove zip/mp3 files after being downloaded by the user
//}
