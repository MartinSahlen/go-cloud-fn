# go-cloud-fn

A Go thing designed to test Google cloud functions https://cloud.google.com/functions/docs/quickstart.

Write your code using go and compile it to a express.js handler function required by Google cloud functions.

This project enables you to create full-blown API's with routing et etc. It uses gopherjs and some polyfills for `net/http` + syscalls, see https://github.com/gopherjs/gopherjs/issues/518 for more information.

Run `npm run live` to set up a fully auto-reloading dev server!

Run `npm run build` to compile the project.

then, you will run:

`gcloud alpha functions deploy helloGO --stage-bucket <your_bucket> --trigger-http`

## License

Copyright © 2017 Martin Sahlen

Distributed under the MIT License
