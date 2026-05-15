<?php
declare(strict_types=1);

// EarthquakeCatalog SDK utility: result_body

class EarthquakeCatalogResultBody
{
    public static function call(EarthquakeCatalogContext $ctx): ?EarthquakeCatalogResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
