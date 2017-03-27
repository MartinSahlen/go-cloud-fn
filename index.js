const spawn = require('child_process').spawn

//Handle Background events according to spec
function backgroundHandler(data) {
  return new Promise((resolve, reject) => {
    const p = spawn('./function', [], {});
    p.stdin.setEncoding('utf-8');
    p.stdin.write(JSON.stringify(data));
    p.stdin.end();
    p.stderr.on('data', (err) => {
      console.error(err.toString());
    })
    p.stdout.on('data', (out) => {
      console.log(out.toString());
    })
    p.on('close', (code) => {
      if (code !== 0) {
        reject();
      } else {
        resolve();
      }
    });
  });
}

exports.helloPubSub = function(event) {
  return backgroundHandler(event.data)
}

exports.helloBucket = function(event) {
  return backgroundHandler(event.data)
}
