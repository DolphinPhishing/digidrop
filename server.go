package main

import (
	"context"
	"digidrop/ent"
	"digidrop/ent/filemiddleware"
	"digidrop/ent/user"
	"digidrop/graphql/auth"
	"digidrop/graphql/graph"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

// tempURLHandler Checks ENT to verify that the url results in a file
func tempURLHandler(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlID := ctx.Param("url_id")
		fmt.Printf("FILE URL ID: %s\n", urlID)
		fileInfo, err := client.FileMiddleware.Query().Where(
			filemiddleware.URLIDEQ(urlID),
		).Only(ctx)
		if err != nil {
			fmt.Println("FILE ERR")
			ctx.AbortWithStatus(404)
			return
		}
		ctx.File(fileInfo.FilePath)
		_, err = fileInfo.Update().SetAccessed(true).Save(ctx)
		if err != nil {
			fmt.Println("FILE SET ACCESSED ERR")
			ctx.AbortWithStatus(404)
			return
		}
		ctx.Next()
	}
}

// Defining the Graphql handler
func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.New(graph.NewSchema(client))

	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New(1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func meHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		entUser := c.Request.Context().Value(auth.UserCtxKey)
		log.Printf("%v", entUser)
		if entUser == nil {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(200, entUser)
		}
	}
}

func importUsers(client *ent.Client, ctx context.Context, filepath string) error {
	// Import users from json file
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	var jsonUsers []ent.User
	err = json.Unmarshal(byteValue, &jsonUsers)
	if err != nil {
		return err
	}

	for _, jsonUser := range jsonUsers {
		entUser, err := client.User.Query().Where(user.NameEQ(jsonUser.Name)).Only(ctx)
		if err != nil {
			if err == err.(*ent.NotFoundError) {
				_, err = client.User.Create().
					SetName(jsonUser.Name).
					SetType(jsonUser.Type).
					Save(ctx)
				if err != nil {
					return err
				}
				continue
			} else {
				return err
			}
		}
		_, err = entUser.Update().SetType(jsonUser.Type).Save(ctx)
		if err != nil {
			return err
		}
	}

	entUsers, err := client.User.Query().All(ctx)
	if err != nil {
		return err
	}

	jsonByteArray, err := json.MarshalIndent(entUsers, "", "  ")
	if err != nil {
		return err
	}
	binaryName := "export.json"

	err = ioutil.WriteFile(binaryName, jsonByteArray, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	client := ent.SQLLiteOpen("file:test.sqlite?_loc=auto&cache=shared&_fk=1")

	ctx := context.Background()
	defer ctx.Done()
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := importUsers(client, ctx, "users.json"); err != nil {
		log.Fatalf("failed creating user resources: %v", err)
	}

	absolutePath, err := filepath.Abs(graph.DownloadFolder)
	if err != nil {
		log.Fatalf("failed finding download folder path: %v", err)
	}
	if err = os.MkdirAll(absolutePath, 0755); err != nil {
		log.Fatalf("failed creating download folder: %v", err)
	}

	router := gin.Default()

	cors_urls := []string{"http://localhost", "http://localhost:3000"}
	if env_value, exists := os.LookupEnv("CORS_ALLOWED_ORIGINS"); exists {
		cors_urls = strings.Split(env_value, ",")
	}

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cors_urls,
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = defaultPort
	}

	gqlHandler := graphqlHandler(client)

	router.GET("/api/download/:url_id", tempURLHandler(client))
	router.POST("/api/auth", auth.LoginHandler(client))
	api := router.Group("/api")
	api.Use(auth.Middleware(client))

	api.GET("/me", meHandler())
	api.POST("/query", gqlHandler)
	api.GET("/query", gqlHandler)
	api.GET("/playground", playgroundHandler())
	router.Run(port)

}
