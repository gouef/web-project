package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func AutoTemplate(c *gin.Context, controller interface{}, handlerFunc interface{}) {
	fullFuncName := runtime.FuncForPC(reflect.ValueOf(handlerFunc).Pointer()).Name()
	parts := strings.Split(fullFuncName, ".")
	if len(parts) < 3 {
		log.Printf("[ERROR] Invalid function name: %s", fullFuncName)
		c.String(http.StatusInternalServerError, "Invalid function name format")
		return
	}

	// Sestavení cesty k šabloně
	templatePath, err := buildTemplatePath(c.Param("lang"), parts)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		c.String(http.StatusInternalServerError, "Failed to build template path")
		return
	}

	// Vykreslení šablony
	c.HTML(http.StatusOK, templatePath, gin.H{
		"Title": "Dynamická šablona",
		"H1":    "Automatické šablony",
	})
}

// Pomocná funkce pro sestavení cesty k šabloně
func buildTemplatePath(lang string, parts []string) (string, error) {
	if lang == "" {
		lang = "cs" // Výchozí jazyk
	}

	// Cesta k balíčku a sestavení názvů
	packagePath := strings.Join(parts[:len(parts)-2], "/")
	controllerName := strings.ToLower(parts[len(parts)-2])
	methodName := strings.ToLower(parts[len(parts)-1])

	// Sestavení výsledné cesty
	templateName := fmt.Sprintf("%s/%s/%s.gohtml", packagePath, controllerName, methodName)
	return fmt.Sprintf("%s/%s", lang, templateName), nil
}

func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	languages := []string{"cs", "en"}

	for _, lang := range languages {
		templates, err := filepath.Glob(fmt.Sprintf("templates/%s/**/*.gohtml", lang))
		if err != nil {
			log.Printf("[ERROR] Failed to load templates: %v", err)
			continue
		}
		// Mapování šablon do rendereru
		for _, tmpl := range templates {
			name := strings.TrimPrefix(tmpl, fmt.Sprintf("templates/%s/", lang))
			r.AddFromFiles(name, tmpl)
		}
	}
	return r
}
