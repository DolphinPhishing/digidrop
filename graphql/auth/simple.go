package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"digidrop/ent"
	"digidrop/ent/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const CompromisedPath = "digidrop.compromised"

type contextKey struct {
	name string
}

var UserCtxKey = &contextKey{"user"}

func LoginHandler(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hostname, ok := os.LookupEnv("GRAPHQL_HOSTNAME")
		if !ok {
			hostname = "keeper.brentwood.nj"
		}

		cookie_name, ok := os.LookupEnv("COOKIE_NAME")
		if !ok {
			cookie_name = "dead-auth"
		}

		userToken := ctx.PostForm("token")
		if userToken == "" {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		userUUID, err := uuid.Parse(userToken)
		if err != nil {
			fmt.Printf("Error: unable to parse user token: %v\n", err)
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		_, err = client.User.Get(ctx, userUUID)
		if err != nil {
			fmt.Printf("Error: unable to find user: %v\n", err)
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		} else {
			ctx.SetCookie(cookie_name, userToken, 60*60*24, "/", hostname, false, false)
			ctx.JSON(http.StatusOK, "")
		}
	}
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookie_name, ok := os.LookupEnv("COOKIE_NAME")
		if !ok {
			cookie_name = "dead-auth"
		}

		authCookie, err := ctx.Cookie(cookie_name)
		if err != nil || authCookie == "" {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		userUUID, err := uuid.Parse(authCookie)
		if err != nil {
			fmt.Printf("Error: unable to parse user cookie: %v\n", err)
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		entUser, err := client.User.Get(ctx, userUUID)
		if err != nil {
			fmt.Printf("Error: unable to find user: %v\n", err)
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		if entUser.Type == user.TypeUSER {
			_, err := os.Stat(CompromisedPath)
			if err == nil {
				ctx.AbortWithStatus(http.StatusGone)
				return
			}
		}

		// put it in context
		c := context.WithValue(ctx.Request.Context(), UserCtxKey, entUser)
		ctx.Request = ctx.Request.WithContext(c)

		ctx.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*ent.User, error) {
	raw, ok := ctx.Value(UserCtxKey).(*ent.User)
	if ok {
		return raw, nil
	}
	return nil, errors.New("unable to get authuser from context")

}
