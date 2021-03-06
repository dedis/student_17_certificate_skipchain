package main

import "gopkg.in/urfave/cli.v1"

/*
This holds the cli-commands so the main-file is less cluttered.
*/

var commandAdmin, commandID, commandConfig, commandKeyvalue, commandSSH, commandFollow, commandCert cli.Command

func init() {
	commandAdmin = cli.Command{
		Name:  "admin",
		Usage: "admin options",
		Subcommands: []cli.Command{
			{
				Name:      "link",
				Usage:     "links admin to cothority",
				ArgsUsage: "IP address [PIN]",
				Action:    adminLink,
			},
			{
				Name:      "store",
				Usage:     "stores the authentication data in cothority",
				ArgsUsage: "file or string with auth data IP address",
				Action:    adminStore,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "type,t",
						Usage: `type of authentication it wants to store
							: PoP, etc.`,
					},
				},
			},
			{
				Name:      "add",
				Usage:     "adds public keys in cothority",
				ArgsUsage: "string with keys IP address",
				Action:    adminAdd,
			},
		},
	}
	commandID = cli.Command{
		Name:  "id",
		Usage: "working on the identity",
		Subcommands: []cli.Command{
			{
				Name:      "create",
				Aliases:   []string{"cr"},
				Usage:     "start a new identity",
				ArgsUsage: "group(public.toml) file(token.toml) or string with auth data  [id-name]",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:  "thr,threshold",
						Usage: "the threshold necessary to add a block",
						Value: 2,
					},
					cli.StringFlag{
						Name:  "type,t",
						Usage: "type of client authentication: PoP, PIN",
					},
					cli.StringFlag{
						Name:  "cred,credentials",
						Usage: "auth data : PoP-token file or PIN-string",
					},
				},
				Action: idCreate,
			},
			{
				Name:    "keypair",
				Aliases: []string{"kp"},
				Usage:   "create keypair",
				Action:  idKeyPair,
			},
			{
				Name:      "connect",
				Aliases:   []string{"co"},
				Usage:     "connect to an existing identity",
				ArgsUsage: "group id [id-name]",
				Action:    idConnect,
			},
			{
				Name:    "del",
				Aliases: []string{"rm"},
				Usage:   "delete an identity",
				Action:  idDel,
			},
			{
				Name:    "check",
				Aliases: []string{"ch"},
				Usage:   "check the health of the cothority",
				Action:  idCheck,
			},
			{
				Name:    "qrcode",
				Aliases: []string{"qr"},
				Usage:   "print out the qrcode of the identity-skipchain and a node for contact",
				Action:  idQrcode,
			},
		},
	}
	commandConfig = cli.Command{
		Name:  "config",
		Usage: "updating and voting on config",
		Subcommands: []cli.Command{
			{
				Name:    "propose",
				Aliases: []string{"l"},
				Usage:   "propose the new config",
				Action:  configPropose,
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "fetch the latest config",
				Action:  configUpdate,
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list existing config and proposed",
				Action:  configList,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "p,propose",
						Usage: "will also show proposed config",
					},
					cli.BoolFlag{
						Name:  "d,details",
						Usage: "also show the values of the keys",
					},
				},
			},
			{
				Name:      "vote",
				Aliases:   []string{"v"},
				Usage:     "vote on existing config",
				ArgsUsage: "[yn]",
				Action:    configVote,
			},
		},
	}
	commandKeyvalue = cli.Command{
		Name:    "keyvalue",
		Aliases: []string{"kv"},
		Usage:   "storing and retrieving key/value pairs",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list all values",
				Action:  kvList,
			},
			{
				Name:      "value",
				Aliases:   []string{"v"},
				Usage:     "return the value of a key",
				ArgsUsage: "key",
				Action:    kvValue,
			},
			{
				Name:      "add",
				Aliases:   []string{"a"},
				Usage:     "add a new key/value pair",
				ArgsUsage: "key value",
				Action:    kvAdd,
			},
			{
				Name:      "addWeb",
				Usage:     "add a web-site to a skipchain",
				Aliases:   []string{"a"},
				ArgsUsage: "path/page.html",
				Action:    kvAddWeb,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "inline",
						Usage: "inline all images, css and scripts",
					},
				},
			},
			{
				Name:      "del",
				Aliases:   []string{"rm"},
				Usage:     "delete a value",
				ArgsUsage: "key",
				Action:    kvDel,
			},
		},
	}
	commandSSH = cli.Command{
		Name:  "ssh",
		Usage: "handling your ssh-keys",
		Subcommands: []cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "adds a new entry to the config",
				Action:  sshAdd,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "a,alias",
						Usage: "alias to use for that entry",
					},
					cli.StringFlag{
						Name:  "u,user",
						Usage: "user for that connection",
					},
					cli.StringFlag{
						Name:  "p,port",
						Usage: "port for the connection",
					},
					cli.IntFlag{
						Name:  "sec,security",
						Usage: "how many bits for the key-creation",
						Value: 2048,
					},
				},
			},
			{
				Name:      "del",
				Aliases:   []string{"rm"},
				Usage:     "deletes an entry from the config",
				ArgsUsage: "alias_or_host",
				Action:    sshDel,
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "shows all entries for this device",
				Action:  sshLs,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "a,all",
						Usage: "show entries for all devices",
					},
				},
			},
			{
				Name:    "rotate",
				Aliases: []string{"r"},
				Usage:   "renews all keys - only active once the vote passed",
				Action:  sshRotate,
			},
			{
				Name:    "sync",
				Aliases: []string{"tc"},
				Usage:   "sync config and blockchain - interactive",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "tob,toblockchain",
						Usage: "force copy of config-file to blockchain",
					},
					cli.StringFlag{
						Name:  "toc,toconfig",
						Usage: "force copy of blockchain to config-file",
					},
				},
				Action: sshSync,
			},
		},
	}
	commandFollow = cli.Command{
		Name:    "follow",
		Aliases: []string{"f"},
		Usage:   "follow skipchains",
		Subcommands: []cli.Command{
			{
				Name:      "add",
				Aliases:   []string{"a"},
				Usage:     "add a new skipchain",
				ArgsUsage: "group ID service-name",
				Action:    followAdd,
			},
			{
				Name:      "del",
				Aliases:   []string{"rm"},
				Usage:     "delete a skipchain",
				ArgsUsage: "ID",
				Action:    followDel,
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list all skipchains and keys",
				Action:  followList,
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update all skipchains",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:  "p,poll",
						Value: 0,
						Usage: "poll every n seconds",
					},
				},
				Action: followUpdate,
			},
		},
	}
	commandCert =cli.Command{
		Name: "cert",
		Aliases:[]string{"c"},
		Usage: "Manage Certificates",
		Subcommands: []cli.Command{
			{
				Name:	   "request",
				Aliases:   []string{"q"},
				Usage: 	   "request a certificate to letsencrypt and store it to the skipchain",		
				ArgsUsage: "name of the directory",
				Action:     certRequest,
			},
			{
				Name:	   "add",
				Aliases:   []string{"a"},
				Usage:     "add a key/cert pair", 
				ArgsUsage: "key/cert pair",
				Action:	    certStore,
			},
			{
				Name:	   "verify",
				Aliases:    []string{"v"},
				Usage:	   "verify the certificate against the root certificate",
				ArgsUsage: "key",
				Action:     certVerify,
			},
			{
				Name:	   "renew",
				Aliases:   []string{"u"},
				Usage:     "renew a certificate", 
				ArgsUsage: "the key",
				Action:	    certRenew,
			},
			{
				Name:	   "list",
				Aliases:   []string{"l"},
				Usage:     "list the certificate", 
				ArgsUsage: "",
				Action:	    certList,
			},
			{
				Name:	   "revoke",
				Aliases:   []string{"k"},
				Usage:     "revoke and delete a certificate", 
				ArgsUsage: "the key",
				Action:	    certRevoke,
			},
			{
				Name:	   "retrieve",
				Aliases:   []string{"r"},
				Usage:     "retrieve the certificate of a given key",
				ArgsUsage: "key and optionnaly a file name to write the retrieved cert ",
				Action:	   certRetrieve,				 	 	 	 	},
		},
	}	
}
