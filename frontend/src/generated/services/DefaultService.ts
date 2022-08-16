/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { App2 } from '../models/App2';
import type { AppDetails } from '../models/AppDetails';
import type { AppList } from '../models/AppList';
import type { CheckerResult } from '../models/CheckerResult';
import type { FEConfig } from '../models/FEConfig';
import type { MessageResult } from '../models/MessageResult';
import type { post_message_request } from '../models/post_message_request';
import type { stringList } from '../models/stringList';
import type { TransportConfig } from '../models/TransportConfig';
import type { UserList } from '../models/UserList';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class DefaultService {

    /**
     * Get Config of the Service
     * @returns FEConfig OK
     * @throws ApiError
     */
    public static getConfig(): CancelablePromise<FEConfig> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/config/frontend',
        });
    }

    /**
     * Get Health of the Service
     * @returns CheckerResult OK
     * @throws ApiError
     */
    public static getHealth(): CancelablePromise<CheckerResult> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/health',
            errors: {
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Get A List of Applications
     * @returns AppList OK
     * @throws ApiError
     */
    public static getApps(): CancelablePromise<AppList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/apps/',
        });
    }

    /**
     * Create a Application
     * @returns App2 OK
     * @throws ApiError
     */
    public static createApp({
        requestBody,
    }: {
        requestBody?: AppDetails,
    }): CancelablePromise<App2> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/v1/apps/',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                406: `Not Acceptable`,
                408: `Request Timeout`,
                413: `Request Entity Too Large`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Post Message to the Service
     * @returns MessageResult OK
     * @throws ApiError
     */
    public static postMessage({
        application,
        requestBody,
    }: {
        /**
         * Application Name
         */
        application: string,
        requestBody?: post_message_request,
    }): CancelablePromise<MessageResult> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/v1/message/{application}',
            path: {
                'application': application,
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

    /**
     * Get A List of Transports
     * @returns stringList OK
     * @throws ApiError
     */
    public static getTransports(): CancelablePromise<stringList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/transports/',
        });
    }

    /**
     * Get A List of Users
     * @returns UserList OK
     * @throws ApiError
     */
    public static getUsers(): CancelablePromise<UserList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/users/',
        });
    }

    /**
     * Get A List of Transports for a User
     * @returns stringList OK
     * @throws ApiError
     */
    public static getUserTransports({
        userid,
    }: {
        userid: number,
    }): CancelablePromise<stringList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/users/{userid}/transports/',
            path: {
                'userid': userid,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Get Details for a Transport for a User
     * @returns TransportConfig OK
     * @throws ApiError
     */
    public static getUserTransportDetails({
        transportid,
        userid,
    }: {
        transportid: string,
        userid: number,
    }): CancelablePromise<TransportConfig> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/v1/users/{userid}/transports/{transportid}/',
            path: {
                'transportid': transportid,
                'userid': userid,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                422: `Unprocessable Entity`,
            },
        });
    }

}
