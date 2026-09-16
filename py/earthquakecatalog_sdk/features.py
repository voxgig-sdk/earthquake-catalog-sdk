# EarthquakeCatalog SDK feature factory

from earthquakecatalog_sdk.feature.base_feature import EarthquakeCatalogBaseFeature
from earthquakecatalog_sdk.feature.ratelimit_feature import EarthquakeCatalogRatelimitFeature
from earthquakecatalog_sdk.feature.retry_feature import EarthquakeCatalogRetryFeature
from earthquakecatalog_sdk.feature.test_feature import EarthquakeCatalogTestFeature
from earthquakecatalog_sdk.feature.timeout_feature import EarthquakeCatalogTimeoutFeature


_FEATURES = {
    "base": lambda: EarthquakeCatalogBaseFeature(),
    "ratelimit": lambda: EarthquakeCatalogRatelimitFeature(),
    "retry": lambda: EarthquakeCatalogRetryFeature(),
    "test": lambda: EarthquakeCatalogTestFeature(),
    "timeout": lambda: EarthquakeCatalogTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
