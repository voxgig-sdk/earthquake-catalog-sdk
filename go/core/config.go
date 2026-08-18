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
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
								"parts": []any{
									"query",
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
								"parts": []any{
									"count",
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
								"parts": []any{
									"catalogs",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/contributors",
								"parts": []any{
									"contributors",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
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
								"parts": []any{
									"application.json",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/application.wadl",
								"parts": []any{
									"application.wadl",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/version",
								"parts": []any{
									"version",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
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
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
