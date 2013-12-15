package govh_test

import (
	"github.com/Toorop/govh/ovh"
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

/*func TestReturnTrue(t *testing.T) {
	if x := returnTrue(); x != true {
		t.Errorf("returnTrue ne retourne pas true mais")
	}
}*/

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type TSuite struct{}

var _ = Suite(&TSuite{})

// OVH_APPLICATION_KEY
func (s *TSuite) TestGetAuthEnvNoAK(c *C) {
	os.Clearenv()
	os.Setenv("OVH_APPLICATION_SECRET", "AS")
	os.Setenv("OVH_CONSUMER_KEY", "CK")
	_, err := ovh.GetAuthEnv()
	c.Assert(err, ErrorMatches, "OVH_APPLICATION_KEY not found in environnement")
}

// OVH_APPLICATION_SECRET
func (s *TSuite) TestGetAuthEnvNoAS(c *C) {
	os.Clearenv()
	os.Setenv("OVH_APPLICATION_KEY", "AK")
	os.Setenv("OVH_CONSUMER_KEY", "CK")
	_, err := ovh.GetAuthEnv()
	c.Assert(err, ErrorMatches, "OVH_APPLICATION_SECRET not found in environnement")
}

// OVH_CONSUMER_KEY
func (s *TSuite) TestGetAuthEnvNoCK(c *C) {
	os.Clearenv()
	os.Setenv("OVH_APPLICATION_SECRET", "AS")
	os.Setenv("OVH_APPLICATION_SECRET", "AS")
	_, err := ovh.GetAuthEnv()
	c.Assert(err, ErrorMatches, "OVH_CONSUMER_SECRET not found in environnement")
}

func (s *TSuite) TestGetAuthEnv(c *C) {
	os.Clearenv()
	os.Setenv("OVH_APPLICATION_KEY", "AK")
	os.Setenv("OVH_APPLICATION_SECRET", "AS")
	os.Setenv("OVH_CONSUMER_KEY", "CK")
	auth, err := ovh.GetAuthEnv()
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, ovh.Auth{AppKey: "AK", AppSecret: "AS", ConsumerKey: "CK"})
}
