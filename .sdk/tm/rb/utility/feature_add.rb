# EarthquakeCatalog SDK utility: feature_add
module EarthquakeCatalogUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
