/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { msgListResponseList } from '../models/msgListResponseList';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class MessagesService {

    /**
     * Get A List of Messages
     * @returns msgListResponseList OK
     * @throws ApiError
     */
    public static getMessages(): CancelablePromise<msgListResponseList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/messages',
        });
    }

}
