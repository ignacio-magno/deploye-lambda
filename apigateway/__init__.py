import boto3

api_id = "m0gqzb21qk"
region_name= "us-west-2"
client = boto3.client('apigateway', region_name=region_name)
method = "GET" # this string must always be in uppercase
account_id = boto3.client('sts').get_caller_identity()['Account']
function_name = "arn:aws:lambda:us-west-2:378009899806:function:contabilidad-client-book-remuneration"

# this is path for create method and all other thins
path = "/contabilidad/remuneration/book"

# get client id and assign to account_id
AuthorizationType="COGNITO_USER_POOLS"
AuthorizationId="7w7418"

