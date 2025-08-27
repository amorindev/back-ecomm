package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/api/core"
	"github.com/amorindev/go-tmpl/pkg/app/ecomm/products/domain"
	cShared "github.com/amorindev/go-tmpl/pkg/shared/api/core"
	dShared "github.com/amorindev/go-tmpl/pkg/shared/domain"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(20 << 20) // 20MB
	if err != nil {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "invalid form"))
		return
	}

	// Get the product JSON
	productJSON := r.FormValue("product")
	if productJSON == "" {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "missing product json"))
		return
	}

	var req core.CreateProductReq
	if err := json.Unmarshal([]byte(productJSON), &req); err != nil {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "invalid product json"))
		return
	}

	if err := req.Validate(); err != nil {
		cShared.RespondError(w, err)
		return
	}

	// Main image
	files := map[string][]byte{}
	contentTypes := map[string]string{}

	if headers, ok := r.MultipartForm.File["main_image"]; ok {
		for _, header := range headers {
			// Try to open the uploaded file
			file, err := header.Open()
			if err != nil {
				msg := fmt.Sprintf("failed to open file %s: %v", header.Filename, err)
				cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, msg))
				return
			}

			defer file.Close()

			// Read file content
			buf, err := io.ReadAll(file)
			if err != nil {
				msg := fmt.Sprintf("failed to read file %s: %v", header.Filename, err)
				cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInternalServerError, msg))
				return
			}

			// Store file content and content type in maps
			files[header.Filename] = buf
			contentTypes[header.Filename] = header.Header.Get("Content-Type")
		}

	} else {
		cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, "main_image is required"))
		return
	}

	// Handle item images (if they exist)
	for i := 0; i < len(req.ProductItems); i++ {
		field := fmt.Sprintf("item_image_%d", i)
		if headers, ok := r.MultipartForm.File[field]; ok {
			for _, header := range headers {
				// Try to open the uploaded item image
				file, err := header.Open()
				if err != nil {
					msg := fmt.Sprintf("failed to open item image %s: %v", header.Filename, err)
					cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInvalidParams, msg))
					return
				}
				defer file.Close()
				buf, err := io.ReadAll(file)
				if err != nil {
					msg := fmt.Sprintf("failed to read item image %s: %v", header.Filename, err)
					cShared.RespondError(w, dShared.NewAppError(dShared.ErrCodeInternalServerError, msg))
					return
				}

				//  Store item image content and content type in maps
				files[header.Filename] = buf
				contentTypes[header.Filename] = header.Header.Get("Content-Type")
			}
		}
	}

	product := &domain.Product{
		Name:        req.Name,
		Desc:        req.Desc,
		CategoryID:  req.CategoryID,
		FilePath:    req.FilePath,
		File:        files[req.FilePath],
		ContentType: contentTypes[req.FilePath],
	}

	// Product with variants
	if len(req.ProductItems) > 1 {
		for _, it := range req.ProductItems {
			pItem := &domain.ProductItem{
				Price:        it.Price,
				QtyInStock:   it.QtyInStock,
				FilePath:     it.FilePath,
				File:         files[it.FilePath],
				ContentType:  contentTypes[it.FilePath],
				VarOptionIDs: it.VarOptionIDs,
			}
			product.ProductItems = append(product.ProductItems, pItem)
		}
	} else {
		// Product without variants -> a single item by default
		pItem := &domain.ProductItem{
			Price:      req.ProductItems[0].Price,
			QtyInStock: req.ProductItems[0].QtyInStock,
			// same as main image: file path, filename and content
			FilePath:    req.FilePath,
			File:        files[req.FilePath],
			ContentType: contentTypes[req.FilePath],
		}
		product.ProductItems = append(product.ProductItems, pItem)
	}

	if err := h.ProductSrv.Create(context.Background(), product); err != nil {
		cShared.RespondError(w, err)
		return
	}

	resp := struct {
		ID string `json:"id"`
	}{
		ID: product.ID.(string),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
