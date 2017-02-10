cd $GOPATH/src/github.com/gopherjs/gopherjs/node-syscall/
npm install --global node-gyp
node-gyp rebuild
mkdir -p ~/.node_libraries/
cp build/Release/syscall.node ~/.node_libraries/syscall.node
