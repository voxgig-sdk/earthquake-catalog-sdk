<?php
declare(strict_types=1);

// EarthquakeCatalog SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

EarthquakeCatalogUtility::setRegistrar(function (EarthquakeCatalogUtility $u): void {
    $u->clean = [EarthquakeCatalogClean::class, 'call'];
    $u->done = [EarthquakeCatalogDone::class, 'call'];
    $u->make_error = [EarthquakeCatalogMakeError::class, 'call'];
    $u->feature_add = [EarthquakeCatalogFeatureAdd::class, 'call'];
    $u->feature_hook = [EarthquakeCatalogFeatureHook::class, 'call'];
    $u->feature_init = [EarthquakeCatalogFeatureInit::class, 'call'];
    $u->fetcher = [EarthquakeCatalogFetcher::class, 'call'];
    $u->make_fetch_def = [EarthquakeCatalogMakeFetchDef::class, 'call'];
    $u->make_context = [EarthquakeCatalogMakeContext::class, 'call'];
    $u->make_options = [EarthquakeCatalogMakeOptions::class, 'call'];
    $u->make_request = [EarthquakeCatalogMakeRequest::class, 'call'];
    $u->make_response = [EarthquakeCatalogMakeResponse::class, 'call'];
    $u->make_result = [EarthquakeCatalogMakeResult::class, 'call'];
    $u->make_point = [EarthquakeCatalogMakePoint::class, 'call'];
    $u->make_spec = [EarthquakeCatalogMakeSpec::class, 'call'];
    $u->make_url = [EarthquakeCatalogMakeUrl::class, 'call'];
    $u->param = [EarthquakeCatalogParam::class, 'call'];
    $u->prepare_auth = [EarthquakeCatalogPrepareAuth::class, 'call'];
    $u->prepare_body = [EarthquakeCatalogPrepareBody::class, 'call'];
    $u->prepare_headers = [EarthquakeCatalogPrepareHeaders::class, 'call'];
    $u->prepare_method = [EarthquakeCatalogPrepareMethod::class, 'call'];
    $u->prepare_params = [EarthquakeCatalogPrepareParams::class, 'call'];
    $u->prepare_path = [EarthquakeCatalogPreparePath::class, 'call'];
    $u->prepare_query = [EarthquakeCatalogPrepareQuery::class, 'call'];
    $u->result_basic = [EarthquakeCatalogResultBasic::class, 'call'];
    $u->result_body = [EarthquakeCatalogResultBody::class, 'call'];
    $u->result_headers = [EarthquakeCatalogResultHeaders::class, 'call'];
    $u->transform_request = [EarthquakeCatalogTransformRequest::class, 'call'];
    $u->transform_response = [EarthquakeCatalogTransformResponse::class, 'call'];
});
