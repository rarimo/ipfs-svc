package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/Dmytro-Hladkykh/ipfs-svc/internal/service/helpers"
	"github.com/Dmytro-Hladkykh/ipfs-svc/internal/service/ipfs"
	"github.com/Dmytro-Hladkykh/ipfs-svc/internal/service/requests"
	"github.com/Dmytro-Hladkykh/ipfs-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UploadJSON(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUploadJSON(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	cfg := helpers.Config(r)
	ipfsClient := ipfs.NewClient(ipfs.ClientOpts{
		NodeAddress: cfg.IPFS().NodeAddress,
		Timeout:     cfg.IPFS().Timeout,
	})

	hash, err := ipfsClient.UploadJSON(r.Context(), request.Data.Attributes.Metadata)
	if err != nil {
		log := helpers.Log(r).WithError(err)

		switch {
		case errors.Is(err, context.DeadlineExceeded):
			log.Error("ipfs upload timeout")
			ape.RenderErr(w, problems.RequestTimeout())
		default:
			log.Error("failed to upload to ipfs")
			ape.RenderErr(w, problems.InternalError())
		}
		return
	}

	response := resources.UploadJsonResponse{
		Key: resources.Key{
			ID:   hash,
			Type: resources.UPLOAD_JSON,
		},
		Attributes: resources.UploadJsonResponseAttributes{
			Hash: hash,
		},
	}

	ape.Render(w, resources.UploadJsonResponseResponse{
		Data: response,
	})
}
