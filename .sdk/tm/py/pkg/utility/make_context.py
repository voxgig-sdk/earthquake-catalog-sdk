# EarthquakeCatalog SDK utility: make_context

from projectname_sdk.core.context import EarthquakeCatalogContext


def make_context_util(ctxmap, basectx):
    return EarthquakeCatalogContext(ctxmap, basectx)
