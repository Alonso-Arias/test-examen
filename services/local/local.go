package local

import (
	"context"
	"strings"

	errs "github.com/alonso/test/errors"
	"github.com/alonso/test/log"
	"github.com/alonso/test/services/model"
)

var loggerf = log.LoggerJSON().WithField("package", "services")

type LocalService struct {
}

type GetLocalsByComunaRequest struct {
	Data    model.GetLocales
	Commune string `json:"commune"`
}

type GetLocalsByComunaResponse struct {
	Locals []model.Local `json:"locals"`
}

func (ls LocalService) GetLocalsByCommune(ctx context.Context, in GetLocalsByComunaRequest) (GetLocalsByComunaResponse, error) {

	log := loggerf.WithField("service", "LocalService").WithField("func", "GetUsers")

	log.Debug("Init")
	log.Debugf("Request :%v ", in)
	result := []model.Local{}

	for _, pf := range in.Data {
		// se usa EqualFold para no tener problemas de sensiblidad de cases en los string
		if strings.EqualFold(pf.ComunaNombre, in.Commune) {
			local := model.Local{
				Name:    pf.LocalNombre,
				Address: pf.LocalDireccion,
				Phone:   pf.LocalTelefono,
			}
			result = append(result, local)
		}
	}
	// valida el resultado final si encontro coincidencias
	if len(result) == 0 {
		return GetLocalsByComunaResponse{}, errs.NotFound
	}
	log.Debug("End")
	return GetLocalsByComunaResponse{Locals: result}, nil

}
