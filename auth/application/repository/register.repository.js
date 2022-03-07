const db = require('diskdb');

class RegisterRepository {
    constructor() {
        this.AddUser = this.AddUser.bind(this);
        this.CheckUsername = this.CheckUsername.bind(this);
    }

    AddUser(user) {
        return db.users.save(user);
    }

    CheckUsername(username) {
        const cekUsername = db.users.find({ username: username });
        return cekUsername.length <= 0;

    }
}

module.exports = RegisterRepository;