package testutils

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/ent/dialect"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TC struct {
	Container testcontainers.Container
	Dialect   string
	URI       string
}

// TODO: troubleshoot support for cockroachdb; currently complains about dropping an index during tests

// func getCRDB(ctx context.Context, image string, opts ...testcontainers.ContainerCustomizer) (testcontainers.Container, string, error) {
// 	defaultImg := "docker.io/cockroachdb/cockroach"
// 	imgTag := "latest"

// 	if strings.Contains(image, ":") {
// 		p := strings.SplitN(image, ":", 2) //nolint:gomnd
// 		imgTag = p[1]
// 	}

// 	req := testcontainers.ContainerRequest{
// 		Image:        fmt.Sprintf("%s:%s", defaultImg, imgTag),
// 		ExposedPorts: []string{"26257/tcp", "8080/tcp"},
// 		WaitingFor:   wait.ForHTTP("/health").WithPort("8080"),
// 		Cmd:          []string{"start-single-node", "--insecure"},
// 	}

// 	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: req,
// 		Started:          true,
// 	})
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	mappedPort, err := container.MappedPort(ctx, "26257")
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	hostIP, err := container.Host(ctx)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	uri := fmt.Sprintf("postgres://root@%s:%s/defaultdb?sslmode=disable", hostIP, mappedPort.Port())

// 	return container, uri, nil
// }

func getPG(ctx context.Context, image string, opts ...testcontainers.ContainerCustomizer) (postgres.PostgresContainer, string, error) {
	defaultImg := "docker.io/postgres"
	imgTag := "alpine"

	if strings.Contains(image, ":") {
		p := strings.SplitN(image, ":", 2) //nolint:gomnd
		imgTag = p[1]
	}

	uriFunc := func(host string, port nat.Port) string {
		return fmt.Sprintf("postgres://postgres:postgres@%s:%s/postgres?sslmode=disable", host, port.Port())
	}

	opts = append(opts,
		testcontainers.WithImage(fmt.Sprintf("%s:%s", defaultImg, imgTag)),
		testcontainers.WithWaitStrategy(wait.ForSQL(nat.Port("5432"), "postgres", uriFunc)),
		postgres.WithPassword("postgres"),
	)

	container, err := postgres.RunContainer(ctx, opts...)
	if err != nil {
		return postgres.PostgresContainer{}, "", err
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return postgres.PostgresContainer{}, "", err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return postgres.PostgresContainer{}, "", err
	}

	uri := fmt.Sprintf("postgres://postgres:postgres@%s:%s/postgres?sslmode=disable", hostIP, mappedPort.Port())

	return *container, uri, nil
}

func getTestDB(ctx context.Context, u string) (TC, error) {
	switch {
	// case strings.HasPrefix(u, "cockroach"), strings.HasPrefix(u, "cockroachdb"), strings.HasPrefix(u, "crdb"):
	// 	container, uri, err := getCRDB(ctx, u)
	// 	return TC{Container: container, URI: uri, Dialect: dialect.Postgres}, err
	case strings.HasPrefix(u, "postgres"):
		container, uri, err := getPG(ctx, u)
		return TC{Container: container, URI: uri, Dialect: dialect.Postgres}, err
	default:
		return TC{}, newURIError(u)
	}
}

// GetTestURI returns the dialect, connection string and if used a testcontainer for database connectivity in tests
func GetTestURI(ctx context.Context, u string) *TC {
	switch {
	case u == "":
		// return dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1"
		return &TC{Dialect: dialect.SQLite, URI: "file:ent?mode=memory&cache=shared&_fk=1"}
	case strings.HasPrefix(u, "sqlite://"):
		// return dialect.SQLite, strings.TrimPrefix(u, "sqlite://")
		return &TC{Dialect: dialect.SQLite, URI: strings.TrimPrefix(u, "sqlite://")}
	case strings.HasPrefix(u, "postgres://"), strings.HasPrefix(u, "postgresql://"):
		// return dialect.Postgres, u
		return &TC{Dialect: dialect.Postgres, URI: u}
	case strings.HasPrefix(u, "docker://"):
		container, err := getTestDB(ctx, strings.TrimPrefix(u, "docker://"))
		if err != nil {
			panic(err)
		}

		return &container
	default:
		panic("invalid DB URI, uri: " + u)
	}
}
