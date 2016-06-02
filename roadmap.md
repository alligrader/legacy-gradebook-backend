# Roadmap

The roadmap is my way of issue tracking. It gives me an idea of what we've accomplished and what still needs done. It starts with a high level overview of the Epics that need completed. An Epic is a series of issues (i.e. stories) that make up a complete set of functionality. After the high level overview of the Epics, the rest of the document is devoted to issue tracking.

# Step One: Models

Currently, all models are incomplete. Once I added privileges and roles, everything that I had done was undone. It's kind of demoralizing, so I haven't touched it in a little while.

Right now there are tests and implementations of Teacher and Student objects. However, with the addition of the ACL, there should no longer be Teacher and Student objects. Instead, there should only be a User table.

I did the first step to resolving this, which is to move the model tests into their own package. Then, I should make a new empty package for database tests, and then cherrypick over the testing code from the old package. I need to implement the privileges tests and database code, and then I can finish the implementation of the ORM on all of the other objects.

# Step Two: Submissions

Right now, there's a route that handles the upload of zip files to the server. However, I would much prefer to have an interface that either stores the zip files locally (for testing purposes) or stores them in the Google Cloud (for production). I haven't done too much planning on the implementation details yet, just the high level overview.

# Step Three: Docker Containers

We need to get the docker containers working like they should. We have most of the code done, but it's untested.

# Step Four: Github Integration

We need to listen for Git hooks and handle them approperately. For that, we'll use [this](https://github.com/google/go-github) library. Note that it requires use of [the OAuth library](https://github.com/golang/oauth2). It doesn't look awful, but testing it is going to be a huge chore.

Natalie's working on this one. It requires experimentation.

We need to add the access token to the organization table, and to add the Stripe Payment token to the organization table. I also need create an unguessable state string, probabling using time.Now() and bcrypt.

## Other

- I need to test the `tasks` package
- We need to test the route that accepts a POST request with the zip file.
- We need to create the module that serializes files and puts them on the Google filesystem.
- Then we need to consider under what conditions we move the files onto the Google Cloud Storage. (Only keep the first 3 projects at a time)
- Then, we need to parse the XML output and store it in the database.

- Finally, we need to expose an endpoint to return the XML as JSON, along with the associated lines of code.

- I need to create a fixture file for tests

# TODO Issue Tracker

- [ ] Implement the tests for privileges and roles

- [ ] Get the privileges tests to pass

- [ ] Cherrypick the models tests into an updated testing package

- [ ] Upgrade to Golang 1.7 Beta

- [ ] Add the Dredd tests for each of the routes in the `routes` document.

- [ ] Add the handler and subrouters for each of the routes int he `routes` document.

- [ ] Implement the database code for marshalling privileges and roles

- [ ] Add an organization table

- [ ] Implement a set of integration tests using Response Recorder from `"net/http/httptest"`
- [ ] Projects need a due date

- [ ] Projects need a late submission policy

- [ ] Courses have Announcements created by Teachers

- [ ] Test the `tasks` package

- [ ] Test the route that accepts a POST request with the zip file.

- [ ] Create the module that serializes files and puts them on the Google filesystem.

- [ ] Consider under what conditions we move the files onto the Google Cloud Storage. (Only keep the first 3 projects at a time)

- [ ] Parse the XML output and store it in the database.

- [ ] Expose an endpoint to return the XML as JSON, along with the associated lines of code.

- [ ] Create a fixture file for tests with the right Java code and the right XML output.

- [ ] Add a record for Stripe access tokens

- [ ] Write the SQL for the Course, Student, and Teacher tables so that you can test the success of Assignment

- [ ] Set the unguessable state string using bcrypt and time.Now()

- [ ] Implement a invite-links table

- [ ] Implement a session database object.

- [ ] Connect the session database object to Gorilla

- [ ] Implement the `tasks` package to build and execute the Docker containers

- [ ] Implement the `git hooks` package to receive information from github.

- [ ] Implement the tests for the model

- [ ] Make sure the .Equals model methods is consistent accross all models (checks ID)

- [x] Move the model tests into their own package

- [x] Correct the models to have arrays when referencing foreign keys in a one-to-many relationship.

- [x] Investigate privileges and roles in DB

- [x] Put oauth in it's own package somewhere

- [x] Implement Equals on each of the models
    - This is dependent on having the models corrected.

- [x] In the oauth package, set the gorilla session store and the GetProviderName() function to work with gorilla mux

- [x] Make Person a private struct

- [x] Set up the tables for access control

- [x] Convert oauth into it's own package

- [x] Make sure that Goose doesn't return an error when it's already at the newest migration

- [x] Clean up the `postinstall` scripts and the rest of the Vagrant install bash.

- [x] Tear out the test code that uses the Shell to load the schema and replace it with Goose.

- [x] Make sure that Goose installs and runs on Vagrant start

- [x] After that, we need to add a message to the message bus to run checkstyle.

- [x] After that, we need to add a message to the message bus to run findbugs.

------------------------

## Icebox

- [ ] Implement server_data.go in `db`

- [ ] Make sure that each of the models implements the correct set of interfaces for `api2go`. 

### Iceboxed Epics

# Step Two: Interfaces

I need to make sure that the models implement each of the JSON:API interfaces for `api2go`.
This includes, but is not limited to, the following interfaces:

0. [type MarshalIdentifier](https://godoc.org/github.com/manyminds/api2go/jsonapi#MarshalIdentifier)
0. [type MarshalIncludedRelations](https://godoc.org/github.com/manyminds/api2go/jsonapi#MarshalIncludedRelations)
0. [type MarshalLinkedRelations](https://godoc.org/github.com/manyminds/api2go/jsonapi#MarshalLinkedRelations)
0. [type MarshalReferences](https://godoc.org/github.com/manyminds/api2go/jsonapi#MarshalReferences)
0. [type EntityNamer](https://godoc.org/github.com/manyminds/api2go/jsonapi#EntityNamer)
0. [type UnmarshalIdentifier](https://godoc.org/github.com/manyminds/api2go/jsonapi#UnmarshalIdentifier)
0. [type UnmarshalToManyRelations](https://godoc.org/github.com/manyminds/api2go/jsonapi#UnmarshalToManyRelations)
0. [type UnmarshalToOneRelations](https://godoc.org/github.com/manyminds/api2go/jsonapi#UnmarshalToOneRelations)

When I complete this, I need to add tests to make sure I can send JSON that creates new records in the database.

