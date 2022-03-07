const RegisterController = require("./controller/register.controller");

module.exports = (app) => {
    const registerController = new RegisterController();

    app.post("/register", registerController.RegisterUser);
};