package voxgigearthquakecatalogsdk

import (
	"github.com/voxgig-sdk/earthquake-catalog-sdk/go/core"
	"github.com/voxgig-sdk/earthquake-catalog-sdk/go/entity"
	"github.com/voxgig-sdk/earthquake-catalog-sdk/go/feature"
	_ "github.com/voxgig-sdk/earthquake-catalog-sdk/go/utility"
)

// Type aliases preserve external API.
type EarthquakeCatalogSDK = core.EarthquakeCatalogSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type EarthquakeCatalogEntity = core.EarthquakeCatalogEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type EarthquakeCatalogError = core.EarthquakeCatalogError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewEarthquakeDataEntityFunc = func(client *core.EarthquakeCatalogSDK, entopts map[string]any) core.EarthquakeCatalogEntity {
		return entity.NewEarthquakeDataEntity(client, entopts)
	}
	core.NewServiceInformationEntityFunc = func(client *core.EarthquakeCatalogSDK, entopts map[string]any) core.EarthquakeCatalogEntity {
		return entity.NewServiceInformationEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewEarthquakeCatalogSDK = core.NewEarthquakeCatalogSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
