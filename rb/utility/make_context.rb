# EarthquakeCatalog SDK utility: make_context
require_relative '../core/context'
module EarthquakeCatalogUtilities
  MakeContext = ->(ctxmap, basectx) {
    EarthquakeCatalogContext.new(ctxmap, basectx)
  }
end
