/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $passwordLoginResult = {
    properties: {
        refreshtoken: {
            type: 'string',
            isRequired: true,
        },
        sessiontoken: {
            type: 'string',
            isRequired: true,
        },
        status: {
            type: 'string',
            isRequired: true,
        },
    },
} as const;
