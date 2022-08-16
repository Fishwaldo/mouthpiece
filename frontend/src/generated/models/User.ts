/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type User = {
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
};

