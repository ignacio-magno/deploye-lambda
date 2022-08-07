# documentation
this code work with folders, need to be executed in the folder where exist the code to deploy.
once the code is deployed, next go to deploy lambda function, the lambda function path is calculated basade in the ubication of the path, the path base is the folder where exist the file env.json.
with the path base, next calculate the path of the api basade in the directories from base to folder where is deployed the code.
ej: 
    execute in /code/path1/path2/{method}
    in folder code exist the file env.json
    then the path base is /code/path1/path2
    the method is the last name of the folder, can see get, post, update, etc.

this code need files, first, in the folder where is the code, need file environment.json with the next struct
```json
{
    "lambda-function-name": "string",
    "environment_variables": [
        "key": "value",
        "key2": "value2",
    ]
}
```
other file need is env.json, ubicated in the base folder, this file is the next
```
    path_base_api is the base path for the api resource, in the developed sum path_base_api+path_generated_by_folders
```
```json
{
        "path_credentials_mongo": "path",
        "path_base_api": "path",
}
```
