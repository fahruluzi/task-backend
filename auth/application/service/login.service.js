const LoginRepository = require("../repository/login.repository");

class LoginService {
    constructor() {
        this.Login = this.Login.bind(this);

        this.loginRepository = new LoginRepository();
    }

    async Login(data) {
        let loginUser = this.loginRepository.GetUserByPhonePassword(data.phone, data.password);
        if (loginUser.length > 0) {
            return {
                success: true,
                code: 200,
                data : loginUser[0]
            }
        }

        return {
            success: false,
            code: 401
        }
    }
}

module.exports = LoginService;