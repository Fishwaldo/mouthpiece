/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type FEConfig = {
    /**
     * An optional URL to a JSON Schema document describing this resource
     */
    $schema?: string;
    /**
     * Provider OAuth Config for Frontend
     */
    oauthproviders: Record<string, {
        /**
         * OAuth Client ID
         */
        clientid: string;
    }>;
};

