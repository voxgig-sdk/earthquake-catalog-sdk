<?php
declare(strict_types=1);

// EarthquakeCatalog SDK utility: prepare_body

class EarthquakeCatalogPrepareBody
{
    public static function call(EarthquakeCatalogContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
