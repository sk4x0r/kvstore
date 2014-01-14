import os
from random import randint
import subprocess
'''
print("Starting server")
subprocess.Popen("go run listener.go")
#os.system("go run listener.go")
'''

INPUT_FILE='input2.txt'
SET_COMMAND="go run client.go set "
with open(INPUT_FILE) as fp:
    for line in fp:
		cmd=SET_COMMAND+line.strip()
		os.system(cmd)

GET_COMMAND="go run client.go get "
with open(INPUT_FILE) as fp:
    for line in fp:
		cmd=GET_COMMAND+line.strip().split()[0]
		os.system(cmd)

DELETE_COMMAND="go run client.go delete "
with open(INPUT_FILE) as fp:
    for line in fp:
		if randint(0,1)==1:
			print("deleting")
			cmd=DELETE_COMMAND+line.strip().split()[0]
			os.system(cmd)

GET_COMMAND="go run client.go get "
with open(INPUT_FILE) as fp:
    for line in fp:
		cmd=GET_COMMAND+line.strip().split()[0]
		os.system(cmd)
