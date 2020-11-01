package commandHandler

type CommandHandler interface {
	HandleCommand(command interface{}) (interface{}, error)
}
