package main
import "fmt"

type Command interface{
 runCommand(cmd string)
}
type Execute struct{
	user string
}
func (c *Execute) runCommand(cmd string) string{
	Runtime.getRuntime().exec(cmd)
	fmt.Println("Command is executed "+cmd)
	return cmd
}

