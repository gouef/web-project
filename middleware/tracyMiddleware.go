package middleware

import (
	"bytes"
	"fmt"
	"github.com/gouef/router"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type TracyData struct {
	Latency      string
	CurrentRoute string
	Routes       []TracyRoute
}

type TracyRoute struct {
	Actual  bool
	Name    string
	Pattern string
	Method  string
}

func TracyMiddleware(r *router.Router) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Měření doby trvání požadavku
		startTime := time.Now()

		// Zachycení odpovědi do bufferu pomocí custom responseWriter
		responseBuffer := &bytes.Buffer{}
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			buffer:         responseBuffer,
		}

		// Pokračování k dalšímu middleware nebo handleru
		c.Next()

		// Měření latence
		latency := time.Since(startTime)

		var formattedLatency string
		switch {
		case latency > time.Second:
			formattedLatency = fmt.Sprintf("%.2f s", float64(latency)/float64(time.Second))
		case latency > time.Millisecond:
			formattedLatency = fmt.Sprintf("%.2f ms", float64(latency)/float64(time.Millisecond))
		case latency > time.Microsecond:
			formattedLatency = fmt.Sprintf("%.2f µs", float64(latency)/float64(time.Microsecond))
		default:
			formattedLatency = fmt.Sprintf("%.2f ns", float64(latency)/float64(time.Nanosecond))
		}

		// Logování latence
		log.Printf("Time: %s", formattedLatency)

		// Získání Content-Type odpovědi
		contentType := writer.Header().Get("Content-Type")

		// Pokud je odpověď HTML, přidáme diagnostický panel
		if writer.Status() == http.StatusOK && contentType == "text/html; charset=utf-8" {
			// Získání aktuální routy
			currentRoute := c.FullPath()

			// Získání seznamu všech rout
			var routes []TracyRoute
			for _, route := range r.GetRoutes() {
				routes = append(routes, TracyRoute{
					Actual:  currentRoute == route.GetPattern(),
					Name:    route.GetName(),
					Pattern: route.GetPattern(),
					Method:  route.GetMethod().String(),
				})
			}

			// Příprava dat pro Tracy panel
			tracyData := TracyData{
				Latency:      formattedLatency,
				CurrentRoute: currentRoute,
				Routes:       routes,
			}

			// Generování Tracy panelu
			tracyPanelHTML, err := generateTracyPanelHTML(tracyData)

			if err != nil {
				log.Println("Error generating Tracy panel HTML:", err)
				return
			}
			// Přidáme diagnostický HTML panel na konec odpovědi
			writer.buffer.WriteString(tracyPanelHTML)
		}

		// Připojíme všechny data zpět do původního odpovědního writeru
		_, err := c.Writer.Write(responseBuffer.Bytes())
		if err != nil {
			log.Println("Error writing response:", err)
		}

		// Logování status kódu
		status := c.Writer.Status()
		log.Printf("Status: %d", status)
	}
}

func generateTracyPanelHTML(data TracyData) (string, error) {
	// Načtení šablony Tracy panelu
	tpl, err := template.New("tracyPanel").ParseFiles("middleware/tracy_panel.gohtml", "middleware/tracy_panel_routes.gohtml")
	if err != nil {
		return "", err
	}

	// Vytvoření strings.Builder pro efektivní generování textu
	var builder strings.Builder

	// Vytvoření výsledného HTML
	err = tpl.Execute(&builder, data)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

// responseWriter zachycuje odpověď do bufferu
type responseWriter struct {
	gin.ResponseWriter
	buffer *bytes.Buffer
}

func (w *responseWriter) Write(data []byte) (int, error) {
	// Zapište data do bufferu a vraťte výsledek
	return w.buffer.Write(data)
}

func (w *responseWriter) WriteHeaderNow() {
	// Není potřeba, můžeme to ponechat prázdné
	if !w.Written() {
		w.ResponseWriter.WriteHeaderNow()
	}
}

func (w *responseWriter) WriteHeader(statusCode int) {
	// Tímto způsobem můžeme zapsat správný status kód
	w.ResponseWriter.WriteHeader(statusCode)
}
