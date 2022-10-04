/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $msgListResponse = {
    properties: {
        appid: {
            type: 'number',
            isRequired: true,
            format: 'int32',
        },
        fields: {
            type: 'dictionary',
            contains: {
                type: 'string',
            },
            isRequired: true,
        },
        id: {
            type: 'string',
            isRequired: true,
        },
        message: {
            type: 'string',
            isRequired: true,
        },
        severity: {
            type: 'number',
            isRequired: true,
            format: 'int32',
        },
        shortmsg: {
            type: 'string',
            isRequired: true,
        },
        timestamp: {
            type: 'string',
            isRequired: true,
            format: 'date-time',
        },
        topic: {
            type: 'string',
            isRequired: true,
        },
    },
} as const;
