package tasks

import (
	"bytes"
	"io"
	"os"
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

// CreateDockerfileForCheckStyle starts with the Checkstyle container image,
// which is hosted on a private artifact repository,
// and Copies onto the image the directory we fetched from Github.
// src -- the source directory for the build
func CreateDockerfileForCheckStyle(src string) io.Reader {

	var (
		buffer     bytes.Buffer                   // Create a buffer to write our archive to
		tarBuf     = tar.NewWriter(buffer)        // Create a tarfile
		dockerfile = fmt.Sprintf(Dockerfile, src) // Create a new tar archive.
		now        = time.Now()
	)

	defer tarBuff.Close()

	header := &tar.Header{
		Name:       "Dockerfile",
		Size:       len(dockerfile),
		ModTime:    t,
		AccessTime: t,
		ChangeTime: t,
	}

	tarBuf.WriteHeader(header)
	tarBuf.Write([]byte(dockerfile))

	return buffer
}

// BuildCheckStyleImage builds the CheckStyle image created from CreateDockerfileForCheckStyle
func BuildCheckStyleImage(dockerfile io.Reader, imageName string) (io.Writer, string, error) {

	var output bytes.Buffer
	var tagName string = GenerateTagname() // TODO how the fuck am i generate the tagname?

	imageOps := docker.BuildImageOptions{
		Name:         "", // TODO checkstyle + repo-name + sha
		InputStream:  dockerfile,
		OutputStream: output, // the output buffer
		ContextDir:   "",     // TODO path to th temp directory
	}

	if err := client.BuildImage(imageOps); err != nil {
		return nil, err
	}
	return output, nil
}

// BuildCheckStyleContainer builds the checkstyle container from the image
func BuildCheckStyleContainer(imageName string) *docker.Container {
	config := getImageConfig()

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

func getImageConfig(client *docker.Client, imageTag string) *docker.Config {

	img, err := client.InspecImage(imageTag)
	if err != nil {
		log.Error(err)
	}
	return img.Config
}

// ExecCheckStyleImage takes the container image we created with BuildCheckStyleImage and runs the container.
func ExecCheckStyleImage(client *docker.Client, port int, container *dockerContainer) {
	ports := make(make[docker.Port][docker.PortBinding])
	portStr := docker.Port(fmt.Sprintf("%d/tcp", 8000))

	ports[portStr] = docker.PortBindings{
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

// StripDockerIO strips the first and last line of the output from the container. This leaves us with an XML document.
func StripDockerIO() {}

// ParseXMDocument parses the XML returned from running the CheckStyle container and converts it into a struct
func ParseXMLDocument() {}

// ConvertDocumentToJSON transforms the XML document to a JSON document with similar tags.
func ConvertDocumentToJSON() {}

// PushJSONToMachinery pushes the JSON data to the RabbitMQ message bush via Machinery.
func PushJSONToMachinery() {}

const (
	Dockerfile = `
	FROM CheckStyle
	CPY %s /repo
	`
)
