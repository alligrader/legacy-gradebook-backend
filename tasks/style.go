package tasks

import (
	"archive/tar"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/fsouza/go-dockerclient"
)

// Create a Dockerfile
// This Dockerfile starts with a Checkstyle container image
// Then it runs CP and copies the temporary directory to it.

// Then, run the Docker image.
// Collect the results.
// String the first and last lines.
// Parse it as XML
// Return that XML document.
// Push that XML document to a new queue
// Or Convert the Document to JSON, then push that ?

type DockerManager struct {
	*docker.Client
}

// CreateDockerfileForCheckStyle starts with the Checkstyle container image,
// which is hosted on a private artifact repository,
// and Copies onto the image the directory we fetched from Github.
// src -- the source directory for the build
func GetClient() *docker.Client {
	c, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (client *DockerManager) CreateDockerfileForCheckStyle(src string) io.Reader {

	var (
		buffer     bytes.Buffer                   // Create a buffer to write our archive to
		tarBuf     = tar.NewWriter(&buffer)       // Create a tarfile
		dockerfile = fmt.Sprintf(Dockerfile, src) // Create a new tar archive.
		now        = time.Now()
	)

	defer tarBuf.Close()

	header := &tar.Header{
		Name:       "Dockerfile",
		Size:       int64(len(dockerfile)),
		ModTime:    now,
		AccessTime: now,
		ChangeTime: now,
	}

	tarBuf.WriteHeader(header)
	tarBuf.Write([]byte(dockerfile))

	return &buffer
}

// BuildCheckStyleImage builds the CheckStyle image created from CreateDockerfileForCheckStyle
func (client *DockerManager) BuildCheckStyleImage(dockerfile io.Reader, imageName string) (io.Writer, error) {

	var output bytes.Buffer
	var tagName string = GenerateTagname() // TODO how the fuck am i generate the tagname?
	_ = tagName

	imageOps := docker.BuildImageOptions{
		Name:         "", // TODO checkstyle + repo-name + sha
		InputStream:  dockerfile,
		OutputStream: &output, // the output buffer
		ContextDir:   "",      // TODO path to th temp directory
	}

	if err := client.BuildImage(imageOps); err != nil {
		return nil, err
	}
	return &output, nil
}

// TODO imple
func GenerateTagname() string {
	return ""
}

// BuildCheckStyleContainer builds the checkstyle container from the image
func (client *DockerManager) BuildCheckStyleContainer(imageName string) *docker.Container {
	config := client.getImageConfig(imageName)

	containerOps := docker.CreateContainerOptions{
		Name:   imageName,
		Config: config,
	}

	container, err := client.CreateContainer(containerOps)
	if err != nil {
		log.Error(err)
	}

	return container
}

func (client *DockerManager) getImageConfig(imageTag string) *docker.Config {

	img, err := client.InspectImage(imageTag)
	if err != nil {
		log.Error(err)
	}
	return img.Config
}

// ExecCheckStyleImage takes the container image we created with BuildCheckStyleImage and runs the container.
func (client *DockerManager) ExecCheckStyleImage(port int, container *docker.Container) {
	ports := make(map[docker.Port][]docker.PortBinding)
	portStr := docker.Port(fmt.Sprintf("%d/tcp", 8000))

	ports[portStr] = []docker.PortBinding{
		{
			HostPort: fmt.Sprintf("%d", port),
			HostIP:   "0.0.0.0",
		},
	}

	err := client.StartContainer(container.ID, &docker.HostConfig{
		PortBindings: ports,
	})

	if err != nil {
		log.Error(err)
	}
}

type CheckstylePayload struct {
	XMLName xml.Name `xml:"checkstyle"`
	Version string   `xml:"version,attr"`
	File    CheckstyleFile
}

type CheckstyleFile struct {
	XMLName xml.Name          `xml:"file"`
	Name    string            `xml:"name,attr"`
	Errors  []CheckstyleError `xml:"error"`
}

type CheckstyleError struct {
	Line     string `xml:"line,attr"`
	Severity string `xml:"severity,attr"`
	Message  string `xml:"message,attr"`
	Source   string `xml:"source,attr"`
}

// ParseXMDocument parses the XML returned from running the CheckStyle container and converts it into a struct
// This is what needs to be implemented for issue #16 ! ! !
func ParseXMLDocument(testOutput []byte) *CheckstylePayload {
	var newPayload CheckstylePayload

	err := xml.Unmarshal(testOutput, &newPayload)
	if err != nil {
		log.Error(err)
	}

	return &newPayload
}

// PersistStyleDetection serializes the object containing all of the style warnings as a database record.
func PersistStyleDetection() {}

const (
	Dockerfile = `
	FROM java:7
	COPY %s/* /usr/src/workdir
	COPY ./.checkstyle/* /usr/src/workdir
	WORKDIR /usr/src/workdir
	`
)
