
import { Context } from './Context'


class EarthquakeCatalogError extends Error {

  isEarthquakeCatalogError = true

  sdk = 'EarthquakeCatalog'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  EarthquakeCatalogError
}

