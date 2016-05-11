package ziphandler

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/alligrader/gradebook-backend/tasks"
	"github.com/alligrader/gradebook-backend/util"

	log "github.com/Sirupsen/logrus"
	_ "github.com/mholt/archiver"
	"github.com/spf13/viper"
)

func init() {
	util.Configure()
	util.ConfigureLogger()
	storagePath = viper.GetString("STORAGE_PATH")
}

var (
	storagePath string
)

// HandleZipUpload responds to a zip upload by saving the file to the filesystem and adding a new record to the database that maps the submission ID to the filepath.
// Handles the POST request with the zipfile attached.
func HandleZipUpload(w http.ResponseWriter, r *http.Request) {

	// TODO Get the student's peoplesoft
	// Write a middleware that validates peoplesoft numbers

	file, _, err := r.FormFile("submission")
	if err != nil {
		log.Error(err)
		// generate 400 error
	}
	defer file.Close()

	// Get the student's name
	submitter := r.FormValue("name")

	// TODO Create a submission record
	submissionID := GenerateSubmissionID(submitter)
	storeZip(file, submissionID)

	// TODO Push new tasks to the message bus
	tasks.PushFindbugs(submissionID)
	tasks.PushCheckstyle(submissionID)
}

func GenerateSubmissionID(submitter string) int {
	return 1
}

func storeZip(file multipart.File, submissionID int) error {
	// Copy the file to the storage path
	subID := strconv.Itoa(submissionID) + ".zip"
	fullpath := filepath.Join(storagePath, subID)
	f, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}
