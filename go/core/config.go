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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "geometry",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maxAllowed",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "properties",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "type",
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
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "alertlevel",
											"orig": "alertlevel",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "callback",
											"orig": "callback",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "catalog",
											"orig": "catalog",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "contributor",
											"orig": "contributor",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "2014-01-02",
											"kind": "query",
											"name": "endtime",
											"orig": "endtime",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "eventid",
											"orig": "eventid",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "earthquake",
											"kind": "query",
											"name": "eventtype",
											"orig": "eventtype",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "quakeml",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "includeallmagnitude",
											"orig": "includeallmagnitude",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "includeallorigin",
											"orig": "includeallorigin",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "includearrival",
											"orig": "includearrival",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": "false",
											"kind": "query",
											"name": "includedeleted",
											"orig": "includedeleted",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "includesuperseded",
											"orig": "includesuperseded",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "jsonerror",
											"orig": "jsonerror",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "kmlanimated",
											"orig": "kmlanimated",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": "age",
											"kind": "query",
											"name": "kmlcolorby",
											"orig": "kmlcolorby",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "latitude",
											"orig": "latitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "longitude",
											"orig": "longitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "magnitudetype",
											"orig": "magnitudetype",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "maxcdi",
											"orig": "maxcdi",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 1000,
											"kind": "query",
											"name": "maxdepth",
											"orig": "maxdepth",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "maxgap",
											"orig": "maxgap",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 90,
											"kind": "query",
											"name": "maxlatitude",
											"orig": "maxlatitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 180,
											"kind": "query",
											"name": "maxlongitude",
											"orig": "maxlongitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "maxmagnitude",
											"orig": "maxmagnitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "maxmmi",
											"orig": "maxmmi",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 180,
											"kind": "query",
											"name": "maxradius",
											"orig": "maxradius",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 20001.6,
											"kind": "query",
											"name": "maxradiuskm",
											"orig": "maxradiuskm",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "maxsig",
											"orig": "maxsig",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "mincdi",
											"orig": "mincdi",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": -100,
											"kind": "query",
											"name": "mindepth",
											"orig": "mindepth",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "minfelt",
											"orig": "minfelt",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "mingap",
											"orig": "mingap",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": -90,
											"kind": "query",
											"name": "minlatitude",
											"orig": "minlatitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": -180,
											"kind": "query",
											"name": "minlongitude",
											"orig": "minlongitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 5,
											"kind": "query",
											"name": "minmagnitude",
											"orig": "minmagnitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "minsig",
											"orig": "minsig",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 204,
											"kind": "query",
											"name": "nodata",
											"orig": "nodata",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": "time",
											"kind": "query",
											"name": "orderby",
											"orig": "orderby",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "productcode",
											"orig": "productcode",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "moment-tensor",
											"kind": "query",
											"name": "producttype",
											"orig": "producttype",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "all",
											"kind": "query",
											"name": "reviewstatus",
											"orig": "reviewstatus",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "2014-01-01",
											"kind": "query",
											"name": "starttime",
											"orig": "starttime",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "updatedafter",
											"orig": "updatedafter",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/query",
								"segments": []any{
									map[string]any{
										"lit": "query",
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"query",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "alertlevel",
											"orig": "alertlevel",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "catalog",
											"orig": "catalog",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "contributor",
											"orig": "contributor",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "endtime",
											"orig": "endtime",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "eventtype",
											"orig": "eventtype",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "text",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "latitude",
											"orig": "latitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "longitude",
											"orig": "longitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "magnitudetype",
											"orig": "magnitudetype",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 1000,
											"kind": "query",
											"name": "maxdepth",
											"orig": "maxdepth",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 90,
											"kind": "query",
											"name": "maxlatitude",
											"orig": "maxlatitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 180,
											"kind": "query",
											"name": "maxlongitude",
											"orig": "maxlongitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "maxmagnitude",
											"orig": "maxmagnitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 180,
											"kind": "query",
											"name": "maxradius",
											"orig": "maxradius",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 20001.6,
											"kind": "query",
											"name": "maxradiuskm",
											"orig": "maxradiuskm",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": -100,
											"kind": "query",
											"name": "mindepth",
											"orig": "mindepth",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": -90,
											"kind": "query",
											"name": "minlatitude",
											"orig": "minlatitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": -180,
											"kind": "query",
											"name": "minlongitude",
											"orig": "minlongitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "minmagnitude",
											"orig": "minmagnitude",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": "all",
											"kind": "query",
											"name": "reviewstatus",
											"orig": "reviewstatus",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "starttime",
											"orig": "starttime",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "updatedafter",
											"orig": "updatedafter",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/count",
								"segments": []any{
									map[string]any{
										"lit": "count",
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"count",
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
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/catalogs",
								"segments": []any{
									map[string]any{
										"lit": "catalogs",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"catalogs",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/contributors",
								"segments": []any{
									map[string]any{
										"lit": "contributors",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"contributors",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/application.json",
								"segments": []any{
									map[string]any{
										"lit": "application.json",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"application.json",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/application.wadl",
								"segments": []any{
									map[string]any{
										"lit": "application.wadl",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"application.wadl",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/version",
								"segments": []any{
									map[string]any{
										"lit": "version",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"version",
								},
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
