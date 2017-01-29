package main

import (
	"fmt"
	"testing"
)

func TestValidParseConfig(t *testing.T) {
	json := `{
  "config": [
    {
      "name": "vagrant",
      "hosts": [
        {
          "hostname": "127.0.0.1",
          "port": 22,
          "username": "vagrant",
          "password": "vagrant-pass"
        }
      ],
      "commands": [
	    {
			"action": "ssh",
			"content": "ps aux"
		},
	    {
			"action": "scp",
			"content": "/var/log/messages"
		}
      ]
    }
  ]
}`

	config, err := ParseConfig([]byte(json))

	if err != nil {
		t.Fatal("Parse Failed", err)
	}

	fmt.Println(config)

	if len(config.Config) != 1 {
		t.Fatal("Config length was not 1")
	}

	entry := config.Config[0]

	if entry.Name != "vagrant" {
		t.Fatal("Entry.Name was not vagrant")
	}

	host := entry.Hosts[0]

	if host.Hostname != "127.0.0.1" {
		t.Fatal("Host.Hostname was not 127.0.0.1")
	}
	if host.Username != "vagrant" {
		t.Fatal("Host.Username was not vagrant")
	}
	if host.Password != "vagrant-pass" {
		t.Fatal("Host.Password was not vagrant-pass")
	}
	if host.KeyFile != "" {
		t.Fatal("Host.KeyFile was not empty")
	}

}

func TestEmptyHostsParseConfig(t *testing.T) {
	json := `{
  "config": [
    {
      "name": "vagrant",
      "commands": [
	    {
			"action": "ssh",
			"content": "ps aux"
		},
	    {
			"action": "scp",
			"content": "/var/log/messages"
		}
      ]
    }
  ]
}`

	config, err := ParseConfig([]byte(json))

	if err != nil {
		t.Fatal("Parse Failed", err)
	}

	fmt.Println(config)

	if len(config.Config) != 1 {
		t.Fatal("Config length was not 1")
	}

	entry := config.Config[0]

	if entry.Name != "vagrant" {
		t.Fatal("Entry.Name was not vagrant")
	}

	if len(entry.Hosts) != 0 {
		t.Fatal("host was not empty")
	}
}
