# EarthquakeCatalog SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

EarthquakeCatalogUtility.registrar = ->(u) {
  u.clean = EarthquakeCatalogUtilities::Clean
  u.done = EarthquakeCatalogUtilities::Done
  u.make_error = EarthquakeCatalogUtilities::MakeError
  u.feature_add = EarthquakeCatalogUtilities::FeatureAdd
  u.feature_hook = EarthquakeCatalogUtilities::FeatureHook
  u.feature_init = EarthquakeCatalogUtilities::FeatureInit
  u.fetcher = EarthquakeCatalogUtilities::Fetcher
  u.make_fetch_def = EarthquakeCatalogUtilities::MakeFetchDef
  u.make_context = EarthquakeCatalogUtilities::MakeContext
  u.make_options = EarthquakeCatalogUtilities::MakeOptions
  u.make_request = EarthquakeCatalogUtilities::MakeRequest
  u.make_response = EarthquakeCatalogUtilities::MakeResponse
  u.make_result = EarthquakeCatalogUtilities::MakeResult
  u.make_point = EarthquakeCatalogUtilities::MakePoint
  u.make_spec = EarthquakeCatalogUtilities::MakeSpec
  u.make_url = EarthquakeCatalogUtilities::MakeUrl
  u.param = EarthquakeCatalogUtilities::Param
  u.prepare_auth = EarthquakeCatalogUtilities::PrepareAuth
  u.prepare_body = EarthquakeCatalogUtilities::PrepareBody
  u.prepare_headers = EarthquakeCatalogUtilities::PrepareHeaders
  u.prepare_method = EarthquakeCatalogUtilities::PrepareMethod
  u.prepare_params = EarthquakeCatalogUtilities::PrepareParams
  u.prepare_path = EarthquakeCatalogUtilities::PreparePath
  u.prepare_query = EarthquakeCatalogUtilities::PrepareQuery
  u.result_basic = EarthquakeCatalogUtilities::ResultBasic
  u.result_body = EarthquakeCatalogUtilities::ResultBody
  u.result_headers = EarthquakeCatalogUtilities::ResultHeaders
  u.transform_request = EarthquakeCatalogUtilities::TransformRequest
  u.transform_response = EarthquakeCatalogUtilities::TransformResponse
}
