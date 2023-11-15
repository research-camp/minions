package router

type Metrics interface {
	Pull() map[string]interface{}
	Hit()
	Miss()
	Up()
	Down()
}
