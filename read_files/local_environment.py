import json
import os
import sys
from . import path_local_environment

# verifique if file environment.json exist
def file_exists():
    return os.path.isfile(path_local_environment)

# read file environment.json and get variables
def get_env_file():
    with open(path_local_environment) as f:
        data = json.load(f)
        return data

# return function wich recieve key and get value of data
def get_variables_from_local_environment(key):

    if file_exists():

        data = get_env_file()

        return data[key]
    else:
        print('file local environment not exist')
        sys.exit()

