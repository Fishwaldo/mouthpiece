/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type CheckerResult = {
    /**
     * An optional URL to a JSON Schema document describing this resource
     */
    $schema?: string;
    details?: Record<string, {
        error?: string;
        status: string;
        timestamp?: string;
    }>;
    status: string;
};

