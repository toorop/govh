package govh

import (
	"errors"
	"os"
)

// init
func LoadKeyringFromEnv() (k Keyring, err error) {

	k.AppKey = os.Getenv("OVH_APPLICATION_KEY")
	k.AppSecret = os.Getenv("OVH_APPLICATION_SECRET")
	k.ConsumerKey = os.Getenv("OVH_CONSUMER_KEY")
	if len(k.AppKey) == 0 {
		err = errors.New("OVH_APPLICATION_KEY not found in environnement")
	}
	if len(k.AppSecret) == 0 {
		err = errors.New("OVH_APPLICATION_SECRET not found in environnement")
	}
	/*if len(k.ConsumerKey) == 0 {
		err = errors.New("OVH_CONSUMER_SECRET not found in environnement")
	}*/
	return
}

func checkKeyring(k Keyring) bool {
	return false
}
