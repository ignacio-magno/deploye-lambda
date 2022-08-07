import json
import sys

from apigateway.create_method import get_resource
from . import client, api_id

# =========================================================== Create or update coors of resource ===========================================================
method = "OPTIONS"
# get integration from resource with id
def get_integration(id_path):
    try:
        response = client.get_integration(restApiId=api_id, resourceId=id_path, httpMethod=method)
        return response
    except :
        return ""

# to call this function is need previous developed resource
def create_or_update_cors(id_resource):

     # try obtain method from the resource with id, if no exist print method not exist, else, asign method to variable resource_method
    resource_method = ""
    print(get_resource(id_resource))
    try:
        resource_method = get_resource(id_resource)['resourceMethods'][method] 
    except KeyError:
        print("method not exist")
    

    if not resource_method == "":
        # print blue creating method
        print("\033[94mcreating method: " + method + "\033[0m")
        # create method in resource
        client.put_method(restApiId=api_id, resourceId=id_resource, httpMethod=method, authorizationType="NONE")
        print("method created")
        create_or_update_cors(id_resource)
    else:
        print("method exist")
        # consulte with y/n if want to delete method
        if input("Do you want to delete method? (y/n)") == "y":
            # delete method in resource
            client.delete_method(restApiId=api_id, resourceId=id_resource, httpMethod=method)
            print("method deleted")
            # recall create method, but now no exist method
            create_or_update_cors(id_resource)

    # create integration type mock
    try:
        # set integration cors configuration
        response = client.put_integration(restApiId=api_id, resourceId=id_resource, httpMethod=method, type="MOCK",
            integrationHttpMethod="POST",)
    # if error print error
    except Exception as e:
        # print integration type mock already exist
        print("\033[91mintegration type mock already exist\033[0m")

    allow_origin = ""
    # create integration response
    try:
        # print in blue "set allow origin"
        print("\033[94mset allow origin\033[0m")
        allow_origin = input()

        # add cors configuration


        headers = {
            "allowOrigins":'["Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"]',
            "allowMethods":'["OPTIONS,POST"]',
      #      "method.request.header.Access-Control-Allow-Origin": '["'+allow_origin+'"]',
      #      "method.request.header.Content-Type":'["application/json"]'
        }

        # cors configuration api gateway
        cors_configuration = {
       }
        # set integration response
        response = client.put_integration_response(restApiId=api_id, resourceId=id_resource, httpMethod=method, statusCode="200",
            responseTemplates=cors_configuration)
    # if error print error
    except Exception as e:
        # print integration response already exist
        print("\033[91mintegration response already exist\033[0m")
        print(e)

        # consulte if want update integration
        if input("Do you want to update integration? (y/n)") == "y":
            # delete integration response
            client.delete_integration_response(restApiId=api_id, resourceId=id_resource, httpMethod=method, statusCode="200")

            # update integration response
            response = client.put_integration_response(restApiId=api_id, resourceId=id_resource, httpMethod=method, statusCode="200",
                responseTemplates=cors_configuration)
            print("integration updated")

