package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// Automatická detekce názvu šablony včetně složky, controllera a metody
func AutoTemplate(c *gin.Context, controller interface{}, handlerFunc interface{}) {
	// Získání celé cesty k funkci (např. "github.com/gouef/web-project/controllers/users.DefaultController.Index")
	fullFuncName := runtime.FuncForPC(reflect.ValueOf(handlerFunc).Pointer()).Name()

	// Extrahujeme balíček a název metody
	parts := strings.Split(fullFuncName, ".")
	if len(parts) < 3 {
		c.String(http.StatusInternalServerError, "Invalid function name: %s", fullFuncName)
		return
	}

	// Cesta k balíčku, např. "controllers/users"
	packagePath := strings.Join(parts[:len(parts)-2], "/")

	// Název controlleru, např. "DefaultController"
	controllerName := parts[len(parts)-2]

	// Název metody, např. "Index"
	methodName := parts[len(parts)-1]

	// Převod na lowercase a sestavení cesty
	templateName := fmt.Sprintf("%s/%s/%s.gohtml", packagePath, strings.ToLower(controllerName), strings.ToLower(methodName))

	// Získání jazyka z URL (např. `/cs/` nebo `/en/`)
	lang := c.Param("lang")
	if lang == "" {
		lang = "cs" // Výchozí jazyk
	}

	// Finální cesta k šabloně
	templatePath := fmt.Sprintf("%s/%s", lang, templateName)

	// Vykreslení šablony
	c.HTML(http.StatusOK, templatePath, gin.H{
		"Title": "Dynamická šablona",
		"H1":    "Automatické šablony",
	})
}
func createRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Načtení všech jazykových verzí
	languages := []string{"cs", "en"}

	for _, lang := range languages {
		templates, _ := filepath.Glob(fmt.Sprintf("templates/%s/**/*.gohtml", lang))
		for _, tmpl := range templates {
			name := strings.TrimPrefix(tmpl, fmt.Sprintf("templates/%s/", lang))
			r.AddFromFiles(name, tmpl)
		}
	}

	return r
}
