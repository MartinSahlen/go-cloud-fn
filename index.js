const spawnSync = require('child_process').spawnSync

function backgroundHandler(data) {

  result = spawnSync('./function', [], {
    input: JSON.stringify(data),
    stdio: 'pipe',
  });

  if (result.status !== 0) {
     throw new Error(result.stderr.toString())
  }
  return result.stdout;
}

exports.helloPubSub = function(event) {
  return backgroundHandler(event.data)
}

exports.helloBucket = function(event) {
  return backgroundHandler(event.data)
}
