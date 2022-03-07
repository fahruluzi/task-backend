const validator = require("../../lib/validator");
const JWTLibrary = require("../../lib/jwt");

const LoginService = require("../service/login.service");

class LoginController {
    constructor() {
        this.LoginUser = this.LoginUser.bind(this);

        this.jwtLibrary = new JWTLibrary();

        this.loginService = new LoginService();
    }

    async LoginUser(req, res) {

        const requires = [
            'phone', "password"
        ]
        let errors = validator.checkBody(req.body, requires)
        if (errors) {
            return res.sendError(errors)
        }

        let result = await this.loginService.Login(req.body);
        if (result.code === 401) {
            return res.sendError(null, "Invalid Phone Number or Password!", result.code);
        }

        let now = new Date().getTime();
        let jwt = this.jwtLibrary.GenerateToken(result.data, now);

        return res.sendSuccess({
            "claims": {
                "name": result.data.name,
                "phone": result.data.phone,
                "role": result.data.role,
                "authenticated_at": now
            },
            "token" : jwt
        }, "User Authenticated!");
    }
}

module.exports = LoginController;