package api

import (
	"encoding/json"
	"net/http"

	"github.com/shinroo/fp/src/webapp/internal/repository"
	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/models"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("search request received")

	var req apimodels.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.Logger.Error("failed to unmarshal search request", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to unmarshal search request",
		}, http.StatusBadRequest, w)
		return
	}

	var filters []repository.DogSearchFilter

	if req.Gender != nil {
		filters = append(filters, repository.DogGenderFilter{
			Gender: *req.Gender,
		})
	}

	if req.LifeStage != nil {
		filters = append(filters, repository.DogLifeStageFilter{
			LifeStage: *req.LifeStage,
		})
	}

	if req.Breed != nil {
		filters = append(filters, repository.DogBreedFilter{
			Breed: *req.Breed,
		})
	}

	dogs, err := s.DogRepo.Search(r.Context(), req.Keywords, filters)
	if err != nil {
		s.Logger.Error("failed to search dogs", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to search dogs",
		}, http.StatusInternalServerError, w)
		return
	}

	var spcaIDs []string
	for _, dog := range dogs {
		spcaIDs = append(spcaIDs, dog.SpcaId)
	}

	spcas, err := s.SPCARepo.GetByIDs(r.Context(), spcaIDs)
	if err != nil {
		s.Logger.Error("failed to get SPCAs by IDs", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get SPCAs by IDs",
		}, http.StatusInternalServerError, w)
		return
	}

	var lookup = make(map[string]*models.SPCA)
	for _, spca := range spcas {
		lookup[spca.ID] = spca
	}

	responses.WriteJSON(apimodels.SearchResponse{
		Results:    dogs,
		SPCALookup: lookup,
	}, http.StatusOK, w)
}
