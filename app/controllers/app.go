package controllers
import (
	"github.com/revel/revel"
)
type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	text := "Hello Wordl!"
	return c.Render(text)
}

func (c App) Hoge() revel.Result {
	type Result struct {
		Message string `json:"message"`
	}
	result := Result{
		Message: "hogehoge",
	}
	return c.RenderJSON(result)
}

func (c App) Fuga(number int) revel.Result {
	type Result struct {
		Message string `json:"message"`
		Number  int    `json:"number"`
	}
	number = number * 10
	result := Result{
		Message: "url/:引数名とすると引数を取れるようになります",
		Number:  number,
	}
	return c.RenderJSON(result)
}
