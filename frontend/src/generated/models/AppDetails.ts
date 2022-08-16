/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type AppDetails = {
    /**
     * An optional URL to a JSON Schema document describing this resource
     */
    $schema?: string;
    /**
     * Application Name
     */
    appname: string;
    /**
     * Description of Application
     */
    description: string;
    /**
     * Icon of Application
     */
    icon: string;
    /**
     * Status of Application
     */
    status: AppDetails.status;
    /**
     * URL of Application
     */
    url: string;
};

export namespace AppDetails {

    /**
     * Status of Application
     */
    export enum status {
        ENABLED = 'Enabled',
        DISABLED = 'Disabled',
    }


}

