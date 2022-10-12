/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $Post_App_request = {
    properties: {
        Icon: {
            type: 'string',
            description: `Icon of the Application`,
        },
        URL: {
            type: 'string',
            description: `URL of the Application`,
        },
        description: {
            type: 'string',
            description: `Name of the Application`,
            maxLength: 255,
            pattern: '^[a-zA-Z0-9_]+$',
        },
        name: {
            type: 'string',
            description: `Name of the Application`,
            isRequired: true,
            maxLength: 32,
            minLength: 3,
            pattern: '^[a-zA-Z0-9_]+$',
        },
        status: {
            type: 'Enum',
            isRequired: true,
        },
    },
} as const;
