from . import client, api_id, path

# in each resources find if exist path with name recieved if exist return id else return False
def exist_path_in_resource(resource, path):
    for item in resource:
        if item['path'] == path:
            return item['id']
    return False

# recieve a path and if not exist path in resources create resource and return id
def create_path(path, id_path, resources):
    exist_path = exist_path_in_resource(resources, path)
    if not exist_path:
        # split path with "/" and get the las element and assign to path_to_create
        path_to_create = path.split("/").pop()
        # print creating path
        print("creating path: " + path_to_create)
        # create path in resources
        response = client.create_resource(restApiId=api_id, parentId=id_path, pathPart=path_to_create)
        # return id of path
        return response['id']
    else:
        return exist_path

# api gateway list resources with region us-west-2 and rest api id 
def list_resources(rest_api_id):
    response = client.get_resources(restApiId=rest_api_id)
    return response['items']


# create resource if not exist and return id of resource
def create_resources_if_not_exist():

    resources = list_resources(api_id)

    # split path with "/" and get list of paths
    path_list = path.split("/")

    path_base =  "/"

    id_path_base = exist_path_in_resource(resources, path_base)

    # create recursive path if not exist, the limit es the entire path
    for item in path_list:
        # if is first element then obtain id of path /
        if path_list.index(item) == 0:
            id_path_base = create_path(path_base, id_path_base, resources)

        if path_base == "/":
            path_base = path_base+ item
        else:
            path_base = path_base + "/" + item

        id_path_base = create_path(path_base, id_path_base, resources)
        # print("\n")

    return id_path_base

