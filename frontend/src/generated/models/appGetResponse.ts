/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type appGetResponse = {
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
     * Filters of the Application
     */
    readonly filters?: Array<{
        /**
         * Description of the Filter
         */
        readonly description: string;
        /**
         * ID of the Application
         */
        readonly id: number;
        /**
         * Name of the Filter
         */
        readonly name: string;
        /**
         * Type of the Filter
         */
        readonly type: string;
    }>;
    /**
     * Groups of the Application
     */
    readonly groups?: Array<{
        /**
         * Description of the Filter
         */
        readonly description: string;
        /**
         * ID of the Application
         */
        readonly id: number;
        /**
         * Name of the Filter
         */
        readonly name: string;
    }>;
    /**
     * ID of the Application
     */
    readonly id: number;
    /**
     * Name of the Application
     */
    name: string;
    /**
     * Status of the Application
     */
    status: appGetResponse.status;
};

export namespace appGetResponse {

    /**
     * Status of the Application
     */
    export enum status {
        ENABLED = 'Enabled',
        DISABLED = 'Disabled',
    }


}

