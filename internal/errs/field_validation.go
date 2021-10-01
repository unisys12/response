package errs

import (
	"context"
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/go-playground/validator/v10"
)

func GetHumanReadableError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required_without_all":
		return fmt.Sprintf("The %s field is required if the following fields are not set: %s", fe.Field(), strings.Join(strings.Split(fe.Param(), " "), ", "))
	case "required_without":
		return fmt.Sprintf("The %s field is required if any of the following fields are not set: %s", fe.Field(), strings.Join(strings.Split(fe.Param(), " "), ", "))
	case "url":
		return fmt.Sprintf("The %s field should be a valid URL.", fe.Field())
	case "min":
		return fmt.Sprintf("The %s field must have a minimum length of %s", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("The %s field must have a maximum length of %s", fe.Field(), fe.Param())
	case "length":
		return fmt.Sprintf("The %s field must have a length of %s", fe.Field(), fe.Param())
	case "WellKnownOAuthProviders":
		return fmt.Sprintf("%s is not a well-known OAuth Provider.", fe.Value())
	default:
		return fe.Error()
	}
}

func FieldsHaveValidationErrors(ctx context.Context, err error) bool {
	fieldErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}

	for _, fieldError := range fieldErrors {
		errPath := graphql.GetPath(ctx)
		errPath = append(errPath, ast.PathName(fieldError.StructField()))

		graphql.AddError(ctx, &gqlerror.Error{
			Path:       errPath,
			Message:    GetHumanReadableError(fieldError),
			Extensions: map[string]interface{}{},
		})
	}

	return true
}
