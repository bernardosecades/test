//go:build e2e
// +build e2e

package purchase

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"testing"

	"github.com/bernardosecades/test/internal/purchase/repository/mysql"
	"github.com/bernardosecades/test/internal/purchase/service"
	"github.com/bernardosecades/test/internal/testtools"
	"github.com/bernardosecades/test/proto/go/protobuf/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var clientDB *sql.DB

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	purchaseRepository := mysql.NewMySQLPurchaseRepository(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	purchaseService := service.NewPurchaseService(&purchaseRepository)

	// Create new gRPC server instance
	srv := &Controller{PurchaseService: purchaseService}
	api.RegisterPurchaseServiceServer(s, srv)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestMain(m *testing.M) {
	clientDB = getDBConnection()
	loadFixtures()
	exitVal := m.Run()

	testtools.TruncateTables(clientDB, getTablesToTruncateOnFinish())
	teardown()
	os.Exit(exitVal)
}

func TestMakePurchaseErrorInvalidUuid(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewPurchaseServiceClient(conn)
	resp, err := client.Purchase(ctx, &api.PurchaseRequest{Id: "invalid uuid"})

	assert.Nil(t, err)
	assert.False(t, resp.Success)
	assert.Contains(t, resp.Error, service.ErrInvalidID.Error())
}

func TestMakePurchaseErrorProductNoAvailable(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewPurchaseServiceClient(conn)
	resp, err := client.Purchase(ctx, &api.PurchaseRequest{Id: "1d6f0039-a9a7-4092-9241-f49a9bc0bb03"})

	assert.Nil(t, err)
	assert.False(t, resp.Success)
	assert.Contains(t, resp.Error, service.ErrProductNoAvailable.Error())
}

func TestMakePurchaseSuccess(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewPurchaseServiceClient(conn)
	resp, err := client.Purchase(ctx, &api.PurchaseRequest{Id: "a1b84675-eae8-446c-91ce-770c7b22ce3e"})

	assert.Nil(t, err)
	assert.True(t, resp.Success)
	assert.Empty(t, resp.Error)
}

func getTablesToTruncateOnFinish() []string {
	return []string{"product"}
}

func teardown() {
	testtools.CloseDB(clientDB)
}

func loadFixtures() {
	queries := []string{
		`INSERT INTO product (id, available, created_at, updated_at)
		VALUES
    	('a1b84675-eae8-446c-91ce-770c7b22ce3e', 1 , '2023-01-02 15:20:44', '2023-01-02 15:20:44'),
		('1d6f0039-a9a7-4092-9241-f49a9bc0bb03', 0 , '2023-01-02 15:20:44', '2023-01-02 15:20:44');`,
	}

	testtools.LoadFixtures(clientDB, queries)
}

func getDBConnection() *sql.DB {
	return testtools.OpenDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
