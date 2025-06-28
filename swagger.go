package package_http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	slog "github.com/salamsites/package-log"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"io/ioutil"
	"reflect"
	"strings"
)

func InitSwaggerRoute(router chi.Router, subRouter string) {
	router.Route(subRouter, func(r chi.Router) {
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(subRouter+"/swagger/doc.json"),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		).ServeHTTP)
	})
}

func InitSwagger(logger *slog.Logger, args interface{}) {
	swaggerFilePaths := []string{"docs.go", "swagger.json", "swagger.yaml"}

	for _, filePath := range swaggerFilePaths {
		filePath = "./docs/" + filePath
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			logger.Error("Error reading file %s: %v\n", filePath, err)
			continue
		}

		newContent := changeArgs(content, args)

		err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
		if err != nil {
			logger.Error("Error writing file %s: %v\n", filePath, err)
			continue
		}

		logger.Info("Updated Swagger file %s\n", filePath)
	}

	logger.Info("All Swagger files updated successfully!")
}

func changeArgs(content []byte, args interface{}) string {
	newContent := string(content)
	val := reflect.ValueOf(args)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
		newContent = strings.ReplaceAll(newContent, "{{ ."+field.Name+" }}", fmt.Sprintf("%v", fieldValue))
	}
	return newContent
}
