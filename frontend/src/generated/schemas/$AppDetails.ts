/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $AppDetails = {
    properties: {
        $schema: {
            type: 'string',
            description: `An optional URL to a JSON Schema document describing this resource`,
            format: 'uri',
        },
        appname: {
            type: 'string',
            description: `Application Name`,
            isRequired: true,
            pattern: '^[a-z0-9]+$',
        },
        description: {
            type: 'string',
            description: `Description of Application`,
            isRequired: true,
        },
        icon: {
            type: 'string',
            description: `Icon of Application`,
            isRequired: true,
        },
        status: {
            type: 'Enum',
            isRequired: true,
        },
        url: {
            type: 'string',
            description: `URL of Application`,
            isRequired: true,
        },
    },
} as const;
