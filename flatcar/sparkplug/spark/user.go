package spark

import (
	"github.com/kinvolk/ignition/config/v2_2/types"
	"io/ioutil"
	"strings"
)

func NewUser(name string, groups []string, keyfile string) (types.PasswdUser, error) {
	keys := []types.SSHAuthorizedKey{}
	if keyfile != "" {
		data, err := ioutil.ReadFile(keyfile)
		if err != nil {
			return types.PasswdUser{}, err
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if len(line) < 10 || line[0] == '#' {
				continue
			}
			keys = append(keys, types.SSHAuthorizedKey(line))
		}
	}
	userGroups := []types.Group{}
	for _, group := range groups {
		userGroups = append(userGroups, types.Group(group))
	}
	return types.PasswdUser{
		Name:              name,
		Groups:            userGroups,
		SSHAuthorizedKeys: keys,
	}, nil
}

func NewUserOpt(name string, groups []string, keyfile string) ConfigOpt {
	return ConfigOpt(func(config *types.Config) error {
		user, err := NewUser(name, groups, keyfile)
		if err != nil {
			return err
		}
		users := config.Passwd.Users
		config.Passwd.Users = append(users, user)
		return nil
	})
}
