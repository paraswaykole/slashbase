package controllers

type ConsoleController struct{}

func (ConsoleController) RunCommand(dbConnectionID, cmdString string) string {
	if cmdString == "ping" {
		return "pong"
	}
	return "unknown command"
}
