import json
import os
import sys
from . import path_global_environment


# read file environment.json and get variables
def get_global_env_file():
    with open(path_global_environment) as f:
        data = json.load(f)
        return data

# read file ../../env.json and return function wich recieve key and get value of data
def get_values_from_global_environment(key):
    # if file with path path_credentials not exist, exit
    if not os.path.isfile(path_global_environment):
        print('file global_environment not exist')
        sys.exit() 
    else:
        data = get_global_env_file()

        return data[key]


