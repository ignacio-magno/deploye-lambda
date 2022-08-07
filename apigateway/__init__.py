import os
import boto3
from lambd import lambda_function_name, lambda_client

from read_files import path_global_environment
from read_files.global_environment import get_values_from_global_environment



api_id = "m0gqzb21qk"
region_name= "us-west-2"
client = boto3.client('apigateway', region_name=region_name)
account_id = boto3.client('sts').get_caller_identity()['Account']

# get arn from function lambda
arn_function = lambda_client.get_function(FunctionName=lambda_function_name)['Configuration']['FunctionArn']

# get client id and assign to account_id
AuthorizationType="COGNITO_USER_POOLS"
AuthorizationId="7w7418"


def create_path():
    path = os.getcwd()
    # count many time exist / in path_global_environment
    count = path_global_environment.count('/')

    # split path_global_environment and get the 3 last element
    path = path.split('/')
    complete_path = path[len(path) - count:-1]
    method = path[len(path) - 1]

    complete_path = '/'.join(complete_path)
    complete_path = get_values_from_global_environment("path_base_api")+"/"+complete_path
    # to complete path split with / and remove the last element and join with /
    return "/"+complete_path, method.upper()

# pwd path
path,method = create_path()
print(path)

