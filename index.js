const spawn = require('child_process').spawn

//Handle Background events according to spec
function backgroundHandler(data) {
  return new Promise((resolve, reject) => {
    const p = spawn('./function', [], {});
    var lastMessage = "";
    p.stdin.setEncoding('utf-8');
    p.stderr.on('data', (err) => {
      console.error(err.toString());
    })
    p.stdout.on('data', (out) => {
      console.log(out.toString());
        lastMessage = err;
    })
    p.on('close', (code) => {
      if (code !== 0) {
        reject();
      } else {
        console.log("Finished, last message was: " + JSON.parse(lastMessage));
        resolve();
      }
    });
    p.stdin.write(JSON.stringify(data));
    p.stdin.end();
  });
}

exports.helloPubSub = function(event) {
  return backgroundHandler(event.data)
}

exports.helloBucket = function(event) {
  return backgroundHandler(event.data)
}
