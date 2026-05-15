<?php
declare(strict_types=1);

// EarthquakeCatalog SDK utility: feature_add

class EarthquakeCatalogFeatureAdd
{
    public static function call(EarthquakeCatalogContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
