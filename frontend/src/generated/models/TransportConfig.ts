/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type TransportConfig = {
    /**
     * An optional URL to a JSON Schema document describing this resource
     */
    $schema?: string;
    config: string;
    createdat: string;
    deletedat: {
        time: string;
        valid: boolean;
    };
    id: number;
    transport: string;
    updatedat: string;
};

