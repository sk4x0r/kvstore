import os
from random import randint

DEBUG=True
def printf(s):
	if DEBUG:
		print s
'''
print("Starting server")
subprocess.Popen("go run listener.go")
#os.system("go run listener.go")
'''

INPUT_FILE='input2.txt' #file from which demo data of key value pairs is loaded

#following lines store some data into key value store
SET_COMMAND="go run client.go set "
with open(INPUT_FILE) as fp:
    for line in fp:
		cmd=SET_COMMAND+line.strip()
		os.system(cmd)

#folliwing code try to fetch the data stored by the lines above
GET_COMMAND="go run client.go get "
with open(INPUT_FILE) as fp:
    for line in fp:
		cmd=GET_COMMAND+line.strip().split()[0]
		os.system(cmd)

#these lines delete some randomly selected keys from key value store
DELETE_COMMAND="go run client.go delete "
with open(INPUT_FILE) as fp:
    for line in fp:
		if randint(0,1)==1:
			printf("Deleting the key")
			cmd=DELETE_COMMAND+line.strip().split()[0]
			os.system(cmd)

#all the keys are probed, out of which some are deleted in previous step
GET_COMMAND="go run client.go get "
with open(INPUT_FILE) as fp:
    for line in fp:
		cmd=GET_COMMAND+line.strip().split()[0]
		os.system(cmd)
