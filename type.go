package logging

const (
	typeZap     typ = "zap"
	typeLogrus  typ = "logrus"
	typeBuiltin typ = "builtin"
)

type typ string
