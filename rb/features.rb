# EarthquakeCatalog SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module EarthquakeCatalogFeatures
  def self.make_feature(name)
    case name
    when "base"
      EarthquakeCatalogBaseFeature.new
    when "test"
      EarthquakeCatalogTestFeature.new
    else
      EarthquakeCatalogBaseFeature.new
    end
  end
end
