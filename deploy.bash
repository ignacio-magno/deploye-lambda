#!/bin/bash

path=/home/ignacio/Desktop/github.com/deploye-lambda/
# activate virtualenv python3
source $path./env/bin/activate

# call deploy.py
python3 $path/deploy.py

# end virtual env
deactivate