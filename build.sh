#!/bin/bash
GOOS=linux go build -o target/function
cp index.js target
