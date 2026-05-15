<?php
declare(strict_types=1);

// EarthquakeCatalog SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class EarthquakeCatalogMakeContext
{
    public static function call(array $ctxmap, ?EarthquakeCatalogContext $basectx): EarthquakeCatalogContext
    {
        return new EarthquakeCatalogContext($ctxmap, $basectx);
    }
}
