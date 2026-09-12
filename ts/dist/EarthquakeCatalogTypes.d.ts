export interface EarthquakeData {
    count?: number;
    geometry?: Record<string, any>;
    id?: string;
    maxAllowed?: number;
    properties?: Record<string, any>;
    type?: string;
}
export interface EarthquakeDataLoadMatch {
    alertlevel?: string;
    catalog?: string;
    contributor?: string;
    endtime?: string;
    eventtype?: string;
    format?: string;
    latitude?: number;
    longitude?: number;
    magnitudetype?: string;
    maxdepth?: number;
    maxlatitude?: number;
    maxlongitude?: number;
    maxmagnitude?: number;
    maxradius?: number;
    maxradiuskm?: number;
    mindepth?: number;
    minlatitude?: number;
    minlongitude?: number;
    minmagnitude?: number;
    reviewstatus?: string;
    starttime?: string;
    updatedafter?: string;
}
export interface EarthquakeDataListMatch {
    alertlevel?: string;
    callback?: string;
    catalog?: string;
    contributor?: string;
    endtime?: string;
    eventid?: string;
    eventtype?: string;
    format?: string;
    includeallmagnitude?: boolean;
    includeallorigin?: boolean;
    includearrival?: boolean;
    includedeleted?: string;
    includesuperseded?: boolean;
    jsonerror?: boolean;
    kmlanimated?: boolean;
    kmlcolorby?: string;
    latitude?: number;
    limit?: number;
    longitude?: number;
    magnitudetype?: string;
    maxcdi?: number;
    maxdepth?: number;
    maxgap?: number;
    maxlatitude?: number;
    maxlongitude?: number;
    maxmagnitude?: number;
    maxmmi?: number;
    maxradius?: number;
    maxradiuskm?: number;
    maxsig?: number;
    mincdi?: number;
    mindepth?: number;
    minfelt?: number;
    mingap?: number;
    minlatitude?: number;
    minlongitude?: number;
    minmagnitude?: number;
    minsig?: number;
    nodata?: number;
    offset?: number;
    orderby?: string;
    productcode?: string;
    producttype?: string;
    reviewstatus?: string;
    starttime?: string;
    updatedafter?: string;
}
export interface ServiceInformation {
}
export interface ServiceInformationLoadMatch {
}
export interface ServiceInformationListMatch {
}
