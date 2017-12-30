Flamin.io
========

Welcome to the Flamin.io repository. This readme is a work in progress, so do not expect much.

## Building the frontend

To build the frontend of  this repository (located at `src/flaminio/webroot`), 
use [Yarn](https://yarnpkg.com/lang/en/) and run the following command in the root directory:  
`yarn install`  
Afterwards, for development, run `gulp build`, assuming you have [Gulp](https://gulpjs.com/) globally installed.

## Building the backend

Building the backend requires more effort. First, make sure you have [Go](https://golang.org/) installed.  

Next, make sure you have a [PostgreSQL](https://www.postgresql.org/) server running (version 10 or later) and make sure to install the following extensions:
* `uuid-ossp`
* `citext`
* `btree_gist`

You can easily do this using `CREATE EXTENSION {extension_name}`. Next, you will have
to make sure that a `schema` exists named `flaminio` inside a database which name and acces user you have to configure using the
connection string inside `src/flaminio/database/database.go`. A more flexible setup solution will follow soon^(tm).

Next, you need to build the `src/flaminio` package using `go build`. Now you can run the go app, which should serve the app to
`localhost:8080`. If the server detects an empty database, it will initialize it with a default user, who's credentials are as follows:
* Email: `admin@admin.com`
* Password: `admin`
