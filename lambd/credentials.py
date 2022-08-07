import json
import sys
from read_files.global_environment import get_values_from_global_environment as glo_env
from . import key_path_credential
# obtain credentials for mongo

# read file credentials.json and get variables for key
def get_mongo_var(key):
    try:
        with open(glo_env[key_path_credential]) as f:
            data = json.load(f)
            return data[key]
    except Exception as e:
        print(e)
        print('key or file not found')
        sys.exit()

# funcion call mongo_var and return mongo_username and mongo_password
def get_mongo_user_and_password():
    # get mongo var with key MONGO_USERNAME and assign to variable mongo_username
    mongo_username = get_mongo_var('MONGO_USERNAME')

    # get mongo var with key MONGO_PASSWORD and assign to variable mongo_password
    mongo_password = get_mongo_var('MONGO_PASSWORD')

    # if mongo_username and mongo_password are empty, exit
    if mongo_username == '' or mongo_password == '':
        print('mongo_username or mongo_password is empty')
        sys.exit()
    else:
        return mongo_username, mongo_password

