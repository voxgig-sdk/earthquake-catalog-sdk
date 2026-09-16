# EarthquakeCatalog SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module EarthquakeCatalogFeatures
  def self.make_feature(name)
    case name
    when "base"
      EarthquakeCatalogBaseFeature.new
    when "ratelimit"
      EarthquakeCatalogRatelimitFeature.new
    when "retry"
      EarthquakeCatalogRetryFeature.new
    when "test"
      EarthquakeCatalogTestFeature.new
    when "timeout"
      EarthquakeCatalogTimeoutFeature.new
    else
      EarthquakeCatalogBaseFeature.new
    end
  end
end
