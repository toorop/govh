# OVH API CLI

## How to use

### 1 - Download executable

* Windows : ftp://ftp.toorop.fr/softs/ovh_cli/windows/ovh.exe
* MacOs : ftp://ftp.toorop.fr/softs/ovh_cli/macos/ovh
* Linux : ftp://ftp.toorop.fr/softs/ovh_cli/linux/ovh


### 2 - Run cmd and follow instructions 
Open a terminal, go to the path where ovh binary is and run : 

	./ovh
	
On Linux and MacOs

	./ovh.exe

On windows

And follow the instruction to get and use your consumer key.

### Response
For now all responses (except error) are raw JSON response as returned by the API.
Other formats will (probably) coming.


### Avalaible commands
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


	


