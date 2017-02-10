var bodyParser = require('body-parser')
var express = require('express')
var handlers = require('../index')

var app = express()
app.use(bodyParser.urlencoded({ extended: false }))
app.use(bodyParser.json())

app.use(function(req, res, next){
  console.log(Date() + " " + req.method.toUpperCase() + " " + req.path + " ");
  next();
});

console.log("Starting to mount endpoints and handler functions...")

Object.keys(handlers).forEach(function(key) {
  app.use('/' + key, handlers[key]);
  console.log("Added handler function for " + key);
});

var port = 8080
console.log("Listening on 0.0.0.0:" + port);
app.listen(port)
