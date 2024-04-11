package initializer

import "log/slog"

type Initializer struct {
	Name string
	Func func() error
}

var initList []Initializer

func InitInitializer() {
	for _, v := range initList {
		if err := v.Func(); err != nil {
			slog.Error("initializer error", "name", v.Name, "error", err)
			panic(err)
		}
		slog.Info("initializer finished", "name", v.Name)
	}
}

func Register(name string, f func() error) {
	initList = append(initList, Initializer{
		Name: name,
		Func: f,
	})
}

func init() {
	initList = make([]Initializer, 0)
}
