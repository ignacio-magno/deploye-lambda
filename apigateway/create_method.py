from . import api_id, client, method, AuthorizationType, AuthorizationId
import sys

# =========================================================== Create methods ===========================================================
# get resource with id 
def get_resource(id_path):
    print(id_path)
    response = client.get_resource(restApiId=api_id, resourceId=id_path)
    return response

def create_method(id_resource):
    # try obtain method from the resource with id, if no exist print method not exist, else, asign method to variable resource_method
    resource_method = ""
    try:
        resource_method = get_resource(id_resource)['resourceMethods'][method] 
    except KeyError:
        print("method not exist")
    
    # if resource_method is not equal to "" create method in resource
    if resource_method != "":
        print("method exist")
        # consulte with y/n if want to delete method
        if input("Do you want to delete method? (y/n)") == "y":
            # delete method in resource
            client.delete_method(restApiId=api_id, resourceId=id_resource, httpMethod=method)
            print("method deleted")

            # recall create method, but now no exist method
            create_method()
        # consulte with y/n if want to update method
        if input("Do you want to update method? (y/n)") == "y":
            # consulte with Y/n if want to set authorization
            if input("Do you want to set authorization cognito? (y/N)") == "y":
                # update method in resource
                if AuthorizationType == "" and AuthorizationId == "":
                    print("Authorization not set")
                    sys.exit()
                else:
                    client.put_method(restApiId=api_id, resourceId=id_resource, httpMethod=method, authorizationType=AuthorizationType,authorizerId=AuthorizationId)
            else:
                client.put_method(restApiId=api_id, resourceId=id_resource, httpMethod=method)
            print("method updated") 
            
    else:
        #consulte y/N to create method
        create_method = input("Do you want to create method? (y/N)")

        if create_method == "y":
            # print in blue creating method
            print("\033[94mcreating method: " + method + "\033[0m")
            try:
                if AuthorizationType == "" and AuthorizationId == "":
                    print("Authorization not set")
                    response = client.put_method(restApiId=api_id, resourceId=id_resource, httpMethod=method)
                else:
                    response = client.put_method(restApiId=api_id, resourceId=id_resource, httpMethod=method, authorizationType=AuthorizationType, authorizerId=AuthorizationId)
            except Exception as e:
                print(e)
                # print red error to create method
                print("\033[91merror to create method\033[0m")
                print(response)

            print("\n")
        else:
            # print method not created and exit
            print("method not created")
            sys.exit()

