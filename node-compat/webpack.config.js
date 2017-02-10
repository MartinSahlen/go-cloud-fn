var path = require('path');

module.exports = {
  entry: './node-compat/polyfill.js',
  target: 'node',
  output: {
    path: path.join(__dirname),
    filename: 'polyfill.inc.js'
  }
}
