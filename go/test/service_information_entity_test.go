package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/earthquake-catalog-sdk/go"
	"github.com/voxgig-sdk/earthquake-catalog-sdk/go/core"

	vs "github.com/voxgig-sdk/earthquake-catalog-sdk/go/utility/struct"
)

func TestServiceInformationEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.ServiceInformation(nil)
		if ent == nil {
			t.Fatal("expected non-nil ServiceInformationEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := service_informationBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"list", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "service_information." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		serviceInformationRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.service_information", setup.data)))
		var serviceInformationRef01Data map[string]any
		if len(serviceInformationRef01DataRaw) > 0 {
			serviceInformationRef01Data = core.ToMapAny(serviceInformationRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = serviceInformationRef01Data

		// LIST
		serviceInformationRef01Ent := client.ServiceInformation(nil)
		serviceInformationRef01Match := map[string]any{}

		serviceInformationRef01ListResult, err := serviceInformationRef01Ent.List(serviceInformationRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		_, serviceInformationRef01ListOk := serviceInformationRef01ListResult.([]any)
		if !serviceInformationRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", serviceInformationRef01ListResult)
		}

		// LOAD
		serviceInformationRef01MatchDt0 := map[string]any{}
		serviceInformationRef01DataDt0Loaded, err := serviceInformationRef01Ent.Load(serviceInformationRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if serviceInformationRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func service_informationBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "service_information", "ServiceInformationTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read service_information test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse service_information test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"service_information01", "service_information02", "service_information03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID": idmap,
		"EARTHQUAKECATALOG_TEST_LIVE":      "FALSE",
		"EARTHQUAKECATALOG_TEST_EXPLAIN":   "FALSE",
		"EARTHQUAKECATALOG_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["EARTHQUAKECATALOG_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["EARTHQUAKECATALOG_APIKEY"],
			},
			extra,
		})
		client = sdk.NewEarthquakeCatalogSDK(core.ToMapAny(mergedOpts))
	}

	live := env["EARTHQUAKECATALOG_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["EARTHQUAKECATALOG_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
