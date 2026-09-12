import { EarthquakeCatalogEntityBase } from '../EarthquakeCatalogEntityBase';
import type { EarthquakeCatalogSDK } from '../EarthquakeCatalogSDK';
import type { Control } from '../types';
import type { EarthquakeData, EarthquakeDataLoadMatch, EarthquakeDataListMatch } from '../EarthquakeCatalogTypes';
declare class EarthquakeDataEntity extends EarthquakeCatalogEntityBase<EarthquakeData> {
    constructor(client: EarthquakeCatalogSDK, entopts: any);
    make(this: EarthquakeDataEntity): EarthquakeDataEntity;
    load(this: any, reqmatch?: EarthquakeDataLoadMatch, ctrl?: Control): Promise<EarthquakeDataEntity>;
    list(this: any, reqmatch?: EarthquakeDataListMatch, ctrl?: Control): Promise<EarthquakeDataEntity[]>;
}
export { EarthquakeDataEntity };
