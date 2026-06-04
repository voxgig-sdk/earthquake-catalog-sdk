# EarthquakeCatalog SDK

Search USGS earthquake events by time, location, magnitude, and depth via the FDSN Event Web Service

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About Earthquake Catalog API

The Earthquake Catalog API is operated by the [U.S. Geological Survey](https://earthquake.usgs.gov/) and implements the [FDSN Event Web Service](https://www.fdsn.org/webservices/) specification. It exposes the USGS Comprehensive Earthquake Catalog (ComCat), aggregating events contributed by USGS regional networks and partner agencies.

What you get from the API:

- `query` — search events by `starttime`/`endtime`, bounding box (`minlatitude`, `maxlatitude`, `minlongitude`, `maxlongitude`), circle (`latitude`, `longitude`, `maxradiuskm`), `minmagnitude`/`maxmagnitude`, `mindepth`/`maxdepth`, `eventid`, `eventtype`, `catalog`, `contributor`, plus `limit` (up to 20,000) and `offset` for pagination
- Output formats: QuakeML (default), GeoJSON (with optional JSONP), CSV, KML, XML, and plain text
- `count` — return the number of matching events without fetching them
- `catalogs` and `contributors` — enumerate the available source catalogs and contributing agencies
- `application.json` / `application.wadl` — machine-readable parameter definitions
- `version` — current service version

No authentication is required and CORS is enabled, so the service can be called from a browser. The default time window is the last 30 days; results are capped at 20,000 events per request, so wide queries should be split or paginated using `offset`.

## Try it

**TypeScript**
```bash
npm install earthquake-catalog
```

**Python**
```bash
pip install earthquake-catalog-sdk
```

**PHP**
```bash
composer require voxgig/earthquake-catalog-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/earthquake-catalog-sdk/go
```

**Ruby**
```bash
gem install earthquake-catalog-sdk
```

**Lua**
```bash
luarocks install earthquake-catalog-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { EarthquakeCatalogSDK } from 'earthquake-catalog'

const client = new EarthquakeCatalogSDK({})

// List all earthquakedatas
const earthquakedatas = await client.EarthquakeData().list()
```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o earthquake-catalog-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "earthquake-catalog": {
      "command": "/abs/path/to/earthquake-catalog-mcp"
    }
  }
}
```

## Entities

The API exposes 2 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **EarthquakeData** | Earthquake event records returned by `query` (and counted by `count`) with parameters for time, location, magnitude, depth, and catalog filters; available as QuakeML, GeoJSON, CSV, KML, XML, or text. | `/query` |
| **ServiceInformation** | Service metadata endpoints including `version`, `catalogs`, `contributors`, `application.json`, and `application.wadl` describing supported parameter values and the WADL interface. | `/catalogs` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from earthquakecatalog_sdk import EarthquakeCatalogSDK

client = EarthquakeCatalogSDK({})

# List all earthquakedatas
earthquakedatas, err = client.EarthquakeData(None).list(None, None)

# Load a specific earthquakedata
earthquakedata, err = client.EarthquakeData(None).load(
    {"id": "example_id"}, None
)
```

### PHP

```php
<?php
require_once 'earthquakecatalog_sdk.php';

$client = new EarthquakeCatalogSDK([]);

// List all earthquakedatas
[$earthquakedatas, $err] = $client->EarthquakeData(null)->list(null, null);

// Load a specific earthquakedata
[$earthquakedata, $err] = $client->EarthquakeData(null)->load(
    ["id" => "example_id"], null
);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/earthquake-catalog-sdk/go"

client := sdk.NewEarthquakeCatalogSDK(map[string]any{})

// List all earthquakedatas
earthquakedatas, err := client.EarthquakeData(nil).List(nil, nil)
```

### Ruby

```ruby
require_relative "EarthquakeCatalog_sdk"

client = EarthquakeCatalogSDK.new({})

# List all earthquakedatas
earthquakedatas, err = client.EarthquakeData(nil).list(nil, nil)

# Load a specific earthquakedata
earthquakedata, err = client.EarthquakeData(nil).load(
  { "id" => "example_id" }, nil
)
```

### Lua

```lua
local sdk = require("earthquake-catalog_sdk")

local client = sdk.new({})

-- List all earthquakedatas
local earthquakedatas, err = client:EarthquakeData(nil):list(nil, nil)

-- Load a specific earthquakedata
local earthquakedata, err = client:EarthquakeData(nil):load(
  { id = "example_id" }, nil
)
```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = EarthquakeCatalogSDK.test()
const result = await client.EarthquakeData().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = EarthquakeCatalogSDK.test(None, None)
result, err = client.EarthquakeData(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = EarthquakeCatalogSDK::test(null, null);
[$result, $err] = $client->EarthquakeData(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.EarthquakeData(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = EarthquakeCatalogSDK.test(nil, nil)
result, err = client.EarthquakeData(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:EarthquakeData(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the Earthquake Catalog API

- Upstream: [https://earthquake.usgs.gov/fdsnws/event/1/](https://earthquake.usgs.gov/fdsnws/event/1/)

- Data are produced by the U.S. Geological Survey and are generally in the public domain in the United States
- Attribution to USGS is appreciated when redistributing or visualizing the data
- The service follows the FDSN Event Web Service specification; consult USGS terms for any agency-contributed catalogs
- No API key or registration is required

---

Generated from the Earthquake Catalog API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
