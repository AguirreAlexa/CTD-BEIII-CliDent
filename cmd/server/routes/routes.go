package routes

import (
	"database/sql"

	"github.com/AguirreAlexa/clinica/cmd/server/handler"
	"github.com/AguirreAlexa/clinica/internal/dentistas"
	"github.com/AguirreAlexa/clinica/internal/pacientes"
	"github.com/AguirreAlexa/clinica/internal/turnos"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.buildDentistaRoutes()
}

func (r *router) buildDentistaRoutes() {

	repository := dentistas.NewRepository((r.db))
	service := dentistas.NewService(repository)
	handlerDentista := handler.NewDentista(service)

	odon := r.rg.Group("/dentista")
	{
		odon.GET("/:id", handlerDentista.GetByID())
		odon.POST("/", handlerDentista.Save())
		odon.PATCH("/:id", handlerDentista.Patch())
		odon.PUT("/:id", handlerDentista.Put())
		odon.DELETE("/:id", handlerDentista.Delete())
	}
}

func (r *router) buildPacienteRoutes() {

	repository := pacientes.NewRepository((r.db))
	service := pacientes.NewService(repository)
	handlerPaciente := handler.NewPaciente(service)
	
	pac := r.rg.Group("/pacientes")
	
	{
		pac.POST("/", handlerPaciente.Create())
		pac.GET("/:id", handlerPaciente.Get())
		pac.DELETE("/:id", handlerPaciente.Delete())
		pac.PUT("/:id", handlerPaciente.Put())
		pac.PATCH("/:id", handlerPaciente.Patch())
	}
}
	
func (r *router) buildTurnoRoutes() {
	
	repository := turnos.NewRepository((r.db))
	service := turnos.NewService(repository)
	handlerTurno := handler.NewTurno(service)
	
	tur := r.rg.Group("/turnos")
	
	{
		tur.POST("/", handlerTurno.Create())
		tur.GET("/:id", handlerTurno.Get())
		tur.GET("/dni/:id", handlerTurno.GetByPacienteID())
		tur.DELETE("/:id", handlerTurno.Delete())
		tur.PUT("/:id", handlerTurno.Put())
		tur.PATCH("/:id", handlerTurno.Patch())
	}
}

