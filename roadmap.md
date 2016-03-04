# Roadmap

I want to lay out what needs to be done, and what the prioritization on each of the tasks is.
This document will serve as the roadmap for feature work, as well as a complete list of TODOs left in the application.

# Step One: Models

First, I have to correct the models to make sure that I use an array where I'm supposed to.
Next, make sure that all of the models have the correct fields (aren't missing any)
Make sure that Goose can generate them.

Second, I have to implement the model tests so that we can "red-green-refactor" our way to success with the SQL queries (DB package).

Write a "WithCleanDB()" function that looks like this:

    ```golang
    func WithCleanDB(func a ) {
        Migrate()
        defer Clean()
        a()
    }

    ```

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

# Step Three: Repo Creation

I was able to log in to Github using OAuth. Now, I need to make sure I request the right permissions, and store the access token in the database. I also need create an unguessable state string, probabling using time.Now() and bcrypt.

I need to implement the code to interface with the Github API. Luckily, there's [this](https://github.com/google/go-github). Note that it requires use of [the OAuth library](https://github.com/golang/oauth2). It doesn't look awful, but testing it is going to be a huge chore.

Luckily, the end goal is only to programmatically create a set of repos with the right properties. In particular, I need to have a teacher submit a repo name and a course. Then, I get each of the student from the course, and create a new team within the org for each student in the course. Then, I create a new repo for each of the students, and add the student's team to the repo.

# Step Four: Git Hooks

Next, I need to implement a micro service that listens for Git hooks and creates messages in the queue to test a new commit. I might want to use [this](https://github.com/phayes/hookserve), which is at least partially designed to work for Github.

# TODO

- Set the unguessable state string using bcrypt and time.Now()
    - This should probable mean that oauth is going in it's own package somewhere

- In the oauth package, set the gorilla session store and the GetProviderName() function to work with gorilla mux

- Make sure that Goose doesn't return an error when it's already at the newest migration

- Make sure that Goose installs and runs on Vagrant start

- Implement the `tasks` package to build and execute the Docker containers

- Implement the `git hooks` package to receive information from github.

- Implement the tests for the model

- Correct the models to have arrays when referencing foreign keys in a one-to-many relationship.

- Implement Equals on each of the models
    - This is dependent on having the models corrected.

- Complete enough SQL to get all of the tests to pass

- Implement a set of integration tests using Response Recorder from `"net/http/httptest"`

- Make sure that each of the models implements the correct set of interfaces for `api2go`. 

- Implement misc JSON:API functions.

- Implement server_data.go in `db`
