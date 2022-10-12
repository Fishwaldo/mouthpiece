/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { appGetResponse } from '../models/appGetResponse';
import type { appListResponseList } from '../models/appListResponseList';
import type { patch_App_request2 } from '../models/patch_App_request2';
import type { Post_App_request } from '../models/Post_App_request';
import type { Put_App_request } from '../models/Put_App_request';

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

    /**
     * Create a App
     * @returns appGetResponse OK
     * @throws ApiError
     */
    public static postApp({
        requestBody,
    }: {
        requestBody?: Post_App_request,
    }): CancelablePromise<appGetResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/apps/new',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                408: `Request Timeout`,
                413: `Request Entity Too Large`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Get App Details
     * @returns appGetResponse OK
     * @throws ApiError
     */
    public static getApp({
        id,
    }: {
        id: number,
    }): CancelablePromise<appGetResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/apps/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Patch App
     * Partial update operation supporting both JSON Merge Patch & JSON Patch updates.
     * @returns appGetResponse OK
     * @throws ApiError
     */
    public static patchApp({
        id,
        requestBody,
    }: {
        /**
         * ID of the Application
         */
        id: number,
        requestBody?: patch_App_request2,
    }): CancelablePromise<appGetResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/api/apps/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json-patch+json',
            errors: {
                304: `Not Modified`,
                400: `Bad Request`,
                404: `Not Found`,
                408: `Request Timeout`,
                413: `Request Entity Too Large`,
                415: `Unsupported Media Type`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Update a App
     * @returns appGetResponse OK
     * @throws ApiError
     */
    public static putApp({
        id,
        requestBody,
    }: {
        /**
         * ID of the Application
         */
        id: number,
        requestBody?: Put_App_request,
    }): CancelablePromise<appGetResponse> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/apps/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                408: `Request Timeout`,
                413: `Request Entity Too Large`,
                422: `Unprocessable Entity`,
            },
        });
    }

}
