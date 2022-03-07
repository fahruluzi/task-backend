class ProfileController {
    constructor() {
        this.ProfileUser = this.ProfileUser.bind(this);
    }

    async ProfileUser(req, res) {
        return res.sendSuccess(req.user, "User Authenticated!", 200);
    }
}

module.exports = ProfileController;