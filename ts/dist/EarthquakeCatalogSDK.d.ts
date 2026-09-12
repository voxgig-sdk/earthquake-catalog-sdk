import { EarthquakeDataEntity } from './entity/EarthquakeDataEntity';
import { ServiceInformationEntity } from './entity/ServiceInformationEntity';
export type * from './EarthquakeCatalogTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { EarthquakeCatalogEntityBase } from './EarthquakeCatalogEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class EarthquakeCatalogSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    EarthquakeData(entopts?: Record<string, any>): EarthquakeDataEntity;
    ServiceInformation(entopts?: Record<string, any>): ServiceInformationEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): EarthquakeCatalogSDK;
    tester(testopts?: any, sdkopts?: any): EarthquakeCatalogSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof EarthquakeCatalogSDK;
export { stdutil, config, BaseFeature, EarthquakeCatalogEntityBase, EarthquakeCatalogSDK, SDK, };
