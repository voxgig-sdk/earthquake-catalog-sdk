<?php
declare(strict_types=1);

// EarthquakeCatalog SDK base feature

class EarthquakeCatalogBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(EarthquakeCatalogContext $ctx, array $options): void {}
    public function PostConstruct(EarthquakeCatalogContext $ctx): void {}
    public function PostConstructEntity(EarthquakeCatalogContext $ctx): void {}
    public function SetData(EarthquakeCatalogContext $ctx): void {}
    public function GetData(EarthquakeCatalogContext $ctx): void {}
    public function GetMatch(EarthquakeCatalogContext $ctx): void {}
    public function SetMatch(EarthquakeCatalogContext $ctx): void {}
    public function PrePoint(EarthquakeCatalogContext $ctx): void {}
    public function PreSpec(EarthquakeCatalogContext $ctx): void {}
    public function PreRequest(EarthquakeCatalogContext $ctx): void {}
    public function PreResponse(EarthquakeCatalogContext $ctx): void {}
    public function PreResult(EarthquakeCatalogContext $ctx): void {}
    public function PreDone(EarthquakeCatalogContext $ctx): void {}
    public function PreUnexpected(EarthquakeCatalogContext $ctx): void {}
}
