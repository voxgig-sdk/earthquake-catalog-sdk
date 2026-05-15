<?php
declare(strict_types=1);

// EarthquakeCatalog SDK utility: feature_hook

class EarthquakeCatalogFeatureHook
{
    public static function call(EarthquakeCatalogContext $ctx, string $name): void
    {
        if (!$ctx->client) {
            return;
        }
        $features = $ctx->client->features ?? null;
        if (!$features) {
            return;
        }
        foreach ($features as $f) {
            if (method_exists($f, $name)) {
                $f->$name($ctx);
            }
        }
    }
}
