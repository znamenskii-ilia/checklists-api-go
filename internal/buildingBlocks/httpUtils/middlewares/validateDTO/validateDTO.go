package validateDTO

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func New[T any]() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var dto T
			err := json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			validate := validator.New(validator.WithRequiredStructEnabled())
			err = validate.Struct(dto)
			if err != nil {
				var validationErrs validator.ValidationErrors
				if errors.As(err, &validationErrs) {
					errs := []string{}
					for _, e := range validationErrs {
						errs = append(errs, fmt.Sprintf("%s: %s", strings.ToLower(e.Field()), e.Tag()))
					}
					http.Error(w, strings.Join(errs, "\n"), http.StatusBadRequest)
				} else {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				return
			}

			cxt := context.WithValue(r.Context(), "validatedDTO", dto)
			next.ServeHTTP(w, r.WithContext(cxt))
		})
	}
}
