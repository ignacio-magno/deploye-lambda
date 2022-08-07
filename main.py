import sys
from apigateway.create_options import create_or_update_cors
from lambd.deploy import lambda_function_exists, start_lambda_deploy
from apigateway.create_path import create_resources_if_not_exist as deploy_path
from apigateway.create_method import create_method as deploy_method
from apigateway.create_integration import create_integration as deploy_integration

is_in_windows = sys.platform.startswith('win')

# print is in windows
if is_in_windows:
    print('is in windows')
    sys.exit()

def main():
    print("hello")

    # consulte if want deploy lambda function
    deploy = input('Do you want to deploy lambda function? (y/N) ')
    if deploy == 'y':
        start_lambda_deploy()
    else:
        print("okay continue")
    
   
    if lambda_function_exists():
        print('lambda already exists')

        # consulte if want to crete path
        create_path = input('Do you want to create apy? (y/N) ')
        if create_path == 'y':
            id_resource = deploy_path()
            deploy_method(id_resource)
            deploy_integration(id_resource)

    id_resource = deploy_path()
    # consulte if want deploy option method
    set_cors = input('Do you want to deploy method options? (y/N) ')
    if set_cors == 'y':
        create_or_update_cors(id_resource)
    

main()