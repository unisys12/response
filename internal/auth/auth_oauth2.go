package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/responserms/response/internal/config"

	"github.com/rs/zerolog"

	"github.com/jeremywohl/flatten"
	"golang.org/x/oauth2"
)

type extraHeadersTemplateData struct {
	ClientID string
}

func handleOAuth2UserDetails(log zerolog.Logger, provider config.OAuth2Provider) func(ctx context.Context, config oauth2.Config, token *oauth2.Token) (map[string]string, error) {
	return func(ctx context.Context, config oauth2.Config, token *oauth2.Token) (map[string]string, error) {
		req, err := http.NewRequest(http.MethodGet, provider.UserinfoURL, nil)
		if err != nil {
			return nil, fmt.Errorf("http.NewRequest: %w", err)
		}

		for key, t := range provider.UserinfoExtraHeaders {
			var buf bytes.Buffer
			if err := t.Execute(&buf, &extraHeadersTemplateData{ClientID: config.ClientID}); err != nil {
				return nil, err
			}

			req.Header.Set(key, buf.String())
		}

		res, err := provider.Config.Client(ctx, token).
			Do(req)

		if err != nil {
			log.Error().
				Err(err).
				Msg("Unable to return Userinfo from the OAuth 2.0 oAuth2Provider. Does the oAuth2Provider require extra config?")

			return nil, fmt.Errorf("oAuth2Provider.Client.Get: %w", err)
		}

		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
		}

		var data map[string]interface{}
		if err := json.Unmarshal(bytes, &data); err != nil {
			log.Error().
				Err(err).
				Msg("Unable to unmarshal Userinfo from oAuth2Provider's response.")

			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}

		flattened, err := flatten.Flatten(data, "", flatten.DotStyle)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Unable to flatten JSON response from oAuth2Provider.")

			return nil, fmt.Errorf("flatten.Flatten: %w", err)
		}

		result := map[string]string{}
		for k, v := range flattened {
			if str, ok := v.(string); ok {
				result[k] = str
			}
		}

		return result, nil
	}
}
