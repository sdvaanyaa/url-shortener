package delete_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/sdvaanyaa/url-shortener/internal/http-server/handlers/url/delete"
	"github.com/sdvaanyaa/url-shortener/internal/http-server/handlers/url/delete/mocks"
	"github.com/sdvaanyaa/url-shortener/internal/lib/api/response"
	"github.com/sdvaanyaa/url-shortener/internal/lib/logger/handlers/slogdiscard"
	"github.com/sdvaanyaa/url-shortener/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestDeleteHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		mockError error
		respError string
	}{
		{
			name:  "Success",
			alias: "test_alias",
		},
		{
			name:      "Empty alias",
			alias:     "",
			respError: "invalid request",
		},
		{
			name:      "URL not found",
			alias:     "not_exists",
			mockError: storage.ErrURLNotFound,
			respError: "url not found",
		},
		{
			name:      "Internal error",
			alias:     "internal_err",
			mockError: errors.New("unexpected error"),
			respError: "internal server error",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			urlDeleterMock := mocks.NewURLDeleter(t)

			if tc.alias != "" {
				urlDeleterMock.On("DeleteURL", tc.alias).
					Return(tc.mockError).
					Once()
			}

			r := chi.NewRouter()
			r.Delete("/{alias}", delete.New(slogdiscard.NewDiscardLogger(), urlDeleterMock))

			req, err := http.NewRequest(http.MethodDelete, "/"+tc.alias, nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if tc.alias == "" {
				require.Equal(t, http.StatusNotFound, rr.Code)
			} else {
				require.Equal(t, http.StatusOK, rr.Code)

				var resp response.Response
				require.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))

				// Проверяем поле Error
				require.Equal(t, tc.respError, resp.Error)
			}
		})
	}
}
