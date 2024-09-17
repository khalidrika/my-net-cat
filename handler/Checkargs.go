package natc

func Checkargs(args []string) string {
	var port string
	switch len(args) {
	case 2:
		port = args[1]
	case 1:
		port = "8989"
	default:
		return "error"
	}
	return port
}
