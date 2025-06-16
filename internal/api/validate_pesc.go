package api

import (
	"app/internal/utils"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ValidatePESCXML validates uploaded XML against a PESC schema
// @Summary Validate XML Application using PESC XSD
// @Tags Applications
// @Accept xml
// @Produce json
// @Param file formData file true "XML File"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /applications/validate [post]
func ValidatePESCXML(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "❌ Missing XML file"})
		return
	}

	tempFile, err := ioutil.TempFile("", "upload-*.xml")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Could not create temp file"})
		return
	}
	defer os.Remove(tempFile.Name())

	if err := c.SaveUploadedFile(file, tempFile.Name()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "❌ Failed to save XML"})
		return
	}

	xsdPath := "/app/internal/xmlschema/application.xsd"
	err = utils.ValidateXMLWithXSD(tempFile.Name(), xsdPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "✅ XML is valid according to PESC schema"})
}
