import { Context } from './Context';
declare class EarthquakeCatalogError extends Error {
    isEarthquakeCatalogError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { EarthquakeCatalogError };
