/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { create_request } from '../models/create_request';
import type { login_request } from '../models/login_request';
import type { logoutResponse } from '../models/logoutResponse';
import type { meResponse } from '../models/meResponse';
import type { passwordLoginResult } from '../models/passwordLoginResult';
import type { patResponse } from '../models/patResponse';
import type { refresh_request } from '../models/refresh_request';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class AuthService {

    /**
     * Logout of Application
     * @returns logoutResponse OK
     * @throws ApiError
     */
    public static logout(): CancelablePromise<logoutResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/auth/logout',
        });
    }

    /**
     * Get information about the current user
     * @returns meResponse OK
     * @throws ApiError
     */
    public static me(): CancelablePromise<meResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/auth/me',
        });
    }

    /**
     * Login to Application with Username/Password
     * @returns passwordLoginResult OK
     * @throws ApiError
     */
    public static login({
        requestBody,
    }: {
        requestBody?: login_request,
    }): CancelablePromise<passwordLoginResult> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/auth/password',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                406: `Not Acceptable`,
                408: `Request Timeout`,
                413: `Request Entity Too Large`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Create a Personal Access Token
     * @returns patResponse OK
     * @throws ApiError
     */
    public static create({
        requestBody,
    }: {
        requestBody?: create_request,
    }): CancelablePromise<patResponse> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/auth/pat',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                406: `Not Acceptable`,
                408: `Request Timeout`,
                413: `Request Entity Too Large`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Refresh a token
     * @returns passwordLoginResult OK
     * @throws ApiError
     */
    public static refresh({
        requestBody,
    }: {
        requestBody?: refresh_request,
    }): CancelablePromise<passwordLoginResult> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/auth/tokenrefresh',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                406: `Not Acceptable`,
                408: `Request Timeout`,
                413: `Request Entity Too Large`,
                422: `Unprocessable Entity`,
            },
        });
    }

}
