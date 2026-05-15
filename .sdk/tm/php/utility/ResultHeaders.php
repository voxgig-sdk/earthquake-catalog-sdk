<?php
declare(strict_types=1);

// EarthquakeCatalog SDK utility: result_headers

class EarthquakeCatalogResultHeaders
{
    public static function call(EarthquakeCatalogContext $ctx): ?EarthquakeCatalogResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
