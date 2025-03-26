package apimodels

type CreateSimilarityAlertRequest struct {
	SimilarityThreshold float32 `json:"similarity_threshold"`
}

type CreateSimilarityAlertResponse struct{}
