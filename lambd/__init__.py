import boto3
from read_files.local_environment import get_variables_from_local_environment as loc_env
from read_files.global_environment import get_values_from_global_environment as glo_env

# init lambda client with region us-west-2
lambda_client = boto3.client('lambda', region_name='us-west-2')

# build files names
handler_name="main.go"
handler_zip="code.zip"
hanlder_binary_name="main"

key_path_credential = 'path_credentials_mongo'


# get env with key lambda-function-name and assign to variable lambda_function_name
lambda_function_name = loc_env('lambda-function-name')
role_name = lambda_function_name +"_role"
path_credentials = glo_env(key_path_credential)

name_policy_logs = "put-logs"



