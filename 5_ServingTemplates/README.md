===========================

Ex1) Problem Statement

# ListenAndServe on port 8080 of localhost

For the default route "/"
Have a func called "foo" 
which writes to the response "foo ran"

For the route "/dog/"
Have a func called "dog" 
which parses a template called "dog.gohtml"
and writes to the response "<h1>This is from dog</h1>"
and also shows a picture of a dog when the template is executed.

Use "http.ServeFile"
to serve the file "dog.jpeg"

===========================

Ex2) Problem Statement

# Serve the files in the "starting-files" folder

Use "http.FileServer"

===========================

Ex3) Problem Statement

# Serve the files in the "starting-files" folder

## To get your images to serve, use only this:

``` Go
	fs := http.FileServer(http.Dir("public"))
```

Hint: look to see what type FileServer returns, then think it through.

===========================

Ex4) Problem Statement

# Serve the files in the "starting-files" folder

## To get your images to serve, use:

``` Go
	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler
```

Constraint: you are not allowed to change the route being used for images in the template file

===========================

Ex5) Problem Statement

# Serve the files in the "starting-files" folder

## To get your images to serve, use:
	
``` Go
	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler
```

Constraint: you are not allowed to change the route being used for images in the template file


===========================

Ex6) Problem Statement

# Starting with the code in the "starting-files" folder:

## wire this program up so that it works

ParseGlob in an init function

Use HandleFunc for each of the routes

Combine apply & applyProcess into one func called "apply"

Inside the func "apply", use this code to create the logic to respond differently to a POST method and a GET method

``` go
if req.Method == http.MethodPost {
    		// code here
    		return
    	}
```
===========================