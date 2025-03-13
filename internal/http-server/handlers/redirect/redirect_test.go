package redirect_test

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/sdvaanyaa/url-shortener/internal/http-server/handlers/redirect"
	"github.com/sdvaanyaa/url-shortener/internal/http-server/handlers/redirect/mocks"
	"github.com/sdvaanyaa/url-shortener/internal/lib/api"
	"github.com/sdvaanyaa/url-shortener/internal/lib/logger/handlers/slogdiscard"
	"github.com/sdvaanyaa/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestRedirectHandler(t *testing.T) {
	cases := []struct {
		name        string
		alias       string
		url         string
		respError   string
		mockError   error
		expectedURL string
		expectError bool
	}{
		{
			name:        "Success",
			alias:       "test_alias",
			url:         "https://www.google.com/",
			expectedURL: "https://www.google.com/",
		},
		{
			name:        "Success Yandex",
			alias:       "ya_music",
			url:         "https://music.yandex.ru/",
			expectedURL: "https://music.yandex.ru/",
		},
		{
			name:        "Empty alias",
			alias:       "",
			respError:   "invalid request",
			expectError: true,
		},
		{
			name:        "URL not found",
			alias:       "not_exists",
			mockError:   storage.ErrURLNotFound,
			respError:   "not found",
			expectError: true,
		},
		{
			name:        "Internal error",
			alias:       "internal_err",
			mockError:   errors.New("db error"),
			respError:   "internal server error",
			expectError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			urlGetterMock := mocks.NewURLGetter(t)

			if tc.alias != "" { // Для кейса с пустым alias мок не нужен
				urlGetterMock.On("GetURL", tc.alias).
					Return(tc.url, tc.mockError).Once()
			}

			r := chi.NewRouter()
			r.Get("/{alias}", redirect.New(slogdiscard.NewDiscardLogger(), urlGetterMock))

			ts := httptest.NewServer(r)
			defer ts.Close()

			redirectedToURL, err := api.GetRedirect(ts.URL + "/" + tc.alias)

			if tc.expectError {
				require.Error(t, err)
				if tc.respError == "invalid request" {
					// Для пустого alias получаем ошибку статуса
					assert.Contains(t, err.Error(), api.ErrInvalidStatusCode.Error())
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedURL, redirectedToURL)
			}
		})
	}
}
