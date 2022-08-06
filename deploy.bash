#!/bin/bash

# activate virtualenv python3
source ./env/bin/activate

# call deploy.py
python3 deploy.py

# end virtual env
deactivate