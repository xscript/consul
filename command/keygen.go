package command

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// KeygenCommand is a Command implementation that generates an encryption
// key for use in `consul agent`.
type KeygenCommand struct {
	BaseCommand
}

func (c *KeygenCommand) Run(args []string) int {
	c.InitFlagSet()
	if err := c.FlagSet.Parse(args); err != nil {
		return 1
	}

	key := make([]byte, 16)
	n, err := rand.Reader.Read(key)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error reading random data: %s", err))
		return 1
	}
	if n != 16 {
		c.UI.Error(fmt.Sprintf("Couldn't read enough entropy. Generate more entropy!"))
		return 1
	}

	c.UI.Output(base64.StdEncoding.EncodeToString(key))
	return 0
}

func (c *KeygenCommand) Help() string {
	c.InitFlagSet()
	return c.HelpCommand(`
Usage: consul keygen

  Generates a new encryption key that can be used to configure the
  agent to encrypt traffic. The output of this command is already
  in the proper format that the agent expects.

`)
}

func (c *KeygenCommand) Synopsis() string {
	return "Generates a new encryption key"
}
