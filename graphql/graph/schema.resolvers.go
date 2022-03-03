package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"digidrop/ent"
	"digidrop/graphql/auth"
	"digidrop/graphql/graph/generated"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jkomyno/nanoid"
)

func (r *fileMiddlewareResolver) User(ctx context.Context, obj *ent.FileMiddleware) (string, error) {
	entUser, err := obj.QueryFileMiddlewareToUser().Only(ctx)
	if err != nil {
		return "nil", err
	}
	return entUser.Name, nil
}

func (r *mutationResolver) Upload(ctx context.Context, file graphql.Upload) (*ent.FileMiddleware, error) {
	currentUser, err := auth.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	urlId, err := nanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 7)
	if err != nil {
		return nil, fmt.Errorf("failed to generate file URLID: %v", err)
	}

	absolutePath, err := filepath.Abs(DownloadFolder + "/" + strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)) + "_" + urlId + filepath.Ext(file.Filename))
	if err != nil {
		return nil, err
	}

	// Create the file
	out, err := os.Create(absolutePath)
	if err != nil {
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, file.File)

	if err != nil {
		return nil, err
	}

	return r.client.FileMiddleware.Create().
		SetURLID(urlId).
		SetFilePath(absolutePath).
		SetFileMiddlewareToUser(currentUser).
		Save(ctx)
}

func (r *mutationResolver) Compromise(ctx context.Context) (bool, error) {
	absolutePath, err := filepath.Abs(CompromisedPath)
	if err != nil {
		return false, err
	}

	// Create the file
	_, err = os.Create(absolutePath)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*ent.FileMiddleware, error) {
	return r.client.FileMiddleware.Query().All(ctx)
}

func (r *queryResolver) Compromised(ctx context.Context) (bool, error) {
	_, err := os.Stat(CompromisedPath)
	if err == nil {
		return true, nil
	}
	return false, nil
}

// FileMiddleware returns generated.FileMiddlewareResolver implementation.
func (r *Resolver) FileMiddleware() generated.FileMiddlewareResolver {
	return &fileMiddlewareResolver{r}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type fileMiddlewareResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
