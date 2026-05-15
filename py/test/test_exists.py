# ProjectName SDK exists test

import pytest
from earthquakecatalog_sdk import EarthquakeCatalogSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = EarthquakeCatalogSDK.test(None, None)
        assert testsdk is not None
