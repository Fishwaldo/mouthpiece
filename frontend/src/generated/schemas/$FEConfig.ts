/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $FEConfig = {
    properties: {
        $schema: {
            type: 'string',
            description: `An optional URL to a JSON Schema document describing this resource`,
            format: 'uri',
        },
        oauthproviders: {
            type: 'dictionary',
            contains: {
                properties: {
                    clientid: {
                        type: 'string',
                        description: `OAuth Client ID`,
                        isRequired: true,
                    },
                },
            },
            isRequired: true,
        },
    },
} as const;
