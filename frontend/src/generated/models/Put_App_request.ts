/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type Put_App_request = {
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
    name: string;
    /**
     * Status of the Application
     */
    status: Put_App_request.status;
};

export namespace Put_App_request {

    /**
     * Status of the Application
     */
    export enum status {
        ENABLED = 'Enabled',
        DISABLED = 'Disabled',
    }


}

