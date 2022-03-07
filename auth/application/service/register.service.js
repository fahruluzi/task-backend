const randomstring = require("randomstring");

const RegisterRepository = require("../repository/register.repository");

class RegisterService {
    constructor() {
        this.Register = this.Register.bind(this);

        this.registerRepository = new RegisterRepository();
    }

    async Register(data) {
        data.password = randomstring.generate(4);

        if (this.registerRepository.CheckUsername(data.username)) {
            // register user
            let addUser = this.registerRepository.AddUser(data)
            if (addUser) {
                return {
                    success: true,
                    code: 200,
                    data: addUser
                }
            }

            return {
                success: false,
                code: 406
            }
        }

        return {
            success: false,
            code: 400
        }
    }
}

module.exports = RegisterService;