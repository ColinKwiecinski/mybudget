export default {
    base: "https://api.justinlim.me",
    testbase: "https://localhost:4000",
    handlers: {
        users: "/users",
        myuser: "/users/me",
        myuserAvatar: "/v1/users/me/avatar",
        sessions: "/v1/sessions",
        sessionsMine: "/v1/sessions/mine",
        resetPasscode: "/v1/resetcodes",
        passwords: "/v1/passwords/"
    }
}