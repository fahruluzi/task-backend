const RegisterController = require("./controller/register.controller");
const LoginController = require("./controller/login.controller");

module.exports = (app) => {
    const registerController = new RegisterController();
    const loginController = new LoginController();

    app.post("/register", registerController.RegisterUser);
    app.post("/login", loginController.LoginUser);
};