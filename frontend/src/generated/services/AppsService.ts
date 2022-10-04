/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { appListResponseList } from '../models/appListResponseList';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class AppsService {

    /**
     * Get A List of Applications
     * @returns appListResponseList OK
     * @throws ApiError
     */
    public static getApps(): CancelablePromise<appListResponseList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/apps',
        });
    }

}
