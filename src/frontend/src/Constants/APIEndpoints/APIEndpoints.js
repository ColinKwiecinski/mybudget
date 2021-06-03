export default {
    base: "https://api.justinlim.me",
    testbase: "https://localhost:4000",
    handlers: {
        users: "/users",
        myuser: "/users/me",
<<<<<<< HEAD
        myuserAvatar: "/users/me/avatar",
        sessions: "/sessions",
        sessionsMine: "/sessions/mine",
        resetPasscode: "/resetcodes",
        passwords: "/passwords/"
=======
        myuserAvatar: "/v1/users/me/avatar",
        sessions: "/v1/sessions",
        sessionsMine: "/v1/sessions/mine",
        resetPasscode: "/v1/resetcodes",
        passwords: "/v1/passwords/"
>>>>>>> 878a3515870739a5995bdf9065769c41c137310f
    }
}