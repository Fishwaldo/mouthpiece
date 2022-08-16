/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type App = {
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
    status: App.status;
    updatedat: string;
    /**
     * URL of Application
     */
    url: string;
};

export namespace App {

    /**
     * Status of Application
     */
    export enum status {
        ENABLED = 'Enabled',
        DISABLED = 'Disabled',
    }


}

