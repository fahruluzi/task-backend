const authMiddleware = require("../middlewares/auth");

const RegisterController = require("./controller/register.controller");
const LoginController = require("./controller/login.controller");
const ProfileController = require("./controller/profile.controller");

module.exports = (app) => {
    const registerController = new RegisterController();
    const loginController = new LoginController();
    const profileController = new ProfileController();

    app.post("/register", registerController.RegisterUser);
    app.post("/login", loginController.LoginUser);

    app.use(authMiddleware)
    app.get("/profile", profileController.ProfileUser);
};