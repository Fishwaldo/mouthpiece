/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { msgListResponse } from '../models/msgListResponse';
import type { msgResponse } from '../models/msgResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class MessagesService {

    /**
     * Get A List of Messages
     * @returns msgListResponse OK
     * @throws ApiError
     */
    public static getMessages({
        page,
        size,
        orderBy,
        orderDir,
    }: {
        page?: number,
        size?: number,
        orderBy?: string,
        orderDir?: string,
    }): CancelablePromise<msgListResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/messages',
            query: {
                'page': page,
                'size': size,
                'orderBy': orderBy,
                'orderDir': orderDir,
            },
            errors: {
                400: `Bad Request`,
                422: `Unprocessable Entity`,
            },
        });
    }

    /**
     * Get a Single Message
     * @returns msgResponse OK
     * @throws ApiError
     */
    public static getMessage({
        msgid,
    }: {
        msgid: string,
    }): CancelablePromise<msgResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/messages/{msgid}',
            path: {
                'msgid': msgid,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
                422: `Unprocessable Entity`,
            },
        });
    }

}
