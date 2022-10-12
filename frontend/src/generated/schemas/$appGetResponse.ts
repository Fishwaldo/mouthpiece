/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $appGetResponse = {
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
        filters: {
            type: 'array',
            contains: {
                properties: {
                    description: {
                        type: 'string',
                        description: `Description of the Filter`,
                        isReadOnly: true,
                        isRequired: true,
                    },
                    id: {
                        type: 'number',
                        description: `ID of the Application`,
                        isReadOnly: true,
                        isRequired: true,
                        format: 'int32',
                    },
                    name: {
                        type: 'string',
                        description: `Name of the Filter`,
                        isReadOnly: true,
                        isRequired: true,
                    },
                    type: {
                        type: 'string',
                        description: `Type of the Filter`,
                        isReadOnly: true,
                        isRequired: true,
                    },
                },
            },
            isReadOnly: true,
        },
        groups: {
            type: 'array',
            contains: {
                properties: {
                    description: {
                        type: 'string',
                        description: `Description of the Filter`,
                        isReadOnly: true,
                        isRequired: true,
                    },
                    id: {
                        type: 'number',
                        description: `ID of the Application`,
                        isReadOnly: true,
                        isRequired: true,
                        format: 'int32',
                    },
                    name: {
                        type: 'string',
                        description: `Name of the Filter`,
                        isReadOnly: true,
                        isRequired: true,
                    },
                },
            },
            isReadOnly: true,
        },
        id: {
            type: 'number',
            description: `ID of the Application`,
            isReadOnly: true,
            isRequired: true,
            format: 'int32',
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
