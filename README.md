# go-cloud-fn

A Go thing designed to test Google cloud functions https://cloud.google.com/functions/docs/quickstart.

Write your code using go and compile it to server-side node required by Google cloud functions.

This project enables you to create full-blown API's with routing support. More to come soon!

Run `npm run build` to compile the project.

then, you will run:

`gcloud alpha functions deploy helloGO --stage-bucket <your_bucket> --trigger-http`

## License

Copyright © 2017 Martin Sahlen

Distributed under the MIT License
