import { EarthquakeCatalogEntityBase } from '../EarthquakeCatalogEntityBase';
import type { EarthquakeCatalogSDK } from '../EarthquakeCatalogSDK';
import type { Control } from '../types';
import type { ServiceInformation, ServiceInformationLoadMatch, ServiceInformationListMatch } from '../EarthquakeCatalogTypes';
declare class ServiceInformationEntity extends EarthquakeCatalogEntityBase<ServiceInformation> {
    constructor(client: EarthquakeCatalogSDK, entopts: any);
    make(this: ServiceInformationEntity): ServiceInformationEntity;
    load(this: any, reqmatch?: ServiceInformationLoadMatch, ctrl?: Control): Promise<ServiceInformationEntity>;
    list(this: any, reqmatch?: ServiceInformationListMatch, ctrl?: Control): Promise<ServiceInformationEntity[]>;
}
export { ServiceInformationEntity };
