import json
import os
import sys
import boto3
from . import lambda_client, handler_name, handler_zip, hanlder_binary_name, lambda_function_name, role_name, name_policy_logs

# find if exist lambda function with name lambda_function_name
def lambda_function_exists():
    try:
        lambda_client.get_function(FunctionName=lambda_function_name)
        return True
    except Exception as e:
        return False

# this functino list all policy in a role and return a list of policy names
def list_role_policy(role_name):
    client = boto3.client('iam')
    response = client.list_role_policies(RoleName=role_name)
    return response['PolicyNames']

# recieve array of policy names and for each policy name delete policy from role
def delete_role_policy(policy_names):
    client = boto3.client('iam')
    for policy_name in policy_names:
        client.delete_role_policy(RoleName=role_name, PolicyName=policy_name)

# recieve name role and delete role
def delete_role(role_name):
    client = boto3.client('iam')
    client.delete_role(RoleName=role_name)
 
# rebuild main.go and zip code with name code.zip
def rebuild_code():
    # go build main.go
    os.system('go build '+ handler_name)
    # to zip main
    os.system('zip -r '+ handler_zip +' '+ hanlder_binary_name)
    
    # if main.zip not exist, exit
    if not os.path.exists(handler_zip):
        print(handler_zip + ' not exist')
        sys.exit()

# delete the zip file with name code.zip
def delete_code():
    os.remove(handler_zip)
    os.remove(hanlder_binary_name)


# recieve path file and find if file exist
def file_exists(path):
    return os.path.isfile(path)

# delete cloud watch log group with name log_group_name
def delete_log_group(log_group_name):
    client = boto3.client('logs')
    client.delete_log_group(logGroupName=log_group_name)

def start_lambda_deploy():
    # print lambda_function_name
    print(lambda_function_name)

    # if lambda exist print lambda already exists, else print lambda not exists and create lambda function
    if lambda_function_exists():
        # print with green color "lambda already exists"
        print('\033[92m' + 'lambda already exists' + '\033[0m')
        # consulte with y/n if want to update lambda function
        update = input('Do you want to update lambda function? (y/N) ')
        if update == 'y':
            # rebuild code
            rebuild_code()
            # update lambda function with new code from code.zip
            lambda_client.update_function_code(FunctionName=lambda_function_name, ZipFile=open(handler_zip, 'rb').read())

        # consulte with y/N if want add policy to role
        add_policy = input('Do you want to add policy to role? (y/N) ')
        if add_policy == 'y':
            # consulte name policy
            name_policy = input('Enter name policy: ')
            # consulte "the name of policy ubicate in forlder roles"
            name_policy_file = input('Enter name policy file: ')

            #  create path policy with name "roles"+name_policy_file+".json"
            path_policy = "roles/"+name_policy_file+".json"

            # if file with path path_policy not exist, exit
            if not file_exists(path_policy):
                print('file with path' +path_policy + ' not exist')
                sys.exit()

            # get all policy names in role
            policy_names = list_role_policy(role_name)

            # if polici_names contain name_policy, print with green color "policy already exists"
            if name_policy in policy_names:
                print('\033[92m' + 'policy already exists' + '\033[0m')

                # get from lambda function policy with name name_policy
                response = lambda_client.get_policy(PolicyName=name_policy)

                # print response
                print(response)

                # consulte if want to update policy
                update_policy = input('Do you want to update policy? (y/N) ')
                if update_policy == 'y':
                    # update policy with new policy from path_policy
                    lambda_client.update_policy(PolicyName=name_policy, PolicyDocument=open(path_policy, 'r').read())

        # consulte if want to delete lambda function, policies and roles
        delete = input('Do you want to delete lambda function, policies and roles? (y/N) ')
        if delete == 'y':
            # get policies names in role
            policy_names = list_role_policy(role_name)

            # for each policy name delete policy from role
            delete_role_policy(policy_names)

            # delete role
            delete_role(role_name)

            # delete log group
            delete_log_group(lambda_function_name)

            # delete lambda function
            lambda_client.delete_function(FunctionName=lambda_function_name)
    else:
        # print with red color "lambda not exists"
        print('\033[91m' + 'lambda not exists' + '\033[0m')

        # consulte with y/n if want to create lambda function with defaults values
        create = input('Do you want to create lambda function with defaults values? (y/N) ')
        if create == 'y':
            path_trust_policy = "/roles/trust-policy.json"
            path_put_logs_policy = "/roles/put-logs.json"

            # build code
            rebuild_code()

            # create a new iam role with basic lambda permissions and assign to variable output_role
            output_role = os.popen('aws iam create-role --role-name ' + role_name + ' --assume-role-policy-document file:/'+ path_trust_policy).read()

            # from json output_role, get Role.Arn and assign to variable role_arn
            role_arn = json.loads(output_role)['Role']['Arn']

            # assign role_arn to file json environment.json with key role_arn_policy
            with open('environment.json', 'r') as f:
                data = json.load(f)
                data['role_arn_policy'] = role_arn
                with open('environment.json', 'w') as f:
                    json.dump(data, f)

            # put policy to role with name lambda_function_name policy name name_policy_logs
            os.system('aws iam put-role-policy --role-name ' + role_name + ' --policy-name ' + name_policy_logs+' --policy-document file:/'+ path_put_logs_policy)

            print('Deploying lambda function ' + lambda_function_name)
            # deploy lambda function with name lambda_function_name, handler main, zip file main.zip, runtime golang1.x, memory 128MB, 
            # timeout 30 seconds, with role role_name
            os.system('aws lambda create-function --function-name ' + lambda_function_name + ' --runtime go1.x --memory-size 128 --timeout 30 --role ' + role_arn + ' --handler main --zip-file fileb://'+ handler_zip)


    #elimina el archivo main.zip
    delete_code()

    # print ended lambda function
    print('\033[92m' + 'ended lambda function' + '\033[0m')