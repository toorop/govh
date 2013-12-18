# OVH API CLI

## How to use

### 1 - Download executable
* [Windows](https://github.com/Toorop/govh/blob/master/cli/bin/windows/ovh.exe)
* [Mac](https://github.com/Toorop/govh/blob/master/cli/bin/macos/ovh)
* [Linux](https://github.com/Toorop/govh/blob/master/cli/bin/linux/ovh)

### 2 - Open a terminal and launch it

### 3 - Follow instruction to get & use  auth token
More detailed procedure will come... 

## Response
For now all responses (except error) are raw JSON response as returned by the API.
Other formats will (probably) coming.


## Avalaible commands
We will consider Linux|MacOs version, just replace *./ovh* by *./ovh.exe* if you are using Windows (~~or replace your OS~~) 
  
### IP
#### List IP
	./ovh ip list
Will return all your IP
You can provide a third argument defining the type of IP returned. For exemple, if you only want IP attached to tour dedicated server, run the command :

	./ovh ip list dedicated
	
Available type are :

* cdn
* dedicated
* hosted_ssl
* loadBalancing
* mail
* pcc
* pci
* vpn
* vps
* xdsl

#### IP load balancing
	./ovh ip lb list
Will return your loadBalancing services


	


