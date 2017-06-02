name: sqlx Concepts
video: 
description: 
level: Beginner
topic: Go
# sqlx Concepts
## sqlx Concepts

---
name: sqlx
video: 
thumb:
github:
# sqlx
## 

Why sqlx?

Now that you've seen how to do it the hard way, we're going to learn how to use sqlx, which is a set of extensions on top of db/sql intended to make it easier to use without being magic or surprising.

Features that make it worthwhile:

- Marshal rows into structs (with embedded struct support), maps, and slices
- Named parameter support including prepared statements
- Get and Select to go quickly from query to struct/slice
- Just the first one is enough reason to use sqlx - scanning fields into structs is a pain!
---
name: sqlx Concepts
video: 
thumb:
github:
# sqlx Concepts
## 


Connecting

A DB instance is not a connection, but an abstraction representing a Database. This is why creating a DB does not return an error and will not panic. It maintains a connection pool internally, and will attempt to connect when a connection is first needed. You can create an sqlx.DB via Open or by creating a new sqlx DB handle from an existing sql.DB via NewDb

	var db *sqlx.DB
	// exactly the same as the built-in
	db = sqlx.Open("sqlite3", ":memory:")

	// from a pre-existing sql.DB; note the required driverName
	db = sqlx.NewDb(sql.Open("sqlite3", ":memory:"), "sqlite3")
---
name: Connect
video: 
thumb:
github:
# Connect
## 


You may want to connect and ensure that your database is available at the same time:

	var err error
	// open and connect at the same time:
	db, err = sqlx.Connect("sqlite3", ":memory:")

	// open and connect at the same time, panicing on error
	db = sqlx.MustConnect("sqlite3", ":memory:")
---
name: Exercise-SQLX
video: 
thumb:
github:
# Exercise
## 


Read the rest of the overview of sqlx 
.link http://jmoiron.github.io/sqlx/ sqlx overview
---
name: sqlx Exercise
video: 
thumb:
github:
# sqlx Exercise
## 


Now let's look at a sqlx version of the first example we did.

	cd $GOPATH/src/github.com/bketelsen/sqlx/sqlxquery
	view code together
	make dockerrun
---
name: sqlx Exercise2
video: 
thumb:
github:
# sqlx Exercise
## 


Now let's look at a sqlx version of the first example we did.

	cd $GOPATH/src/github.com/bketelsen/sqlx/sqlxquery
	view code together
	make dockerrun
---
name: Exec Assignment
video: 
thumb:
github:
# Exec Assignment
## 


github.com/bketelsen/sqlx/sqlxexec

Finish the insertEmployee() function to insert Elvis into the employee database.
