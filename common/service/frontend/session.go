package frontend

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"

	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/auth"
	"github.com/pydio/cells/common/config"
	"github.com/pydio/cells/common/log"
)

const (
	SessionTimeoutMinutes = 24
)

var sessionStore *sessions.CookieStore

func GetSessionStore() sessions.Store {
	if sessionStore == nil {
		val := config.ApplicationConfig.Values("frontend", "session", "secureKey").String()
		var key []byte
		if val == "" {
			key = securecookie.GenerateRandomKey(64)
			val = base64.StdEncoding.EncodeToString(key)
			config.Set(val, "frontend", "session", "secureKey")
			config.Save(common.PYDIO_SYSTEM_USERNAME, "Generating session random key")
		} else {
			key, _ = base64.StdEncoding.DecodeString(val)
		}
		sessionStore = sessions.NewCookieStore([]byte(val))
		sessionStore.Options = &sessions.Options{
			Path:     "/a/frontend",
			MaxAge:   60 * SessionTimeoutMinutes,
			HttpOnly: true,
		}
		if config.ApplicationConfig.Values("cert", "proxy", "ssl").Bool() {
			sessionStore.Options.Secure = true
		}
		urlVal := config.ApplicationConfig.Values("defaults", "url").String()
		if parsed, e := url.Parse(urlVal); e == nil {
			sessionStore.Options.Domain = strings.Split(parsed.Host, ":")[0]
		}
	}
	return sessionStore
}

// NewSessionWrapper creates a Http middleware checking if a cookie is passed
// along and if this cookie contains a proper JWT.
// The excludes parameter may be used to namely ignore specific "METHOD:/URI“ for this check.
func NewSessionWrapper(h http.Handler, excludes ...string) http.Handler {

	jwtVerifier := auth.DefaultJWTVerifier()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, excluded := range excludes {
			split := strings.Split(excluded, ":")
			method := split[0]
			uri := split[1]
			if strings.ToLower(method) == strings.ToLower(r.Method) && strings.HasPrefix(r.RequestURI, uri) {
				log.Logger(r.Context()).Debug("Skipping session wrapper by exclusion:" + excluded)
				h.ServeHTTP(w, r)
				return
			}
		}

		sessionName := "pydio"
		if h, ok := r.Header["X-Pydio-Minisite"]; ok {
			sessionName = sessionName + "-" + strings.Join(h, "")
		}
		if h, ok := r.Header[common.XPydioFrontendSessionUuid]; ok {
			sessionName = sessionName + "-" + strings.Join(h, "")
		}
		session, err := GetSessionStore().Get(r, sessionName)
		if err != nil && !strings.Contains(err.Error(), "securecookie: the value is not valid") {
			log.Logger(r.Context()).Error("Cannot retrieve session", zap.Error(err))
		}

		if value, ok := session.Values["access_token"]; ok {
			ctx := r.Context()
			if ctx, _, err = jwtVerifier.Verify(ctx, value.(string)); err == nil {
				// Update context
				log.Logger(ctx).Debug("Found token in session " + sessionName)
				r = r.WithContext(ctx)
			}
		}
		h.ServeHTTP(w, r)
	})
}
