# Simple Gin web service

A basic RESTful web service API with Go and using the Gin framework

### API Endpoints

It implements the folowing endpoints:

/album

* GET - get a list of all albums in JSON format
* POST - add a new album from the request data sent in JSON format
 
/albums/:id

* GET - get an album by its ID in JSON format
