# ServiceInformation entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from earthquakecatalog_sdk import EarthquakeCatalogSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestServiceInformationEntity:

    def test_should_create_instance(self):
        testsdk = EarthquakeCatalogSDK.test(None, None)
        ent = testsdk.ServiceInformation(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "service_information": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = EarthquakeCatalogSDK.test(seed, None)
        seen = list(base.ServiceInformation(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from config import make_config
        cfg = make_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = EarthquakeCatalogSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.ServiceInformation(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_run_basic_flow(self):
        setup = _service_information_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["list", "load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "service_information." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        service_information_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.service_information")))
        service_information_ref01_data = None
        if len(service_information_ref01_data_raw) > 0:
            service_information_ref01_data = helpers.to_map(service_information_ref01_data_raw[0][1])

        # LIST
        service_information_ref01_ent = client.ServiceInformation(None)
        service_information_ref01_match = {}

        service_information_ref01_list_result = service_information_ref01_ent.list(service_information_ref01_match, None)
        assert isinstance(service_information_ref01_list_result, list)

        # LOAD
        service_information_ref01_match_dt0 = {}
        service_information_ref01_data_dt0_loaded = service_information_ref01_ent.load(service_information_ref01_match_dt0, None)
        assert service_information_ref01_data_dt0_loaded is not None



def _service_information_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/service_information/ServiceInformationTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = EarthquakeCatalogSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["service_information01", "service_information02", "service_information03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID": idmap,
        "EARTHQUAKECATALOG_TEST_LIVE": "FALSE",
        "EARTHQUAKECATALOG_TEST_EXPLAIN": "FALSE",
    })

    idmap_resolved = helpers.to_map(
        env.get("EARTHQUAKECATALOG_TEST_SERVICE_INFORMATION_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("EARTHQUAKECATALOG_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            {
            },
            extra or {},
        ])
        client = EarthquakeCatalogSDK(helpers.to_map(merged_opts))

    _live = env.get("EARTHQUAKECATALOG_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("EARTHQUAKECATALOG_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
