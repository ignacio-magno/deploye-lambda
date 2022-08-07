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
def get_variables_from_local_environment():

    if file_exists():

        data = get_env_file()

        def get_val_of_data (key):
            return data[key]

        return get_val_of_data
    else:
        print('file local environment not exist')
        sys.exit()

