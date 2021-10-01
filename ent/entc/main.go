package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaPath("../graphql/generated/ent.graphql"),
	)

	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	entcOptions := []entc.Option{
		entc.Extensions(ex),
		entc.TemplateDir("./template"),
	}

	err = entc.Generate("./schema", &gen.Config{
		Schema:  "./schema",
		Target:  "../internal/ent",
		Package: "github.com/responserms/response/internal/ent",
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeaturePrivacy,
			gen.FeatureLock,
			gen.FeatureUpsert,
			gen.FeatureModifier,
		},
	}, entcOptions...)

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
