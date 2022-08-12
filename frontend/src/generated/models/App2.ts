/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type App2 = {
    /**
     * An optional URL to a JSON Schema document describing this resource
     */
    $schema?: string;
    /**
     * Application Name
     */
    appname: string;
    associatedusers: Array<{
        createdat: string;
        deletedat: {
            time: string;
            valid: boolean;
        };
        email: string;
        firstname: string;
        id: number;
        lastname: string;
        transports?: Array<{
            config: string;
            createdat: string;
            deletedat: {
                time: string;
                valid: boolean;
            };
            id: number;
            transport: string;
            updatedat: string;
        }>;
        updatedat: string;
    }>;
    createdat: string;
    deletedat: {
        time: string;
        valid: boolean;
    };
    /**
     * Description of Application
     */
    description: string;
    filters: Array<{
        createdat: string;
        deletedat: {
            time: string;
            valid: boolean;
        };
        id: number;
        name: string;
        updatedat: string;
    }>;
    /**
     * Icon of Application
     */
    icon: string;
    id: number;
    /**
     * Status of Application
     */
    status: App2.status;
    updatedat: string;
    /**
     * URL of Application
     */
    url: string;
};

export namespace App2 {

    /**
     * Status of Application
     */
    export enum status {
        ENABLED = 'Enabled',
        DISABLED = 'Disabled',
    }


}

