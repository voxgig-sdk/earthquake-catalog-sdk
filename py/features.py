# EarthquakeCatalog SDK feature factory

from feature.base_feature import EarthquakeCatalogBaseFeature
from feature.test_feature import EarthquakeCatalogTestFeature


def _make_feature(name):
    features = {
        "base": lambda: EarthquakeCatalogBaseFeature(),
        "test": lambda: EarthquakeCatalogTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
