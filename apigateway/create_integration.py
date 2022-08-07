import uuid
from . import client, api_id, method, account_id, arn_function, path
import boto3

# =========================================================== Create integration ===========================================================
# get integration from resource with id
def get_integration(id_path):
    try:
        response = client.get_integration(restApiId=api_id, resourceId=id_path, httpMethod=method)
        return response
    except :
        return ""

# create integration for resource with id
def create_integration(id_resource):
    # get integration with id_path_base
    integration = get_integration(id_resource)

    # if integration not exist create integration in resource
    if integration == "":
        print("creating integration")

        # create integration in resource with integration type aws proxy and lambda function
        response = client.put_integration(restApiId=api_id, resourceId=id_resource, httpMethod=method, type="AWS_PROXY", 
            integrationHttpMethod="POST", 
            uri = "arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/"+arn_function+"/invocations",
        )
        print(response)
        print("\n")
    else:
        print("integration exist")
        # consulte with y/N if want update integration
        if input("Do you want to update integration? (y/N)") == "y":
            # update integration in resource with integration type aws proxy and lambda function
            response = client.update_integration(restApiId=api_id, resourceId=id_resource, httpMethod=method, type="AWS_PROXY",
                integrationHttpMethod="POST",
                uri = "arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/"+arn_function+"/invocations",
            )

            print(response)
            print("\n")
        else:
            print("integration not updated")
            print("\n")

    # =========================================================== put policy statement to lambda function ===========================================================
    # create arn for method request
    arn = "arn:aws:execute-api:us-west-2:"+account_id+":"+api_id+"/*/"+method+ path

    # put policy statement to lambda function type aws, service api gatewat, with lambda invoke
    def put_policy_statement(arn):
        client = boto3.client('lambda')
        response = client.add_permission(
            FunctionName=arn_function,
            StatementId=uuid.uuid4().hex,
            Action="lambda:InvokeFunction",
            Principal="apigateway.amazonaws.com",
            SourceArn=arn
        )
        return response

    put_policy_statement(arn)