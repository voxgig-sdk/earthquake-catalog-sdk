<?php
declare(strict_types=1);

// EarthquakeCatalog SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class EarthquakeCatalogFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new EarthquakeCatalogBaseFeature();
            case "test":
                return new EarthquakeCatalogTestFeature();
            default:
                return new EarthquakeCatalogBaseFeature();
        }
    }
}
