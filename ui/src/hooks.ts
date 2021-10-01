import type { GetSession, Request } from '@sveltejs/kit'

export const getSession: GetSession = async function (req: Request) {
    const res = await fetch(`http://${req.host}/auth/user`, {
        headers: {
            cookie: req.headers.cookie,
        }
    })

    if (res.ok) {
        return {
            auth: true,
            user: await res.json()
        }
    }

    console.log(await res.json())

    return {
        auth: false,
    }
}