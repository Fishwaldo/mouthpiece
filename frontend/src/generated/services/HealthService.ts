/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CheckerResult } from '../models/CheckerResult';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class HealthService {

    /**
     * Get Health of the Service
     * @returns CheckerResult OK
     * @throws ApiError
     */
    public static getHealth(): CancelablePromise<CheckerResult> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/health',
            errors: {
                500: `Internal Server Error`,
            },
        });
    }

}
