1. Go Dependencies Used in the project
     github.com/go-chi/chi v1.5.4 // indirect
	 github.com/jackc/pgpassfile v1.0.0 // indirect
	 github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	 github.com/jackc/pgx/v5 v5.3.1 // indirect
	 golang.org/x/crypto v0.6.0 // indirect
	 golang.org/x/text v0.7.0 // indirect

2.  To install the dependencies we have to first enable go modules.To enable the go modules we have to use syntax go 
    mod init.
    example: go mod init github.com/navneetshukl/elitmus-test-extension

3. To run the project you have to first install all these dependencies.To install any dependency use syntax:
         go get +(github repository of the dependency)
         example: ge get github.com/go-chi/chi v1.5.4

4.  There are various routes present in the project. 
    a.) "/" (Method=Get) Get the form to enter all the data of the user.
    b.) "/" (Method=Post) Post the user detail to the server.
    c.) "/save" (Method=Get) To show the WebCam.
    d.) "/save" (Method=Post) To Save the image to the DataBase.
    e.) "/display" (Method=Get) To display the image.
    f.) "/getData" (Method=Get) To Get all the data of the user.

5.  Folders:-
    a.) cmd-> This folder contains the main.go file.This is the entry point of the project.
    b.) database-> this folder contains the database.go file.This file contains all the code related to 
        connecting,inserting and getting  the data from the database.
    c.) render-> This  Folder contains the render.go file.This file contains the code related to rendering the html 
        file.
    d.) routes-> This folder contains routes.go file.This files contains all the functions which performs different 
        functions on different routes.
    e.) form.html-> This is HTML code for rendering the form so that user can input his data.
    f.) index.html-> This is HTML file which contains the code to display the webcam and sending the image to the 
        server via AJAX.
    g.) manifest.json-> This is the basic code to create google chrome extension.
    h.) popup.html-> This is HTML file which is default HTML which open when a extension loads. 
       