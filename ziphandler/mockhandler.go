package ziphandler

import (
	"fmt"
	"net/http"
)

// HandleZipUpload responds to a zip upload by saving the file to the filesystem and adding a new record to the database that maps the submission ID to the filepath.
// Handles the POST request with the zipfile attached.
func MockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html)
}

var html = `
<html>
<head>
  <!--Import Google Icon Font-->
  <link href="http://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
  <!--Import materialize.css-->
</head>

<body>
  <nav>
    <div class="nav-wrapper">
      <a href="/" class="brand-logo">Alligrader</a>
    </div>
  </nav>
  <div class="row">
    <div class="col s12 offset-s1">
      <div class="card-panel white hoverable"><!--Todo make pink-->
        <div>
          <h1>Upload your Things</h1>
        </div>
        <div class="row">
          <div class="input-field col s8 offset-s2">
            <form method="post" action="/zip/upload" enctype="multipart/form-data">
              <input id="name" name="name" type="text" class="validate">
              <label for="name">Name</label>
            </div>
          </div>
         <div class="row">
            <div class="file-field input-field">
              <div class="btn col s2">
                <span>File</span>
                <input type="file" id= "submission" name="submission">
              </div>
              <div class="file-path-wrapper col s9">
                <input class="file-path validate" type="text">
              </div>
            </div>
          </div>
          <div class="row">
            <button class="btn waves-effect waves-light col s2 offset-s5" type="submit" name="action">Submit
            </button>
          </div>
        </form>
        <div class="row"></div>
      </div>
    </div>
  </div>
  <!--  <div class="row" />
  <img class="col s5 offset-s3" src="http://d55ohm6038bug.cloudfront.net/June2014-Bojack/images/flamingline.gif" />-->
</div>
</body>
</html>
`
