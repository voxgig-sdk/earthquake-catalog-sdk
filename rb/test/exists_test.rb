# EarthquakeCatalog SDK exists test

require "minitest/autorun"
require_relative "../EarthquakeCatalog_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = EarthquakeCatalogSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
