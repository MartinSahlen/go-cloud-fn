package main

import (
	"bytes"
	"html/template"
)

type config struct {
	Production     bool
	TargetDir      string
	ExecutableName string
	TriggerHTTP    bool
	StageBucket    string
	FunctionName   string
}

func parseTemplate(templateString string, data interface{}) (string, error) {
	tmpl, err := template.New("name").Parse(templateString)
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	err = tmpl.Execute(&b, data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

var buildSh = `GOOS=linux go build -o {{.TargetDir}}/{{.ExecutableName}}`
var deploySh = `gcloud beta functions deploy {{.FunctionName}} {{if .TriggerHTTP}}--trigger-http{{end}} --stage-bucket {{.StageBucket}} --local-path {{.TargetDir}}`
var httpJs = `const spawnSync = require('child_process').spawnSync;

exports.{{.FunctionName}} = function(req, res) {
  var requestBody;

  switch (req.get('content-type')) {
    case 'application/json':
      requestBody = JSON.stringify(req.body);
      break;
    case 'application/x-www-form-urlencoded':
      //The body parser for cloud functions does this, so just play along
      req.set('content-type', 'application/json')
      requestBody = JSON.stringify(req.body);
      break;
    case 'application/octet-stream':
      requestBody = req.body;
      break;
    case 'text/plain':
      requestBody = req.body;
      break;
  }

  var fullUrl = req.protocol + '://' + req.get('host') + req.originalUrl;

  var httpRequest = {
    'body': requestBody,
    'headers': req.headers,
    'method': req.method,
    'remote_addr': req.ip,
    'url': fullUrl
  };

  result = spawnSync('./function', [], {
    input: JSON.stringify(httpRequest),
    stdio: 'pipe',
  });

  if (result.status !== 0) {
     console.log(result.stderr.toString());
     res.status(500);
     return;
  }

  data = JSON.parse(result.stdout);
  res.status(data.status_code);
  res.set(data.headers)
  res.send(data.body);
};`
