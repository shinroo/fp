package api

import (
	"encoding/json"
	"net/http"

	"github.com/shinroo/fp/src/webapp/pkg/apimodels"
	"github.com/shinroo/fp/src/webapp/pkg/responses"
)

func (s *Server) Profile(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("profile request received")

	switch r.Method {
	case http.MethodGet:
		s.getProfile(w, r)
	case http.MethodPost:
		s.updateProfile(w, r)
	}
}

func (s *Server) getProfile(w http.ResponseWriter, r *http.Request) {
	session, err := s.SessionRepo.GetSessionByToken(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		s.Logger.Error("failed to get session", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get session",
		}, http.StatusInternalServerError, w)
		return
	}

	profile, err := s.ProfileRepo.GetProfileByAccountID(r.Context(), session.AccountID)
	if err != nil {
		s.Logger.Error("failed to get profile", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get profile",
		}, http.StatusInternalServerError, w)
		return
	}

	responses.WriteJSON(apimodels.GetProfileResponse{
		Profile: profile,
	}, http.StatusOK, w)
}

func (s *Server) updateProfile(w http.ResponseWriter, r *http.Request) {
	var req apimodels.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.Logger.Error("failed to unmarshal update profile request", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to unmarshal update profile request",
		}, http.StatusBadRequest, w)
		return
	}

	session, err := s.SessionRepo.GetSessionByToken(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		s.Logger.Error("failed to get session", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get session",
		}, http.StatusInternalServerError, w)
		return
	}

	profile, err := s.ProfileRepo.GetProfileByAccountID(r.Context(), session.AccountID)
	if err != nil {
		s.Logger.Error("failed to get profile", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to get profile",
		}, http.StatusInternalServerError, w)
		return
	}

	profile.AffectionateWithFamily = req.AffectionateWithFamily
	profile.GoodWithYoungChildren = req.GoodWithYoungChildren
	profile.GoodWithOtherDogs = req.GoodWithOtherDogs
	profile.SheddingLevel = req.SheddingLevel
	profile.CoatGroomingFrequency = req.CoatGroomingFrequency
	profile.DroolingLevel = req.DroolingLevel
	profile.CoatLength = req.CoatLength
	profile.OpennessToStrangers = req.OpennessToStrangers
	profile.PlayfulnessLevel = req.PlayfulnessLevel
	profile.WatchdogProtectiveNature = req.WatchdogProtectiveNature
	profile.AdaptabilityLevel = req.AdaptabilityLevel
	profile.TrainabilityLevel = req.TrainabilityLevel
	profile.EnergyLevel = req.EnergyLevel
	profile.BarkingLevel = req.BarkingLevel
	profile.MentalStimulationNeeds = req.MentalStimulationNeeds
	profile.Latitude = req.Latitude
	profile.Longitude = req.Longitude

	err = s.ProfileRepo.UpdateProfile(r.Context(), profile)
	if err != nil {
		s.Logger.Error("failed to update profile", "error:", err)
		responses.WriteJSON(apimodels.ErrorResponse{
			Message: "failed to update profile",
		}, http.StatusInternalServerError, w)
		return
	}

	responses.WriteJSON(apimodels.UpdateProfileRequest{}, http.StatusOK, w)
}
