# path of file environment values, this file is in same folder of execution script
import os


path_local_environment = 'environment.json'

# find file env.json in current folder, if no exist then search in parent folder
def find_env_json(path="env.json"):
    if not os.path.exists(path):
        find_env_json("../"+path)
    else:
        return path


path_global_environment = find_env_json

