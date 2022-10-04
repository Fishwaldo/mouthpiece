/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $CheckerResult = {
    properties: {
        details: {
            type: 'dictionary',
            contains: {
                properties: {
                    error: {
                        type: 'string',
                    },
                    status: {
                        type: 'string',
                        isRequired: true,
                    },
                    timestamp: {
                        type: 'string',
                        format: 'date-time',
                    },
                },
            },
        },
        status: {
            type: 'string',
            isRequired: true,
        },
    },
} as const;
