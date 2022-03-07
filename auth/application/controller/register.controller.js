const validator = require("../../lib/validator");

const RegisterService = require("../service/register.service");

class RegisterController {
    constructor() {
        this.RegisterUser = this.RegisterUser.bind(this);

        this.registerService = new RegisterService();
    }

    async RegisterUser(req, res) {

        const requires = [
            'username', "name", "phone", "role"
        ]
        let errors = validator.checkBody(req.body, requires)
        if (errors) {
            return res.sendError(errors)
        }

        let register = await this.registerService.Register(req.body);
        if (register.code === 406) {
            res.sendError(null, "Failed Insert Database!", register.code);
        } else if (register.code === 400) {
            res.sendError(null, "Username was already used!", register.code);
        }

        res.sendSuccess(register.data, "User created!", register.code);
    }
}

module.exports = RegisterController;