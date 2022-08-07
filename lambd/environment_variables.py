import sys
from lambd.credentials import get_mongo_user_and_password
from read_files.local_environment import get_variables_from_local_environment as loc_env
# from file with path local_environment import all values from key "environment_variables"

key = 'environment_variables'

# read file local_environment.json and get variables for key
def get_var():
    try:
        return loc_env(key)
    except Exception as e:
        print(e)
        print('key or file not found')
        sys.exit()
    
def get_environment_variables():
    obj = get_var()
    print(obj)

    # craete variable json
    json = {"Variables":{}}
    # iterate over array obj and  assign to variable json
    for item in obj:
        # for each key in item
        for key in item:
            # assign to variable json key and value
            json['Variables'][key] = item[key]

   
    user, password = get_mongo_user_and_password()

    json["Variables"]["MONGO_USERNAME"] = user
    json["Variables"]["MONGO_PASSWORD"] = password
    print(json)
    return json