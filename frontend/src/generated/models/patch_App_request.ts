/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type patch_App_request = {
    /**
     * Icon of the Application
     */
    Icon?: string;
    /**
     * URL of the Application
     */
    URL?: string;
    /**
     * Name of the Application
     */
    description?: string;
    /**
     * Name of the Application
     */
    name?: string;
    /**
     * Status of the Application
     */
    status?: patch_App_request.status;
};

export namespace patch_App_request {

    /**
     * Status of the Application
     */
    export enum status {
        ENABLED = 'Enabled',
        DISABLED = 'Disabled',
    }


}
