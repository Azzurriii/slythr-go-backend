package handlers

import (
	"net/http"
	"strings"

	"github.com/Azzurriii/slythr/internal/application/dto/analysis"
	"github.com/Azzurriii/slythr/internal/application/services"
	"github.com/gin-gonic/gin"
)

type DynamicAnalysisHandler struct {
	service *services.DynamicAnalysisService
}

func NewDynamicAnalysisHandler(service *services.DynamicAnalysisService) *DynamicAnalysisHandler {
	return &DynamicAnalysisHandler{
		service: service,
	}
}

// AnalyzeContract godoc
// @Summary Analyze Solidity contract using AI for security vulnerabilities
// @Description Performs dynamic security analysis on Solidity source code using AI/LLM
// @Tags dynamic-analysis
// @Accept json
// @Produce json
// @Param request body analysis.AnalyzeRequest true "Contract source code"
// @Router /dynamic-analysis [post]
func (h *DynamicAnalysisHandler) AnalyzeContract(c *gin.Context) {
	var req analysis.AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.respondError(c, http.StatusBadRequest, "Invalid request format", err)
		return
	}

	if strings.TrimSpace(req.SourceCode) == "" {
		h.respondError(c, http.StatusBadRequest, "Source code cannot be empty", nil)
		return
	}

	result, err := h.service.AnalyzeContract(c.Request.Context(), req.SourceCode)
	if err != nil {
		h.respondError(c, http.StatusInternalServerError, "Failed to analyze contract", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetDynamicAnalysis godoc
// @Summary Get dynamic analysis result by source hash
// @Description Retrieves dynamic analysis result from cache or database using source hash
// @Tags dynamic-analysis
// @Accept json
// @Produce json
// @Param sourceHash path string true "Source hash"
// @Router /dynamic-analysis/{sourceHash} [get]
func (h *DynamicAnalysisHandler) GetDynamicAnalysis(c *gin.Context) {
	sourceHash := c.Param("sourceHash")

	if strings.TrimSpace(sourceHash) == "" {
		h.respondError(c, http.StatusBadRequest, "Source hash cannot be empty", nil)
		return
	}

	result, err := h.service.GetDynamicAnalysis(c.Request.Context(), sourceHash)
	if err != nil {
		h.respondError(c, http.StatusNotFound, "Dynamic analysis not found", err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *DynamicAnalysisHandler) respondError(c *gin.Context, status int, message string, err error) {
	response := gin.H{
		"success": false,
		"message": message,
	}

	if err != nil {
		response["error"] = err.Error()
	}

	c.JSON(status, response)
}
