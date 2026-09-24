package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "EarthquakeCatalog",
			"slug": "earthquake-catalog",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://earthquake.usgs.gov/fdsnws/event/1",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"earthquake_data": map[string]any{},
				"service_information": map[string]any{},
			},
		},
		"entity": map[string]any{
			"earthquake_data": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "count",
						"title": "Count",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "geometry",
						"title": "Geometry",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maxAllowed",
						"title": "Max Allowed",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "properties",
						"title": "Properties",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "earthquake_data",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/query",
								"segments": []any{
									map[string]any{
										"lit": "query",
									},
								},
								"parts": []any{
									"query",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "alertlevel",
											"orig": "alertlevel",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "callback",
											"orig": "callback",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "catalog",
											"orig": "catalog",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "contributor",
											"orig": "contributor",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "endtime",
											"orig": "endtime",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2014-01-02",
										},
										map[string]any{
											"name": "eventid",
											"orig": "eventid",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "eventtype",
											"orig": "eventtype",
											"type": "`$STRING`",
											"kind": "query",
											"example": "earthquake",
										},
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "quakeml",
										},
										map[string]any{
											"name": "includeallmagnitude",
											"orig": "includeallmagnitude",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": false,
										},
										map[string]any{
											"name": "includeallorigin",
											"orig": "includeallorigin",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": false,
										},
										map[string]any{
											"name": "includearrival",
											"orig": "includearrival",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": false,
										},
										map[string]any{
											"name": "includedeleted",
											"orig": "includedeleted",
											"type": "`$STRING`",
											"kind": "query",
											"example": "false",
										},
										map[string]any{
											"name": "includesuperseded",
											"orig": "includesuperseded",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": false,
										},
										map[string]any{
											"name": "jsonerror",
											"orig": "jsonerror",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": false,
										},
										map[string]any{
											"name": "kmlanimated",
											"orig": "kmlanimated",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"example": false,
										},
										map[string]any{
											"name": "kmlcolorby",
											"orig": "kmlcolorby",
											"type": "`$STRING`",
											"kind": "query",
											"example": "age",
										},
										map[string]any{
											"name": "latitude",
											"orig": "latitude",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "longitude",
											"orig": "longitude",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "magnitudetype",
											"orig": "magnitudetype",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "maxcdi",
											"orig": "maxcdi",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "maxdepth",
											"orig": "maxdepth",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 1000,
										},
										map[string]any{
											"name": "maxgap",
											"orig": "maxgap",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "maxlatitude",
											"orig": "maxlatitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 90,
										},
										map[string]any{
											"name": "maxlongitude",
											"orig": "maxlongitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 180,
										},
										map[string]any{
											"name": "maxmagnitude",
											"orig": "maxmagnitude",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "maxmmi",
											"orig": "maxmmi",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "maxradius",
											"orig": "maxradius",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 180,
										},
										map[string]any{
											"name": "maxradiuskm",
											"orig": "maxradiuskm",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 20001.6,
										},
										map[string]any{
											"name": "maxsig",
											"orig": "maxsig",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "mincdi",
											"orig": "mincdi",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "mindepth",
											"orig": "mindepth",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": -100,
										},
										map[string]any{
											"name": "minfelt",
											"orig": "minfelt",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "mingap",
											"orig": "mingap",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "minlatitude",
											"orig": "minlatitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": -90,
										},
										map[string]any{
											"name": "minlongitude",
											"orig": "minlongitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": -180,
										},
										map[string]any{
											"name": "minmagnitude",
											"orig": "minmagnitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 5,
										},
										map[string]any{
											"name": "minsig",
											"orig": "minsig",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "nodata",
											"orig": "nodata",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 204,
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
											"example": 1,
										},
										map[string]any{
											"name": "orderby",
											"orig": "orderby",
											"type": "`$STRING`",
											"kind": "query",
											"example": "time",
										},
										map[string]any{
											"name": "productcode",
											"orig": "productcode",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "producttype",
											"orig": "producttype",
											"type": "`$STRING`",
											"kind": "query",
											"example": "moment-tensor",
										},
										map[string]any{
											"name": "reviewstatus",
											"orig": "reviewstatus",
											"type": "`$STRING`",
											"kind": "query",
											"example": "all",
										},
										map[string]any{
											"name": "starttime",
											"orig": "starttime",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2014-01-01",
										},
										map[string]any{
											"name": "updatedafter",
											"orig": "updatedafter",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"alertlevel",
										"callback",
										"catalog",
										"contributor",
										"endtime",
										"eventid",
										"eventtype",
										"format",
										"includeallmagnitude",
										"includeallorigin",
										"includearrival",
										"includedeleted",
										"includesuperseded",
										"jsonerror",
										"kmlanimated",
										"kmlcolorby",
										"latitude",
										"limit",
										"longitude",
										"magnitudetype",
										"maxcdi",
										"maxdepth",
										"maxgap",
										"maxlatitude",
										"maxlongitude",
										"maxmagnitude",
										"maxmmi",
										"maxradius",
										"maxradiuskm",
										"maxsig",
										"mincdi",
										"mindepth",
										"minfelt",
										"mingap",
										"minlatitude",
										"minlongitude",
										"minmagnitude",
										"minsig",
										"nodata",
										"offset",
										"orderby",
										"productcode",
										"producttype",
										"reviewstatus",
										"starttime",
										"updatedafter",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/count",
								"segments": []any{
									map[string]any{
										"lit": "count",
									},
								},
								"parts": []any{
									"count",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "alertlevel",
											"orig": "alertlevel",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "catalog",
											"orig": "catalog",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "contributor",
											"orig": "contributor",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "endtime",
											"orig": "endtime",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "eventtype",
											"orig": "eventtype",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"example": "text",
										},
										map[string]any{
											"name": "latitude",
											"orig": "latitude",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "longitude",
											"orig": "longitude",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "magnitudetype",
											"orig": "magnitudetype",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "maxdepth",
											"orig": "maxdepth",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 1000,
										},
										map[string]any{
											"name": "maxlatitude",
											"orig": "maxlatitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 90,
										},
										map[string]any{
											"name": "maxlongitude",
											"orig": "maxlongitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 180,
										},
										map[string]any{
											"name": "maxmagnitude",
											"orig": "maxmagnitude",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "maxradius",
											"orig": "maxradius",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 180,
										},
										map[string]any{
											"name": "maxradiuskm",
											"orig": "maxradiuskm",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": 20001.6,
										},
										map[string]any{
											"name": "mindepth",
											"orig": "mindepth",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": -100,
										},
										map[string]any{
											"name": "minlatitude",
											"orig": "minlatitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": -90,
										},
										map[string]any{
											"name": "minlongitude",
											"orig": "minlongitude",
											"type": "`$NUMBER`",
											"kind": "query",
											"example": -180,
										},
										map[string]any{
											"name": "minmagnitude",
											"orig": "minmagnitude",
											"type": "`$NUMBER`",
											"kind": "query",
										},
										map[string]any{
											"name": "reviewstatus",
											"orig": "reviewstatus",
											"type": "`$STRING`",
											"kind": "query",
											"example": "all",
										},
										map[string]any{
											"name": "starttime",
											"orig": "starttime",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updatedafter",
											"orig": "updatedafter",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"alertlevel",
										"catalog",
										"contributor",
										"endtime",
										"eventtype",
										"format",
										"latitude",
										"longitude",
										"magnitudetype",
										"maxdepth",
										"maxlatitude",
										"maxlongitude",
										"maxmagnitude",
										"maxradius",
										"maxradiuskm",
										"mindepth",
										"minlatitude",
										"minlongitude",
										"minmagnitude",
										"reviewstatus",
										"starttime",
										"updatedafter",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"service_information": map[string]any{
				"fields": []any{},
				"name": "service_information",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/catalogs",
								"segments": []any{
									map[string]any{
										"lit": "catalogs",
									},
								},
								"parts": []any{
									"catalogs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contributors",
								"segments": []any{
									map[string]any{
										"lit": "contributors",
									},
								},
								"parts": []any{
									"contributors",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/application.json",
								"segments": []any{
									map[string]any{
										"lit": "application.json",
									},
								},
								"parts": []any{
									"application.json",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/application.wadl",
								"segments": []any{
									map[string]any{
										"lit": "application.wadl",
									},
								},
								"parts": []any{
									"application.wadl",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/version",
								"segments": []any{
									map[string]any{
										"lit": "version",
									},
								},
								"parts": []any{
									"version",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
