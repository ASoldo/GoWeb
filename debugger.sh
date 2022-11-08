#!/bin/bash

dlv debug -l 127.0.0.1:38697 --headless ./main.go --only-same-user=false
