package application

type Application struct {
	Name string
	Mode string
}

func NewApplication(name string, mode string) *Application {

	return &Application{
		Name: name,
		Mode: mode,
	}

}


