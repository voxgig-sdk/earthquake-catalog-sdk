package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewEarthquakeDataEntityFunc func(client *EarthquakeCatalogSDK, entopts map[string]any) EarthquakeCatalogEntity

var NewServiceInformationEntityFunc func(client *EarthquakeCatalogSDK, entopts map[string]any) EarthquakeCatalogEntity

