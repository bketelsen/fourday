name: Effective Error Management
video: 
description: 
level: Intermediate
topic: Go
# Effective Error Management
## Effective Error Management

---
name: Effective Error Management
video: 
thumb:
github:
# Effective Error Management
## 


In this section you'll learn how to create and handle errors effectively.  

Topics:

- Creating Errors
- Handling Errors
- Differences between libraries and applications
- Guidelines

You'll end up with a clear guideline for error handling, in both your library code and your applications.
---
name: Philosophy: Go and Errors
video: 
thumb:
github:
# Philosophy: Go and Errors
## 


Proverb:
	Errors are just values.

.caption _Rob_Pike_ at [[https://www.youtube.com/watch?v=PAAkCSZUG1c][Gopherfest SV, 2015]]



.link https://go-proverbs.github.io/
---
name: Philosophy: Go and Errors2
video: 
thumb:
github:
# Philosophy: Go and Errors
## 


What does that mean?

Functionally, it means that *error* is an interface in Go.  Any type that implements the function 

	Error() string

can be used as an *error* type.
---
name: Philosophy: Go and Errors3
video: 
thumb:
github:
# Philosophy: Go and Errors
## 


Conceptually, it means that you have a lot of freedom and flexibility to do error handling.
---
name: Creating Errors - errors.New()
video: 
thumb:
github:
# Creating Errors - errors.New()
## 

You can use the `errors` package to create a descriptive error on the fly
	return errors.New("Login Failed")

- Simple
- Least amount flexibility
---
name: Creating Errors - Wrapping an Error
video: 
thumb:
github:
# Creating Errors - Wrapping an Error
## 


- You can use fmt.Errorf() to wrap an error with more context
	return fmt.Errorf("Login failed:", err)

- Adds some context
- Harder to determine original error type
---
name: Creating Errors - Exporting Named Errors
video: 
thumb:
github:
# Creating Errors - Exporting Named Errors
## 


- You can create known error types and return them 

	return ErrInvalidLogin

.code errmgmt/includes/errs/errtype.go /START OMIT/,/END OMIT/

- Easy to compare and test
- Requires discipline to remember to define these when you find new error conditions
---
name: Creating Errors - Exporting Custom Types 
video: 
thumb:
github:
# Creating Errors - Exporting Custom Types 
## 


- You can create a custom error type

	Code next slide

- Great flexibility
- Good when you might want additional error state stored in the error (time)
- I've used these in web applications/API's where I want to send a very consistent message back to the user, serialized as JSON.
---
name: Creating Errors - Exporting Custom Types 2
video: 
thumb:
github:
# Creating Errors - Exporting Custom Types 
## 


.code errmgmt/includes/errs/errstruct.go /START OMIT/,/END OMIT/
---
name: Creating Errors - Your Own Types
video: 
thumb:
github:
# Creating Errors - Your Own Types
## 


- You can implement the `error` interface on your domain types

.code errmgmt/includes/errs/owntype.go /START OMIT/,/END OMIT/

- Creative
- Requires setting your type's ErrMsg equivalent when an error occurs
- Great for a state machine
---
name: Handling Errors
video: 
thumb:
github:
# Handling Errors
## 


There are really three ways to test error conditions

- Compare with known values
- Assert error type
- Ignore and pass up the call stack
---
name: Handling Errors - Sentinel Errors 
video: 
thumb:
github:
# Handling Errors - Sentinel Errors 
## 


If an error is your own package, or exported from an external package, you can test the error in your code.

Dave Cheney calls these "Sentinel errors" *[1]

	if err == io.EOF


[1] http://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
---
name: Handling Errors - Sentinel Errors2
video: 
thumb:
github:
# Handling Errors - Sentinel Errors
## 


.code errmgmt/includes/errs/knowntype.go /START OMIT/,/END OMIT/
---
name: Handling Errors - Sentinel Errors3
video: 
thumb:
github:
# Handling Errors - Sentinel Errors
## 


- No context available to the caller, just a known error condition
- High coupling between caller and called packages
- Defeated if you add extra context by wrapping the error:
	return fmt.Errorf("DB Error:", ErrDBUnavailable)
	// NOT == ErrDBUnavailable now

- Many packages export error types, especially in network, OS, and database operations
---
name: Handling Errors - Assert with Custom Type
video: 
thumb:
github:
# Handling Errors - Assert with Custom Type
## 


Using custom types provides you with more flexibility in your error handling logic.


.code errmgmt/includes/errs/structcheck.go /START OMIT/,/END OMIT/
---
name: Handling Errors - Assert with Custom Types2
video: 
thumb:
github:
# Handling Errors - Assert with Custom Types
## 


- Testing for error type lets you embed useful context in your error
- Error handling can be very specific
- Still highly coupled, sometimes this is OK and preferred
---
name: Handling Errors - Ignore and Pass Up The Stack
video: 
thumb:
github:
# Handling Errors - Ignore and Pass Up The Stack
## 


	return db.Login(u,p)

- Not My Problem(tm)
- Do this deep internally, not in exported Functions/Methods
---
name: Libraries are different
video: 
thumb:
github:
# Libraries are different
## 


- internal error handling vs external library API

Libraries should export known error types for conditions that are static

	sql.ErrNoRows
	io.EOF
	bytes.ErrTooLarge 

- Return underlying error if there's nothing the library can do to fix things

	Network calls fail - you can't fix this, return the error
	ex: log/syslog#Dial()

- Wrap error 
	
	NO.
	return fmt.Errorf("reading from database:", err)
	// can't test for known error without strings.Contains() 
	// Destroyed original context
---
name: Handling Errors - General Advice
video: 
thumb:
github:
# Handling Errors - General Advice
## 


- Only handle errors once - NO

.code errmgmt/includes/errs/handleonce.go /START OMIT/,/END OMIT/

If the caller logs the error too, you've got possibly two log lines with confusingly similar messages.
---
name: Handling Errors - General Advice2
video: 
thumb:
github:
# Handling Errors - General Advice
## 


- Only handle errors once - BETTER

.code errmgmt/includes/errs/handleonce2.go /START OMIT/,/END OMIT/

Since you're not doing anything to resolve the error in your function, return it unchanged.
---
name: Wrapping Errors For Clarity
video: 
thumb:
github:
# Wrapping Errors For Clarity
## 


In early 2016, Dave Cheney released *github.com/pkg/errors*

This package allows you to have the best of all worlds, by letting you wrap errors in contextual information while preserving the original error type.

	// open config file
	_, err := ioutil.ReadAll(r)
		if err != nil {
		return errors.Wrap(err, "read failed")
	}

Callers that print the error will get the wrapped error with your prefix:

	read failed: file not found
---
name: Wrapping Errors For Clarity2
video: 
thumb:
github:
# Wrapping Errors For Clarity
## 


Every function in the stack can wrap the error with their own context:

	return errors.Wrap(err, "opening configuration file")

	// opening configuration file: read failed: file not found
---
name: Wrapping Errors For Clarity3
video: 
thumb:
github:
# Wrapping Errors For Clarity
## 


You can retrieve the original cause of the error by using the errors.Cause() function:

	originalError := errors.Cause(err)
---
name: Wrapping Errors For Clarity - example
video: 
thumb:
github:
# Wrapping Errors For Clarity - example
## 


.code errmgmt/includes/errs/openfile.go /START OMIT/,/END OMIT/
---
name: Wrapping Errors For Clarity4
video: 
thumb:
github:
# Wrapping Errors For Clarity
## 


github.com/pkg/errors is under consideration for inclusion in the standard library.  (which is rare!)

Use it.
---
name: Exercise-errmgnt
video: 
thumb:
github:
# Exercise
## 


Discuss pros and cons of error handling.

- Do you have examples of projects that do it well?
- Do you have examples of confusing projects?
