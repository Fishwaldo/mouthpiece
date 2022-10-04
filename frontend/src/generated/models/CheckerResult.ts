/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type CheckerResult = {
    details?: Record<string, {
        error?: string;
        status: string;
        timestamp?: string;
    }>;
    status: string;
};

