/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type ErrorModel = {
    /**
     * A human-readable explanation specific to this occurrence of the problem.
     */
    detail?: string;
    /**
     * Optional list of individual error details
     */
    errors?: Array<{
        /**
         * Where the error occured, e.g. 'body.items[3].tags' or 'path.thing-id'
         */
        location?: string;
        /**
         * Error message text
         */
        message?: string;
        /**
         * The value at the given location
         */
        value?: any;
    }>;
    /**
     * A URI reference that identifies the specific occurence of the problem.
     */
    instance?: string;
    /**
     * HTTP status code
     */
    status?: number;
    /**
     * A short, human-readable summary of the problem type. This value should not change between occurances of the error.
     */
    title?: string;
    /**
     * A URI reference to human-readable documentation for the error.
     */
    type?: string;
};

