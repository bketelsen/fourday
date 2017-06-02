name: Intro to SQL Development in Go
video: 
description: 
level: Beginner
topic: Go
# Intro to SQL Development in Go
## Intro to SQL Development in Go

---
name: Database access in Go
video: 
thumb:
github:
# Database access in Go
## 


The standard library provides the db/sql package which is a very low level abstraction for connecting to common sql databases.

To use db/sql you have to have a database driver that implements db/sql's driver interface.

.link https://github.com/golang/go/wiki/SQLDrivers Common Drivers

Note the compatibility test at the bottom of the wiki - try to stick with drivers that pass.
---
name: MySQL
video: 
thumb:
github:
# MySQL
## 


We'll use github.com/go-sql-driver/mysql in these exercises.

.link http://go-database-sql.org/index.html Overview of db/sql 

Written by Baron Schwartz (VividCortex, Author of High Performance MySQL)
---
name: db/sql Exercise
video: 
thumb:
github:
# db/sql Exercise
## 


Exercise using standard library db/sql package

	cd $GOPATH/src/github.com/bketelsen/sqlx/dbsql
	
	look at source code

	make dockerrun
