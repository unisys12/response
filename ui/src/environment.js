import { Environment } from '$houdini'

export default new Environment(async function ({ text, variables = {} }) {
    // send the request to the api
    const result = await this.fetch('/api/graphql', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: text,
            variables,
        }),
    })

    // parse the result as json
    return await result.json()
})