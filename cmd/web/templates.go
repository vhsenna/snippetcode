package main

import "github.com/vhsenna/snippetcode/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
