package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rarimo/ipfs-svc/internal/service/helpers"
	"github.com/rarimo/ipfs-svc/internal/service/ipfs"
	"github.com/rarimo/ipfs-svc/internal/service/requests"
	"github.com/rarimo/ipfs-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	cfg := helpers.Config(r)

	maxFileSize := cfg.IPFS().MaxFileSize

	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		ape.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to parse multipart form: %w", err))...)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(fmt.Errorf("failed to get image file from request: %w", err))...)
		return
	}
	defer file.Close()

	_, err = requests.NewUploadFile(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ipfsClient := ipfs.NewClient(ipfs.ClientOpts{
		NodeAddress: cfg.IPFS().NodeAddress,
		Timeout:     cfg.IPFS().Timeout,
	})

	hash, err := ipfsClient.UploadFile(r.Context(), file)
	if err != nil {
		log := helpers.Log(r).WithError(err)

		switch {
		case errors.Is(err, context.DeadlineExceeded):
			log.Error("ipfs upload timeout")
			ape.RenderErr(w, problems.RequestTimeout())
		default:
			log.Error("failed to upload image to ipfs")
			ape.RenderErr(w, problems.InternalError())
		}
		return
	}

	response := resources.UploadFileResponse{
		Key: resources.Key{
			ID:   hash,
			Type: resources.FILE,
		},
		Attributes: resources.UploadFileResponseAttributes{
			Hash: hash,
		},
	}

	ape.Render(w, resources.UploadFileResponseResponse{
		Data: response,
	})
}
